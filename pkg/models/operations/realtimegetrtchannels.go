package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type RealtimeGetRTChannelsPathParams struct {
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimeGetRTChannelsRequest struct {
	PathParams RealtimeGetRTChannelsPathParams
}

type RealtimeGetRTChannelsResponse struct {
	ContentType           string
	GetRTChannelsResponse *shared.GetRTChannelsResponse
	Status                *shared.Status
	StatusCode            int
	RawResponse           *http.Response
}
