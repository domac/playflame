package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/domac/playflame/handler"
)

//端口
const hostPort = ":9090"

//go-torch -u http://localhost:9090 -t 30
//kapok -d=35 -c=1000  http://localhost:9090/advance

func main() {
	flag.Parse()

	//高级接口
	http.HandleFunc("/advance", handler.WithAdvanced(handler.Simple))

	//简单接口
	http.HandleFunc("/simple", handler.Simple)
	http.HandleFunc("/", index)

	fmt.Println("Starting Server on", hostPort)
	if err := http.ListenAndServe(hostPort, nil); err != nil {
		log.Fatalf("HTTP Server Failed: %v", err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "text/html")
	io.WriteString(w, "<h2>Links</h2>\n<ul>")
	for _, link := range []string{"/advance", "/simple"} {
		fmt.Fprintf(w, `<li><a href="%v">%v</a>`, link, link)
	}
	io.WriteString(w, "</ul>")
}
