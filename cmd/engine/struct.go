package main

var (
	Config = Cfg{}
)

type Payload struct {
	Jwt string `json:"jwt"`
}

type Cfg struct {
	Engine struct {
		BrowserHswThreadCount int `toml:"browserHswThreadCount"`
		BrowserCount          int `toml:"browserCount"`
		Rotation              int `toml:"rotation"`
		RotationWait          int `toml:"rotation_wait"`
	} `toml:"engine"`
	Browser struct {
		Name      string `toml:"name"`
		Useragent string `toml:"useragent"`
		Os        string `toml:"os"`
		Headless  bool   `toml:"headless"`
	} `toml:"browser"`
	Mock struct {
		EnableHsj     bool   `toml:"enable_hsj"`
		InjectHsw     bool   `toml:"inject_hsw"`
		InjectSpoof   bool   `toml:"inject_spoof"`
		BlockRegister bool   `toml:"block_register"`
		HswVersion    string `toml:"hsw_version"`
	} `toml:"mock"`
	Server struct {
		Port int `toml:"port"`
	} `toml:"server"`
	Debug struct {
		Gotos []string `toml:"gotos"`
	} `toml:"debug"`
}
