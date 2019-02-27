package main

import (
	"os"
	"io/ioutil"
	"log"
	"time"
)

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

func getUTCTime() string {
	return time.Now().Format("Mon, 2 Jan 2006 15:04:05 GMT")
}

// return values are body, status, msg, string
func getResponseItem(method string, reqPath string) ([]byte, int, string, string) {
	path := "./views"
	var statusCode int
	var message string
	var location string
	switch {
	case method == "GET" && reqPath == "/":
		path = path + "/root.html"
		statusCode = 301
		message = "Moved Permanently"
		location = "http://localhost:8080/index"
	case method == "GET" && reqPath == "/index":
		path = path + "/index.html"
		statusCode = 200
		message = "OK"
	default:
		// ここで謎のエラーハンドル
		path = path + "/errors/404.html"
		statusCode = 404
		message = "Not Found"
	}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		log.Fatal("can not open ", path)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("error was occured in reading file")
	}

	return buf, statusCode, message, location
}