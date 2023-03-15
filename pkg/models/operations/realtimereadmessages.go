package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type RealtimeReadMessagesRequest struct {
	Channel   string  `pathParam:"style=simple,explode=false,name=channel"`
	End       *string `queryParam:"style=form,explode=true,name=end"`
	Event     *string `queryParam:"style=form,explode=true,name=event"`
	Limit     *int64  `queryParam:"style=form,explode=true,name=limit"`
	Project   string  `pathParam:"style=simple,explode=false,name=project"`
	SessionID *string `queryParam:"style=form,explode=true,name=session_id"`
	SocketID  *string `queryParam:"style=form,explode=true,name=socket_id"`
	Start     *string `queryParam:"style=form,explode=true,name=start"`
}

type RealtimeReadMessagesResponse struct {
	ContentType          string
	ReadMessagesResponse *shared.ReadMessagesResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
