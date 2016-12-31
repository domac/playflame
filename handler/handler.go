package handler

import (
	"fmt"
	"github.com/domac/playflame/stats"
	"github.com/varstr/uaparser"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Simple(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello VIP!")
}

//高级封装
func WithAdvanced(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		tags := getStatsTags(r)
		stats.IncCounter("handler.received", tags, 1)
		h(w, r)
		duration := time.Since(start)
		stats.RecordTimer("handler.latency", tags, duration)
	}
}

func getStatsTags(r *http.Request) map[string]string {
	userBrowser, userOS := parseUserAgent(r.UserAgent())
	stats := map[string]string{
		"browser":  userBrowser,
		"os":       userOS,
		"endpoint": filepath.Base(r.URL.Path),
	}
	host, err := os.Hostname()
	if err == nil {
		if idx := strings.IndexByte(host, '.'); idx > 0 {
			host = host[:idx]
		}
		stats["host"] = host
	}
	return stats
}

func parseUserAgent(uaString string) (browser, os string) {
	ua := uaparser.Parse(uaString)

	if ua.Browser != nil {
		browser = ua.Browser.Name
	}
	if ua.OS != nil {
		os = ua.OS.Name
	}

	return browser, os
}
