// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"net/http"
)

type TigrisListCollectionsRequest struct {
	// Optionally specify a database branch name to perform operation on
	Branch *string `queryParam:"style=form,explode=true,name=branch"`
	// Project name whose db is under target to list collections.
	Project string `pathParam:"style=simple,explode=false,name=project"`
}

type TigrisListCollectionsResponse struct {
	ContentType string
	// OK
	ListCollectionsResponse *shared.ListCollectionsResponse
	// Default error response
	Status      *shared.Status
	StatusCode  int
	RawResponse *http.Response
}
