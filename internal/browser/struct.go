package browser

import (
	"github.com/playwright-community/playwright-go"
)

var (
	BUTTON_LOGIN   = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.ctaContainer-5100Xg > button`
	CHECKBOX_TOS   = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.formContainer-1Mw7aR > div > div > div`
	USERNAME_INPUT = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.formContainer-1Mw7aR > form > input`
	BUTTON_SUBMIT  = `#app-mount > div > div > div.grid-3d2PVT.heroBackground-itJzsn > div.row-31oovZ.heroContainer-3YphMz > div > div.formContainer-1Mw7aR > form > button > img`
	HCAPTCHA_FRAME = `//iframe[contains(@title,'Widget containing checkbox for hCaptcha security challenge')]`
)

type Instance struct {
	Pw      *playwright.Playwright
	Br      playwright.Browser
	Page    playwright.Page
	Frame   playwright.Frame
	Manager chan struct{}
}
