package timeutils

import "fmt"

func AddAllTime(s []string) string {
	sum := "00:00:00"

	for _, v := range s {
		sum = AddTime(sum, v)
	}

	return sum
}

func AddTime(a string, b string) string {
	t1 := ExtractTime(a)
	t2 := ExtractTime(b)

	sum := MyTime{
		Hour:   t1.Hour + t2.Hour,
		Minute: t1.Minute + t2.Minute,
		Second: t1.Second + t2.Second,
	}

	if sum.Second > 59 {
		sum.Second -= 60
		sum.Minute++
	}
	if sum.Minute > 59 {
		sum.Minute -= 60
		sum.Hour++
	}

	s := fmt.Sprintf("%02d:%02d:%02d", sum.Hour, sum.Minute, sum.Second)

	return s
}
