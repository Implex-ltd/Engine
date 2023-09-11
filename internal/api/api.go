package api

import (
	"errors"

	gologin "github.com/Implex-ltd/engine/internal/api/gologin"
	gologinraw "github.com/Implex-ltd/engine/internal/api/gologinraw"
	hidenium "github.com/Implex-ltd/engine/internal/api/hidenium"
)

func NewGologinConstructor(config *Config) (Browser, error) {
	return gologin.NewGologin(config.UserAgent, config.Os)
}

func NewGologinrawConstructor(config *Config) (Browser, error) {
	return gologinraw.NewGologinRaw(config.UserAgent, config.Os)
}

func NewHideniumConstructor(config *Config) (Browser, error) {
	return hidenium.NewHidenium(config.UserAgent, config.Os)
}

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

// implement
func NewBrowserAPI(br Browser) *BrowserAPI {
	return &BrowserAPI{
		Browser: br,
	}
}

func (api *BrowserAPI) Delete() error {
	return api.Browser.Delete()
}

func (api *BrowserAPI) Close() error {
	return api.Browser.Close()
}

func (api *BrowserAPI) Start() (string, error) {
	return api.Browser.Start()
}

func (api *BrowserAPI) Create() error {
	return api.Browser.Create()
}
