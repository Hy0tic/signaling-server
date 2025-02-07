// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gofiber/fiber/v2"
	"github.com/oapi-codegen/runtime"
)

// PostLivekitGenerateTokenForHostRoomJSONBody defines parameters for PostLivekitGenerateTokenForHostRoom.
type PostLivekitGenerateTokenForHostRoomJSONBody struct {
	// Room The name of the room.
	Room string `json:"room"`

	// Username The name of the user hosting the room.
	Username string `json:"username"`
}

// GenerateTokenForJoinRoomJSONBody defines parameters for GenerateTokenForJoinRoom.
type GenerateTokenForJoinRoomJSONBody struct {
	// Room The name of the room.
	Room string `json:"room"`

	// Username The name of the user joining the room.
	Username string `json:"username"`
}

// GetLivekitGetUsersInRoomParams defines parameters for GetLivekitGetUsersInRoom.
type GetLivekitGetUsersInRoomParams struct {
	// Room The name of the room to retrieve users from.
	Room string `form:"room" json:"room"`
}

// GetLivekitRoomCheckParams defines parameters for GetLivekitRoomCheck.
type GetLivekitRoomCheckParams struct {
	RoomName string `form:"roomName" json:"roomName"`
	Username string `form:"username" json:"username"`
}

// PostLivekitGenerateTokenForHostRoomJSONRequestBody defines body for PostLivekitGenerateTokenForHostRoom for application/json ContentType.
type PostLivekitGenerateTokenForHostRoomJSONRequestBody PostLivekitGenerateTokenForHostRoomJSONBody

// GenerateTokenForJoinRoomJSONRequestBody defines body for GenerateTokenForJoinRoom for application/json ContentType.
type GenerateTokenForJoinRoomJSONRequestBody GenerateTokenForJoinRoomJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Generate a token for a host in a room.
	// (POST /livekit/generateTokenForHostRoom)
	PostLivekitGenerateTokenForHostRoom(c *fiber.Ctx) error
	// Generate a token for a user to join a room.
	// (POST /livekit/generateTokenForJoinRoom)
	GenerateTokenForJoinRoom(c *fiber.Ctx) error
	// Get a list of users in a room.
	// (GET /livekit/getUsersInRoom)
	GetLivekitGetUsersInRoom(c *fiber.Ctx, params GetLivekitGetUsersInRoomParams) error
	// Check if a user is in a room.
	// (GET /livekit/roomCheck)
	GetLivekitRoomCheck(c *fiber.Ctx, params GetLivekitRoomCheckParams) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// PostLivekitGenerateTokenForHostRoom operation middleware
func (siw *ServerInterfaceWrapper) PostLivekitGenerateTokenForHostRoom(c *fiber.Ctx) error {

	return siw.Handler.PostLivekitGenerateTokenForHostRoom(c)
}

// GenerateTokenForJoinRoom operation middleware
func (siw *ServerInterfaceWrapper) GenerateTokenForJoinRoom(c *fiber.Ctx) error {

	return siw.Handler.GenerateTokenForJoinRoom(c)
}

// GetLivekitGetUsersInRoom operation middleware
func (siw *ServerInterfaceWrapper) GetLivekitGetUsersInRoom(c *fiber.Ctx) error {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLivekitGetUsersInRoomParams

	var query url.Values
	query, err = url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
	}

	// ------------- Required query parameter "room" -------------

	if paramValue := c.Query("room"); paramValue != "" {

	} else {
		err = fmt.Errorf("Query argument room is required, but not found")
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	err = runtime.BindQueryParameter("form", true, true, "room", query, &params.Room)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter room: %w", err).Error())
	}

	return siw.Handler.GetLivekitGetUsersInRoom(c, params)
}

