package tigris

import (
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/utils"
	"net/http"
	"time"
)

var ServerList = []string{
	"https://api.preview.tigrisdata.cloud",
	"http://localhost:8081",
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// SDK Documentation: # Overview
// This section is organized around HTTP APIs. The APIs accepts JSON requests and returns JSON-encoded responses.The APIs conforms to standard HTTP status codes.
//
// # Errors
// Tigris uses conventional HTTP response codes to indicate the success or failure of an API request.The response will  contain an error code or other information that reveals the reason of the error.
// The error response is in JSON format and is structured like this:
// ```
//
//	{
//	  "error": {
//	    "code": "ALREADY_EXISTS",
//	    "message": "row already exists"
//	  }
//	}
//
// ```
//
// ## Status 2xx
//
//	HTTP Code  | Reason
//	----------------|-------------
//	200 - OK | Everything worked as expected.
//
// ## Status 4xx
// Status codes in the `400-499` range indicate errors that have been caused by the requesting application (e.g., a malformed request body has been sent).
// Retrying such requests with the same request body is pointless and _will_ result in the same status code again. Some `4xx` errors can be handled programmatically. Change the request accordingly before retrying. Below you can find the most frequent 4XX errors and how to fix them.
//
//	HTTP Code  | Tigris Code | Reason
//	----------------|-------------|---------
//	400 - Bad Request | INVALID_ARGUMENT | The request was unacceptable, often due to missing a required parameter. <br>Examples: <li>Missing documents for write request</li><li>Collection or Database name is not following the allowed characters</li>
//	401 - Unauthorized | UNAUTHENTICATED | No valid API key provided. Check that your `api_key` and `api_secret` are correct.
//	403 - Forbidden | PERMISSION_DENIED | The API key doesn't have permissions to perform the request.
//	404 - Not Found | NOT_FOUND | The requested resource doesn't exist. <br>Examples: <li>Database or Collection doesn't exists</li><li>Access Token missing in the request header</li>
//	409 - Conflict | ALREADY_EXISTS | TThe request conflicts with another request (perhaps due to using the same idempotent key). <br>Examples: <li>Trying to insert document again for the primary key that already exists</li><li>Creating a database that already exists</li>
//	429 - Too Many Requests | RESOURCE_EXHAUSTED | Too many requests hit the API too quickly. We recommend an exponential backoff of your requests.
//
// ## Status 5xx
// The 5xx range indicate an error with Tigris servers (these are rare).
//
//	HTTP Code  | Tigris Code | Reason
//	----------------|-------------|---------
//	500 - Internal Server Error | UNKNOWN | Something went wrong on Tigris server side.
//	501 - Not Implemented       | UNIMPLEMENTED | Not supported by the Tigris server and cannot be handled.
//	502 - Bad Gateway           | BAD_GATEWAY | The API key doesn't have permissions to perform the request.
//	503 - Service Unavailable   | UNAVAILABLE | The server is not ready to handle the request. Common causes are a server that is down for maintenance or that is overloaded.
//	504 - Gateway Timeout       | DEADLINE_EXCEEDED | Tigris server is unable to process the request timely. Common causes are that request is expensive, or server is overloaded.
//
// # Idempotent Requests
//
//	Tigris provides idempotency guarantees when an API call is disrupted and either no response is returned or an HTTP
//	Status code `429,500,501,502,503` is returned. These errors ensure that the request is not committed and retrying the same request will have the same effect.
//
// # Limitations
// <li>Do not rely on case to distinguish between databases or collections names.</li> <li>Database Name and Collection Name cannot be empty and can only have the characters matches the regex: <code>^[a-zA-Z]+[a-zA-Z0-9_]+$</code>.</li> <li>Duplicate field names are not allowed. </li> <li>The maximum allowed document size is 100KB.</li> <li>The maximum allowed transaction size is 10MB.</li>
type Tigris struct {
	AppKey     *appKey
	Auth       *auth
	Cache      *cache
	Channel    *channel
	Collection *collection
	Database   *database
	Namespace  *namespace
	Project    *project
	Search     *search
	System     *system
	User       *user

	// Non-idiomatic field names below are to namespace fields from the fields names above to avoid name conflicts
	_defaultClient  HTTPClient
	_securityClient HTTPClient
	_security       *shared.Security
	_serverURL      string
	_language       string
	_sdkVersion     string
	_genVersion     string
}

type SDKOption func(*Tigris)

func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Tigris) {
		sdk._serverURL = serverURL
	}
}

func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Tigris) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk._serverURL = serverURL
	}
}

func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Tigris) {
		sdk._defaultClient = client
	}
}

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *Tigris) {
		sdk._security = &security
	}
}

func New(opts ...SDKOption) *Tigris {
	sdk := &Tigris{
		_language:   "go",
		_sdkVersion: "0.0.3",
		_genVersion: "1.9.2",
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk._defaultClient == nil {
		sdk._defaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk._securityClient == nil {

		if sdk._security != nil {
			sdk._securityClient = utils.ConfigureSecurityClient(sdk._defaultClient, sdk._security)
		} else {
			sdk._securityClient = sdk._defaultClient
		}

	}

	if sdk._serverURL == "" {
		sdk._serverURL = ServerList[0]
	}

	sdk.AppKey = newAppKey(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Auth = newAuth(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Cache = newCache(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Channel = newChannel(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Collection = newCollection(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Database = newDatabase(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Namespace = newNamespace(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Project = newProject(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.Search = newSearch(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.System = newSystem(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	sdk.User = newUser(
		sdk._defaultClient,
		sdk._securityClient,
		sdk._serverURL,
		sdk._language,
		sdk._sdkVersion,
		sdk._genVersion,
	)

	return sdk
}
