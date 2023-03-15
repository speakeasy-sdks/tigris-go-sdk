package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDropCollectionRequest struct {
	DropCollectionRequest shared.DropCollectionRequest `request:"mediaType=application/json"`
	Collection            string                       `pathParam:"style=simple,explode=false,name=collection"`
	Project               string                       `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDropCollectionResponse struct {
	ContentType            string
	DropCollectionResponse *shared.DropCollectionResponse
	Status                 *shared.Status
	StatusCode             int
	RawResponse            *http.Response
}
