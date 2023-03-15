package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type RealtimeGetRTChannelRequest struct {
	Channel string `pathParam:"style=simple,explode=false,name=channel"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimeGetRTChannelResponse struct {
	ContentType          string
	GetRTChannelResponse *shared.GetRTChannelResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
