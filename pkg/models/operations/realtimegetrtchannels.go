package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type RealtimeGetRTChannelsRequest struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimeGetRTChannelsResponse struct {
	ContentType           string
	GetRTChannelsResponse *shared.GetRTChannelsResponse
	Status                *shared.Status
	StatusCode            int
	RawResponse           *http.Response
}
