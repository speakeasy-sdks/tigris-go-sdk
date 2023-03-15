package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementInsertNamespaceMetadataRequest struct {
	InsertNamespaceMetadataRequest shared.InsertNamespaceMetadataRequest `request:"mediaType=application/json"`
	MetadataKey                    string                                `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementInsertNamespaceMetadataResponse struct {
	ContentType                     string
	InsertNamespaceMetadataResponse *shared.InsertNamespaceMetadataResponse
	Status                          *shared.Status
	StatusCode                      int
	RawResponse                     *http.Response
}
