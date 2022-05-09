package timeutils

import (
	"strconv"
	"strings"
)

type MyTime struct {
	Hour   int
	Minute int
	Second int
}

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
