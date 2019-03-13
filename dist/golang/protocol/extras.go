package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

const (
	// Version of the Tokenized protocol.
	Version = uint8(0)

	// OpReturn (OP_RETURN) is a script opcode is used to mark a transaction
	// output as invalid, and can be used to add data to a TX.
	OpReturn = 0x6a

	// OpPushdata1 represent the OP_PUSHDATA1 opcode.
	OpPushdata1 = byte(0x4c)

	// OpPushdata2 represents the OP_PUSHDATA2 opcode.
	OpPushdata2 = byte(0x4d)

	// OpPushdata4 represents the OP_PUSHDATA4 opcode.
	OpPushdata4 = byte(0x4e)

	// OpPushdata1Max is the maximum number of bytes that can be used in the
	// OP_PUSHDATA1 opcode.
	OpPushdata1Max = 255

	// OpPushdata2Max is the maximum number of bytes that can be used in the
	// OP_PUSHDATA2 opcode.
	OpPushdata2Max = 65535
)

// PayloadMessage is the interface for messages that are derived from
// payloads, such as asset types.
type PayloadMessage interface {
	io.Writer
	Type() string
	Serialize() ([]byte, error)
}

// NewPayloadMessageFromCode returns the approriate PayloadMessage for the
// given code.
func NewPayloadMessageFromCode(code []byte) (PayloadMessage, error) {
	s := string(code)

	switch s {
	case CodeShareCommon:
		return NewShareCommon(), nil
	}

	return nil, fmt.Errorf("No asset type for code %s", code)
}

// OpReturnMessage implements a base interface for all message types.
type OpReturnMessage interface {
	PayloadMessage
	String() string
	PayloadMessage() (PayloadMessage, error)
}

// New returns a new message, as an OpReturnMessage, from the OP_RETURN
// payload.
func New(b []byte) (OpReturnMessage, error) {
	code, err := Code(b)
	if err != nil {
		return nil, err
	}

	t, ok := TypeMapping[code]
	if !ok {
		return nil, fmt.Errorf("Unknown code :  %v", code)
	}

	if _, err := t.Write(b); err != nil {
		return nil, err
	}

	return t, nil
}

// Code returns the identifying code from the OP_RETURN payload.
func Code(b []byte) (string, error) {
	if len(b) < 9 || b[0] != OpReturn {
		return "", errors.New("Not an OP_RETURN payload")
	}

	offset := 7

	if b[1] < 0x4c {
		offset = 6
	}

	return string(b[offset : offset+2]), nil
}

// NewHeaderForCode returns a new Header with the given code and size.
func NewHeaderForCode(code []byte, size int) (*Header, error) {
	// work out which opcode to use depending on size of the data.
	opcode := OpPushdata1

	if size > OpPushdata2Max {
		opcode = OpPushdata4
	} else if size > OpPushdata1Max {
		opcode = OpPushdata2
	}

	lenPayload, err := uintToBytes(uint64(size))
	if err != nil {
		return nil, err
	}

	h := Header{
		ProtocolID:       ProtocolID,
		OpPushdata:       opcode,
		LenActionPayload: lenPayload,
		Version:          Version,
		ActionPrefix:     code,
	}

	return &h, nil
}

func WriteVarChar(buf *bytes.Buffer, value string, maxSize uint64) error {
	var err error
	if maxSize < 256 {
		err = write(buf, uint8(len(value)))
	} else if maxSize < 65536 {
		err = write(buf, uint16(len(value)))
	} else if maxSize < 4294967296 {
		err = write(buf, uint32(len(value)))
	} else {
		err = write(buf, uint64(len(value)))
	}
	if err != nil {
		return err
	}

	return write(buf, []byte(value))
}

func ReadVarChar(buf *bytes.Buffer, maxSize uint64) (string, error) {
	var err error
	var size uint64
	if maxSize < 256 {
		var size8 uint8
		err = read(buf, &size8)
		size = uint64(size8)
	} else if maxSize < 65536 {
		var size16 uint16
		err = read(buf, &size16)
		size = uint64(size16)
	} else if maxSize < 4294967296 {
		var size32 uint32
		err = read(buf, &size32)
		size = uint64(size32)
	} else {
		err = read(buf, &size)
	}
	if err != nil {
		return "", err
	}

	data := make([]byte, size)
	err = readLen(buf, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFixedChar(buf *bytes.Buffer, value string, size uint64) error {
	if uint64(len(value)) > size {
		return errors.New(fmt.Sprintf("FixedChar too long %d > %d", len(value), size))
	}
	err := write(buf, []byte(value))
	if err != nil {
		return err
	}

	// Pad with zeroes
	if uint64(len(value)) < size {
		padCount := size - uint64(len(value))
		empty := make([]byte, padCount)
		for i := uint64(0); i < padCount; i++ {
			empty[i] = 0
		}
		err := write(buf, empty)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadFixedChar(buf *bytes.Buffer, size uint64) (string, error) {
	var err error
	data := make([]byte, size)
	err = readLen(buf, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
