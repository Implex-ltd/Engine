package hidenium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Implex-ltd/engine/internal/utils"
	"io"
	"net/http"
	"time"
)

var endpoint = "http://127.0.0.1:5555"

func NewHidenium(UserAgent, Os string) (*Hidenium, error) {
	Self := &Hidenium{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}

	w, h := utils.GenerateRandomBrowserSize(500, 3500)

	Self.Config = HideniumCreateBrowserPayload{
		Os:                Os,
		Version:           "116.0.5845.98",
		UserAgent:         UserAgent,
		Canvas:            "noise",
		WebGLImage:        "true",
		WebGLMetadata:     "true",
		AudioContext:      "true",
		ClientRectsEnable: "false",
		NoiseFont:         "true",
		Languages:         "fr-fr;q=0.9",
		Resolution:        fmt.Sprintf("%dx%d", w, h),
	}

	return Self, nil
}

func (H *Hidenium) Create() error {
	Payload, err := json.Marshal(H.Config)
	if err != nil {
		return err
	}

	resp, err := H.Client.Post(fmt.Sprintf("%s/create-profile-custom", endpoint), "application/json", bytes.NewReader(Payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data HideniumCreateBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	H.UUID = data.Content.UUID
	return nil
}

func (H *Hidenium) Start() (string, error) {
	resp, err := H.Client.Get(fmt.Sprintf("%s/openProfile?uuid=%s", endpoint, H.UUID))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var data HideniumStartBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return fmt.Sprintf("http://127.0.0.1:%d", data.Data.RemotePort), nil
}

func (H *Hidenium) Close() error {
	H.Client.Get(fmt.Sprintf("%s/closeProfile?uuid=%s", endpoint, H.UUID))
	if err := H.Delete(); err != nil {
		return err
	}

	return nil
}

func (H *Hidenium) Delete() error {
	client := http.Client{}

	payload, err := json.Marshal(&DeletePayload{
		UUIDBrowser: []string{
			H.UUID,
		},
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/delete-profile", endpoint), bytes.NewReader(payload))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
