package main

import (
	"encoding/json"
	"fmt"
	"github.com/Implex-ltd/engine/internal/crawler"
	"log"
	"math/rand"
	"os"
	"os/signal"
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
	pool []*browser.Instance
	mt   sync.Mutex
	curr = 0
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

func initBrowser() {
	c := goccm.New(Config.Engine.BrowserCount)

	i := 0.0
	j := 0
	for {
		c.Wait()

		go func(i float64) {
			defer c.Done()

			br, err := api.GetBrowser(Config.Browser.Name, Config.Browser.Useragent, Config.Browser.Os)
			if err != nil {
				log.Println(err)
				return
			}

			cdp, err := br.Start()
			if err != nil {
				log.Println(err)
				return
			}

			client, err := browser.NewInstance(&browser.InstanceConfig{
				Mock:     Config.Mock.BlockRegister,
				IsRaw:    Config.Browser.Name == "gologinraw",
				Spoof:    Config.Mock.InjectSpoof,
				Headless: Config.Browser.Headless,
				Threads:  Config.Engine.BrowserHswThreadCount,
				Version:  Config.Mock.HswVersion,
				CDP:      cdp,
				Path:     cdp,
				Inject:   Config.Mock.InjectHsw,
				Hsj:      Config.Mock.EnableHsj,
				API:      br,
			})
			if err != nil {
				log.Println("NewInstance", err)
				return
			}

			defer func(client *browser.Instance) {
				err := client.CloseInstance()
				if err != nil {
					log.Println(err)
				}
			}(client)

			mt.Lock()
			pool = append(pool, client)
			mt.Unlock()

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
	currBrowser := next()

	pow, err := currBrowser.Hsw(b.Jwt, 10*time.Second)
	if err != nil {
		currBrowser.Online = false
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

	log.Printf("solved (%dms) (%d): %s\n", time.Since(t).Milliseconds(), len(pow), pow[:50])
	return c.Send([]byte(pow))
}

func debug() {
	log.Println("ctrl+c to exit.")

	gotos := []string{
		"https://abrahamjuliot.github.io/creepjs/",
	}
	lock := true

	c := crawler.NewCrawler(Config.Mock.HswVersion, Config.Browser.Name, Config.Browser.Useragent, Config.Browser.Os, gotos, lock, Config.Mock.InjectSpoof, Config.Mock.EnableHsj, Config.Browser.Headless, Config.Mock.InjectHsw)

	out, err := c.Run()
	log.Println(out)

	if err != nil {
		panic(err)
	}
}

func cgfHandler(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf(`{"useragent": "%s"}`, Config.Browser.Useragent))
}

func engine() {
	go initBrowser()

	app := fiber.New()
	app.Post("/n", solveHandler)
	app.Post("/config", cgfHandler)

	err := app.Listen(fmt.Sprintf(`:%d`, Config.Server.Port))
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}

func main() {
	err := playwright.Install()
	if err != nil {
		panic(err)
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))

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

	var wg sync.WaitGroup

	go func() {
		<-interrupt
		fmt.Println("\nExit...")

		for _, br := range pool {
			wg.Add(1)
			go func(b *browser.Instance) {
				defer wg.Done()

				err := b.CloseInstance()
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
