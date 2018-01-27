package controller

import (
	"net/url"
	"strconv"
)

func getURLParameterAsInt(rURL *url.URL, param string, defaultVal int) int {
	query := rURL.Query()
	val, _ := strconv.Atoi(query.Get(param))
	if val == 0 {
		val = defaultVal
	}

	return val
}
