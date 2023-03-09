package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type SearchCreateOrReplacePathParams struct {
	Index   string `pathParam:"style=simple,explode=false,name=index"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type SearchCreateOrReplaceRequest struct {
	PathParams SearchCreateOrReplacePathParams
	Request    shared.CreateOrReplaceDocumentRequest `request:"mediaType=application/json"`
}

type SearchCreateOrReplaceResponse struct {
	ContentType                     string
	CreateOrReplaceDocumentResponse *shared.CreateOrReplaceDocumentResponse
	Status                          *shared.Status
	StatusCode                      int
	RawResponse                     *http.Response
}
