package main

var (
	Config = Cfg{}
)

type Cfg struct {
	Performances struct {
		Goroutines int  `toml:"goroutines"`
		Debug      bool `toml:"debug"`
	} `toml:"performances"`
	Engine struct {
		BrowserHswThreadCount int `toml:"browserHswThreadCount"`
		BrowserCount          int `toml:"browserCount"`
	} `toml:"engine"`
}

type Payload struct {
	Jwt string `json:"jwt"`
}