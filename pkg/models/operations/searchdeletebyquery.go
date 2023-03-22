// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchDeleteByQueryRequest struct {
	DeleteByQueryRequest shared.DeleteByQueryRequest `request:"mediaType=application/json"`
	// The index name of the documents that needs deletion.
	Index string `pathParam:"style=simple,explode=false,name=index"`
	// The project name.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchDeleteByQueryResponse struct {
	ContentType string
	// OK
	DeleteByQueryResponse *shared.DeleteByQueryResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
