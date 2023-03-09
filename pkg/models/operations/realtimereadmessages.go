package operations

import (
	"net/http"
	"tigris-core/pkg/models/shared"
)

type RealtimeReadMessagesPathParams struct {
	Channel string `pathParam:"style=simple,explode=false,name=channel"`
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimeReadMessagesQueryParams struct {
	End       *string `queryParam:"style=form,explode=true,name=end"`
	Event     *string `queryParam:"style=form,explode=true,name=event"`
	Limit     *int64  `queryParam:"style=form,explode=true,name=limit"`
	SessionID *string `queryParam:"style=form,explode=true,name=session_id"`
	SocketID  *string `queryParam:"style=form,explode=true,name=socket_id"`
	Start     *string `queryParam:"style=form,explode=true,name=start"`
}

type RealtimeReadMessagesRequest struct {
	PathParams  RealtimeReadMessagesPathParams
	QueryParams RealtimeReadMessagesQueryParams
}

type RealtimeReadMessagesResponse struct {
	ContentType          string
	ReadMessagesResponse *shared.ReadMessagesResponse
	Status               *shared.Status
	StatusCode           int
	RawResponse          *http.Response
}
