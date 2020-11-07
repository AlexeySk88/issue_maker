package helpers

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var imgRegExp = regexp.MustCompile(`\S+\.(jpg|jpeg|jpe|png|tif|tiff|bmp)`)
var uploadRegExp = regexp.MustCompile(`!\[\S+]\(/uploads/[a-z0-9]+/\S+\)`)

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

func FindAllImageLinks(arr string) []string {
	uploadStr := strings.Join(uploadRegExp.FindAllString(arr, -1), ",")
	images := imgRegExp.FindAllString(arr, -1)

	var imgResult []string
	for _, image := range images {
		if !strings.Contains(uploadStr, image) {
			imgResult = append(imgResult, image)
		}
	}
	return imgResult
}

func replaceLabelForRestParam(param string) string {
	labels := strings.Split(param, GetLabelsSeparator())
	var params []string
	for _, label := range labels {
		params = append(params, url.QueryEscape(label))
	}
	return strings.Join(params, GetLabelsSeparator())
}
