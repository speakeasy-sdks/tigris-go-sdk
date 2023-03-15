package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type RealtimeMessagesRequest struct {
	MessagesRequest shared.MessagesRequest `request:"mediaType=application/json"`
	Channel         string                 `pathParam:"style=simple,explode=false,name=channel"`
	Project         string                 `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimeMessagesResponse struct {
	ContentType      string
	MessagesResponse *shared.MessagesResponse
	Status           *shared.Status
	StatusCode       int
	RawResponse      *http.Response
}
