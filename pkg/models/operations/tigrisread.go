package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisReadRequest struct {
	ReadRequest shared.ReadRequest `request:"mediaType=application/json"`
	Collection  string             `pathParam:"style=simple,explode=false,name=collection"`
	Project     string             `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisReadResponse struct {
	ContentType           string
	StatusCode            int
	RawResponse           *http.Response
	Status                *shared.Status
	StreamingReadResponse *shared.StreamingReadResponse
}
