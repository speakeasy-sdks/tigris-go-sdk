package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementUpdateUserMetadataPathParams struct {
	MetadataKey string `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementUpdateUserMetadataRequest struct {
	PathParams ManagementUpdateUserMetadataPathParams
	Request    shared.UpdateUserMetadataRequest `request:"mediaType=application/json"`
}

type ManagementUpdateUserMetadataResponse struct {
	ContentType                string
	StatusCode                 int
	RawResponse                *http.Response
	Status                     *shared.Status
	UpdateUserMetadataResponse *shared.UpdateUserMetadataResponse
}
