// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateOrUpdateIndexRequest struct {
	// search index name.
	Name *string `json:"name,omitempty"`
	// If set to `true` then a conflict with HTTP Status code 409 is returned if an index already exists. The default is false.
	OnlyCreate *bool `json:"only_create,omitempty"`
	// Tigris project name.
	Project *string `json:"project,omitempty"`
	// schema of the index. The schema specifications are same as JSON schema specification defined <a href="https://json-schema.org/specification.html" title="here">here</a>.<p></p> Schema example: `{  "title": "ecommerce_index",  "description": "an ecommerce store search index",  "properties": {    "name": {      "description": "Name of the product",      "type": "string",      "maxLength": 128    },    "brand": {      "description": "Brand of the product",      "type": "string"    },    "price": {      "description": "Price of the product",      "type": "number"    }  } }`
	Schema *string `json:"schema,omitempty"`
}

func (o *CreateOrUpdateIndexRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateOrUpdateIndexRequest) GetOnlyCreate() *bool {
	if o == nil {
		return nil
	}
	return o.OnlyCreate
}

func (o *CreateOrUpdateIndexRequest) GetProject() *string {
	if o == nil {
		return nil
	}
	return o.Project
}

func (o *CreateOrUpdateIndexRequest) GetSchema() *string {
	if o == nil {
		return nil
	}
	return o.Schema
}
