package akhttp

import (
	"strings"
	"net/url"
)

type Request struct {
	path string
	method string
	query map[string]string
	field map[string]string
}

func MakeRequest(buf []byte) Request {
	requestLines := readLines(buf)
	requestFstLine := strings.Split(requestLines[0], " ")
	method, reqPath := requestFstLine[0], requestFstLine[1]
	field := map[string]string{}
	for i := 1; i <= len(requestLines); i++ {
		key, value := parseField(requestLines[i])
		field[key] = value
	}
	
	return Request{
		path: reqPath, 
		method: method, 
		field: field,
	}
}


func readLines(bytes []byte) []string {
	strSlice := make([]string, 0)
	start := 0
	for i := 0; bytes[i] != 0 ; i++ {
		if bytes[i] == 10 {
			// println(string(bytes[start:i]))
			strSlice = append(strSlice, string(bytes[start:i]))
			start = i + 1
		}
	}
	return strSlice
}

func parseField(line string) (string, string) {
	strs := strings.Split(line, ": ")
	return strs[0], strs[1]
}


func urlDecode(query string) string {
	str, _ = url.QueryUnescape(query)
	

}