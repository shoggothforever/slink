package controller

import (
	_ "errors"
	"math"
	"time"
)

/*
若用户没有输入自定义的短链接则通过base62编码方法生成短链接
*/
func GenShort(shortUrl string) string {
	ans := ""
	if shortUrl == "" {
		//未给定短链接, 通过当前时间的纳秒数生成新的短链接存储
		temp := time.Now().UnixNano() % int64(math.Pow(62, 6))
		ans := ""
		for {
			if temp == 0 {
				break
			}

			now := temp % 62
			if now >= 0 && now <= 25 { //generate A-Z
				ans = ans + string(65+now)
			} else if now >= 26 && now <= 51 { //generate a-z
				ans = ans + string(71+now)
			} else if now >= 52 && now <= 61 { //generate 0-9
				ans = ans + string(now-4)
			}
			temp /= 62
		}
		return "bit.do/" + ans
	} else {
		for _, s := range shortUrl {
			if (s <= 57 && s >= 48) || (s <= 90 && s >= 65) || (s <= 122 && s >= 97) {
				ans = ans + string(s)
			}

		}
		return "bit.do/" + ans
	}
}
