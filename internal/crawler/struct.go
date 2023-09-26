package crawler

type Crawler struct {
	lock, Screenshot bool
	Threads          int
	Gotos            []string
}
