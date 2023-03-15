package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchDeleteIndexRequest struct {
	DeleteIndexRequest shared.DeleteIndexRequest `request:"mediaType=application/json"`
	Name               string                    `pathParam:"style=simple,explode=false,name=name"`
	Project            string                    `pathParam:"style=simple,explode=false,name=project"`
}

type SearchDeleteIndexResponse struct {
	ContentType         string
	DeleteIndexResponse *shared.DeleteIndexResponse
	Status              *shared.Status
	StatusCode          int
	RawResponse         *http.Response
}
