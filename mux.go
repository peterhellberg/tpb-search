package main

import (
	"mime"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func staticFileRouter() *mux.Router {
	r := mux.NewRouter()

	r.StrictSlash(true)

	fh := tpbFileHandler{}

	// static
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fh))

	// application pages
	appPages := []string{
		"/search",
		"/metrics",
		"/about",
	}

	for _, p := range appPages {
		r.PathPrefix(p).Handler(RewriteURL("index.html", fh))
	}

	r.Handle("/", http.RedirectHandler("/search/syntax/", 302))

	return r
}

type tpbFileHandler struct{}

func (fh tpbFileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path

	w.Header().Add("Content-Type", mime.TypeByExtension(filepath.Ext(name)))

	data, err := Asset(name)

	if err == nil {
		w.Write(data)
	}
}

func RewriteURL(to string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = to
		h.ServeHTTP(w, r)
	})
}

func muxVariableLookup(req *http.Request, name string) string {
	return mux.Vars(req)[name]
}

func docIDLookup(req *http.Request) string {
	return muxVariableLookup(req, "docID")
}

func indexNameLookup(req *http.Request) string {
	return "tpb"
}
