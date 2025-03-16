// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package gen

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	strictgin "github.com/oapi-codegen/runtime/strictmiddleware/gin"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// ヘルスチェックAPI
	// (GET /v1/healthcheck)
	Healthcheck(c *gin.Context)
	// ユーザの新規登録API
	// (POST /v1/users)
	CreateUser(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// Healthcheck operation middleware
func (siw *ServerInterfaceWrapper) Healthcheck(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.Healthcheck(c)
}

// CreateUser operation middleware
func (siw *ServerInterfaceWrapper) CreateUser(c *gin.Context) {

	c.Set(BearerAuthScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateUser(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/v1/healthcheck", wrapper.Healthcheck)
	router.POST(options.BaseURL+"/v1/users", wrapper.CreateUser)
}

type AlreadyExistsResponse struct {
}

type BadRequestResponse struct {
}

type InternalServerErrorResponse struct {
}

type UnauthorizedResponse struct {
}

type HealthcheckRequestObject struct {
}

type HealthcheckResponseObject interface {
	VisitHealthcheckResponse(w http.ResponseWriter) error
}

type Healthcheck200JSONResponse HealthCheck

func (response Healthcheck200JSONResponse) VisitHealthcheckResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type CreateUserRequestObject struct {
	Body *CreateUserJSONRequestBody
}

type CreateUserResponseObject interface {
	VisitCreateUserResponse(w http.ResponseWriter) error
}

type CreateUser201JSONResponse CreateUserResponse

func (response CreateUser201JSONResponse) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type CreateUser400Response = BadRequestResponse

func (response CreateUser400Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(400)
	return nil
}

type CreateUser401Response = UnauthorizedResponse

func (response CreateUser401Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(401)
	return nil
}

type CreateUser409Response = AlreadyExistsResponse

func (response CreateUser409Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(409)
	return nil
}

type CreateUser500Response = InternalServerErrorResponse

func (response CreateUser500Response) VisitCreateUserResponse(w http.ResponseWriter) error {
	w.WriteHeader(500)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// ヘルスチェックAPI
	// (GET /v1/healthcheck)
	Healthcheck(ctx *gin.Context, request HealthcheckRequestObject) (HealthcheckResponseObject, error)
	// ユーザの新規登録API
	// (POST /v1/users)
	CreateUser(ctx *gin.Context, request CreateUserRequestObject) (CreateUserResponseObject, error)
}

type StrictHandlerFunc = strictgin.StrictGinHandlerFunc
type StrictMiddlewareFunc = strictgin.StrictGinMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// Healthcheck operation middleware
func (sh *strictHandler) Healthcheck(ctx *gin.Context) {
	var request HealthcheckRequestObject

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.Healthcheck(ctx, request.(HealthcheckRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "Healthcheck")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(HealthcheckResponseObject); ok {
		if err := validResponse.VisitHealthcheckResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}

// CreateUser operation middleware
func (sh *strictHandler) CreateUser(ctx *gin.Context) {
	var request CreateUserRequestObject

	var body CreateUserJSONRequestBody
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Error(err)
		return
	}
	request.Body = &body

	handler := func(ctx *gin.Context, request interface{}) (interface{}, error) {
		return sh.ssi.CreateUser(ctx, request.(CreateUserRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "CreateUser")
	}

	response, err := handler(ctx, request)

	if err != nil {
		ctx.Error(err)
		ctx.Status(http.StatusInternalServerError)
	} else if validResponse, ok := response.(CreateUserResponseObject); ok {
		if err := validResponse.VisitCreateUserResponse(ctx.Writer); err != nil {
			ctx.Error(err)
		}
	} else if response != nil {
		ctx.Error(fmt.Errorf("unexpected response type: %T", response))
	}
}
