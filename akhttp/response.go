package akhttp

import (
	"strconv"
	"time"
)

type Response struct {
	protocol string
	status int
	message string
	location *string
	body string
}

func MakeResponseHeader(status int, message string, time string, location string) []byte {
	buf := make([]byte, 0)
	fstLine := "HTTP/1.1 " + strconv.Itoa(status) + " " + message + "\r\n"
	buf = append(buf, []byte(fstLine)...)
	// 2006年1月2日15時4分5秒 フォーマットの例文
	buf = append(buf, []byte("Date: " + time + "\r\n")...)
	buf = append(buf, []byte("Server: GolangServer\r\n")...)
	buf = append(buf, []byte("Connection: close\r\n")...)
	buf = append(buf, []byte("Content-Type: text/html\r\n")...)
	
	if status == 301 {
		buf = append(buf, []byte("Location: " + location + "\r\n")...)
	}

	buf = append(buf, []byte("\r\n")...)


	return buf
}


// func MakeResonse(req Request) Response {

// }

func (res *Response) ToByteArray() []byte {
	buf := make([]byte, 0)
	fstLine := "HTTP/1.1 " + strconv.Itoa(res.status) + " " + res.message + "\r\n"
	buf = append(buf, []byte(fstLine)...)
	// 2006年1月2日15時4分5秒 フォーマットの例文

	buf = append(buf, []byte("Date: " + time.Now().Format("Mon, 2 Jan 2006 15:04:05 GMT") + "\r\n")...)
	buf = append(buf, []byte("Server: GolangServer\r\n")...)
	buf = append(buf, []byte("Connection: close\r\n")...)
	buf = append(buf, []byte("Content-Type: text/html\r\n")...)
	
	if res.status == 301 && &res.location != nil {
		buf = append(buf, []byte("Location: " + *(res.location) + "\r\n")...)
	}

	buf = append(buf, []byte("\r\n")...)

	buf = append(buf, []byte(res.body)...)

	return buf
}