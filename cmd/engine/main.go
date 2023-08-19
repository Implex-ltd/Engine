package main

import (
	"fmt"
	"github.com/Implex-ltd/engine/internal/browser"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/playwright-community/playwright-go"
	"github.com/zenthangplus/goccm"
)

var (
	pool []*browser.Instance
	mt   sync.Mutex
	curr = 0
)

func next() *browser.Instance {
	mt.Lock()
	defer mt.Unlock()

	if curr >= len(pool) {
		curr = 0
	}

	c := pool[curr]
	curr++

	return c
}

func initBrowser() {
	c := goccm.New(Config.Engine.BrowserCount)

	for {
		c.Wait()

		go func() {
			defer c.Done()

			client, err := browser.NewInstance(true, false, Config.Engine.BrowserHswThreadCount)
			if err != nil {
				log.Println(err)
				return
			}

			defer client.CloseInstance()

			if err := client.NavigateToDiscord(); err != nil {
				log.Println(err)
				return
			}

			if err := client.TriggerCaptcha(); err != nil {
				log.Println(err)
				return
			}

			log.Println("Hooked!")

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
			select {}
		}()
	}
}

func solveHandler(c *fiber.Ctx) error {
	token := c.Query("req")

	if token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token is missing",
		})
	}

	if len(token) < 61 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Token is invalid",
		})
	}

	log.Println("recv", token[:61])

	t := time.Now()
	pow, err := next().Hsw(token)
	if err != nil {
		return c.Send([]byte(""))
	}

	log.Printf("solved (%dms): %s\n", time.Since(t).Milliseconds(), pow[:50])
	return c.Send([]byte(pow))
}

func engine() {
	go initBrowser()

	app := fiber.New()

	app.Get("/n", solveHandler)

	err := app.Listen(":4000")
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
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

	client, err := browser.NewInstance(true, headless, Config.Engine.BrowserHswThreadCount)
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
		"https://browserleaks.com/webgl",
		"https://browserleaks.com/webrtc",
		"https://browserleaks.com/canvas",
		"https://browserleaks.com/webgl",
		"https://browserleaks.com/tls",
		"https://browserleaks.com/javascript",
		"https://browserleaks.com/fonts",
		"https://browserleaks.com/ip",
		"https://bot.sannysoft.com/",
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

func main() {
	if _, err := toml.DecodeFile("../../../../scripts/config.toml", &Config); err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "debug":
		debug()
	case "engine":
		engine()
	default:
		panic("invalid args. use: debug, engine")
	}
}
