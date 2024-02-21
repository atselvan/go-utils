package httputil

import (
	"net/url"

	"github.com/atselvan/go-utils/utils/config"
	"github.com/jarcoal/httpmock"
)

const (
	ServerStartupMsg          = "Starting the API server"
	ServerStartupSuccessMsg   = "The API server has started and is listening on %s"
	ServerUrlMsg              = "API server url: %s"
	ServerAPIDocsMsg          = "API server swagger docs url: %s"
	ApiHealthPath             = "/health"
	ContentTypeHeaderKey      = "Content-Type"
	AcceptHeaderKey           = "Accept"
	AuthorizationHeaderKey    = "Authorization"
	ApplicationJsonMIMEType   = "application/json"
	TextPlainMIMEType         = "text/plain"
	XWWWFromUrlEncodeMIMEType = "application/x-www-form-urlencoded"
	SwaggerPath               = "/swagger/*any"
	SwaggerSpecPathFormat     = "%s/swagger/doc.json"
	SwaggerUriPathFormat      = "%s/swagger/index.html"
)

// RestMsg represents a message returned by a REST API
type RestMsg struct {
	Message string `json:"message"`
}

// RestErr represents an error returned by a REST API
type RestErr struct {
	Error string `json:"error"`
}

// GetServerURL formats and returns the server url.
func GetServerURL(sc *config.ServerConfig, apiBasePath string) *url.URL {
	var host string
	if sc.Protocol == "http" {
		host = sc.Host + ":" + sc.Port
	} else {
		host = sc.Host
	}
	serverURL := &url.URL{
		Scheme: sc.Protocol,
		Host:   host,
		Path:   apiBasePath,
	}
	return serverURL
}

// NewStringToJsonResponder is a custom httpmock.Responder that takes the status code and a json string body
// and creates a responder for a http mock. This is a useful function when unit testing rest API responses.
func NewStringToJsonResponder(statusCode int, body string) httpmock.Responder {
	response := httpmock.NewBytesResponse(statusCode, []byte(body))
	response.Header.Set(ContentTypeHeaderKey, ApplicationJsonMIMEType)
	return httpmock.ResponderFromResponse(response)
}
