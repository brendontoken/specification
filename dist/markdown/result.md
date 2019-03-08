
# Result Action

Result Action -  Once a vote has been completed the results are published.

The following breaks down the construction of a Result Action. The action is constructed by building a single string from each of the elements in order.

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
		<td class="g6">Header[]</td>
		<td class="g6">Header Array</td>
		<td class="g6">-</td>
		<td class="g6">-</td>
		<td class="g6">Common header data for all messages</td>
		<td class="g6">Header</td>
		<td class="g7"></td>
	</tr>

	<tr>
		<td class="g10">Text Encoding</td>
		<td class="g10">TextEncoding</td>
		<td class="g10">1</td>
		<td class="g10" style="word-break:break-all">0</td>
		<td class="g10"> 0 = ASCII, 1 = UTF-8, 2 = UTF-16, 3 = Unicode.  Encoding applies to all &#39;text&#39; data types. All &#39;string&#39; types will always be encoded with ASCII.  Where string is selected, all fields will be ASCII.</td>
		<td class="g10">uint8</td>
		<td class="g11">Can be changed by Issuer or Operator at their discretion.</td>
	</tr>

	<tr>
		<td class="g10">Asset Type</td>
		<td class="g10">AssetType</td>
		<td class="g10">3</td>
		<td class="g10" style="word-break:break-all">SHC</td>
		<td class="g10">eg. Share, Bond, Ticket</td>
		<td class="g10">string</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10">Asset ID</td>
		<td class="g10">AssetID</td>
		<td class="g10">32</td>
		<td class="g10" style="word-break:break-all">apm2qsznhks23z8d83u41s8019hyri3i</td>
		<td class="g10">Randomly generated base58 string.  Each Asset ID should be unique.  However, a Asset ID is always linked to a Contract that is identified by the public address of the Contract wallet. The Asset Type can be the leading bytes - a convention - to make it easy to identify that it is a token by humans. If its a Contract vote then can be null.</td>
		<td class="g10">string</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10">Proposal</td>
		<td class="g10">Proposal</td>
		<td class="g10">1</td>
		<td class="g10" style="word-break:break-all">0</td>
		<td class="g10">1 for a Proposal, 0 for an initiative that is requesting changes to specific subfields for modification. If this field is true, the subfields should be empty.  The smart contract cannot interpret the results of a vote when Proposal = 1.  All meaning is interpreted by the token owners and smart contract simply facilates the record keeping.  When Proposal = 0, the smart contract always assumes the first choice is a &#39;yes&#39;, or &#39;pass&#39;, if the threshold is met, and will process the proposed changes accordingly.</td>
		<td class="g10">bool</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10"></td>
		<td class="g10">ProposedChangesCount</td>
		<td class="g10">1</td>
		<td class="g10" style="word-break:break-all">0</td>
		<td class="g10"></td>
		<td class="g10">uint8</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10">Proposed Changes</td>
		<td class="g10">ProposedChanges</td>
		<td class="g10">0</td>
		<td class="g10" style="word-break:break-all"></td>
		<td class="g10">Each element contains details of which fields to modify, or delete. Because the number of fields in a Contract and Asset is dynamic due to some fields being able to be repeated, the index value of the field needs to be calculated against the Contract or Asset the changes are to apply to. In the event of a Vote being created from this Initiative, the changes will be applied to the version of the Contract or Asset at that time.</td>
		<td class="g10">Amendment[]</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10">Vote Txn ID</td>
		<td class="g10">VoteTxnID</td>
		<td class="g10">32</td>
		<td class="g10" style="word-break:break-all">f2318be9fb3f73e53a29868beae46b42911c2116f979a5d3284face90746cb37</td>
		<td class="g10">Link to the Vote Action txn.</td>
		<td class="g10">sha256</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10">VoteOptionsCount</td>
		<td class="g10">VoteOptionsCount</td>
		<td class="g10">1</td>
		<td class="g10" style="word-break:break-all">1</td>
		<td class="g10">Number of Vote Options to follow.</td>
		<td class="g10">uint8</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10">Option X Tally</td>
		<td class="g10">OptionXTally</td>
		<td class="g10">8</td>
		<td class="g10" style="word-break:break-all">3000</td>
		<td class="g10">Number of valid votes counted for Option X</td>
		<td class="g10">uint64</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10">Result</td>
		<td class="g10">Result</td>
		<td class="g10">0</td>
		<td class="g10" style="word-break:break-all">2</td>
		<td class="g10">Length 1-255 bytes. 0 is not valid. The Option with the most votes. In the event of a draw for 1st place, all winning options are listed. </td>
		<td class="g10">nvarchar8</td>
		<td class="g11"></td>
	</tr>

	<tr>
		<td class="g10">Timestamp</td>
		<td class="g10">Timestamp</td>
		<td class="g10">8</td>
		<td class="g10" style="word-break:break-all">1551767413250187179</td>
		<td class="g10">Timestamp in nanoseconds of when the smart contract created the action.</td>
		<td class="g10">timestamp</td>
		<td class="g11">Cannot be changed by issuer, operator. Smart contract controls.</td>
	</tr>

</table>