package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementInsertUserMetadataRequest struct {
	InsertUserMetadataRequest shared.InsertUserMetadataRequest `request:"mediaType=application/json"`
	MetadataKey               string                           `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementInsertUserMetadataResponse struct {
	ContentType                string
	InsertUserMetadataResponse *shared.InsertUserMetadataResponse
	Status                     *shared.Status
	StatusCode                 int
	RawResponse                *http.Response
}
