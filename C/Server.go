package C

import (
	"fmt"
	"log"
	"net/http"
)

type Http struct {
	app  *http.ServeMux
	port int
}

func (h *Http) Get(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	h.app.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fn(w, r)
		}
	})
}

func (h *Http) Post(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	h.app.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			fn(w, r)
		}
	})
}

func (h *Http) Put(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	h.app.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			fn(w, r)
		}
	})
}

func (h *Http) Delete(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	h.app.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			fn(w, r)
		}
	})
}

func (h *Http) Start(i int, fn func()) {
	fn()
	h.port = i
	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), h.app)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (h *Http) ServeStatic() {
	fs := http.FileServer(http.Dir("static"))
	h.app.Handle("/static/", http.StripPrefix("/static/", fs))
}

func New() *Http {
	return &Http{
		app: http.NewServeMux(),
	}
}
