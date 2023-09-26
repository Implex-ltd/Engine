package pool

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Implex-ltd/engine/internal/api"
	"github.com/Implex-ltd/engine/internal/browser"
	"github.com/Implex-ltd/engine/internal/config"
	"github.com/zenthangplus/goccm"
)

func NewPool(size int) (*Pool, error) {
	return &Pool{
		Size: size,
		Mut:  &sync.Mutex{},
		Curr: 0,
	}, nil
}

func (P *Pool) NextWorker() (*browser.Instance, error) {
	P.Mut.Lock()
	defer P.Mut.Unlock()

	lp := len(P.Workers)

	if lp == 0 {
		return nil, fmt.Errorf("no workers available")
	}

	if P.Curr >= lp {
		P.Curr = 0
	}

	c := P.Workers[P.Curr]
	P.Curr++

	if !c.Online {
		P.Workers = append(P.Workers[:P.Curr-1], P.Workers[P.Curr:]...)
		return nil, fmt.Errorf("worker is offline")
	}

	return c, nil
}

func (P *Pool) AddWorker(w *browser.Instance) {
	P.Mut.Lock()
	defer P.Mut.Unlock()

	P.Workers = append(P.Workers, w)
}

func (P *Pool) Worker(i int) error {
	br, err := api.GetBrowser(config.Config.Browser.Name, config.Config.Browser.Useragent, config.Config.Browser.Os)
	if err != nil {
		return err
	}

	cdp, err := br.Start()
	if err != nil {
		return err
	}

	client, err := browser.NewInstance(&browser.InstanceConfig{
		Mock:     config.Config.Mock.BlockRegister,
		IsRaw:    config.Config.Browser.Name == "gologinraw",
		Spoof:    config.Config.Mock.InjectSpoof,
		Headless: config.Config.Browser.Headless,
		Threads:  config.Config.Engine.BrowserHswThreadCount,
		Version:  config.Config.Mock.HswVersion,
		CDP:      cdp,
		Path:     cdp,
		Inject:   config.Config.Mock.InjectHsw,
		Hsj:      config.Config.Mock.EnableHsj,
		API:      br,
	})
	if err != nil {
		log.Println("non handled browser failure")
		return err
	}

	defer func(client *browser.Instance) {
		err := client.CloseInstance()
		if err != nil {
			log.Println(err)
		}
	}(client)

	if err := client.NavigateToDiscord(); err != nil {
		return err
	}

	if err := client.TriggerCaptcha(); err != nil {
		return err
	}

	P.AddWorker(client)
	client.Online = true

	t := time.NewTicker(time.Second)
	st := time.Now()

	defer t.Stop()

	for range t.C {
		if !client.Online {
			return fmt.Errorf("browser crashed, restarting")
		}

		elapsedSeconds := time.Since(st).Seconds()
		gracefulRestartThreshold := float64(config.Config.Engine.Rotation + i)

		if elapsedSeconds > gracefulRestartThreshold {
			client.Online = false
			log.Println("gracefully restarting")
			return nil
		}
	}

	return nil
}

func (P *Pool) Run() {
	c := goccm.New(P.Size)
	i := 0

	for {
		c.Wait()

		go func(i int) {
			defer c.Done()
			if err := P.Worker(i); err != nil {
				log.Println("worker exited with error:", err)
			}
		}(i)

		i++
		if i > P.Size {
			i = 0
		}
	}
}
