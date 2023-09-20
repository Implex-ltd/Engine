package browser

import (
	"github.com/playwright-community/playwright-go"
)
/*
var (
	BUTTON_LOGIN   = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.ctaContainer-5100Xg > button`
	CHECKBOX_TOS   = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.formContainer-1Mw7aR > div > div > div`
	USERNAME_INPUT = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.formContainer-1Mw7aR > form > input`
	BUTTON_SUBMIT  = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.formContainer-1Mw7aR > form > button > img`
	HCAPTCHA_FRAME = `//iframe[contains(@title,'Widget contenant une case à cocher pour le défi de sécurité hCaptcha')]`
)
*/
// /invite

var (
	BUTTON_LOGIN   = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.ctaContainer-5100Xg > button`
	CHECKBOX_TOS   = `#app-mount > div.appAsidePanelWrapper-ev4hlp > div.notAppAsidePanel-3yzkgB > div.app-3xd6d0 > div > div > div > div > form > div.centeringWrapper-dGnJPQ > div.block-3uVSn4.marginTop40-Q4o1tS > div.flex-2S1XBF.horizontal-112GEH.justifyStart-2Mwniq.alignCenter-14kD11.noWrap-hBpHBz.marginTop20-2T8ZJx > label > input`
	USERNAME_INPUT = `#uid_6`
	BUTTON_SUBMIT  = `#app-mount > div.appAsidePanelWrapper-ev4hlp > div.notAppAsidePanel-3yzkgB > div.app-3xd6d0 > div > div > div > div > form > div.centeringWrapper-dGnJPQ > div.block-3uVSn4.marginTop40-Q4o1tS > div:nth-child(4) > button`
	HCAPTCHA_FRAME = `//iframe[contains(@title,'Widget contenant une case à cocher pour le défi de sécurité hCaptcha')]`
)


type Instance struct {
	Pw      *playwright.Playwright
	Br      playwright.Browser
	Page    playwright.Page
	Frame   playwright.Frame
	Manager chan struct{}
	Online  bool
}

type InstanceConfig struct {
	Mock     bool
	IsRaw    bool
	Spoof    bool
	Headless bool
	Inject   bool
	Threads  int
	Version  string
	CDP      string
	Path     string
	Hsj      bool
}
