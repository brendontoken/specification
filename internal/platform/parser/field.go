package parser

import (
	"fmt"
	"strings"
)

type IncludeIf struct {
	Field  string
	Values []string
}

type Field struct {
	Name           string
	Label          string
	Description    string
	Type           string
	Size           uint64
	Options        []string
	IncludeIf      IncludeIf `yaml:"include_if"`
	IncludeIfInt   IncludeIf `yaml:"include_if_int"`
	IncludeIfTrue  string    `yaml:"include_if_true"`
	IncludeIfFalse string    `yaml:"include_if_false"`
	IntValues      []int     `yaml:"int_values"`
	CharValues     []string  `yaml:"char_values"`
	Required       bool
	Example        string `yaml:"example"`
	ExampleHex     string `yaml:"example_hex"`
	Notes          string
	Computed       bool
	DisplayOrder   int
}

func (f Field) FieldName() string {
	return strings.Replace(f.Name, " ", "", -1)
}

func (f Field) FieldDescription() string {
	return strings.Trim(f.Description, " ")
}

func (f Field) DescriptionAbbr() string {
	return fmt.Sprintf(f.Description[:90])
}

func (f Field) FormType() string {
	if f.Name == "Payload" {
		return "json.RawAction"
	}

	if f.Type == "[]byte" || f.Type == "byte" {
		return "string"
	}

	return f.Type
}

func (f Field) GoTags() string {
	return fmt.Sprintf("`json:\"%s,omitempty\"`", f.SnakeCase())
}

func (f Field) SnakeCase() string {
	return SnakeCase(f.Name)
}

func (f Field) IsData() bool {
	return f.Type == "bin16"
}

func (f Field) IsBytes() bool {
	return f.FieldGoType() == "[]byte" && !f.IsData()
}

func (f Field) IsVarChar() bool {
	return f.Type == "varchar"
}

func (f Field) IsFixedChar() bool {
	return f.Type == "fixedchar" && f.Size != 1
}

func (f Field) IsVarBin() bool {
	return f.Type == "varbin"
}

func (f Field) IsNumeric() bool {
	s := strings.ToLower(f.Type)

	return strings.HasPrefix(s, "uint") ||
		strings.HasPrefix(s, "float") ||
		strings.HasPrefix(s, "int")
}

func (f Field) IsFloat() bool {
	s := strings.ToLower(f.Type)

	return strings.HasPrefix(f.Type, s)
}

func (f Field) IsArrayType() bool {
	return strings.HasSuffix(f.Type, "[]")
}

func (f Field) SingularDisplayType() string {
	return strings.Replace(f.Type, "[]", "", 1)
}

func (f Field) Length() uint64 {
	if strings.HasSuffix(f.Type, "[]") && f.Size == 0 {
		return 8
	}
	return f.Size
}

func (f Field) TypeURLCode() string {
	if f.IsResource() {
		typeName := f.Type
		if strings.HasSuffix(typeName, "[]") {
			typeName = typeName[:len(typeName)-2]
		}
		var name string
		switch typeName {
		case "Role":
			name = "roles"
		case "EntityType":
			name = "entities"
		case "Polity":
			name = "polities"
		case "CurrencyType":
			name = "currencies"
		case "RejectionCode":
			name = "rejections"
		case "Tag":
			name = "tags"
		}
		return "resource-" + KebabCase(name)
	}
	return "type-" + KebabCase(f.SingularType())
}

func (f Field) SingularType() string {
	return strings.Replace(f.FieldGoType(), "[]", "", 1)
}

func (f Field) FieldGoType() string {
	return GoType(f.Type, f.Size, f.Options)
}

func (f Field) IsInternalTypeArray() bool {
	return f.IsInternalType() && strings.HasSuffix(f.Type, "[]")
}

func (f Field) IsNativeTypeArray() bool {
	return !f.IsInternalType() && strings.HasSuffix(f.Type, "[]")
}

func (f Field) IsResource() bool {
	return IsResource(f.Type)
}

func (f Field) IsResourceTypeArray() bool {
	return IsResource(f.Type) && strings.HasSuffix(f.Type, "[]")
}

func (f Field) IsInternalType() bool {
	return IsInternalType(f.Type, f.Size)
}

func (f Field) IsNativeType() bool {
	return !IsInternalType(f.Type, f.Size) && !strings.HasSuffix(f.Type, "[]")
}

func (f Field) IsComplexType() bool {
	return f.IsInternalType()
}

func (f Field) Trimmable() bool {
	if f.Name == "Payload" {
		return false
	}

	if f.Type == "STRING" && f.Length() > 1 {
		return true
	}

	if f.Type == "SHA" {
		return true
	}

	return false
}
