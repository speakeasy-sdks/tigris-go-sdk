package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type RealtimePresenceRequest struct {
	Channel string `pathParam:"style=simple,explode=false,name=channel"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimePresenceResponse struct {
	ContentType      string
	PresenceResponse *shared.PresenceResponse
	Status           *shared.Status
	StatusCode       int
	RawResponse      *http.Response
}
