package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisCreateOrUpdateCollectionPathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisCreateOrUpdateCollectionRequest struct {
	PathParams TigrisCreateOrUpdateCollectionPathParams
	Request    shared.CreateOrUpdateCollectionRequest `request:"mediaType=application/json"`
}

type TigrisCreateOrUpdateCollectionResponse struct {
	ContentType                      string
	CreateOrUpdateCollectionResponse *shared.CreateOrUpdateCollectionResponse
	Status                           *shared.Status
	StatusCode                       int
	RawResponse                      *http.Response
}
