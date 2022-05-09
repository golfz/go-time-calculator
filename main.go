package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MyTime struct {
	Hour   int
	Minute int
	Second int
}

func main() {
	fmt.Println("Hello, world.")
}

//func AddTime(a string, b string) string {
//	var h1, h2, m1, m2, s1, s2 int
//
//	return t1 + t2
//}

func ExtractTime(s string) MyTime {
	s = strings.TrimSpace(s)
	ss := strings.Split(s, ":")
	t := MyTime{}

	if len(ss) >= 1 {
		i := len(ss) - 1
		t.Second, _ = strconv.Atoi(ss[i])
	}
	if len(ss) >= 2 {
		var i int
		if len(ss) == 2 {
			i = 0
		} else {
			i = len(ss) - 2
		}
		t.Minute, _ = strconv.Atoi(ss[i])
	}
	if len(ss) == 3 {
		i := 0
		t.Hour, _ = strconv.Atoi(ss[i])
	}

	return t
}
