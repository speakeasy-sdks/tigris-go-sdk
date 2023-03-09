package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type RealtimeMessagesPathParams struct {
	Channel string `pathParam:"style=simple,explode=false,name=channel"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimeMessagesRequest struct {
	PathParams RealtimeMessagesPathParams
	Request    shared.MessagesRequest `request:"mediaType=application/json"`
}

type RealtimeMessagesResponse struct {
	ContentType      string
	MessagesResponse *shared.MessagesResponse
	Status           *shared.Status
	StatusCode       int
	RawResponse      *http.Response
}
