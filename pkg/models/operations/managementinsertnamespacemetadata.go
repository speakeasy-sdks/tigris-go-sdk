package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type ManagementInsertNamespaceMetadataPathParams struct {
	MetadataKey string `pathParam:"style=simple,explode=false,name=metadataKey"`
}

type ManagementInsertNamespaceMetadataRequest struct {
	PathParams ManagementInsertNamespaceMetadataPathParams
	Request    shared.InsertNamespaceMetadataRequest `request:"mediaType=application/json"`
}

type ManagementInsertNamespaceMetadataResponse struct {
	ContentType                     string
	InsertNamespaceMetadataResponse *shared.InsertNamespaceMetadataResponse
	Status                          *shared.Status
	StatusCode                      int
	RawResponse                     *http.Response
}
