package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementUpdateNamespaceMetadataRequest struct {
	UpdateNamespaceMetadataRequest shared.UpdateNamespaceMetadataRequest `request:"mediaType=application/json"`
	MetadataKey                    string                                `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementUpdateNamespaceMetadataResponse struct {
	ContentType                     string
	StatusCode                      int
	RawResponse                     *http.Response
	Status                          *shared.Status
	UpdateNamespaceMetadataResponse *shared.UpdateNamespaceMetadataResponse
}
