package livekit

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	livekit "github.com/livekit/protocol/livekit"
	lksdk "github.com/livekit/server-sdk-go/v2"
)


var (
	hostUrl     string
	apiKey      string
	apiSecret   string
	roomService *lksdk.RoomServiceClient
)

func init() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not load .env file, will use environment variables.")
	}

	// Assign values from environment variables
	hostUrl = os.Getenv("LIVEKIT_WS_URL")
	apiKey = os.Getenv("LIVEKIT_API_KEY")
	apiSecret = os.Getenv("LIVEKIT_API_SECRET")

	if hostUrl == "" || apiKey == "" || apiSecret == "" {
		log.Fatalln("Missing LiveKit credentials in environment variables")
	}

	// Initialize RoomServiceClient
	roomService = lksdk.NewRoomServiceClient(hostUrl, apiKey, apiSecret)
}

func UsernameTaken(name, room string) bool {
	roomExist := RoomExist(room)

	if !roomExist {
		return false
	}

    res, _ := roomService.ListParticipants(context.Background(), &livekit.ListParticipantsRequest{
		Room: room,
	})

	for _, p := range res.Participants {
		if p.Name == name {
			return true
		}
	}

	return false
}

func RoomExist(room string) bool {
	res, err := roomService.ListRooms(context.Background(), &livekit.ListRoomsRequest{})
	if err != nil {
		panic("QUE PASA")
	}

	rooms := res.GetRooms()

	for _, r := range rooms {
		if r.Name == room {
			return true
		}
	}

	return false
}


