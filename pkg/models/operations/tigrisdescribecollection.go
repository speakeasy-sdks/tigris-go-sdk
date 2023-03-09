package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDescribeCollectionPathParams struct {
	Collection string `pathParam:"style=simple,explode=false,name=collection"`
	Project    string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDescribeCollectionRequest struct {
	PathParams TigrisDescribeCollectionPathParams
	Request    shared.DescribeCollectionRequest `request:"mediaType=application/json"`
}

type TigrisDescribeCollectionResponse struct {
	ContentType                string
	DescribeCollectionResponse *shared.DescribeCollectionResponse
	Status                     *shared.Status
	StatusCode                 int
	RawResponse                *http.Response
}
