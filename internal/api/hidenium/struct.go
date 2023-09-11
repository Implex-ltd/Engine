package hidenium

import (
	"net/http"
	"time"
)

type Hidenium struct {
	Config HideniumCreateBrowserPayload
	Client http.Client
	UUID   string
}

type DeletePayload struct {
	UUIDBrowser []string `json:"uuid_browser"`
}

type HideniumCreateBrowserPayload struct {
	Os                string `json:"os"`
	Version           string `json:"version"`
	UserAgent         string `json:"userAgent"`
	Canvas            string `json:"canvas"`
	WebGLImage        string `json:"webGLImage"`
	AudioContext      string `json:"audioContext"`
	WebGLMetadata     string `json:"webGLMetadata"`
	ClientRectsEnable string `json:"clientRectsEnable"`
	NoiseFont         string `json:"noiseFont"`
	Languages         string `json:"languages"`
	Resolution        string `json:"resolution"`
}

type HideniumCreateBrowserResponse struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content struct {
		UUID   string `json:"uuid"`
		Config struct {
			Name        string `json:"name"`
			Version     string `json:"version"`
			Os          string `json:"os"`
			Browser     string `json:"browser"`
			Fingerprint struct {
				Navigator struct {
					UserAgent           string `json:"userAgent"`
					OsVersion           string `json:"osVersion"`
					BrowserVersion      string `json:"browserVersion"`
					Browser             string `json:"browser"`
					HardwareConcurrency int    `json:"hardwareConcurrency"`
					Platform            string `json:"platform"`
					MaxTouchPoints      int    `json:"maxTouchPoints"`
					DeviceMemory        int    `json:"deviceMemory"`
					DoNotTrack          string `json:"doNotTrack"`
					Oscpu               any    `json:"oscpu"`
					BuildID             any    `json:"buildID"`
					BrowserVersionName  string `json:"browserVersionName"`
					Device              string `json:"device"`
					SourceVersion       string `json:"sourceVersion"`
				} `json:"navigator"`
				Screen struct {
					Width      int `json:"width"`
					Height     int `json:"height"`
					PixelRatio int `json:"pixelRatio"`
				} `json:"screen"`
				Webgl2 struct {
					Canvas       string `json:"canvas"`
					AudioContext struct {
						Enable     bool    `json:"enable"`
						NoiseValue float64 `json:"noiseValue"`
					} `json:"audioContext"`
					WebGLImage         bool    `json:"WebGLImage"`
					WebGLMeta          bool    `json:"WebGLMeta"`
					ClientRectsEnable  bool    `json:"clientRectsEnable"`
					ClientRectsNoice   float64 `json:"ClientRectsNoice"`
					WebglNoiseValue    float64 `json:"WebglNoiseValue"`
					UnmaskedVendor     string  `json:"unmaskedVendor"`
					UnmaskedRenderer   string  `json:"unmaskedRenderer"`
					Extensions         []any   `json:"extensions"`
					SupportedFunctions []any   `json:"supportedFunctions"`
					Params             []any   `json:"params"`
					Antialiasing       bool    `json:"antialiasing"`
				} `json:"webgl2"`
				Webgl struct {
					UnmaskedVendor   string `json:"unmaskedVendor"`
					UnmaskedRenderer string `json:"unmaskedRenderer"`
				} `json:"webgl"`
				Languages struct {
					AcceptLanguage string `json:"acceptLanguage"`
				} `json:"languages"`
				Fonts struct {
					IsEnabled bool `json:"isEnabled"`
					Value     []struct {
						Name      string `json:"name"`
						Required  int    `json:"required"`
						IsChecked int    `json:"isChecked"`
					} `json:"value"`
				} `json:"fonts"`
				MediaDevices struct {
					EnableMasking bool   `json:"enableMasking"`
					Enable        bool   `json:"enable"`
					VideoInputs   int    `json:"videoInputs"`
					AudioInputs   int    `json:"audioInputs"`
					AudioOutputs  int    `json:"audioOutputs"`
					UID           string `json:"uid"`
				} `json:"mediaDevices"`
				WebRtcLocalIps string `json:"webRtcLocalIps"`
				Ports          struct {
					IsEnabled      bool   `json:"isEnabled"`
					PortsToProtect string `json:"portsToProtect"`
				} `json:"ports"`
				DisableImage bool   `json:"disableImage"`
				BrowserType  string `json:"browser_type"`
				Timezone     struct {
					IsProxy      bool   `json:"isProxy"`
					NameTimezone string `json:"nameTimezone"`
				} `json:"timezone"`
				Webrtc struct {
					FillBasedOnIP  bool   `json:"fill_based_on_ip"`
					LocalIps       string `json:"localIps"`
					LocalIPMasking bool   `json:"local_ip_masking"`
					Mode           string `json:"mode"`
					PublicIP       bool   `json:"publicIP"`
					PublicIP0      bool   `json:"public_ip"`
				} `json:"webrtc"`
				GeoLocation struct {
					IsProxy   bool   `json:"isProxy"`
					Mode      string `json:"mode"`
					Accuracy  int    `json:"accuracy"`
					Latitude  string `json:"latitude"`
					Longitude string `json:"longitude"`
				} `json:"geoLocation"`
				Proxy struct {
					IsProxy          bool   `json:"IsProxy"`
					ProxyHost        string `json:"ProxyHost"`
					ProxyPort        string `json:"ProxyPort"`
					ProxyType        int    `json:"ProxyType"`
					ProxyUsername    string `json:"ProxyUsername"`
					ProxyPassword    string `json:"ProxyPassword"`
					IsDNS            bool   `json:"IsDNS"`
					IPDNS            any    `json:"IpDNS"`
					Geolocation      string `json:"Geolocation"`
					Webrtc           string `json:"Webrtc"`
					TimeZone         string `json:"TimeZone"`
					CheckBeforeStart bool   `json:"checkBeforeStart"`
				} `json:"proxy"`
				Storage         any     `json:"storage"`
				StartURL        string  `json:"StartURL"`
				UID             any     `json:"Uid"`
				CanvasCntPixels int     `json:"canvasCntPixels"`
				CanvasModel     any     `json:"CanvasModel"`
				CanvasNoise     float64 `json:"canvasNoise"`
				Plugins         struct {
					AllEnable   bool `json:"all_enable"`
					FlashEnable bool `json:"flash_enable"`
				} `json:"plugins"`
				Exxtension any `json:"exxtension"`
				MacAddress struct {
					IsEnabled bool   `json:"isEnabled"`
					Value     string `json:"value"`
				} `json:"macAddress"`
				ComputerName struct {
					IsEnabled bool   `json:"isEnabled"`
					Value     string `json:"value"`
				} `json:"computerName"`
			} `json:"fingerprint"`
			BrowserType string   `json:"browser_type"`
			Fonts       []string `json:"fonts"`
		} `json:"config"`
		Directory     string    `json:"directory"`
		FileName      string    `json:"file_name"`
		CanBeRunning  bool      `json:"can_be_running"`
		Proxy         string    `json:"proxy"`
		CreatedAt     time.Time `json:"created_at"`
		CreatedAtDiff string    `json:"created_at_diff"`
		ID            int       `json:"id"`
		FolderID      string    `json:"folder_id"`
		CheckOpen     any       `json:"check_open"`
		Name          string    `json:"name"`
		Version       int       `json:"version"`
		BrowserType   string    `json:"browser_type"`
		Os            string    `json:"os"`
		Note          string    `json:"note"`
		LastOpen      any       `json:"last_open"`
		LastOpenDiff  any       `json:"last_open_diff"`
		Transfer      string    `json:"transfer"`
		SourceVersion string    `json:"source_version"`
		SharesRole    bool      `json:"shares_role"`
		SharedUUID    string    `json:"shared_uuid"`
		Status        any       `json:"status"`
		StatusChil    int       `json:"status_chil"`
		Tags          []any     `json:"tags"`
		StatusID      any       `json:"status_id"`
		Option        []any     `json:"option"`
		Disk          string    `json:"disk"`
	} `json:"content"`
}

type HideniumStartBrowserResponse struct {
	Status string `json:"status"`
	Data   struct {
		RemotePort  int    `json:"remote_port"`
		ProfilePath string `json:"profile_path"`
		ExecutePath string `json:"execute_path"`
	} `json:"data"`
}
