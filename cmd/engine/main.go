package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Implex-ltd/engine/internal/api"
	"github.com/Implex-ltd/engine/internal/browser"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/playwright-community/playwright-go"
	"github.com/zenthangplus/goccm"
)

var (
	pool     []*browser.Instance
	browsers []api.Browser
	mt       sync.Mutex
	curr     = 0
)

func next() *browser.Instance {
	for {
		if len(pool) == 0 {
			continue
		}

		mt.Lock()
		if curr >= len(pool) {
			curr = 0
		}

		c := pool[curr]
		curr++
		mt.Unlock()

		if !c.Online {
			continue
		}

		return c
	}
}

func GetBrowser() (api.Browser, error) {
	browser, err := api.NewBrowserByName(Config.Browser.Name, &api.Config{
		UserAgent: Config.Browser.Useragent,
		Os:        Config.Browser.Os,
	})

	if err != nil {
		return nil, err
	}

	if err := browser.Create(); err != nil {
		return nil, err
	}

	browsers = append(browsers, browser)

	return browser, nil
}

func initBrowser() {
	c := goccm.New(Config.Engine.BrowserCount)

	i := 0.0
	j := 0
	for {
		c.Wait()

		go func(i float64) {
			defer c.Done()

			br, err := GetBrowser()
			if err != nil {
				log.Println(err)
				return
			}

			cdp, err := br.Start()
			if err != nil {
				log.Println(err)
				return
			}

			defer br.Close()

			client, err := browser.NewInstance(&browser.InstanceConfig{
				Mock:     Config.Mock.BlockRegister,
				IsRaw:    Config.Browser.Name == "gologinraw",
				Spoof:    Config.Mock.InjectSpoof,
				Headless: Config.Browser.Headless,
				Threads:  Config.Engine.BrowserHswThreadCount,
				Version:  Config.Mock.HswVersion,
				CDP:      cdp,
				Path:     cdp,
			})
			if err != nil {
				log.Println("NewInstance", err)
				return
			}

			defer client.CloseInstance()

			if err := client.NavigateToDiscord(); err != nil {
				log.Println("NavigateToDiscord", err)
				return
			}

			if err := client.TriggerCaptcha(); err != nil {
				log.Println("TriggerCaptcha", err)
				return
			}

			log.Println("Hooked!")
			client.Online = true

			mt.Lock()
			pool = append(pool, client)
			mt.Unlock()

			defer func() {
				mt.Lock()
				defer mt.Unlock()

				for i, c := range pool {
					if c == client {
						pool = append(pool[:i], pool[i+1:]...)
						break
					}
				}
			}()

			t := time.NewTicker(time.Second)
			st := time.Now()
			defer t.Stop()

			for client.Online {
				select {
				case <-t.C:
					if time.Since(st).Seconds() > (float64(Config.Engine.Rotation) + i) {
						client.Online = false
						log.Println("gracefully restarting")
						break
					}

					if !client.Online {
						client.Online = false
						log.Println("browser crashed, restarting")
						break
					}
				}
			}
		}(i)

		i += float64(Config.Engine.RotationWait)
		j++

		if j >= Config.Engine.BrowserCount {
			i = 0.0
			j = 0
		}
	}
}

func solveHandler(c *fiber.Ctx) error {
	var b Payload
	if err := json.Unmarshal(c.Body(), &b); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Review your input",
			"data":    err.Error(),
		})
	}

	if b.Jwt == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Review your input",
			"data":    fmt.Errorf("please provide jwt token"),
		})
	}

	if len(b.Jwt) < 61 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Review your input",
			"data":    fmt.Errorf("token is invalid"),
		})
	}

	t := time.Now()
	browser := next()

	pow, err := browser.Hsw(b.Jwt, 10*time.Second)
	if err != nil {
		browser.Online = false
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "cant eval",
			"data":    err.Error(),
		})
	}

	if len(pow) < 50 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "error",
			"data":    fmt.Errorf("n lenght is invalid"),
		})
	}

	log.Printf("solved (%dms): %s\n", time.Since(t).Milliseconds(), pow[:50])
	return c.Send([]byte(pow))
}

func crawl(url string, headless bool) {
	var name string
	var head string

	if headless {
		head = "headless"
	} else {
		head = "headfull"
	}

	if strings.Contains(url, "browserleaks") {
		name = strings.Split(url, "browserleaks.com/")[1]
	} else {
		name = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(url, "://")[1], ".", ""), "/", ""), ":", "")
	}

	br, err := GetBrowser()
	if err != nil {
		log.Println(err)
		return
	}

	cdp, err := br.Start()
	if err != nil {
		log.Println(err)
		return
	}

	defer br.Close()

	client, err := browser.NewInstance(&browser.InstanceConfig{
		Mock:     Config.Mock.BlockRegister,
		IsRaw:    Config.Browser.Name == "gologinraw",
		Spoof:    Config.Mock.InjectSpoof,
		Headless: Config.Browser.Headless,
		Threads:  Config.Engine.BrowserHswThreadCount,
		Version:  Config.Mock.HswVersion,
		CDP:      cdp,
		Path:     cdp,
	})
	if err != nil {
		log.Println(err)
		return
	}

	defer client.CloseInstance()
	client.Page.Goto(url)

	log.Println("Loaded", url)
	defer log.Println("Done", url)

	time.Sleep(5 * time.Second)
	client.Page.Screenshot(playwright.PageScreenshotOptions{
		FullPage: playwright.Bool(true),
		Path:     playwright.String(fmt.Sprintf("./img/%s/%s.png", head, name)),
	})

	select {}
}

func debug() {
	log.Println("ctrl+c to exit.")

	gotos := []string{
		"https://abrahamjuliot.github.io/creepjs/",
	}

	c := goccm.New(len(gotos))

	for _, page := range gotos {
		c.Wait()

		go func(url string) {
			defer c.Done()
			go crawl(url, false)
			crawl(url, true)
		}(page)
	}

	c.WaitAllDone()
	log.Println("Done.")
}

func engine() {
	go initBrowser()

	app := fiber.New()
	app.Post("/n", solveHandler)

	err := app.Listen(fmt.Sprintf(`:%d`, Config.Server.Port))
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}

func main() {
	playwright.Install()
	rand.Seed(time.Now().UnixNano())

	if _, err := toml.DecodeFile("../../scripts/config.toml", &Config); err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "debug":
		go debug()
	case "engine":
		go engine()
	default:
		panic("invalid args. use: debug, engine")
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Use a sync.WaitGroup to wait for cleanup before exiting
	var wg sync.WaitGroup

	go func() {
		<-interrupt
		fmt.Println("\nExit...")

		for _, br := range browsers {
			wg.Add(1)
			go func(b api.Browser) {
				defer wg.Done()

				err := b.Close()
				if err != nil {
					log.Println("cant close: ", err)
				}
			}(br)
		}

		wg.Wait()
		os.Exit(0)
	}()

	select {}
}
