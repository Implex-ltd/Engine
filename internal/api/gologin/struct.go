package gologin

import (
	"net/http"
	"time"
)

type GologinCreateBrowserPayload struct {
	DevicePixelRatio int    `json:"devicePixelRatio"`
	Name             string `json:"name"`
	Notes            string `json:"notes"`
	BrowserType      string `json:"browserType"`
	OS               string `json:"os"`
	//StartURL              string        `json:"startUrl"`
	//GoogleServicesEnabled bool          `json:"googleServicesEnabled"`
	//LockEnabled           bool          `json:"lockEnabled"`
	//DebugMode             bool          `json:"debugMode"`
	Navigator Navigator `json:"navigator"`
	//GeoProxyInfo          GeoProxyInfo  `json:"geoProxyInfo"`
	//Storage               Storage       `json:"storage"`
	ProxyEnabled bool  `json:"proxyEnabled"`
	Proxy        Proxy `json:"proxy"`
	//DNS                   string        `json:"dns"`
	//Plugins               Plugins       `json:"plugins"`
	Timezone      Timezone      `json:"timezone"`
	AudioContext  AudioContext  `json:"audioContext"`
	Canvas        AudioContext  `json:"canvas"`
	Fonts         Fonts         `json:"fonts"`
	MediaDevices  MediaDevices  `json:"mediaDevices"`
	WebRTC        WebRTC        `json:"webRTC"`
	WebGL         WebGL         `json:"webGL"`
	ClientRects   AudioContext  `json:"clientRects"`
	WebGLMetadata WebGLMetadata `json:"webGLMetadata"`
	WebglParams   any           `json:"webglParams"`
	//Profile               string        `json:"profile"`
	//GoogleClientID        string        `json:"googleClientId"`
	//UpdateExtensions      bool          `json:"updateExtensions"`
	//ChromeExtensions      []string      `json:"chromeExtensions"`
	A        []any `json:"userExtensionsToNewProfiles"`
	B        []any `json:"extensionsToNewProfiles"`
	AutoLang bool  `json:"autoLang"`
}

type AudioContext struct {
	Mode string `json:"mode"`
	//Noise int64  `json:"noise"`
}

type Fonts struct {
	Families      []string `json:"families"`
	EnableMasking bool     `json:"enableMasking"`
	EnableDOMRect bool     `json:"enableDomRect"`
}

type GeoProxyInfo struct {
}

type MediaDevices struct {
	VideoInputs  int64 `json:"videoInputs"`
	AudioInputs  int64 `json:"audioInputs"`
	AudioOutputs int64 `json:"audioOutputs"`
	//EnableMasking bool  `json:"enableMasking"`
}

type Navigator struct {
	UserAgent  string `json:"userAgent"`
	Resolution string `json:"resolution"`
	Language   string `json:"language"`
	Platform   string `json:"platform"`
	//DoNotTrack          bool   `json:"doNotTrack"`
	HardwareConcurrency string `json:"hardwareConcurrency"`
	DeviceMemory        string `json:"deviceMemory"`
	MaxTouchPoints      int64  `json:"maxTouchPoints"`
}

type xNavigator struct {
	UserAgent           string `json:"userAgent"`
	Resolution          string `json:"resolution"`
	Language            string `json:"language"`
	Platform            string `json:"platform"`
	DoNotTrack          bool   `json:"doNotTrack"`
	HardwareConcurrency int    `json:"hardwareConcurrency"`
	DeviceMemory        int    `json:"deviceMemory"`
	MaxTouchPoints      int64  `json:"maxTouchPoints"`
}

type Plugins struct {
	EnableVulnerable bool `json:"enableVulnerable"`
	EnableFlash      bool `json:"enableFlash"`
}

type Proxy struct {
	Mode string `json:"mode"`
	//Host     string `json:"host"`
	//Port     int64  `json:"port"`
	//Username string `json:"username"`
	//Password string `json:"password"`
}

type Storage struct {
	Local      bool `json:"local"`
	Extensions bool `json:"extensions"`
	Bookmarks  bool `json:"bookmarks"`
	History    bool `json:"history"`
	Passwords  bool `json:"passwords"`
	Session    bool `json:"session"`
}

type Timezone struct {
	Enabled       bool   `json:"enabled"`
	FillBasedOnIP bool   `json:"fillBasedOnIp"`
	Timezone      string `json:"timezone"`
}

type WebGL struct {
	Mode string `json:"mode"`
	//GetClientRectsNoise int64  `json:"getClientRectsNoise"`
	//Noise               int64  `json:"noise"`
}

type WebGLMetadata struct {
	Mode     string `json:"mode"`
	Vendor   string `json:"vendor"`
	Renderer string `json:"renderer"`
}

