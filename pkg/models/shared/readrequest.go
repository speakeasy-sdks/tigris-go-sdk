// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ReadRequestFields - To read specific fields from a document. Default is all.
type ReadRequestFields struct {
}

// ReadRequestFilter - Returns documents matching this filter. A filter can simply be a key, value pair where a key is the field name and the value would be the value for this field. Tigris also allows complex filtering by passing logical expressions. Logical filters are applied on two or more fields using `$or` and `$and`. A few examples of filters: <li> To read a user document where the id has a value 1: ```{"id": 1 }``` <li> To read all the user documents where the key "id" has a value 1 or 2 or 3: `{"$or": [{"id": 1}, {"id": 2}, {"id": 3}]}` Filter allows setting collation on an individual field level. To set collation for all the fields see options. The detailed documentation of the filter is <a href="https://docs.tigrisdata.com/overview/query#specification-1" title="here">here</a>.
type ReadRequestFilter struct {
}

type ReadRequest struct {
	// Optionally specify a database branch name to perform operation on
	Branch *string `json:"branch,omitempty"`
	// To read specific fields from a document. Default is all.
	Fields *ReadRequestFields `json:"fields,omitempty"`
	// Returns documents matching this filter. A filter can simply be a key, value pair where a key is the field name and the value would be the value for this field. Tigris also allows complex filtering by passing logical expressions. Logical filters are applied on two or more fields using `$or` and `$and`. A few examples of filters: <li> To read a user document where the id has a value 1: ```{"id": 1 }``` <li> To read all the user documents where the key "id" has a value 1 or 2 or 3: `{"$or": [{"id": 1}, {"id": 2}, {"id": 3}]}` Filter allows setting collation on an individual field level. To set collation for all the fields see options. The detailed documentation of the filter is <a href="https://docs.tigrisdata.com/overview/query#specification-1" title="here">here</a>.
	Filter *ReadRequestFilter `json:"filter,omitempty"`
	// Options that can be used to modify the results, for example "limit" to control the number of documents returned by the server.
	Options *ReadRequestOptions `json:"options,omitempty"`
	// Array of fields and corresponding sort orders to order the results. Ex: 1 `[{ "salary": "$desc" }]`, Ex: 2  `[{ "salary": "$asc"}]`
	Sort *string `json:"sort,omitempty"`
}

func (o *ReadRequest) GetBranch() *string {
	if o == nil {
		return nil
	}
	return o.Branch
}

func (o *ReadRequest) GetFields() *ReadRequestFields {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ReadRequest) GetFilter() *ReadRequestFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *ReadRequest) GetOptions() *ReadRequestOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *ReadRequest) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}
