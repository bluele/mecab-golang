# mecab-golang

A Golang wrapper for [mecab](//mecab.googlecode.com/svn/trunk/mecab/doc/index.html).

## Dependencies

mecab-golang needs [mecab 0.99](//mecab.googlecode.com/svn/trunk/mecab/doc/index.html#install).

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
$ export CGO_CFLAGS="-I`mecab-config --inc-dir`"
$ go get github.com/bluele/mecab-golang
```

## Examples

```go
package main

import (
  "fmt"
  "github.com/bluele/mecab-golang"
  "strings"
)

func parseToNode(m *mecab.MeCab) {
  tg, err := m.NewTagger()
  if err != nil {
    panic(err)
  }
  defer tg.Destroy()
  lt, err := m.NewLattice("すもももももももものうち")
  if err != nil {
    panic(err)
  }
  defer lt.Destroy()

  node := tg.ParseToNode(lt)
  for {
    features := strings.Split(node.Feature(), ",")
    if features[0] == "名詞" {
      fmt.Println(fmt.Sprintf("%s %s", node.Surface(), node.Feature()))
    }
    if node.Next() != nil {
      break
    }
  }
}

func main() {
  m, err := mecab.New("-Owakati")
  if err != nil {
    panic(err)
  }
  defer m.Destroy()
  parseToNode(m)
}
```

# Author

**Jun Kimura**

* <http://github.com/bluele>
* <junkxdev@gmail.com>
