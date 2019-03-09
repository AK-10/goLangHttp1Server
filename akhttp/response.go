package akhttp

import (
	"strconv"
)

type response struct {
	protocol string
	status int
	message string
	others map[string]string
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