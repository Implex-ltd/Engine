package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	gologinendpoint = "http://127.0.0.1:36912"
	token           = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI2NGZjNWFlYTZhMWI0ZjQ2YmJhMzEwMTgiLCJ0eXBlIjoiZGV2Iiwiand0aWQiOiI2NGZjNWFmMDA1MDJkYjA3MTk0MDE3MDYifQ.d6DWwRrNLzYym8aJGQuoZBAHNxPH0CRSahKbRENqebU"
)

func NewGologin() *Gologin {
	return &Gologin{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (G *Gologin) ApplyConfig(Config GologinCreateBrowserPayload) {
	G.Config = Config
}

func (G *Gologin) GetFingerprint() (*GologinGetFingerprintResponse, error) {
	req, err := http.NewRequest("GET", "https://api.gologin.com/browser/fingerprint?os=mac&isM1=true", nil)
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
		return nil, err
	}

	return &data, nil
}

func (G *Gologin) Create() (string, error) {
	Payload, err := json.Marshal(G.Config)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.gologin.com/browser", bytes.NewReader(Payload))
	if err != nil {
		return "", err
	}

	req.Header.Add("authorization", fmt.Sprintf(`Bearer %s`, token))
	req.Header.Add("content-type", "application/json")

	resp, err := G.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data GologinCreateBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data.ID, nil
}

func (G *Gologin) Start(uuid string) (string, error) {
	fmt.Println(uuid)
	resp, err := G.Client.Post(fmt.Sprintf("%s/browser/start-profile", gologinendpoint), "application/json", strings.NewReader(fmt.Sprintf(`{"profileId": "%s","sync": true}`, uuid)))
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

	time.Sleep(5 * time.Second)

	return data.WsUrl, nil
}

func (G *Gologin) Close(uuid string) error {
	fmt.Println(uuid)
	resp, err := G.Client.Post(fmt.Sprintf("%s/browser/stop-profile", gologinendpoint), "application/json", strings.NewReader(fmt.Sprintf(`{"profileId": "%s"}`, uuid)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
