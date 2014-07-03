package form

import (
	"regexp"
)

type jsonformat struct {
	ErrMsg string
}

func Validate(keyword, content string) string {
	switch keyword {
	case "username":
		reg, _ := regexp.Compile(`^(\w)+$`)
		if b := reg.MatchString(content); !b {
			return tojson("invalidusername")
		}
	case "password":
		reg, _ := regexp.Compile(`^(\w)+$`)
		if b := reg.MatchString(content); !b {
			return tojson("invalidpassword")
		}
	case "email":
		reg, _ := regexp.Compile(`^(\w)+\@(\w)+\.(com|cn|net|org)$`)
		if b := reg.MatchString(content); !b {
			return tojson("invalidemail")
		}
	default:
		return tojson("no keywords matched")
	}
	return ""
}
