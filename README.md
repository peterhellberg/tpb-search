# TPB Search

[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](https://github.com/peterhellberg/tpb#license)

Locally index and search database dumps from
[The Pirate Bay](http://en.wikipedia.org/wiki/The_Pirate_Bay) using the
[bleve](https://github.com/couchbaselabs/bleve) text indexing library.

Based on the example applications
[beer-search](https://github.com/blevesearch/beer-search) and
[bleve-explorer](https://github.com/blevesearch/bleve-explorer)

## Requirements

You need to download a database dump in the following format:

```
1234567|Name of the first file|54321|0|1|f2b2c2e4a786d3924b8922454772d784118e6421
2345678|Name of the second file|9876|1|1|27a12d50782e1412bcdec126224da29afb23c293
```

You can probably find a database dump using [Google](http://lmgtfy.com/?q=%22thepiratebay-dump%22+%22.txt.gz%22+--bitsnoop)

## Compilation

```
git clone git@github.com:peterhellberg/tpb.git
cd tpb
make deps
make
```

## Usage

```
Usage of ./tpb:
  -batchSize=800: batch size for indexing
  -d="thepiratebay-dump-2014-09-18.txt": the tpb dump file to use
  -h="localhost": http listen host
  -index="tpb.index": index path
  -p=1337: http listen port
  -static="static/": path to the static content
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
