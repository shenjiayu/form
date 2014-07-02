package form

import (
	"errors"
	"regexp"
)

func Validate(keyword, content string) error {
	switch keyword {
	case "username":
		reg, _ := regexp.Compile(`^(\w)+$`)
		if b := reg.MatchString(content); !b {
			return errors.New("invalidusername")
		}
	case "password":
		reg, _ := regexp.Compile(`^(\w)+$`)
		if b := reg.MatchString(content); !b {
			return errors.New("invalidpassword")
		}
	case "email":
		reg, _ := regexp.Compile(`^(\w)+\@(\w)+\.(com|cn|net|org)$`)
		if b := reg.MatchString(content); !b {
			return errors.New("invalidemail")
		}
	default:
		return errors.New("No keywords matched")
	}
	return nil
}
