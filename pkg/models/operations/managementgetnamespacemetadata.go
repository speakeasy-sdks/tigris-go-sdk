package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
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
