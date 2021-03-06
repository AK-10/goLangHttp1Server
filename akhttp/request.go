package akhttp

import (
	"strings"
	"net/url"
	"errors"
	"log"
)

type AKRequest struct {
	path string
	method string
	version string
	query map[string]string
	header map[string]string
	body map[string]string
}

func (req *AKRequest) Print() {
	println(req.path)
	println(req.method)
	println(req.version)
	for k, v := range req.query {
		println(k, v)
	}
	for k, v := range req.header {
		println(k, v)
	}
	for k, v := range req.header {
		println(k, v)
	}
}

func NewRequest() *AKRequest {
	return &AKRequest{}
}

func NewRequestFromBytes(bytes []byte) *AKRequest {
	requestLines := parseByteArray(bytes)
	// for _, line := range requestLines {
	// 	println(line)
	// 	println(len(line))
	// } 
	
	// 一行目は特殊
	method, uri, httpVersion := parseStartLine(requestLines[0])
	headerStrs, bodyStrs := sepHeaderBody(requestLines[1:])
	

	println(method)
	temp := strings.Split(uri, "?")

	// var path string
	// var queryStr string
	// var querys []string
	switch len(temp) {
	case 1:
		path := temp[0]
		return &AKRequest{
			path: path,
			method: method,
			version: httpVersion,
			header: makeMap(headerStrs, ":"),
			body: makeMap(bodyStrs, "="),
		}
	case 2:
		path, queryStr := temp[0], temp[1]
		querys := strings.Split(urlDecode(queryStr), "&") // 怪しい
		return &AKRequest{
			path: path,
			method: method,
			version: httpVersion,
			query: makeMap(querys, "="),
			header: makeMap(headerStrs, ":"),
			body: makeMap(bodyStrs, "="),
		}
	default:
		log.Fatal("invalid request. no uri.")
		// errors.New("invalid request. no uri.")
		return &AKRequest{}
	}
}


func (req *AKRequest) GetHTTPVersion() string {
	return req.version
}

func (req *AKRequest) GetPath() string {
	return req.path
}

func (req *AKRequest) EqualMethodAndPath(method string, path string) bool {
	return req.method == method && req.path == path 
}

func (req *AKRequest) GetBody() map[string]string {
	return req.body
}


func parseStartLine(str string) (string, string, string) {
	startLine := strings.Split(str, " ")
	method := startLine[0]
	uri := startLine[1]
	httpVersion := startLine[2]
	return method, uri, httpVersion
}

func sepHeaderBody(strs []string) ([]string, []string) {
	var emptyLineIdx int
	for i, str := range strs {
		if str == "\r" {
			emptyLineIdx = i
			break
		}
	}
	headerStrs := strs[:emptyLineIdx]
	bodyStr := strs[emptyLineIdx + 1:][0]

	return headerStrs, strings.Split(bodyStr, "&")
}

func makeMap(strs []string, sep string) map[string]string {
	m := map[string]string{}
	for _, str := range strs {
		k, v, err := parseKeyValue(str, sep)
		if err == nil {
			m[k] = v
		}
	}
	return m
}

func parseByteArray(bytes []byte) []string {
	strSlice := make([]string, 0)
	start := 0
	for i := 0; i < len(bytes) ; i++ {
		if bytes[i] == 10 { // \n(10)だったら次へ
			strSlice = append(strSlice, string(bytes[start:i]))
			// println(string(bytes[start:i]))
			start = i + 1
		} 
		if bytes[i] == 0 {
			strSlice = append(strSlice, string(bytes[start:i]))
			break
		}
	}

	return strSlice
}

func parseKeyValue(line string, with string) (string, string, error) {
	strs := strings.SplitN(line, with, 2)
	if len(strs) < 2 {
		return "", "", errors.New("can not split")
	}
	return strings.Trim(strs[0], " "), strings.Trim(strs[1], " "), nil
}

func urlDecode(query string) string {
	str, _ := url.QueryUnescape(query)
	println(str)
	return str
}