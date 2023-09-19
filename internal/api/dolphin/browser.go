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
	token    = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiYmZkZTQ1MWJmNzEzYjQ5MzVmYTI5MjE5OGRiNWZiYjZmMTJkNjU2M2FjNmNlZmM5NTczY2RkMmMwNzEwNzBmYTJlZWQ4MDM2ODA2NDVjMTgiLCJpYXQiOjE2OTUxMTU5OTQuMTAxNzgyLCJuYmYiOjE2OTUxMTU5OTQuMTAxNzg0LCJleHAiOjE2OTc3MDc5OTQuMDkzMzA0LCJzdWIiOiIyNjc5MzU3Iiwic2NvcGVzIjpbXX0.snIlOIe2ivpXDipmFRItomqNMy_n-_4jzTtETKwya5jT5NR_Zy2Fxdl3iXOfH1by2hS-6QFu-GHW38xJkITIujaojToCtec_QS9YSevc-MoJvZHRTpDSqzmJQz3nIbDrvMSFdYWqN_auRhFFSTdi2S9Go3kz-FCa6Bhy9xC76b8fuNm22xVH-b7rWXGjJhYiw6b38dZw4yJdP2LIK_SMYoiKRwkCcYlANRCnqAw1TGPEdDB9CsUW8nrzfmp2-Npv_YQNGkoNdR47iJZZm6eaD0zRdpVqzs0nfMqOf4mRRpYiaEIkP4b9SoZofk1fULWYXYo2F8zYlm4Chw1yDUHX0RLxxMHY93S9rStGkpLC4TyRXq4A_2mVvyZ8oQmeOIppZA_WrBqAH3uo314GYWjxEHm7aiIEb96Ylt5HCQ3TcQRf1Z0xUirTbAvGF6eTQu6pOJB0gfe_wTtbCcFgWSp_R--c6C1WkL-8C6CoC2oxGW3-Elle626lrhodLWMrjkvX9otKz3iISnW3VzPqyFPt7n9lgmRn1VHDi5VN56q0pqTHvkarqydpU4ErovV_J4f2BizfFg-1XKrh5WgPjZtMmznG7g3LYAeTUkp-TqmZKcdWG2OZ4BK2W9uH0J0TBVu2QcVmVg1hL6BDsJ4ksrsr3sut58yGs_GuI_Q73wZ0zfY"
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
