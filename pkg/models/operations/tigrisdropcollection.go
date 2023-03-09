package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDropCollectionPathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDropCollectionRequest struct {
	PathParams TigrisDropCollectionPathParams
	Request    shared.DropCollectionRequest `request:"mediaType=application/json"`
}

type TigrisDropCollectionResponse struct {
	ContentType            string
	DropCollectionResponse *shared.DropCollectionResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
