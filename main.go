package main

import "fmt"

type stringList []string

func (sl stringList) String() {
	for _, v := range sl {
		fmt.Println(v)
	}
}

func newStringList() stringList {
	return stringList{"1", "2"}
}

func main() {
	sl := newStringList()
	sl = append(sl, "3")
	sl.String()
}
