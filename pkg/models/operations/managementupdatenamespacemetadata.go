package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementUpdateNamespaceMetadataPathParams struct {
	MetadataKey string `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementUpdateNamespaceMetadataRequest struct {
	PathParams ManagementUpdateNamespaceMetadataPathParams
	Request    shared.UpdateNamespaceMetadataRequest `request:"mediaType=application/json"`
}

type ManagementUpdateNamespaceMetadataResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	Status                          *shared.Status
	UpdateNamespaceMetadataResponse *shared.UpdateNamespaceMetadataResponse
}