// GetLivekitRoomCheck operation middleware
func (siw *ServerInterfaceWrapper) GetLivekitRoomCheck(c *fiber.Ctx) error {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetLivekitRoomCheckParams

	var query url.Values
	query, err = url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
	}

	// ------------- Required query parameter "roomName" -------------

	if paramValue := c.Query("roomName"); paramValue != "" {

	} else {
		err = fmt.Errorf("Query argument roomName is required, but not found")
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	err = runtime.BindQueryParameter("form", true, true, "roomName", query, &params.RoomName)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter roomName: %w", err).Error())
	}

	// ------------- Required query parameter "username" -------------

	if paramValue := c.Query("username"); paramValue != "" {

	} else {
		err = fmt.Errorf("Query argument username is required, but not found")
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	err = runtime.BindQueryParameter("form", true, true, "username", query, &params.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter username: %w", err).Error())
	}

	return siw.Handler.GetLivekitRoomCheck(c, params)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(m)
	}

	router.Post(options.BaseURL+"/livekit/generateTokenForHostRoom", wrapper.PostLivekitGenerateTokenForHostRoom)

	router.Post(options.BaseURL+"/livekit/generateTokenForJoinRoom", wrapper.GenerateTokenForJoinRoom)

	router.Get(options.BaseURL+"/livekit/getUsersInRoom", wrapper.GetLivekitGetUsersInRoom)

	router.Get(options.BaseURL+"/livekit/roomCheck", wrapper.GetLivekitRoomCheck)

}

type PostLivekitGenerateTokenForHostRoomRequestObject struct {
	Body *PostLivekitGenerateTokenForHostRoomJSONRequestBody
}

type PostLivekitGenerateTokenForHostRoomResponseObject interface {
	VisitPostLivekitGenerateTokenForHostRoomResponse(ctx *fiber.Ctx) error
}

type PostLivekitGenerateTokenForHostRoom200JSONResponse struct {
	// Token The generated authentication token.
	Token *string `json:"token,omitempty"`
}

func (response PostLivekitGenerateTokenForHostRoom200JSONResponse) VisitPostLivekitGenerateTokenForHostRoomResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type PostLivekitGenerateTokenForHostRoom400Response struct {
}

func (response PostLivekitGenerateTokenForHostRoom400Response) VisitPostLivekitGenerateTokenForHostRoomResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type PostLivekitGenerateTokenForHostRoom500Response struct {
}

func (response PostLivekitGenerateTokenForHostRoom500Response) VisitPostLivekitGenerateTokenForHostRoomResponse(ctx *fiber.Ctx) error {
	ctx.Status(500)
	return nil
}

type GenerateTokenForJoinRoomRequestObject struct {
	Body *GenerateTokenForJoinRoomJSONRequestBody
}

type GenerateTokenForJoinRoomResponseObject interface {
	VisitGenerateTokenForJoinRoomResponse(ctx *fiber.Ctx) error
}

type GenerateTokenForJoinRoom200JSONResponse struct {
	// Host Username of the host of the room.
	Host *string `json:"host,omitempty"`

	// Participants List of users in the room.
	Participants *[]string `json:"participants,omitempty"`

	// Token The generated authentication token.
	Token *string `json:"token,omitempty"`
}

func (response GenerateTokenForJoinRoom200JSONResponse) VisitGenerateTokenForJoinRoomResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GenerateTokenForJoinRoom400JSONResponse struct {
	// Error Error message describing the invalid input.
	Error *string `json:"error,omitempty"`
}

func (response GenerateTokenForJoinRoom400JSONResponse) VisitGenerateTokenForJoinRoomResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(400)

	return ctx.JSON(&response)
}

type GenerateTokenForJoinRoom500JSONResponse struct {
	// Error Error message describing the server error.
	Error *string `json:"error,omitempty"`
}

func (response GenerateTokenForJoinRoom500JSONResponse) VisitGenerateTokenForJoinRoomResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(500)

	return ctx.JSON(&response)
}

type GetLivekitGetUsersInRoomRequestObject struct {
	Params GetLivekitGetUsersInRoomParams
}

type GetLivekitGetUsersInRoomResponseObject interface {
	VisitGetLivekitGetUsersInRoomResponse(ctx *fiber.Ctx) error
}

type GetLivekitGetUsersInRoom200JSONResponse struct {
	// Room The room name.
	Room *string `json:"room,omitempty"`

	// Users A list of usernames currently in the room.
	Users *[]string `json:"users,omitempty"`
}

func (response GetLivekitGetUsersInRoom200JSONResponse) VisitGetLivekitGetUsersInRoomResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetLivekitGetUsersInRoom400Response struct {
}

func (response GetLivekitGetUsersInRoom400Response) VisitGetLivekitGetUsersInRoomResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

type GetLivekitGetUsersInRoom500Response struct {
}

