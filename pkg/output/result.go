package output

import (
	"time"

	"github.com/projectdiscovery/katana/pkg/navigation"
)

// Result of the crawling
type Result struct {
	Timestamp time.Time            `json:"timestamp,omitempty"`
	Request   *navigation.Request  `json:"-"`
	Response  *navigation.Response `json:"-"`
	Error     string               `json:"error,omitempty"`

	// HTTPX Fields
	Hashes          map[string]interface{} `json:"hash,omitempty"`
	Port            string                 `json:"port,omitempty"`
	URL             string                 `json:"url,omitempty"`
	Input           string                 `json:"input,omitempty"`
	Title           string                 `json:"title,omitempty"`
	Scheme          string                 `json:"scheme,omitempty"`
	WebServer       string                 `json:"webserver,omitempty"`
	ResponseBody    string                 `json:"body,omitempty"`
	ContentType     string                 `json:"content_type,omitempty"`
	Method          string                 `json:"method,omitempty"`
	Host            string                 `json:"host,omitempty"`
	Path            string                 `json:"path,omitempty"`
	Header          map[string]string      `json:"header,omitempty"`
	RawHeader       string                 `json:"raw_header,omitempty"`
	RequestRaw      string                 `json:"request,omitempty"`
	ResponseTime    string                 `json:"time,omitempty"`
	A               []string               `json:"a,omitempty"`
	Technologies    []string               `json:"tech,omitempty"`
	Words           int                    `json:"words"`
	Lines           int                    `json:"lines"`
	StatusCode      int                    `json:"status_code"`
	ContentLength   int                    `json:"content_length"`
	Failed          bool                   `json:"failed"`
	KnowledgeBase   map[string]interface{} `json:"knowledgebase,omitempty"`
	Resolvers       []string               `json:"resolvers,omitempty"`
}

// HasResponse checks if the result has a valid response
func (r *Result) HasResponse() bool {
	return r.Response != nil && r.Response.Resp != nil
}
