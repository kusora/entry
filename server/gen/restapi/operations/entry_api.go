// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/laincloud/entry/server/gen/restapi/operations/auth"
	"github.com/laincloud/entry/server/gen/restapi/operations/commands"
	"github.com/laincloud/entry/server/gen/restapi/operations/config"
	"github.com/laincloud/entry/server/gen/restapi/operations/container"
	"github.com/laincloud/entry/server/gen/restapi/operations/ping"
	"github.com/laincloud/entry/server/gen/restapi/operations/sessions"
)

// NewEntryAPI creates a new Entry instance
func NewEntryAPI(spec *loads.Document) *EntryAPI {
	return &EntryAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		ContainerAttachContainerHandler: container.AttachContainerHandlerFunc(func(params container.AttachContainerParams) middleware.Responder {
			return middleware.NotImplemented("operation ContainerAttachContainer has not yet been implemented")
		}),
		AuthAuthorizeHandler: auth.AuthorizeHandlerFunc(func(params auth.AuthorizeParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthAuthorize has not yet been implemented")
		}),
		ContainerEnterContainerHandler: container.EnterContainerHandlerFunc(func(params container.EnterContainerParams) middleware.Responder {
			return middleware.NotImplemented("operation ContainerEnterContainer has not yet been implemented")
		}),
		ConfigGetConfigHandler: config.GetConfigHandlerFunc(func(params config.GetConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation ConfigGetConfig has not yet been implemented")
		}),
		CommandsListCommandsHandler: commands.ListCommandsHandlerFunc(func(params commands.ListCommandsParams) middleware.Responder {
			return middleware.NotImplemented("operation CommandsListCommands has not yet been implemented")
		}),
		SessionsListSessionsHandler: sessions.ListSessionsHandlerFunc(func(params sessions.ListSessionsParams) middleware.Responder {
			return middleware.NotImplemented("operation SessionsListSessions has not yet been implemented")
		}),
		AuthLogoutHandler: auth.LogoutHandlerFunc(func(params auth.LogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthLogout has not yet been implemented")
		}),
		PingPingHandler: ping.PingHandlerFunc(func(params ping.PingParams) middleware.Responder {
			return middleware.NotImplemented("operation PingPing has not yet been implemented")
		}),
		SessionsReplaySessionHandler: sessions.ReplaySessionHandlerFunc(func(params sessions.ReplaySessionParams) middleware.Responder {
			return middleware.NotImplemented("operation SessionsReplaySession has not yet been implemented")
		}),
	}
}

/*EntryAPI the entry API */
type EntryAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/vnd.laincloud.entry.v3+json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/vnd.laincloud.entry.v3+json" mime type
	JSONProducer runtime.Producer

	// ContainerAttachContainerHandler sets the operation handler for the attach container operation
	ContainerAttachContainerHandler container.AttachContainerHandler
	// AuthAuthorizeHandler sets the operation handler for the authorize operation
	AuthAuthorizeHandler auth.AuthorizeHandler
	// ContainerEnterContainerHandler sets the operation handler for the enter container operation
	ContainerEnterContainerHandler container.EnterContainerHandler
	// ConfigGetConfigHandler sets the operation handler for the get config operation
	ConfigGetConfigHandler config.GetConfigHandler
	// CommandsListCommandsHandler sets the operation handler for the list commands operation
	CommandsListCommandsHandler commands.ListCommandsHandler
	// SessionsListSessionsHandler sets the operation handler for the list sessions operation
	SessionsListSessionsHandler sessions.ListSessionsHandler
	// AuthLogoutHandler sets the operation handler for the logout operation
	AuthLogoutHandler auth.LogoutHandler
	// PingPingHandler sets the operation handler for the ping operation
	PingPingHandler ping.PingHandler
	// SessionsReplaySessionHandler sets the operation handler for the replay session operation
	SessionsReplaySessionHandler sessions.ReplaySessionHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *EntryAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *EntryAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *EntryAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *EntryAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *EntryAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *EntryAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *EntryAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the EntryAPI
func (o *EntryAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.ContainerAttachContainerHandler == nil {
		unregistered = append(unregistered, "container.AttachContainerHandler")
	}

	if o.AuthAuthorizeHandler == nil {
		unregistered = append(unregistered, "auth.AuthorizeHandler")
	}

	if o.ContainerEnterContainerHandler == nil {
		unregistered = append(unregistered, "container.EnterContainerHandler")
	}

	if o.ConfigGetConfigHandler == nil {
		unregistered = append(unregistered, "config.GetConfigHandler")
	}

	if o.CommandsListCommandsHandler == nil {
		unregistered = append(unregistered, "commands.ListCommandsHandler")
	}

	if o.SessionsListSessionsHandler == nil {
		unregistered = append(unregistered, "sessions.ListSessionsHandler")
	}

	if o.AuthLogoutHandler == nil {
		unregistered = append(unregistered, "auth.LogoutHandler")
	}

	if o.PingPingHandler == nil {
		unregistered = append(unregistered, "ping.PingHandler")
	}

	if o.SessionsReplaySessionHandler == nil {
		unregistered = append(unregistered, "sessions.ReplaySessionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *EntryAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *EntryAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *EntryAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *EntryAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/vnd.laincloud.entry.v3+json":
			result["application/vnd.laincloud.entry.v3+json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *EntryAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/vnd.laincloud.entry.v3+json":
			result["application/vnd.laincloud.entry.v3+json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *EntryAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the entry API
func (o *EntryAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *EntryAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/attach"] = container.NewAttachContainer(o.context, o.ContainerAttachContainerHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/authorize"] = auth.NewAuthorize(o.context, o.AuthAuthorizeHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/enter"] = container.NewEnterContainer(o.context, o.ContainerEnterContainerHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/config"] = config.NewGetConfig(o.context, o.ConfigGetConfigHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/commands"] = commands.NewListCommands(o.context, o.CommandsListCommandsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/sessions"] = sessions.NewListSessions(o.context, o.SessionsListSessionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/logout"] = auth.NewLogout(o.context, o.AuthLogoutHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/ping"] = ping.NewPing(o.context, o.PingPingHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/sessions/{session_id}/replay"] = sessions.NewReplaySession(o.context, o.SessionsReplaySessionHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *EntryAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *EntryAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *EntryAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *EntryAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