func (response GetLivekitGetUsersInRoom500Response) VisitGetLivekitGetUsersInRoomResponse(ctx *fiber.Ctx) error {
	ctx.Status(500)
	return nil
}

type GetLivekitRoomCheckRequestObject struct {
	Params GetLivekitRoomCheckParams
}

type GetLivekitRoomCheckResponseObject interface {
	VisitGetLivekitRoomCheckResponse(ctx *fiber.Ctx) error
}

type GetLivekitRoomCheck200JSONResponse struct {
	RoomExists        *bool `json:"roomExists,omitempty"`
	UsernameAvailable *bool `json:"usernameAvailable,omitempty"`
}

func (response GetLivekitRoomCheck200JSONResponse) VisitGetLivekitRoomCheckResponse(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Content-Type", "application/json")
	ctx.Status(200)

	return ctx.JSON(&response)
}

type GetLivekitRoomCheck400Response struct {
}

func (response GetLivekitRoomCheck400Response) VisitGetLivekitRoomCheckResponse(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Generate a token for a host in a room.
	// (POST /livekit/generateTokenForHostRoom)
	PostLivekitGenerateTokenForHostRoom(ctx context.Context, request PostLivekitGenerateTokenForHostRoomRequestObject) (PostLivekitGenerateTokenForHostRoomResponseObject, error)
	// Generate a token for a user to join a room.
	// (POST /livekit/generateTokenForJoinRoom)
	GenerateTokenForJoinRoom(ctx context.Context, request GenerateTokenForJoinRoomRequestObject) (GenerateTokenForJoinRoomResponseObject, error)
	// Get a list of users in a room.
	// (GET /livekit/getUsersInRoom)
	GetLivekitGetUsersInRoom(ctx context.Context, request GetLivekitGetUsersInRoomRequestObject) (GetLivekitGetUsersInRoomResponseObject, error)
	// Check if a user is in a room.
	// (GET /livekit/roomCheck)
	GetLivekitRoomCheck(ctx context.Context, request GetLivekitRoomCheckRequestObject) (GetLivekitRoomCheckResponseObject, error)
}

type StrictHandlerFunc func(ctx *fiber.Ctx, args interface{}) (interface{}, error)

