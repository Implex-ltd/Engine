package main

var (
	Config = Cfg{}
)

type Cfg struct {
	Engine struct {
		BrowserHswThreadCount int  `toml:"browserHswThreadCount"`
		BrowserCount          int  `toml:"browserCount"`
		Hidenium              bool `toml:"hidenium"`
		Gologin               bool `toml:"gologin"`
		Rotation              int  `toml:"rotation"`
		RotationWait          int  `toml:"rotation_wait"`
	} `toml:"engine"`
	Mock struct {
		Hsw      bool   `toml:"hsw"`
		Spoofing bool   `toml:"spoofing"`
		Version  string `toml:"version"`
	} `toml:"mock"`
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
}

type Payload struct {
	Jwt string `json:"jwt"`
}
