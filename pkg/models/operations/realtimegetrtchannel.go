package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type RealtimeGetRTChannelPathParams struct {
	Channel string `pathParam:"style=simple,explode=false,name=channel"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimeGetRTChannelRequest struct {
	PathParams RealtimeGetRTChannelPathParams
}

type RealtimeGetRTChannelResponse struct {
	ContentType          string
	GetRTChannelResponse *shared.GetRTChannelResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