type WebRTC struct {
	Mode string `json:"mode"`
	//Enabled        bool     `json:"enabled"`
	//Customize      bool     `json:"customize"`
	//LocalIPMasking bool     `json:"localIpMasking"`
	//FillBasedOnIP  bool     `json:"fillBasedOnIp"`
	//PublicIP       string   `json:"publicIp"`
	//LocalIPS       []string `json:"localIps"`
}

// gologin

type Gologin struct {
	Config GologinCreateBrowserPayload
	Client http.Client
	UUID   string
}

type GologinCreateBrowserResponse struct {
	Name        string `json:"name"`
	Role        string `json:"role"`
	ID          string `json:"id"`
	Notes       string `json:"notes"`
	BrowserType string `json:"browserType"`
	LockEnabled bool   `json:"lockEnabled"`
	Timezone    struct {
		Enabled       bool   `json:"enabled"`
		FillBasedOnIP bool   `json:"fillBasedOnIp"`
		Timezone      string `json:"timezone"`
	} `json:"timezone"`
	Navigator struct {
		HardwareConcurrency int    `json:"hardwareConcurrency"`
		DoNotTrack          bool   `json:"doNotTrack"`
		DeviceMemory        int    `json:"deviceMemory"`
		MaxTouchPoints      int    `json:"maxTouchPoints"`
		UserAgent           string `json:"userAgent"`
		Resolution          string `json:"resolution"`
		Language            string `json:"language"`
		Platform            string `json:"platform"`
	} `json:"navigator"`
	Geolocation struct {
		Mode          string `json:"mode"`
		Enabled       bool   `json:"enabled"`
		Customize     bool   `json:"customize"`
		FillBasedOnIP bool   `json:"fillBasedOnIp"`
		Latitude      int    `json:"latitude"`
		Longitude     int    `json:"longitude"`
		Accuracy      int    `json:"accuracy"`
	} `json:"geolocation"`
	DebugMode bool   `json:"debugMode"`
	Os        string `json:"os"`
	Proxy     struct {
		Mode            string `json:"mode"`
		Port            int    `json:"port"`
		AutoProxyRegion string `json:"autoProxyRegion"`
		TorProxyRegion  string `json:"torProxyRegion"`
		Host            string `json:"host"`
		Username        string `json:"username"`
		Password        string `json:"password"`
		ID              string `json:"id"`
	} `json:"proxy"`
	Folders              []any     `json:"folders"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
	ChromeExtensions     []string  `json:"chromeExtensions"`
	UserChromeExtensions []any     `json:"userChromeExtensions"`
	Tags                 []any     `json:"tags"`
	ProxyEnabled         bool      `json:"proxyEnabled"`
}

type GologinStartBrowserResponse struct {
	Status string `json:"status"`
	WsUrl  string `json:"wsUrl"`
}

// get fp
type GologinGetFingerprintResponse struct {
	Navigator               xNavigator    `json:"navigator"`
	Canvas                  Canvas        `json:"canvas"`
	MediaDevices            MediaDevices  `json:"mediaDevices"`
	WebRTC                  WebRTC        `json:"webRTC"`
	WebGLMetadata           WebGLMetadata `json:"webGLMetadata"`
	WebglParams             WebglParams   `json:"webglParams"`
	WebGL                   Canvas        `json:"webGL"`
	ClientRects             Canvas        `json:"clientRects"`
	OS                      string        `json:"os"`
	DevicePixelRatio        float64       `json:"devicePixelRatio"`
	Fonts                   []string      `json:"fonts"`
	ExtensionsToNewProfiles []interface{} `json:"extensionsToNewProfiles"`
}

type Canvas struct {
	Mode string `json:"mode"`
}

type WebglParams struct {
	GlCanvas                string              `json:"glCanvas"`
	SupportedFunctions      []SupportedFunction `json:"supportedFunctions"`
	GlParamValues           []GlParamValue      `json:"glParamValues"`
	Antialiasing            bool                `json:"antialiasing"`
	TextureMaxAnisotropyEXT int64               `json:"textureMaxAnisotropyExt"`
	ShaiderPrecisionFormat  string              `json:"shaiderPrecisionFormat"`
	Extensions              []string            `json:"extensions"`
}

type GlParamValue struct {
	Name  any `json:"name"`
	Value any `json:"value"`
}

type SupportedFunction struct {
	Name      string `json:"name"`
	Supported bool   `json:"supported"`
}

type Name struct {
	String      *string
	StringArray []string
}

type Value struct {
	Integer    *int64
	IntegerMap map[string]int64
	String     *string
}
