package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type SearchDeleteIndexPathParams struct {
	Name    string `pathParam:"style=simple,explode=false,name=name"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchDeleteIndexRequest struct {
	PathParams SearchDeleteIndexPathParams
	Request    shared.DeleteIndexRequest `request:"mediaType=application/json"`
}

type SearchDeleteIndexResponse struct {
	ContentType         string
	DeleteIndexResponse *shared.DeleteIndexResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
