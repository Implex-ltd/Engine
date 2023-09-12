package gologin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Implex-ltd/engine/internal/ratelimiter"
)

var (
	gologinendpoint = "http://127.0.0.1:36912"
	token           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2NGZlOTYxNGI4NDU5YjYxYzc2ZGM0OWIiLCJ0eXBlIjoiZGV2Iiwiand0aWQiOiI2NGZlOTYxZWI4NDU5YjYyNTU2ZGM1NjUifQ.AqN21xvaANsyGV9oRdY6ME9BL-MvFIyFZ-eOCNyC5Ls" //"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2NGZjNWFlYTZhMWI0ZjQ2YmJhMzEwMTgiLCJ0eXBlIjoiZGV2Iiwiand0aWQiOiI2NGZjNWFmMDA1MDJkYjA3MTk0MDE3MDYifQ.d6DWwRrNLzYym8aJGQuoZBAHNxPH0CRSahKbRENqebU"
)

func NewGologin(UserAgent, Os string) (*Gologin, error) {
	Self := &Gologin{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
		Limiter: ratelimiter.NewLimiter(300, time.Minute),
	}

	fp, err := Self.GetFingerprint(Os)
	if err != nil {
		return nil, err
	}

	Self.Config = GologinCreateBrowserPayload{
		DevicePixelRatio: int(fp.DevicePixelRatio),
		Name:             "default_name",
		Notes:            "auto generated",
		OS:               fp.OS,
		BrowserType:      "chrome",
		Navigator: Navigator{
			UserAgent:           UserAgent,
			Resolution:          fp.Navigator.Resolution,
			Language:            "fr-FR",
			Platform:            fp.Navigator.Platform,
			HardwareConcurrency: strconv.Itoa(fp.Navigator.HardwareConcurrency),
			DeviceMemory:        strconv.Itoa(fp.Navigator.DeviceMemory),
			MaxTouchPoints:      0,
		},
		Timezone: Timezone{
			Enabled:       true,
			FillBasedOnIP: true,
		},
		AudioContext: AudioContext{
			Mode: "noise",
		},
		Canvas: AudioContext{
			Mode: "noise",
		},
		Fonts: Fonts{
			EnableMasking: true,
			EnableDOMRect: true,
			Families:      fp.Fonts,
		},
		MediaDevices: fp.MediaDevices,
		WebGL: WebGL{
			Mode: "noise",
			//GetClientRectsNoise: 0,
		},
		ClientRects: AudioContext{
			Mode: "off",
		},
		WebGLMetadata: WebGLMetadata{
			Mode:     "mask",
			Vendor:   fp.WebGLMetadata.Vendor,
			Renderer: fp.WebGLMetadata.Renderer,
		},
		ProxyEnabled: false,
		Proxy: Proxy{
			Mode: "none",
		},
		WebRTC: WebRTC{
			Mode: "real",
		},
		WebglParams: fp.WebglParams,
		A:           []any{},
		B:           []any{},
		AutoLang:    true,
	}

	return Self, nil
}

func (G *Gologin) GetFingerprint(Os string) (*GologinGetFingerprintResponse, error) {
	G.Limiter.Wait()

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.gologin.com/browser/fingerprint?os=%s", Os), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", fmt.Sprintf(`Bearer %s`, token))
	req.Header.Add("content-type", "application/json")

	resp, err := G.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data GologinGetFingerprintResponse
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
		return nil, err
	}

	return &data, nil
}

func (G *Gologin) Create() error {
	G.Limiter.Wait()

	Payload, err := json.Marshal(G.Config)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.gologin.com/browser", bytes.NewReader(Payload))
	if err != nil {
		return err
	}

	req.Header.Add("authorization", fmt.Sprintf(`Bearer %s`, token))
	req.Header.Add("content-type", "application/json")

	resp, err := G.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data GologinCreateBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
		return err
	}

	G.UUID = data.ID
	return nil
}

func (G *Gologin) Start() (string, error) {
	resp, err := G.Client.Post(fmt.Sprintf("%s/browser/start-profile", gologinendpoint), "application/json", strings.NewReader(fmt.Sprintf(`{"profileId": "%s","sync": true}`, G.UUID)))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data GologinStartBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data.WsUrl, nil
}

func (G *Gologin) Close() error {
	resp, err := G.Client.Post(fmt.Sprintf("%s/browser/stop-profile", gologinendpoint), "application/json", strings.NewReader(fmt.Sprintf(`{"profileId": "%s"}`, G.UUID)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := G.Delete(); err != nil {
		return err
	}

	return nil
}

func (G *Gologin) Delete() error {
	G.Limiter.Wait()

	client := http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://api.gologin.com/browser/%s", G.UUID), nil)
	if err != nil {
		return err
	}

	req.Header.Add("authorization", fmt.Sprintf(`Bearer %s`, token))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
