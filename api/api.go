package api

import (
	"context"
	livekit "signaling-server/livekit"
	"github.com/gofiber/fiber/v2"
)

type API struct{}

// GenerateTokenForJoinRoom implements StrictServerInterface.
func (a *API) GenerateTokenForJoinRoom(ctx context.Context, request GenerateTokenForJoinRoomRequestObject) (GenerateTokenForJoinRoomResponseObject, error) {
	room := request.Body.Room
	username := request.Body.Username

	token, err := livekit.CreateToken(room, username, false, true)
	if err != nil {
		panic("error creating token")
	}

	p := &[]string{}
	h := ""

	return GenerateTokenForJoinRoom200JSONResponse{
		Token: &token,
		Host : &h,
		Participants: p,
	}, nil
}

// GetLivekitGetUsersInRoom implements StrictServerInterface.
func (a *API) GetLivekitGetUsersInRoom(ctx context.Context, request GetLivekitGetUsersInRoomRequestObject) (GetLivekitGetUsersInRoomResponseObject, error) {
	panic("unimplemented")
}

// GetLivekitRoomCheck implements StrictServerInterface.
func (a *API) GetLivekitRoomCheck(ctx context.Context, request GetLivekitRoomCheckRequestObject) (GetLivekitRoomCheckResponseObject, error) {

	roomname := request.Params.RoomName
	username := request.Params.Username

	roomExist := livekit.RoomExist(roomname)
	usernameAvailable := !livekit.UsernameTaken(username, roomname)

	return GetLivekitRoomCheck200JSONResponse{
		RoomExists: &roomExist,
		UsernameAvailable: &usernameAvailable,
	}, nil

}

// PostLivekitGenerateTokenForHostRoom implements StrictServerInterface.
func (a *API) PostLivekitGenerateTokenForHostRoom(ctx context.Context, request PostLivekitGenerateTokenForHostRoomRequestObject) (PostLivekitGenerateTokenForHostRoomResponseObject, error) {
	panic("unimplemented")
}

func NewApp() *fiber.App {
	api := &API{}
	app := fiber.New()

	server := NewStrictHandler(api, nil)

	RegisterHandlers(app, server)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from signaling-server")
	})

	return app
}
