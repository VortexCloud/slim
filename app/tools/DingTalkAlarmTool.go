package tools

import (
	"fmt"
	"time"
)

var (
	accessToken string = ""
	secret      string = ""
)

func DingTalkAlarm() {
	milliSecondTimeStamp := time.Now().UnixNano() / 1e6
	string_to_sign := string(milliSecondTimeStamp) + "\n" + secret
	fmt.Println(milliSecondTimeStamp)
	fmt.Println(string_to_sign)
}
