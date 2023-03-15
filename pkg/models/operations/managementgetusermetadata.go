package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementGetUserMetadataRequest struct {
	GetUserMetadataRequest shared.GetUserMetadataRequest `request:"mediaType=application/json"`
	MetadataKey            string                        `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementGetUserMetadataResponse struct {
	ContentType             string
	GetUserMetadataResponse *shared.GetUserMetadataResponse
	Status                  *shared.Status
	StatusCode              int
	RawResponse             *http.Response
}
