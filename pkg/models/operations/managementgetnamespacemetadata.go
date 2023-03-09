package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type ManagementGetNamespaceMetadataPathParams struct {
	MetadataKey string `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementGetNamespaceMetadataRequest struct {
	PathParams ManagementGetNamespaceMetadataPathParams
	Request    shared.GetNamespaceMetadataRequest `request:"mediaType=application/json"`
}

type ManagementGetNamespaceMetadataResponse struct {
	ContentType                  string
	GetNamespaceMetadataResponse *shared.GetNamespaceMetadataResponse
	Status                       *shared.Status
	StatusCode                   int
	RawResponse                  *http.Response
}
