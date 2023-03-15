package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type SearchCreateOrReplaceRequest struct {
	CreateOrReplaceDocumentRequest shared.CreateOrReplaceDocumentRequest `request:"mediaType=application/json"`
	Index                          string                                `pathParam:"style=simple,explode=false,name=index"`
	Project                        string                                `pathParam:"style=simple,explode=false,name=project"`
}

type SearchCreateOrReplaceResponse struct {
	ContentType                     string
	CreateOrReplaceDocumentResponse *shared.CreateOrReplaceDocumentResponse
	Status                          *shared.Status
	StatusCode                      int
	RawResponse                     *http.Response
}
