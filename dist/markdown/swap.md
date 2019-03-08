
# Swap Action

Swap Action -  Two parties (or more) want to swap a token (Atomic Swap) directly for another token.  At a minimum, Bitcoin is used in the txn for paying the necessary network/transaction fees.

The following breaks down the construction of a Swap Action. The action is constructed by building a single string from each of the elements in order.

<table class="waffle">
	<tr style='height:19px;'>
		<th style="width:6%" class="s0">Field</th>
		<th style="width:9%" class="s1">Label</th>
		<th style="width:9%" class="s1">Name</th>
		<th style="width:2%" class="s1">Bytes</th>
		<th style="width:29%" class="s1">Example Values</th>
		<th style="width:26%" class="s1">Comments</th>
		<th style="width:5%" class="s1">Data Type</th>
		<th style="width:14%" class="s2">Amendment Restrictions</th>
	</tr>
	<tr>
		<td class="s5" rowspan="100">Metadata (OP_RETURN Payload)</td>
		<td class="t6">Header[]</td>
		<td class="t6">Header Array</td>
		<td class="t6">-</td>
		<td class="t6">-</td>
		<td class="t6">Common header data for all messages</td>
		<td class="t6">Header</td>
		<td class="t7"></td>
	</tr>

	<tr>
		<td class="t10">Text Encoding</td>
		<td class="t10">TextEncoding</td>
		<td class="t10">1</td>
		<td class="t10" style="word-break:break-all">0</td>
		<td class="t10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all &#39;text&#39; data types. All &#39;string&#39; types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
		<td class="t10">uint8</td>
		<td class="t11">Can be changed by Issuer or Operator at their discretion.</td>
	</tr>

	<tr>
		<td class="t10">Asset Type 1</td>
		<td class="t10">AssetType1</td>
		<td class="t10">3</td>
		<td class="t10" style="word-break:break-all">RRE</td>
		<td class="t10">To be swapped for Asset2.</td>
		<td class="t10">string</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Asset ID 1</td>
		<td class="t10">AssetID1</td>
		<td class="t10">32</td>
		<td class="t10" style="word-break:break-all">ran2qsznhis53z</td>
		<td class="t10">To be swapped for Asset2.</td>
		<td class="t10">string</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Asset Type 2</td>
		<td class="t10">AssetType2</td>
		<td class="t10">3</td>
		<td class="t10" style="word-break:break-all">SHC</td>
		<td class="t10">To be swapped for Asset1.</td>
		<td class="t10">string</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Asset ID 2</td>
		<td class="t10">AssetID2</td>
		<td class="t10">32</td>
		<td class="t10" style="word-break:break-all">apm2qsznhks23z</td>
		<td class="t10">To be swapped for Asset1.</td>
		<td class="t10">string</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Offer Expiry</td>
		<td class="t10">OfferExpiry</td>
		<td class="t10">8</td>
		<td class="t10" style="word-break:break-all">Sun May 06 2018 06:00:00 GMT&#43;1000 (AEST)</td>
		<td class="t10">This prevents either party from holding on to the partially signed message as a form of an option.  Eg. the exchange at this price is valid for 30 mins.</td>
		<td class="t10">time</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Exchange Fee Currency</td>
		<td class="t10">ExchangeFeeCurrency</td>
		<td class="t10">3</td>
		<td class="t10" style="word-break:break-all">AUD</td>
		<td class="t10">BSV, USD, AUD, EUR, etc.</td>
		<td class="t10">string</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Exchange Fee Variable</td>
		<td class="t10">ExchangeFeeVar</td>
		<td class="t10">4</td>
		<td class="t10" style="word-break:break-all">0.005</td>
		<td class="t10">Percent of the value of the transaction</td>
		<td class="t10">float32</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Exchange Fee Fixed</td>
		<td class="t10">ExchangeFeeFixed</td>
		<td class="t10">4</td>
		<td class="t10" style="word-break:break-all">0.01</td>
		<td class="t10">Fixed fee</td>
		<td class="t10">float32</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Exchange Fee Address</td>
		<td class="t10">ExchangeFeeAddress</td>
		<td class="t10">34</td>
		<td class="t10" style="word-break:break-all">1HQ2ULuD7T5ykaucZ3KmTo4i29925Qa6ic</td>
		<td class="t10">Identifies the public address that the exchange fee should be paid to.</td>
		<td class="t10">string</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Qty of Asset 1 Sending Addresses</td>
		<td class="t10">Asset1SenderCount</td>
		<td class="t10">1</td>
		<td class="t10" style="word-break:break-all">0</td>
		<td class="t10">Asset 1 Sending Addresses</td>
		<td class="t10">uint8</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Asset 1 Senders</td>
		<td class="t10">Asset1Senders</td>
		<td class="t10">0</td>
		<td class="t10" style="word-break:break-all"></td>
		<td class="t10">Each element has the quantity of Asset1 tokens to be sent by the input address, which is referred to by the index.</td>
		<td class="t10">QuantityIndex[]</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Qty of Asset 1 Receiving Addresses</td>
		<td class="t10">Asset1ReceiverCount</td>
		<td class="t10">1</td>
		<td class="t10" style="word-break:break-all">0</td>
		<td class="t10"></td>
		<td class="t10">uint8</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Asset 1 Receivers</td>
		<td class="t10">Asset1Receivers</td>
		<td class="t10">0</td>
		<td class="t10" style="word-break:break-all"></td>
		<td class="t10">Each element has the quantity of Asset 1 tokens to be received by the output address, which is referred to by the index.</td>
		<td class="t10">TokenReceiver[]</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Qty of Asset 2 Sending Addresses</td>
		<td class="t10">Asset2SenderCount</td>
		<td class="t10">1</td>
		<td class="t10" style="word-break:break-all">0</td>
		<td class="t10">Asset 2 Sending Addresses</td>
		<td class="t10">uint8</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Asset 2 Senders</td>
		<td class="t10">Asset2Senders</td>
		<td class="t10">0</td>
		<td class="t10" style="word-break:break-all"></td>
		<td class="t10">Each element has the quantity of Asset2 tokens to be sent by the input address, which is referred to by the index.</td>
		<td class="t10">QuantityIndex[]</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Qty of Asset 2 Receiving Addresses</td>
		<td class="t10">Asset2ReceiverCount</td>
		<td class="t10">1</td>
		<td class="t10" style="word-break:break-all">0</td>
		<td class="t10"></td>
		<td class="t10">uint8</td>
		<td class="t11"></td>
	</tr>

	<tr>
		<td class="t10">Asset 2 Receivers</td>
		<td class="t10">Asset2Receivers</td>
		<td class="t10">0</td>
		<td class="t10" style="word-break:break-all"></td>
		<td class="t10">Each element contains the quantity of Asset 2 tokens to be received by the output address, which is referred to by the index.</td>
		<td class="t10">TokenReceiver[]</td>
		<td class="t11"></td>
	</tr>

</table>