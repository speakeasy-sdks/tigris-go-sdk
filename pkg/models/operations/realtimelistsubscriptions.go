package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type RealtimeListSubscriptionsRequest struct {
	Channel  string `pathParam:"style=simple,explode=false,name=channel"`
	Page     *int   `queryParam:"style=form,explode=true,name=page"`
	PageSize *int   `queryParam:"style=form,explode=true,name=page_size"`
	Project  string `pathParam:"style=simple,explode=false,name=project"`
}

type RealtimeListSubscriptionsResponse struct {
	ContentType              string
	ListSubscriptionResponse *shared.ListSubscriptionResponse
	Status                   *shared.Status
	StatusCode               int
	RawResponse              *http.Response
}
