package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementUpdateUserMetadataRequest struct {
	UpdateUserMetadataRequest shared.UpdateUserMetadataRequest `request:"mediaType=application/json"`
	MetadataKey               string                           `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementUpdateUserMetadataResponse struct {
	ContentType                string
	StatusCode                 int
	RawResponse                *http.Response
	Status                     *shared.Status
	UpdateUserMetadataResponse *shared.UpdateUserMetadataResponse
}
