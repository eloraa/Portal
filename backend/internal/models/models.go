package models

import "github.com/gorilla/websocket"

type User struct {
	ID string // Unique user identifier
}

type Room struct {
	ID       string                     // Unique room identifier
	Name     string                     // Room name (e.g., FluffyCookie)
	Creator  string                     // ID of the room creator
	IsPublic bool                       // Visibility (true = public, false = private)
	Password string                     // Password for private rooms
	Members  map[*websocket.Conn]string // Map of WebSocket connections to usernames
}

type Client struct {
	Conn     *websocket.Conn // WebSocket connection
	UserID   string          // Unique user identifier
	Username string          // Username
	AvatarID string          // Avatar identifier
	RoomIDs  []string        // Rooms the user is connected to
}

type Message struct {
	Type    string      `json:"type"`
	RoomID  string      `json:"roomId,omitempty"`
	UserID  string      `json:"userId,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

// Add these new message types
const (
	SignalOffer     = "signal_offer"
	SignalAnswer    = "signal_answer"
	SignalCandidate = "signal_candidate"
)

// Add SignalingPayload struct
type SignalingPayload struct {
	Type       string      `json:"type"`
	SDP        string      `json:"sdp,omitempty"`
	Candidate  interface{} `json:"candidate,omitempty"`
	FromUserID string      `json:"fromUserId"`
	ToUserID   string      `json:"toUserId,omitempty"`
}
