package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var endpoint = "http://127.0.0.1:5555"

func NewHidenium(Config HideniumCreateBrowserPayload) *Hidenium {
	return &Hidenium{
		Config: Config,
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (H *Hidenium) Create() (string, error) {
	Payload, err := json.Marshal(H.Config)
	if err != nil {
		return "", err
	}

	resp, err := H.Client.Post(fmt.Sprintf("%s/create-profile-custom", endpoint), "application/json", bytes.NewReader(Payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	var data HideniumCreateBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	return data.Content.UUID, nil
}

func (H *Hidenium) Start(uuid string) (string, error) {
	fmt.Println(uuid)
	resp, err := H.Client.Get(fmt.Sprintf("%s/openProfile?uuid=%s", endpoint, uuid))
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

	time.Sleep(1 * time.Second)

	return fmt.Sprintf("http://127.0.0.1:%d", data.Data.RemotePort), nil
}

func (H *Hidenium) Close(uuid string) error {
	H.Client.Get(fmt.Sprintf("%s/closeProfile?uuid=%s", endpoint, uuid))
	return nil
}
