package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Implex-ltd/engine/internal/browser"
	"github.com/Implex-ltd/engine/internal/crawler"
	"github.com/Implex-ltd/engine/internal/pool"

	"github.com/Implex-ltd/engine/internal/config"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/playwright-community/playwright-go"
)

var P = &pool.Pool{}

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
	// Todo: add timeout
	var br *browser.Instance
	var err error
	
	for {
		br, err = P.NextWorker()
		if err != nil {
			log.Println(err)
			continue
		}

		break
	}

	pow, err := br.Hsw(b.Jwt, 10*time.Second)
	if err != nil {
		br.Online = false
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

	c := crawler.NewCrawler(gotos, lock)

	out, err := c.Run()
	log.Println(out)

	if err != nil {
		panic(err)
	}
}

func cgfHandler(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf(`{"useragent": "%s"}`, config.Config.Browser.Useragent))
}

func engine() error {
	var err error

	P, err = pool.NewPool(config.Config.Engine.BrowserCount)
	if err != nil {
		return err
	}

	go P.Run()

	app := fiber.New()
	app.Post("/n", solveHandler)
	app.Post("/config", cgfHandler)

	err = app.Listen(fmt.Sprintf(`:%d`, config.Config.Server.Port))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	//os.Setenv("DEBUGP", "1")

	err := playwright.Install()
	if err != nil {
		panic(err)
	}
	rand.New(rand.NewSource(time.Now().UnixNano()))

	if _, err := toml.DecodeFile("../../scripts/config.toml", &config.Config); err != nil {
		panic(err)
	}

	switch os.Args[1] {
	case "debug":
		go debug()
	case "engine":
		go func() {
			if err := engine(); err != nil {
				panic(err)
			}
		}()
	default:
		panic("invalid args. use: debug, engine")
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	/*var wg sync.WaitGroup

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
	}()*/

	select {}
}
