package main

import (
	"expvar"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/blevesearch/bleve"
	bleveHttp "github.com/blevesearch/bleve/http"
)

var (
	port = flag.Int("p", 1337, "http listen port")
	host = flag.String("h", "localhost", "http listen host")
	dump = flag.String("d", "thepiratebay-dump-2014-09-18.txt", "the tpb dump file to use")

	batchSize  = flag.Int("batchSize", 800, "batch size for indexing")
	indexPath  = flag.String("index", "tpb.index", "index path")
	staticPath = flag.String("static", "static/", "path to the static content")

	expvars = expvar.NewMap("metrics")

	bind string
)

func main() {
	flag.Parse()

	expvars.Set("indexes", bleveHttp.IndexStats())

	// open the index
	tpbIndex, err := bleve.Open(*indexPath)
	if err == bleve.ErrorIndexPathDoesNotExist {
		log.Printf("Creating new index...")

		// create a mapping
		indexMapping := buildIndexMapping()

		// use the mapping
		tpbIndex, err = bleve.New(*indexPath, indexMapping)
		if err != nil {
			log.Fatal(err)
		}

		// index data in the background
		go func() {
			err = indexTPB(tpbIndex)
			if err != nil {
				log.Fatal(err)
			}
		}()
	} else if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Opening existing index...")
	}

	// create a router to serve static files
	router := staticFileRouter()

	// add the API
	bleveHttp.RegisterIndexName("tpb", tpbIndex)

	listIndexesHandler := bleveHttp.NewListIndexesHandler()
	router.Handle("/api", listIndexesHandler).Methods("GET")

	getIndexHandler := bleveHttp.NewGetIndexHandler()
	getIndexHandler.IndexNameLookup = indexNameLookup
	router.Handle("/api/tpb", getIndexHandler).Methods("GET")

	searchHandler := bleveHttp.NewSearchHandler("tpb")
	router.Handle("/api/search", searchHandler).Methods("POST")

	listFieldsHandler := bleveHttp.NewListFieldsHandler("tpb")
	router.Handle("/api/fields", listFieldsHandler).Methods("GET")

	docCountHandler := bleveHttp.NewDocCountHandler("tpb")
	router.Handle("/api/count", docCountHandler).Methods("GET")

	docGetHandler := bleveHttp.NewDocGetHandler("tpb")
	docGetHandler.DocIDLookup = docIDLookup
	router.Handle("/api/tpb/{docID}", docGetHandler).Methods("GET")

	debugHandler := bleveHttp.NewDebugDocumentHandler("tpb")
	debugHandler.DocIDLookup = docIDLookup
	router.Handle("/api/tpb/{docID}/_debug", debugHandler).Methods("GET")

	aliasHandler := bleveHttp.NewAliasHandler()
	router.Handle("/api/_aliases", aliasHandler).Methods("POST")

	// start the HTTP server
	http.Handle("/", router)

	bind = fmt.Sprintf("%s:%d", *host, *port)

	log.Printf("Listening on http://%v", bind)
	log.Fatal(http.ListenAndServe(bind, nil))
}
