package crawler

import (
	"fmt"
	"github.com/Implex-ltd/engine/internal/api"
	"github.com/Implex-ltd/engine/internal/browser"
	"github.com/playwright-community/playwright-go"
	"github.com/zenthangplus/goccm"
	"log"
	"strings"
	"sync"
	"time"
)

func NewCrawler(hcversion, brName, ua, os string, gotos []string, lock, spoof, hsj, headless, inject bool) *Crawler {
	return &Crawler{
		Screenshot: false,
		Gotos:      gotos,
		Threads:    len(gotos),
		brName:     brName,
		ua:         ua,
		os:         os,
		spoof:      spoof,
		hsj:        hsj,
		headless:   headless,
		inject:     inject,
		lock:       lock,
	}
}

func (C *Crawler) Crawl(url string) error {
	// Todo: make it into api.getbrowser
	br, err := api.GetBrowser(C.brName, C.ua, C.os)
	if err != nil {
		return err
	}

	cdp, err := br.Start()
	if err != nil {
		return err
	}

	client, err := browser.NewInstance(&browser.InstanceConfig{
		Mock:     false,
		IsRaw:    C.brName == "gologinraw",
		Spoof:    C.spoof,
		Headless: C.headless,
		Threads:  0,
		Version:  C.hcversion,
		CDP:      cdp,
		Path:     cdp,
		Inject:   C.inject,
		Hsj:      C.hsj,
	})
	if err != nil {
		return err
	}

	defer client.CloseInstance()

	time.Sleep(time.Second)

	_, err = client.Page.Goto(url)
	if err != nil {
		return err
	}

	log.Println("Loaded", url)
	defer log.Println("Done", url)

	var name string
	var head string

	if C.headless {
		head = "headless"
	} else {
		head = "headfull"
	}

	if strings.Contains(url, "browserleaks") {
		name = strings.Split(url, "browserleaks.com/")[1]
	} else {
		name = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(url, "://")[1], ".", ""), "/", ""), ":", "")
	}

	if C.Screenshot {
		time.Sleep(5 * time.Second)
		_, err := client.Page.Screenshot(playwright.PageScreenshotOptions{
			FullPage: playwright.Bool(true),
			Path:     playwright.String(fmt.Sprintf("./img/%s/%s.png", head, name)),
		})
		if err != nil {
			return err
		}
	}

	if C.lock {
		select {}
	}

	return nil
}

func (C *Crawler) Run() (errors map[string]error, err error) {
	c := goccm.New(C.Threads)
	mut := sync.Mutex{}

	errors = map[string]error{}

	for _, url := range C.Gotos {
		c.Wait()

		go func(url string) {
			defer c.Done()

			if err := C.Crawl(url); err != nil {
				mut.Lock()
				errors[url] = err
				mut.Unlock()
			}
		}(url)
	}

	c.WaitAllDone()
	return errors, err
}
