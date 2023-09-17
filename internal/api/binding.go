package api

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
