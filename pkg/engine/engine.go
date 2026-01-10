package engine

type Engine interface {
	Crawl(rawInput, rootURL string) error
	Close() error
}
