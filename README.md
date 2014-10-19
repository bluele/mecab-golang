# mecab-golang

A Golang wrapper for [mecab](//mecab.googlecode.com/svn/trunk/mecab/doc/index.html).

## Dependencies

mecab-golang needs [mecab](//mecab.googlecode.com/svn/trunk/mecab/doc/index.html#install). 

Please install it.

## Getting started

Do the following before trying to `go get github.com/bluele/mecab-golang`.

```
$ export CGO_LDFLAGS="-L/path/to/lib -lmecab -lstdc++"

$ export CGO_CFLAGS="-I/path/to/include"

$ go get github.com/bluele/mecab-golang
```

If you installed `mecab-config` (check `which mecab-config`), do the following commands.

```
$ export CGO_LDFLAGS="`mecab-config --libs`"

$ export CGO_FLAGS="`mecab-config --inc-dir`"

$ go get github.com/bluele/mecab-golang
```

## Examples

See [examples](//github.com/bluele/mecab-golang/blob/master/examples).