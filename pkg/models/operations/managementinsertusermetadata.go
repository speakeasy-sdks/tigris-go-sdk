package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementInsertUserMetadataPathParams struct {
	MetadataKey string `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementInsertUserMetadataRequest struct {
	PathParams ManagementInsertUserMetadataPathParams
	Request    shared.InsertUserMetadataRequest `request:"mediaType=application/json"`
}

type ManagementInsertUserMetadataResponse struct {
	ContentType                string
	InsertUserMetadataResponse *shared.InsertUserMetadataResponse
	Status                     *shared.Status
	StatusCode                 int
	RawResponse                *http.Response
}
