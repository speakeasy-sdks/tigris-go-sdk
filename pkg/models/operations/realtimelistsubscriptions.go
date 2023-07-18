// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

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

func (o *RealtimeListSubscriptionsRequest) GetChannel() string {
	if o == nil {
		return ""
	}
	return o.Channel
}

func (o *RealtimeListSubscriptionsRequest) GetPage() *int {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *RealtimeListSubscriptionsRequest) GetPageSize() *int {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *RealtimeListSubscriptionsRequest) GetProject() string {
	if o == nil {
		return ""
	}
	return o.Project
}

type RealtimeListSubscriptionsResponse struct {
	ContentType string
	// OK
	ListSubscriptionResponse *shared.ListSubscriptionResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}

func (o *RealtimeListSubscriptionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RealtimeListSubscriptionsResponse) GetListSubscriptionResponse() *shared.ListSubscriptionResponse {
	if o == nil {
		return nil
	}
	return o.ListSubscriptionResponse
}

func (o *RealtimeListSubscriptionsResponse) GetStatus() *shared.Status {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *RealtimeListSubscriptionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RealtimeListSubscriptionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
