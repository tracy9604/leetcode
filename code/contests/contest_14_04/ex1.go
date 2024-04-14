package contest_14_04

import (
	"strconv"
)

func ValidTime(s string) string {
	tmpHour := ""
	if string(s[0]) == "?" && string(s[1]) == "?" {
		tmpHour += "11"
	} else {
		if string(s[0]) == "?" {
			val, _ := strconv.Atoi(string(s[1]))
			if val > 1 {
				tmpHour += "0" + string(s[1])
			} else {
				tmpHour += "1" + string(s[1])
			}
		} else if string(s[1]) == "?" {
			val, _ := strconv.Atoi(string(s[0]))
			if val >= 1 {
				tmpHour += string(s[0]) + "1"
			} else {
				tmpHour += string(s[0]) + "9"
			}
		}
	}
	if len(tmpHour) == 0 {
		tmpHour = s[0:2]
	}
	tmpHour += ":"
	tmpMinute := ""
	if string(s[3]) == "?" && string(s[4]) == "?" {
		tmpMinute += "59"
	} else {
		if string(s[3]) == "?" {
			tmpMinute += "5" + string(s[4])
		} else if string(s[4]) == "?" {
			tmpMinute += string(s[3]) + "9"
		}
	}
	if len(tmpMinute) == 0 {
		tmpMinute = s[3:5]
	}
	return tmpHour + tmpMinute
}
