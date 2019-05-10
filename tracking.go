// Tracking cursor position in real-time without JavaScript
// Demo: https://twitter.com/davywtf/status/1124146339259002881

package main

import (
	"fmt"
	"net/http"
	"strings"
)

const W = 50
const H = 50

var ch chan string

const head = `<head>
<style>
*{margin:0;padding:0}
html,body{width:100%;height:100%}
p{
width:10px;
height:10px;
display:inline-block;
border-right:1px solid #666;
border-bottom:1px solid #666
}
</style>
</head>`

func main() {
	ch = make(chan string)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	if p == "/" {
		index(w, r)
		return
	} else if p == "/watch" {
		watch(w, r)
		return
	} else {
		if strings.HasPrefix(p, "/c") && strings.HasSuffix(p, ".png") {
			ch <- p[1 : len(p)-4]
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	w.Write([]byte(head))
	flusher.Flush()
	// Send <p> grid
	w.Write([]byte("<body>\n"))
	for i := 0; i < H; i++ {
		w.Write([]byte("<div>"))
		for j := 0; j < W; j++ {
			w.Write([]byte(fmt.Sprintf("<p id=\"c%dx%d\"></p>", i, j)))
		}
		w.Write([]byte("</div>\n"))
	}
	w.Write([]byte("</body>\n"))
	flusher.Flush()
	// Send CSS
	w.Write([]byte("<style>"))
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			id := fmt.Sprintf("c%dx%d", i, j)
			s := fmt.Sprintf("#%s:hover{background:url(\"%s.png\")}", id, id)
			w.Write([]byte(s))
		}
	}
	w.Write([]byte("</style>"))
	flusher.Flush()
}

func watch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}
	w.Write([]byte(head))
	flusher.Flush()
	// Send <p> grid
	w.Write([]byte("<body>\n"))
	for i := 0; i < H; i++ {
		w.Write([]byte("<div>"))
		for j := 0; j < W; j++ {
			w.Write([]byte(fmt.Sprintf("<p id=\"c%dx%d\"></p>", i, j)))
		}
		w.Write([]byte("</div>\n"))
	}
	w.Write([]byte("</body>\n"))
	flusher.Flush()
	// Listen to ch for updates and write to response
	for p := range ch {
		s := fmt.Sprintf("<style>#%s{background:#000}</style>\n", p)
		_, err := w.Write([]byte(s))
		if err != nil {
			return
		}
		flusher.Flush()
	}
}