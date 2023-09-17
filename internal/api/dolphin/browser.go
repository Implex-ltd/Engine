package dolphin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Implex-ltd/engine/internal/utils"
)

var (
	token    = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiN2IxNmZiNzcyZGEwYzIwZWE2NGFiNjdhOGNmYWMyOTQ0ZjkzZjViMTdhYzE1Y2U5YmU1NTIxNzg0MDcwNGY2MTQ4ZmFmZDQwYTAxMzk1NjYiLCJpYXQiOjE2OTQ5NTY2MDIuMTUxMjUyLCJuYmYiOjE2OTQ5NTY2MDIuMTUxMjU0LCJleHAiOjE2OTc1NDg2MDIuMTQzMjY4LCJzdWIiOiIyNjc5MzU3Iiwic2NvcGVzIjpbXX0.paiOcaxKzRz9ZPhz5Ky30aTUZxV-AzqhcwmxOiwnktGu8hyE6ySpGhCpclRhByRvD00d7wlF3dHAGhtsS52rW9mKfM91Ecdo7jmy_LxEbTa8CebJ94Bw7JPl3ZQ7AP79rBZNP6Lo1kDgbmxQ0SRFDrvLpKUEs5OsmX_JFrANvnAj70o9T-GreGQybhu3ZJqZZzG7o3du_fUoM2UsRlLkMt_6xZjvH0qbrOcu_fMuxagLXDZpOTuSd6ns7u097Xd1Rx44Pa7nLKqfR7mK1ciayJHUVHTb1OOGm8l2ivXiNfV7FUpt3buLa3oaY8U4dnGFCJlA287U_CBAEtrePoidwYzkS_o-GjGiCsnTUZTCS3cl_4LZYkZ3CXvtkN0wt9OCCoy8rSHXjTMqfwk95zxQPgJbkZotgr9Tea41hJ5QhndyMmpX0YF5LFN_fmpoPSC4mAs6kdVPIRoz_aHSTfOHj9zS5leSZE6_nLT5RWKlgVWYViAUAw7vepdzRgDX6SRHEeh9YtoeIrKOiU_izHPWoa6D1K3H2KKJCErxwUpoNLdO_6xk_Fgz6ezCU90mN6GMmL6ygRpqiNFN2ZQoZpE5rSbBg0_1FoQqP5zfeiZfJGeVfJdhyfy4-cJYz5e0xwYdwZC5ibFNNtRA1MgZ46iaJ-IIn_DP29jrh1F9TjD2pao"
	endpoint = "http://127.0.0.1:3001/v1.0"
)

func NewDolphin(UserAgent, Os string) (*Dolphin, error) {
	Self := &Dolphin{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}

	pname := "Win32"
	if Os == "mac" {
		pname = "MacIntel"
	} else {
		pname = "windows"
	}

	Self.Config = DolphinCreateBrowserPayload{
		Useragent: Locale{
			Mode:  "manual",
			Value: UserAgent,
		},
		//OSVersion:       "10",
		Platform:    pname,
		Name:        utils.GenerateRandomString(5),
		BrowserType: "anty", // Browser type. Available values: ['anty']. really ?????
		Webrtc: Webrtc{
			Mode: "off",
		},
		ProductSub:  "20030107",
		Vendor:      "Google Inc.",
		AppCodeName: "Mozilla",
		Webgl: Audio{
			Mode: "noise",
		},
		Canvas: Audio{
			Mode: "noise",
		},
		Timezone: Locale{
			Mode: "auto",
		},
		Locale: Locale{
			Mode: "auto",
		},
		WebglInfo: WebglInfo{
			Mode: "noise",
		},
		ClientRect: Audio{
			Mode: "off",
		},
		Memory: CPU{
			Mode: "real",
		},
		CPU: CPU{
			Mode: "real",
		},
		Geolocation: Geolocation{
			Mode: "auto",
		},
	}

	return Self, nil
}

func (D *Dolphin) Create() error {
	Payload, err := json.Marshal(D.Config)
	if err != nil {
		return err
	}

	fmt.Println(string(Payload))

	req, err := http.NewRequest("POST", "https://dolphin-anty-api.com/browser_profiles", bytes.NewReader(Payload))
	if err != nil {
		return err
	}

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("content-type", "application/json")

	resp, err := D.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	var data DolphinCreateBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}

	D.UUID = fmt.Sprintf("%d", data.BrowserProfileID)
	return nil
}

func (D *Dolphin) Start() (string, error) {
	fmt.Println(D.UUID)
	resp, err := D.Client.Get(fmt.Sprintf("%s/browser_profiles/%s/start?automation=1", endpoint, D.UUID))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	var data DolphinStartBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return "", err
	}

	// {"success":true,"automation":{"port":64933,"wsEndpoint":"/devtools/browser/10cd3b85-2eea-46b4-85b9-62235b17e039"}}
	return fmt.Sprintf("ws://127.0.0.1:%d%s", data.Automation.Port, data.Automation.WsEndpoint), nil
}

func (D *Dolphin) Close() error {
	D.Client.Get(fmt.Sprintf("%s/browser_profiles/%s/stop", endpoint, D.UUID))
	if err := D.Delete(); err != nil {
		return err
	}

	return nil
}

func (D *Dolphin) Delete() error {
	client := http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf("https://dolphin-anty-api.com/browser_profiles/%s?forceDelete=1", D.UUID), nil)
	if err != nil {
		return err
	}

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
