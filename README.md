# TPB Search

[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/tpb-search#license)

Locally index and search database dumps from
[The Open Bay](http://openbay.isohunt.to/) using the
[bleve](https://github.com/couchbaselabs/bleve) text indexing library.

Based on the example applications
[beer-search](https://github.com/blevesearch/beer-search) and
[bleve-explorer](https://github.com/blevesearch/bleve-explorer)

![Screenshot](http://assets.c7.se/skitch/TPB_Search-20141221-042424.png)

## Requirements

You need to download an [openbay-db-dump](http://openbay.isohunt.to/files/openbay-db-dump.torrent)

It should have the following format:

```
"Name of the first file"|54321|f2b2c2e4a786d3924b8922454772d784118e6421|8|music|0|0
"Name of the second file"|9876|27a12d50782e1412bcdec126224da29afb23c293|1|movies|0|0
```

## Compilation

```
git clone git@github.com:peterhellberg/tpb-search.git
cd tpb-search
make deps
make
```

## Usage

```
Usage of ./tpb-search:
  -b=800: batch size for indexing
  -d="torrents_mini.csv": the openbay-db-dump to use
  -h="localhost": http listen host
  -i="tpb.index": index path
  -p=1337: http listen port
```

### Building an index of the first 10000 rows in the dump:

```bash
$ GOMAXPROCS=4 make && ./tpb-search -l 10000
rm -f tpb-search assets.go
go-bindata -o assets.go -prefix "assets/" assets/...
go build -ldflags "-X main.buildCommit `git rev-parse --short HEAD`" -o tpb-search .
2014/12/21 04:45:18 Creating new index...
2014/12/21 04:45:18 Listening on http://localhost:1337
2014/12/21 04:45:18 Indexing...
2014/12/21 04:45:19 Indexed 1000 documents in 1.16s (average 1.16ms/doc)
2014/12/21 04:45:19 Indexed 2000 documents in 1.47s (average 0.73ms/doc)
2014/12/21 04:45:20 Indexed 3000 documents in 1.82s (average 0.61ms/doc)
2014/12/21 04:45:20 Indexed 4000 documents in 2.50s (average 0.63ms/doc)
2014/12/21 04:45:21 Indexed 5000 documents in 2.91s (average 0.58ms/doc)
2014/12/21 04:45:21 Indexed 6000 documents in 3.33s (average 0.56ms/doc)
2014/12/21 04:45:22 Indexed 7000 documents in 3.85s (average 0.55ms/doc)
2014/12/21 04:45:22 Indexed 8000 documents in 4.79s (average 0.60ms/doc)
2014/12/21 04:45:23 Indexed 9000 documents in 5.26s (average 0.58ms/doc)
2014/12/21 04:45:23 Indexed 10000 documents in 5.74s (average 0.57ms/doc)
2014/12/21 04:45:24 Finished indexing 10000 documents in 6.04s (average 0.60ms/doc)
2014/12/21 04:45:24 Still listening on http://localhost:1337
```

## Third party packages

 - [bleve](https://godoc.org/github.com/blevesearch/bleve)
 - [bleve/http](https://godoc.org/github.com/blevesearch/bleve/http)
 - [gorilla/mux](https://godoc.org/github.com/gorilla/mux)

## License

*Copyright (C) 2014 Peter Hellberg*

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
