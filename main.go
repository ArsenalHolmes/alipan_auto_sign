package main

import (
	"autoSign/config"
	"autoSign/platform"
	"strings"
)

func main() {
	pushPlusToken := config.ConfigInstance.PushPlusToken
	kkCookie := config.ConfigInstance.KKCookie

	if kkCookie != "" {
		kkCookieList := strings.Split(kkCookie, ",")
		kk := &platform.KK{}
		for _, value := range kkCookieList {
			kk.Run(pushPlusToken, value)
		}
	}
}
