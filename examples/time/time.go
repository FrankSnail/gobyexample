// Go offers extensive support for times and durations;
// here are some examples.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	//获取当前时间
	now := time.Now()
	p(now)
	//时区
	tz, _ := time.LoadLocation("Asia/Shanghai")

	//构造时间
	then := time.Date(
		20023, 11, 17, 20, 34, 58, 651387237, tz)
	p(then)

	// 时间的一些属性
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p("--------时间的格式化和解析--------")
	/*
	  go格式化字符不是用Ymd这些字符表示年月
	  而是用特定的数值，如2006、Mon、Jan等具体值表示
	  这些值都是固定的，不能使用其他的
	*/
	p(now.Format(time.RFC3339))
	p(now.Format("06-"))
	p(now.Format("03:04PM"))
	p(now.Format("Mon Jan _2 15:04:05 2006"))
	p(now.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// 时间比较
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 时间运算
	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	p(then.Add(diff))
	p(then.Add(-diff))
}
