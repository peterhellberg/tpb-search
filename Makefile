BINARY = tpb-search

all: $(BINARY)

$(BINARY): clean assets
	go build -tags leveldb -ldflags "-X main.buildCommit `git rev-parse --short HEAD`" -o $(BINARY) .

assets: clean
	go-bindata -o assets.go -prefix "assets/" assets/...

clean:
	rm -f $(BINARY) assets.go

deps:
	go get -u github.com/gorilla/mux
	go get -u github.com/blevesearch/bleve
	go get -u github.com/jmhodges/levigo

run: $(BINARY)
	./$(BINARY)
