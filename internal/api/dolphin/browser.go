package dolphin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Implex-ltd/engine/internal/utils"
)

var (
	token    = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiIxIiwianRpIjoiZjJmOTk5ZTkxZWUzNTQxNTFmZGU2ZmQ0ODRlZTU4Y2NjMzJmZGExNTFkNzQwZjhlYjBhNDIzNjRkMTk1YWVhMDc2ZGRjZmM4MDQ5MzBlNDgiLCJpYXQiOjE2OTYwNjExODguNDMwNjY1LCJuYmYiOjE2OTYwNjExODguNDMwNjY2LCJleHAiOjE2OTg2NTMxODguNDE4ODM0LCJzdWIiOiIyNzA0NjkzIiwic2NvcGVzIjpbXX0.ZFEsCz6atc9MgrGpPJCZEiu75WJYTn7FPGZRi1whky16rx7ftUd8EButSyPWHIwAU-utWu9BN3Kmwc1U99rlCN79xjuqCoVcBBx0bKXORJjySwdHKhK4A1WMb7N0IebpZzSQTNvX7l_Slhmx2Fn-7FCuWJKLIrjKX8oSqCKNKxru5sdeVsYt-Yj2uHlGEg9-83YwsiBFsa5KPRoHjhSjiyY4uQ1_mnuuuDHmifNI88FOXQ_PAubkrRNa2vpar2CuWrYVorMTghTjY13ny8EpnuTx5XO-L_1qokc24ra-H9rSzplzqPxcpzz2TjtXIpxJgEaMXfpBWTrt7xmQi_4odA3efVL90drZxgTCi7nthIBi_WbSAxDZas_qoJzVOr39oUX4C9J4nTJyP_O5WY6DtQpGdbRoOnAUgGwrOj75EmhzXuan9rdAeozfxIFHT2xNGBXkF2YVaPKQLR9URCkg1ICPk8erGsOj6TzXpoza_Riv7rAmHtjyUgdsiO75SmokKMS8_sSAlfQGCxLxiJeOEpN1wbr3Qm9L363Gu33pcps1OPM1bRgUonDD7aSLTEzatEeUGEWrKccVT6DmFXiNlI7XM0EFuJlplSNzAjmrbw9EHTtYVOgyKcu0VG8nnMUQRlkXobEabIn6CzD4O9ikGHBjEBxQ8pXbs0YdyQjswYE"
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

	vendor, renderer := utils.GetRandomVendorRenderer()

	w, h := utils.GenerateRandomBrowserSize(800, 3500)

	Self.Config = DolphinCreateBrowserPayload{
		Useragent: Locale{
			Mode:  "manual",
			Value: UserAgent,
		},
		Platform:    pname,
		Name:        utils.GenerateRandomString(5),
		BrowserType: "anty",
		Webrtc: Webrtc{
			Mode: "altered",
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
			Mode:  "manual",
			Value: "Europe/Paris",
		},
		Locale: Locale{
			Mode: "fr-FR",
		},
		WebglInfo: WebglInfo{
			Mode:     "manual",
			Vendor:   vendor,
			Renderer: renderer,
		},
		ClientRect: Audio{
			Mode: "real",
		},
		Memory: CPU{
			Mode:  "manual",
			Value: utils.RandomElementInt([]int64{2, 4, 8, 16}),
		},
		CPU: CPU{
			Mode:  "manual",
			Value: utils.RandomElementInt([]int64{2, 4, 6, 8, 12, 16}),
		},
		Geolocation: Geolocation{
			Mode: "auto",
		},
		Screen: Screen{
			Mode:       "manual",
			Resolution: fmt.Sprintf("%dx%d", w, h),
		},
		/*Audio: Audio{
			Mode: "noise",
		},*/
		MediaDevices: MediaDevices{
			Mode:         "manual",
			AudioInputs:  utils.RandomElementInt([]int64{0, 1, 2, 3}),
			VideoInputs:  utils.RandomElementInt([]int64{1, 2, 3}),
			AudioOutputs: utils.RandomElementInt([]int64{0, 1, 2, 3}),
		},
	}

	return Self, nil
}

func (D *Dolphin) Create() error {
	Payload, err := json.Marshal(D.Config)
	if err != nil {
		return err
	}

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

	var data DolphinCreateBrowserResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(string(body))
		return err
	}

	D.UUID = fmt.Sprintf("%d", data.BrowserProfileID)
	return nil
}

func (D *Dolphin) Start() (string, error) {
	resp, err := D.Client.Get(fmt.Sprintf("%s/browser_profiles/%s/start?automation=1", endpoint, D.UUID))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

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
