package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisDescribeDatabaseRequest struct {
	DescribeDatabaseRequest shared.DescribeDatabaseRequest `request:"mediaType=application/json"`
	Project                 string                         `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDescribeDatabaseResponse struct {
	ContentType              string
	DescribeDatabaseResponse *shared.DescribeDatabaseResponse
	Status                   *shared.Status
	StatusCode               int
	RawResponse              *http.Response
}
