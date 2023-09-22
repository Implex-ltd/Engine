package dolphin

import (
	"net/http"
)

type Dolphin struct {
	Config DolphinCreateBrowserPayload
	Client http.Client
	UUID   string
}

type DeletePayload struct {
	UUIDBrowser []string `json:"uuid_browser"`
}

type DolphinCreateBrowserPayload struct {
	Name string `json:"name"`
	//Tags                    []interface{} `json:"tags"`
	BrowserType string `json:"browserType"`
	//MainWebsite             string        `json:"mainWebsite"`
	Useragent  Locale    `json:"useragent"`
	Webrtc     Webrtc    `json:"webrtc"`
	Canvas     Audio     `json:"canvas"`
	Webgl      Audio     `json:"webgl"`
	WebglInfo  WebglInfo `json:"webglInfo"`
	ClientRect Audio     `json:"clientRect"`
	//Notes                   Notes         `json:"notes"`
	Timezone Locale `json:"timezone"`
	Locale   Locale `json:"locale"`
	//Proxy                   interface{}   `json:"proxy"`
	//StatusID                int64         `json:"statusId"`
	Geolocation Geolocation `json:"geolocation"`
	CPU         CPU         `json:"cpu"`
	Memory      CPU         `json:"memory"`
	//Screen                  Screen        `json:"screen"`
	Audio        Audio        `json:"audio"`
	MediaDevices MediaDevices `json:"mediaDevices"`
	Ports        Ports        `json:"ports"`
	//DoNotTrack              int64         `json:"doNotTrack"`
	//Args                    []interface{} `json:"args"`
	//PlatformVersion         string        `json:"platformVersion"`
	//UaFullVersion           string        `json:"uaFullVersion"`
	//Login                   string        `json:"login"`
	//Password                string        `json:"password"`
	AppCodeName string `json:"appCodeName"`
	//PlatformName            string        `json:"platformName"`
	Platform string `json:"platform"`
	//ConnectionDownlink      int64         `json:"connectionDownlink"`
	//ConnectionEffectiveType string        `json:"connectionEffectiveType"`
	//ConnectionRtt           int64         `json:"connectionRtt"`
	//ConnectionSaveData      int64         `json:"connectionSaveData"`
	//CPUArchitecture         string        `json:"cpuArchitecture"`
	//OSVersion               string        `json:"osVersion"`
	//VendorSub               string        `json:"vendorSub"`
	ProductSub string `json:"productSub"`
	Vendor     string `json:"vendor"`
	//Product                 string        `json:"product"`
}

type Audio struct {
	Mode string `json:"mode"`
}

type CPU struct {
	Mode  string `json:"mode"`
	Value int64  `json:"value"`
}

type Geolocation struct {
	Mode string `json:"mode"`
	//Latitude  interface{} `json:"latitude"`
	//Longitude interface{} `json:"longitude"`
	//Accuracy  interface{} `json:"accuracy"`
}

type Locale struct {
	Mode  string `json:"mode"`
	Value string `json:"value"`
}

type MediaDevices struct {
	Mode         string `json:"mode"`
	AudioInputs  int64  `json:"audioInputs"`
	VideoInputs  int64  `json:"videoInputs"`
	AudioOutputs int64  `json:"audioOutputs"`
}

type Notes struct {
	Content interface{} `json:"content"`
	Color   string      `json:"color"`
	Style   string      `json:"style"`
	Icon    interface{} `json:"icon"`
}

type Ports struct {
	Mode      string `json:"mode"`
	Blacklist string `json:"blacklist"`
}

type Screen struct {
	Mode       string `json:"mode"`
	Resolution string `json:"resolution"`
}

type WebglInfo struct {
	Mode          string `json:"mode"`
	Vendor        string `json:"vendor"`
	Renderer      string `json:"renderer"`
	Webgl2Maximum string `json:"webgl2Maximum"`
	Webgpu        string `json:"webgpu"`
}

type Webrtc struct {
	Mode      string      `json:"mode"`
	IPAddress interface{} `json:"ipAddress"`
}

type DolphinCreateBrowserResponse struct {
	Success          int `json:"success"`
	BrowserProfileID int `json:"browserProfileId"`
	Data             struct {
		ID int `json:"id"`
	} `json:"data"`
}

type DolphinStartBrowserResponse struct {
	Success    bool `json:"success"`
	Automation struct {
		Port       int    `json:"port"`
		WsEndpoint string `json:"wsEndpoint"`
	} `json:"automation"`
}
