package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type TigrisDescribeDatabasePathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisDescribeDatabaseRequest struct {
	PathParams TigrisDescribeDatabasePathParams
	Request    shared.DescribeDatabaseRequest `request:"mediaType=application/json"`
}

type TigrisDescribeDatabaseResponse struct {
	ContentType              string
	DescribeDatabaseResponse *shared.DescribeDatabaseResponse
	Status                   *shared.Status
	StatusCode               int
	RawResponse              *http.Response
}
