package main

import (
	"fmt"
	"github.com/bluele/mecab-golang"
	"time"
)

var str = "すもももももももものうち"

func parse(m *mecab.MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(str)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	fmt.Println(tg.Parse(lt))
}

func parseToNode(m *mecab.MeCab) {
	tg, err := m.NewTagger()
	if err != nil {
		panic(err)
	}
	defer tg.Destroy()
	lt, err := m.NewLattice(str)
	if err != nil {
		panic(err)
	}
	defer lt.Destroy()

	node := tg.ParseToNode(lt)

	for {
		fmt.Println(fmt.Sprintf("%s %s", node.Surface(), node.Feature()))
		if node.Next() != nil {
			break
		}
	}
}

func main() {
	m, err := mecab.New("-Ochasen")
	if err != nil {
		panic(err)
	}
	defer m.Destroy()
	number := int(1e5)
	for i := 0; i < number; i++ {
		go parse(m)
		go parseToNode(m)
	}
	time.Sleep(time.Second)
}
