package api

var AvailableBrowsers = map[string]BrowserConstructor{
	"gologin":    NewGologinConstructor,
	"hidenium":   NewHideniumConstructor,
	"gologinraw": NewGologinrawConstructor,
}

type Config struct {
	UserAgent string
	Os        string
}

type Browser interface {
	Delete() error
	Close() error
	Start() (string, error)
	Create() error
}

type BrowserAPI struct {
	Browser Browser
}

type BrowserConstructor func(config *Config) (Browser, error)
