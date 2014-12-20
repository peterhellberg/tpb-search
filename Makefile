BINARY = tpb

all: $(BINARY)

$(BINARY): clean assets
	go build -ldflags "-X main.buildCommit `git rev-parse --short HEAD`" .

assets: clean
	go-bindata -o assets.go -prefix "assets/" assets/...

clean:
	rm -f $(BINARY) assets.go

run: $(BINARY)
	./$(BINARY)
