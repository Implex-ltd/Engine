package main

var (
	Config = Cfg{}
)

type Cfg struct {
	Discord struct {
		Invite             string `toml:"invite"`
		InviteGuildID      string `toml:"inviteGuildID"`
		InviteChannelID    string `toml:"inviteChannelID"`
		SendChannelID      string `toml:"sendChannelID"`
		SendMessage        bool   `toml:"sendMessage"`
		SendMessageContent string `toml:"sendMessageContent"`
		TryJoin            bool   `toml:"tryJoin"`
		Humanize           bool   `toml:"humanize"`
		Version            int    `toml:"version"`
	} `toml:"discord"`
	Hcaptcha struct {
		Lang          string `toml:"lang"`
		Version       string `toml:"version"`
		SubmitTime    int    `toml:"submitTime"`
		HswAddress    string `toml:"hswAddress"`
		SolverAddress string `toml:"SolverAddress"`
		SiteKey       string `toml:"siteKey"`
	} `toml:"hcaptcha"`
	Spoof struct {
		Cfbm             bool   `toml:"cfbm"`
		ProfilePath      string `toml:"profilePath"`
		ProxyType        string `toml:"proxyType"`
		SessionReplaceMe string `toml:"sessionReplaceMe"`
	} `toml:"spoof"`
	Performances struct {
		Goroutines int  `toml:"goroutines"`
		Debug      bool `toml:"debug"`
	} `toml:"performances"`
	Engine struct {
		BrowserHswThreadCount int `toml:"browserHswThreadCount"`
		BrowserCount          int `toml:"browserCount"`
	} `toml:"engine"`
}
