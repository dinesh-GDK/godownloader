package main

import "net/http"

type http_file struct {
	http_response *http.Response
	file_name     string
}

type http_file_list []http_file

func (e http_file_list) Len() int {
	return len(e)
}

func (e http_file_list) Less(i, j int) bool {
	return e[i].http_response.ContentLength < e[j].http_response.ContentLength
}

func (e http_file_list) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}
