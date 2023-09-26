package octo

import (
	"github.com/Implex-ltd/engine/internal/ratelimiter"
	"net/http"
)

type Octo struct {
	Config  OctoCreateBrowserPayload
	Client  http.Client
	Limiter *ratelimiter.Limiter
	UUID    string
}

type DeletePayload struct {
	UUIDBrowser []string `json:"uuid_browser"`
}

type OctoCreateBrowserPayload struct {
	Title string `json:"title"`
	//Description    string         `json:"description"`
	//StartPages     []string       `json:"start_pages"`
	//Tags           []string       `json:"tags"`
	//Proxy          Proxy          `json:"proxy"`
	//StorageOptions StorageOptions `json:"storage_options"`
	//Cookies        []Cooky        `json:"cookies"`
	//Image          string         `json:"image"`
	//Extensions     []interface{}  `json:"extensions"`
	Fingerprint Fingerprint `json:"fingerprint"`
}

type Cooky struct {
	Domain         string  `json:"domain"`
	ExpirationDate float64 `json:"expirationDate"`
	HostOnly       bool    `json:"hostOnly"`
	HTTPOnly       bool    `json:"httpOnly"`
	Name           string  `json:"name"`
	Path           string  `json:"path"`
	SameSite       string  `json:"sameSite"`
	Secure         bool    `json:"secure"`
	Value          string  `json:"value"`
}

type Fingerprint struct {
	OS        string `json:"os"`
	UserAgent string `json:"user_agent"`
	//Screen       string       `json:"screen"`
	//Languages    Geolocation  `json:"languages"`
	Timezone Timezone `json:"timezone"`
	//Geolocation  Geolocation  `json:"geolocation"`
	//CPU          int64        `json:"cpu"`
	//RAM          int64        `json:"ram"`
	Noise Noise `json:"noise"`
	//Webrtc       Geolocation  `json:"webrtc"`
	//DNS          string       `json:"dns"`
	//MediaDevices MediaDevices `json:"media_devices"`
}

type Timezone struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type Geolocation struct {
	Type string `json:"type"`
}

type MediaDevices struct {
	VideoIn  int64 `json:"video_in"`
	AudioIn  int64 `json:"audio_in"`
	AudioOut int64 `json:"audio_out"`
}

type Noise struct {
	Webgl       bool `json:"webgl"`
	Canvas      bool `json:"canvas"`
	Audio       bool `json:"audio"`
	ClientRects bool `json:"client_rects"`
}

type Proxy struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int64  `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type StorageOptions struct {
	Cookies      bool `json:"cookies"`
	Passwords    bool `json:"passwords"`
	Extensions   bool `json:"extensions"`
	Localstorage bool `json:"localstorage"`
	History      bool `json:"history"`
	Bookmarks    bool `json:"bookmarks"`
}

type OctoCreateBrowserResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    struct {
		UUID string `json:"uuid"`
	} `json:"data"`
}

type OctoStartBrowserResponse struct {
	UUID       string `json:"uuid"`
	Headless   bool   `json:"headless"`
	StartTime  int64  `json:"start_time"`
	WsEndpoint string `json:"ws_endpoint"`
}

type OctoStartBrowsePayload struct {
	UUID      string   `json:"uuid"`
	Headless  bool     `json:"headless"`
	DebugPort bool     `json:"debug_port"`
	Flags     []string `json:"flags"`
}

type OctoStopBrowsePayload struct {
	UUID string `json:"uuid"`
}

type OctoDeleteBrowsePayload struct {
	Uuids        []string `json:"uuids"`
	SkipTrashBin bool     `json:"skip_trash_bin"`
}
