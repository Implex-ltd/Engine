package api

import (
	"errors"
	"github.com/Implex-ltd/engine/internal/api/octo"

	dolphin "github.com/Implex-ltd/engine/internal/api/dolphin"
	gologin "github.com/Implex-ltd/engine/internal/api/gologin"
	gologinraw "github.com/Implex-ltd/engine/internal/api/gologinraw"
	hidenium "github.com/Implex-ltd/engine/internal/api/hidenium"
)

func NewBrowserByName(name string, config *Config) (Browser, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	constructor, found := AvailableBrowsers[name]
	if !found {
		return nil, errors.New("unknown browser")
	}

	return constructor(config)
}

func NewGologinConstructor(config *Config) (Browser, error) {
	return gologin.NewGologin(config.UserAgent, config.Os)
}

func NewGologinrawConstructor(config *Config) (Browser, error) {
	return gologinraw.NewGologinRaw(config.UserAgent, config.Os)
}

func NewHideniumConstructor(config *Config) (Browser, error) {
	return hidenium.NewHidenium(config.UserAgent, config.Os)
}

func NewDolphinConstructor(config *Config) (Browser, error) {
	return dolphin.NewDolphin(config.UserAgent, config.Os)
}

func NewOctoConstructor(config *Config) (Browser, error) {
	return octo.NewOcto(config.UserAgent, config.Os)
}

func GetBrowser(name, ua, os string) (Browser, error) {
	browser, err := NewBrowserByName(name, &Config{
		UserAgent: ua,
		Os:        os,
	})

	if err != nil {
		return nil, err
	}

	if err := browser.Create(); err != nil {
		return nil, err
	}

	return browser, nil
}