type StrictMiddlewareFunc func(f StrictHandlerFunc, operationID string) StrictHandlerFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// PostLivekitGenerateTokenForHostRoom operation middleware
func (sh *strictHandler) PostLivekitGenerateTokenForHostRoom(ctx *fiber.Ctx) error {
	var request PostLivekitGenerateTokenForHostRoomRequestObject

	var body PostLivekitGenerateTokenForHostRoomJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.PostLivekitGenerateTokenForHostRoom(ctx.UserContext(), request.(PostLivekitGenerateTokenForHostRoomRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostLivekitGenerateTokenForHostRoom")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(PostLivekitGenerateTokenForHostRoomResponseObject); ok {
		if err := validResponse.VisitPostLivekitGenerateTokenForHostRoomResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GenerateTokenForJoinRoom operation middleware
func (sh *strictHandler) GenerateTokenForJoinRoom(ctx *fiber.Ctx) error {
	var request GenerateTokenForJoinRoomRequestObject

	var body GenerateTokenForJoinRoomJSONRequestBody
	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	request.Body = &body

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GenerateTokenForJoinRoom(ctx.UserContext(), request.(GenerateTokenForJoinRoomRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GenerateTokenForJoinRoom")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GenerateTokenForJoinRoomResponseObject); ok {
		if err := validResponse.VisitGenerateTokenForJoinRoomResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetLivekitGetUsersInRoom operation middleware
func (sh *strictHandler) GetLivekitGetUsersInRoom(ctx *fiber.Ctx, params GetLivekitGetUsersInRoomParams) error {
	var request GetLivekitGetUsersInRoomRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetLivekitGetUsersInRoom(ctx.UserContext(), request.(GetLivekitGetUsersInRoomRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetLivekitGetUsersInRoom")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetLivekitGetUsersInRoomResponseObject); ok {
		if err := validResponse.VisitGetLivekitGetUsersInRoomResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// GetLivekitRoomCheck operation middleware
func (sh *strictHandler) GetLivekitRoomCheck(ctx *fiber.Ctx, params GetLivekitRoomCheckParams) error {
	var request GetLivekitRoomCheckRequestObject

	request.Params = params

	handler := func(ctx *fiber.Ctx, request interface{}) (interface{}, error) {
		return sh.ssi.GetLivekitRoomCheck(ctx.UserContext(), request.(GetLivekitRoomCheckRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetLivekitRoomCheck")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	} else if validResponse, ok := response.(GetLivekitRoomCheckResponseObject); ok {
		if err := validResponse.VisitGetLivekitRoomCheckResponse(ctx); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/9RY4W/bthP9Vwh+agFFjttff0D9LQu61EHSBmk6DAiCgpHO0sUUyZKUGy/I/z4cRduS",
	"Ii+pk23Yp9gSybv3+O7dxXcc1UzzyR3PwWUWjUet+IR/Dh+EZFUtPUpUwLRlDlUhYS98ba1nqNjloa4q",
	"rU6FnV+9Kr03k9EoC48qYeeptsWoBGlGr+mcjxenJylPuEcvgU/4FyyUkKgKdnA25QlfgHVNHvvpOH3P",
	"7xOuDShhkE/423Sc7vOEG+FLR5mPJC5gjn5UgAIrPFzoOahftf2onT/XuqJFRjv/EOZR3OGYUEzUvgTl",
	"MRMBlKdT2ExbJljtwDKvWamdZ4I5AxnOMGNW64qAaEPHoFbTnE/4mXb+pMnpaFtKCbfwvQbnf9H5khLL",
	"tPKgQo7CGBmzGN04SvSOu6yESgQklqJ5hADeRnxdXBclMCUqYHrGfAnrPP3SEN/OW1QF0UrAaOHjRwQK",
	"CD9d01+ced8gQws5n1w2+bXiXK036OsbyDy/7+7wtobwwBmtXIPxzf7+MxgK9ziMbyWYfPDuh8ENZN87",
	"OOhmc7Srswycm9VSLlPi/H8NnO6uqVoIiTlDZWqfsAodFRszwooKPFgXtr4b2voF7AIsA2u1ZT9KlGtg",
	"q6uKcChXV1eVsMuW9pnoSD1IHBUT8YJp09YKO9aoXrbCbnSI/UiFHW3LY/eyEnmOjeedteQzE9JB8qya",
	"g1tRmWBzmVYzsKAy2KOXe+MXqEjiq1+Rm4g3ulTfcg3/gTItB+XzNSa0whzUuZVgevuNMAwRa4T1mKER",
	"yruHgU6wOZh2O9L/YITLwNA4MvUm/n1LdKGHqvGbfuD4QFgrluH7cwxpAxaWx+X1UYaf8Xj69Y/p+BNO",
	"3VSdv8sOp/+fzs3vvx0ev0/Tv9nEdrzs4FUPKfgQLKwC50SxmjCuV+rGtkN2qeiY54TBrbeCbSIypT0T",
	"UuofkO9Gx6PunNBY04/btux/kCjX6gd9njyVk+wsGWgZQaA7EfWyvajXFNYThygc1eJBp0b4Va9VeTIP",
	"Kolo1wUMGMw5+NoqxwSTLQcgy3Esq60F5eVy5QexK0G+tS1t5r5O9OA+USp8cvmU1kG4LXiLsIDoSjPb",
	"BEXa870Gu+QJb3rFysC7Dp20dNW/zKsXde/tLTFgoSS3zp8DbnzwlNtY3cFTnfdJ+m353Jr+PIRrZ+S2",
	"D3Kn0Ru0XTvWjkNcDE+HrWL2isb3ZOu2Dm706LCEbL61EMJbx3C2qjt0zFhwoOI4WOAC1KPCP18HeqD5",
	"bbL9RB9/RrrJ8Fnr4eXfLYMPt+h8W4/XWksQqj3cHSwESnEtYWjZU3RKNLOMeGYWXC39bnrsKSpcXFcC",
	"HUHR4qDTIRNb/2QQm0vrbcIgLVJ2KlCxV8bqvM7o8eu4lMYoK2l+a343EAbT2LTSTFejxZjTnf9suE23",
	"86IgCuJK6i0ewj+y/chx5V4vA35/df9nAAAA//9hyBEbKhEAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
