package crawler

type Crawler struct {
	Screenshot, spoof, hsj, headless, inject, lock bool
	hcversion, brName, ua, os                      string
	Threads                                        int
	Gotos                                          []string
}
