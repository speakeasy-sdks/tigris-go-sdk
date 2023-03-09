package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type ManagementGetUserMetadataPathParams struct {
	MetadataKey string `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementGetUserMetadataRequest struct {
	PathParams ManagementGetUserMetadataPathParams
	Request    shared.GetUserMetadataRequest `request:"mediaType=application/json"`
}

type ManagementGetUserMetadataResponse struct {
	ContentType             string
	GetUserMetadataResponse *shared.GetUserMetadataResponse
	Status                  *shared.Status
	StatusCode              int
	RawResponse             *http.Response
}
