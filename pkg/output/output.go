package output

import (
	"fmt"
	"github.com/josem/redfollow/pkg/redirect"
	"strconv"
)

func PrintRedirects(redirects []redirect.Step) {
	lastIndex := len(redirects) - 1

	for i, e := range redirects {
		fmt.Println(strconv.Itoa(e.Code) + " - " + e.Url)

		if i != lastIndex {
			fmt.Println("		↓")
		}
	}
}
