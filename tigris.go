// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package tigrisgosdk

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/tigris-go-sdk/internal/hooks"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/tigris-go-sdk/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Tigris Cloud
	"https://api.preview.tigrisdata.cloud",
	// Localhost
	"http://localhost:8081",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client            HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// Tigris API Reference: # Overview
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
	// The auth section of API provides OAuth 2.0 APIs. Tigris supports pluggable OAuth provider. Pass the token in the headers for authentication, as an example `-H "Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6I"`(replace it with your token). All API requests must be made over HTTPS. Calls made over plain HTTP will fail. API requests without authentication will also fail.
	Auth *Auth
	// The Observability section has APIs that provides full visibility into the health, metrics, and monitoring of the Server.
	System *System
	// The Management section provide APIs that can be used to manage users, and app_keys.
	Namespace *Namespace
	// A User on the Tigris Platform.
	User *User
	// Every Tigris projects comes with a transactional document database built on FoundationDB, one of the most resilient and battle-tested open source distributed key-value store. A database is created automatically for you when you create a project.
	Project *Project
	// The application keys section provide APIs that can be used to manage application keys for your project. A single project can have one or more application keys.
	AppKey *AppKey
	// The cache section provide APIs that can be used to perform cache operations.
	Cache *Cache
	// The Database section provide APIs that can be used to interact with the database. A single Database can have one or more collections. A database is created automatically for you when you create a project.
	Database *Database
	// The Collections section provide APIs that can be used to manage collections. A collection can have one or more documents.
	Collection *Collection
	// The realtime section provide APIs that can be used realtime operations.
	Channel *Channel
	// The search section provides you APIs that can be used to implement powerful apps with search experiences. You can manage storing documents on your own or you can simply integrate it with your database.
	Search *Search

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*Tigris)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *Tigris) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *Tigris) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *Tigris) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *Tigris) {
		sdk.sdkConfiguration.Client = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(bearerAuth string) SDKOption {
	return func(sdk *Tigris) {
		security := shared.Security{BearerAuth: bearerAuth}
		sdk.sdkConfiguration.Security = withSecurity(&security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *Tigris) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *Tigris) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *Tigris {
	sdk := &Tigris{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "0.0.1",
			SDKVersion:        "0.26.3",
			GenVersion:        "2.291.0",
			UserAgent:         "speakeasy-sdk/go 0.26.3 2.291.0 0.0.1 github.com/speakeasy-sdks/tigris-go-sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Auth = newAuth(sdk.sdkConfiguration)

	sdk.System = newSystem(sdk.sdkConfiguration)

	sdk.Namespace = newNamespace(sdk.sdkConfiguration)

	sdk.User = newUser(sdk.sdkConfiguration)

	sdk.Project = newProject(sdk.sdkConfiguration)

	sdk.AppKey = newAppKey(sdk.sdkConfiguration)

	sdk.Cache = newCache(sdk.sdkConfiguration)

	sdk.Database = newDatabase(sdk.sdkConfiguration)

	sdk.Collection = newCollection(sdk.sdkConfiguration)

	sdk.Channel = newChannel(sdk.sdkConfiguration)

	sdk.Search = newSearch(sdk.sdkConfiguration)

	return sdk
}
