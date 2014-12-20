BINARY = tpb

all: $(BINARY)

$(BINARY): clean assets
	go build -ldflags "-X main.buildCommit `git rev-parse --short HEAD`" .

assets: clean
	go-bindata -o assets.go -prefix "assets/" assets/...

clean:
	rm -f $(BINARY) assets.go

deps:
	go get -u github.com/gorilla/mux
	go get -u github.com/blevesearch/bleve

run: $(BINARY)
	./$(BINARY)
