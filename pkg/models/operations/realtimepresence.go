package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type RealtimePresencePathParams struct {
	Channel string `pathParam:"style=simple,explode=false,name=channel"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimePresenceRequest struct {
	PathParams RealtimePresencePathParams
}

type RealtimePresenceResponse struct {
	ContentType      string
	PresenceResponse *shared.PresenceResponse
	Status           *shared.Status
	StatusCode       int
	RawResponse      *http.Response
}
