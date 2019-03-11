package akhttp

import (
	"strings"
	"net/url"
)

type Request struct {
	path string
	method string
	protocol string
	// query map[string]string
	query map[string]string
	field map[string]string
}

func MakeRequest(buf []byte) Request {
	requestLines := readLines(buf)
	requestFstLine := strings.Split(requestLines[0], " ")
	method := requestFstLine[0]
	protocol := requestFstLine[2]
	decodedURLString := strings.Split(urlDecode(requestFstLine[1]), "?")
	reqPath := decodedURLString[0]
	query := map[string]string{}
	if len(decodedURLString) > 1 {
		querys := strings.Split(decodedURLString[1], "&")
		for i := 0; i < len(querys); i++ {
			key, value := parseKeyValue(querys[i], "=")
			query[key] = value
		}
	}

	field := map[string]string{}
	for _, word := range requestLines {
		println(word)
	}

	for i := 1; i < len(requestLines); i++ {
		key, value := parseKeyValue(requestLines[i], ": ")
		field[key] = value
	}
	
	return Request{
		path: reqPath, 
		method: method, 
		protocol: protocol,
		field: field,
		query: query,
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

func parseKeyValue(line string, with string) (string, string) {
	strs := strings.Split(line, with)
	if len(strs) < 2 {
		return "", "" // fieldの最後に改行があるのでそれの対処. もっとちゃんとやるべき
	}
	return strs[0], strs[1]
}


func urlDecode(query string) string {
	str, _ := url.QueryUnescape(query)
	println(str)
	return str

}