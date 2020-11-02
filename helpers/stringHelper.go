package helpers

import "strings"

func ReplaceForRestParam(p string) string {
	return strings.ReplaceAll(p, " ", "%20")
}
