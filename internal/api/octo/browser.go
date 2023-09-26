package octo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Implex-ltd/engine/internal/ratelimiter"
	"github.com/Implex-ltd/engine/internal/utils"
	"io"
	"net/http"
	"time"
)

var (
	endpoint = "http://localhost:58888"
	token    = "83476531c727476f9cf41d04f56d31ad"

	/*
		Plan	    RPM	RPH  | Real
		base	    50	500  | 8/m
		team	    100	1500
		advanced	200	3000
	*/
	Limiter = ratelimiter.NewLimiter(8, time.Minute)
)

func NewOcto(UserAgent, Os string) (*Octo, error) {
	Self := &Octo{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
		Limiter: Limiter,
	}

	Self.Config = OctoCreateBrowserPayload{
		Title: utils.GenerateRandomString(5),
		Fingerprint: Fingerprint{
			OS:        Os,
			UserAgent: UserAgent,
			//Languages: Geolocation{},
			Timezone: Timezone{
				Type: "manual",
				Data: "Europe/Paris",
			},
			//Geolocation: Geolocation{},
			Noise: Noise{
				Webgl:       true,
				Canvas:      true,
				Audio:       true,
				ClientRects: false,
			},
		},
	}

	return Self, nil
}

func (O *Octo) Create() error {
	Payload, err := json.Marshal(O.Config)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://app.octobrowser.net/api/v2/automation/profiles", bytes.NewReader(Payload))
	if err != nil {
		return err
	}

	req.Header.Set("X-Octo-Api-Token", token)
	req.Header.Set("content-type", "application/json")

	resp, err := O.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data OctoCreateBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	if !data.Success {
		return fmt.Errorf("%s", string(body))
	}

	O.UUID = data.Data.UUID
	return nil
}

func (O *Octo) Start() (string, error) {
	Payload, err := json.Marshal(&OctoStartBrowsePayload{
		UUID:      O.UUID,
		Flags:     []string{},
		DebugPort: true,
	})

	if err != nil {
		return "", err
	}

	resp, err := O.Client.Post(fmt.Sprintf("%s/api/profiles/start", endpoint), "application/json", bytes.NewReader(Payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	var data OctoStartBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data.WsEndpoint, nil
}

func (O *Octo) Close() error {
	Payload, err := json.Marshal(&OctoStopBrowsePayload{
		UUID: O.UUID,
	})
	if err != nil {
		return err
	}

	_, err = O.Client.Post(fmt.Sprintf("%s/api/profiles/stop", endpoint), "application/json", bytes.NewReader(Payload))
	if err := O.Delete(); err != nil {
		return err
	}

	return nil
}

func (O *Octo) Delete() error {
	client := http.Client{}

	payload, err := json.Marshal(&OctoDeleteBrowsePayload{
		Uuids: []string{
			O.UUID,
		},
		SkipTrashBin: true,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", "https://app.octobrowser.net/api/v2/automation/profiles", bytes.NewReader(payload))
	if err != nil {
		return err
	}

	req.Header.Set("X-Octo-Api-Token", token)
	req.Header.Set("content-type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
