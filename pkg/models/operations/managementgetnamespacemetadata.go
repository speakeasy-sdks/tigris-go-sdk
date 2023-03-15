package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementGetNamespaceMetadataRequest struct {
	GetNamespaceMetadataRequest shared.GetNamespaceMetadataRequest `request:"mediaType=application/json"`
	MetadataKey                 string                             `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementGetNamespaceMetadataResponse struct {
	ContentType                  string
	GetNamespaceMetadataResponse *shared.GetNamespaceMetadataResponse
	Status                       *shared.Status
	StatusCode                   int
	RawResponse                  *http.Response
}
