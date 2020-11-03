package helpers

import (
	"fmt"
	"net/url"
	"strings"
)

func ReplaceForRestParam(m *map[string]string) string {
	var params []string
	for k, v := range *m {
		var param string
		if k == "labels" {
			param = fmt.Sprintf("%s=%s", k, replaceLabelForRestParam(v))
		} else {
			param = fmt.Sprintf("%s=%s", k, url.QueryEscape(v))
		}
		params = append(params, param)
	}
	return strings.Join(params, "&")
}

func GetLabelsSeparator() string {
	return ","
}

func replaceLabelForRestParam(param string) string {
	labels := strings.Split(param, GetLabelsSeparator())
	var params []string
	for _, label := range labels {
		params = append(params, url.QueryEscape(label))
	}
	return strings.Join(params, GetLabelsSeparator())
}
