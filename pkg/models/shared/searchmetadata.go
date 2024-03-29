// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SearchMetadata struct {
	// Total number of search results across all pages
	Found *int64 `json:"found,omitempty"`
	// Pagination metadata for SearchResponse
	Page *Page `json:"page,omitempty"`
	// Number representing the total pages of results
	TotalPages *int `json:"total_pages,omitempty"`
}

func (o *SearchMetadata) GetFound() *int64 {
	if o == nil {
		return nil
	}
	return o.Found
}

func (o *SearchMetadata) GetPage() *Page {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *SearchMetadata) GetTotalPages() *int {
	if o == nil {
		return nil
	}
	return o.TotalPages
}
