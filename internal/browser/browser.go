package browser

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_"github.com/Implex-ltd/engine/internal/gologin"
	"github.com/playwright-community/playwright-go"
)

func NewInstance(spoof, headless bool, threads int, cdp string, inject bool, version string) (*Instance, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, err
	}

	browser, err := pw.Chromium.ConnectOverCDP(cdp)
	if err != nil {
		panic(err)
	}

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		ColorScheme: playwright.ColorSchemeNoPreference,
	})
	if err != nil {
		return nil, err
	}

	/*gologpath, err := gologin.Setup()
	if err != nil {
		return nil, err
	}

	context, err := pw.Chromium.LaunchPersistentContext(gologpath, playwright.BrowserTypeLaunchPersistentContextOptions{
		ExecutablePath: playwright.String("C:\\Users\\arm\\Desktop\\MYBROWSER\\gologin\\orbita\\chrome.exe"),
		Headless:       playwright.Bool(false),
		ColorScheme:    playwright.ColorSchemeDark,
		Args: []string{
			"--lang=fr-FR", // fr-FR
			"--disable-encryption",
			"--font-masking-mode=3",
			"--flag-switches-begin",
			"--flag-switches-end",
			"--enable-quic",
			"--enable-tcp-fast-open",
			"--disable-remote-fonts",
			"--disable-background-networking",
			"--disable-extensions",
			"--no-first-run",

			//"--headless=new",
		},
	})
	if err != nil {
		return nil, err
	}*/

	if err := context.Route("**https://discord.com/api/v9/auth/register**", func(r playwright.Route) {
		r.Fulfill(playwright.RouteFulfillOptions{
			Status:      playwright.Int(400),
			ContentType: playwright.String("application/json"),
			Body:        []byte(`{"captcha_key": ["captcha-required"],"captcha_sitekey": "4c672d35-0701-42b2-88c3-78380b0db560","captcha_service": "hcaptcha"}`),
		})
	}); err != nil {
		return nil, err
	}

	if inject {
		hsw, err := os.ReadFile("../../cmd/engine/scripts/hsw.js")
		if err != nil {
			return nil, err
		}

		context.Route(fmt.Sprintf("**https://newassets.hcaptcha.com/c/%s/hsw.js**", version), func(r playwright.Route) {
			log.Println("Injected !")

			r.Fulfill(playwright.RouteFulfillOptions{
				Status: playwright.Int(200),
				Body:   hsw,
			})
		})
	}

	page, err := context.NewPage()
	if err != nil {
		return nil, err
	}

	if spoof {
		if err := page.AddInitScript(playwright.PageAddInitScriptOptions{
			Path: playwright.String("./scripts/spoof.js"),
		}); err != nil {
			return nil, err
		}
	}

	return &Instance{
		Online:  false,
		Pw:      pw,
		Br:      context.Browser(),
		Page:    page,
		Manager: make(chan struct{}, threads),
		Path:    "", //strings.ReplaceAll(gologpath, `\`, `\\`),
	}, nil
}

func (I *Instance) CloseInstance() error {
	var errList []error

	defer func() {
		if I.Path != "" {
			fmt.Println(I.Path)
			fmt.Println(os.RemoveAll(I.Path))
		}
	}()

	if I.Page != nil {
		if err := I.Page.Close(); err != nil {
			errList = append(errList, err)
		}
	}

	/*
		// make crash
		if I.Br != nil {
				if err := I.Br.Close(); err != nil {
					errList = append(errList, err)
				}
			}
	*/

	if I.Pw != nil {
		if err := I.Pw.Stop(); err != nil {
			errList = append(errList, err)
		}
	}

	if len(errList) > 0 {
		return fmt.Errorf("closing instance failed: %v", errList)
	}

	return nil
}

func (I *Instance) NavigateToDiscord() error {
	if _, err := I.Page.Goto("https://discord.com", playwright.PageGotoOptions{
		Timeout:   playwright.Float(10000),
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		return err
	}

	time.Sleep(time.Second * 2)

	bl, err := I.Page.WaitForSelector(BUTTON_LOGIN, playwright.PageWaitForSelectorOptions{
		State:   playwright.WaitForSelectorStateAttached,
		Timeout: playwright.Float(5000),
	})

	if err != nil {
		return err
	}

	if err := bl.Click(playwright.ElementHandleClickOptions{
		Timeout: playwright.Float(3000),
	}); err != nil {
		return err
	}

	if err := I.Page.Type(USERNAME_INPUT, "vichyontop1337", playwright.PageTypeOptions{
		Timeout: playwright.Float(3000),
		Delay:   playwright.Float(0),
	}); err != nil {
		return err
	}

	if err := I.Page.Click(CHECKBOX_TOS, playwright.PageClickOptions{
		Timeout: playwright.Float(3000),
		Delay:   playwright.Float(0),
	}); err != nil {
		return err
	}

	if err := I.Page.Click(BUTTON_SUBMIT, playwright.PageClickOptions{
		Timeout: playwright.Float(3000),
		Delay:   playwright.Float(0),
	}); err != nil {
		return err
	}

	return nil
}

func (I *Instance) TriggerCaptcha() error {
	bl, err := I.Page.WaitForSelector(HCAPTCHA_FRAME, playwright.PageWaitForSelectorOptions{
		State:   playwright.WaitForSelectorStateAttached,
		Timeout: playwright.Float(5000),
	})
	if err != nil {
		return err
	}

	time.Sleep(time.Second * 2)

	if err := bl.Click(playwright.ElementHandleClickOptions{
		Timeout: playwright.Float(5000),
	}); err != nil {
		return err
	}

	timeout := time.After(time.Second * 3)

	ticker := time.NewTicker(time.Millisecond * 250)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			for _, frame := range I.Page.Frames() {
				r, err := frame.Evaluate(`typeof hsw !== "undefined"`)
				if err != nil {
					continue
				}

				if r == true {
					I.Frame = frame
					return nil
				}
			}

		case <-timeout:
			return errors.New("timeout: hsw not found after 5 seconds")
		}
	}
}

func (I *Instance) Hsw(jwt string, timeoutDuration time.Duration) (string, error) {
	I.Manager <- struct{}{}
	defer func() {
		<-I.Manager
	}()

	resultChan := make(chan string)
	errChan := make(chan error)

	go func() {
		answer, err := I.Frame.Evaluate(fmt.Sprintf("spoofall();hsw(`%s`)", jwt), playwright.ElementHandleInputValueOptions{
			Timeout: playwright.Float(10000),
		})
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- fmt.Sprintf("%s", answer)
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errChan:
		fmt.Println(err)
		return "", err
	case <-time.After(timeoutDuration):
		return "", errors.New("evaluation timed out")
	}
}
