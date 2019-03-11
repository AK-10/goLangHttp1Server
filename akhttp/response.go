package akhttp

import (
	"strconv"
	"time"
	"strings"
)

type Response struct {
	protocol string
	status int
	message string
	location *string
	body string
	contentType string
}


func MakeResponse(req Request) Response {
	// req.pathを元にファイルを読み取る
	// ディレクトリトラバーサルの検査もする
	// protocolはreq.protocolをそのまま
	// status, messageはファイルの読み取り結果次第
	// locationはpath次第
	// bodyは読み取ったファイルのstring <- 実はbyteで良いかも
	// contentTypeはファイルの拡張子から

	return Response{
		protocol: req.protocol,
		status: 200,
		message: "OK",
		location: nil,
		body: "a",
	}
}

func (res *Response) ToByteArray() []byte {
	buf := make([]byte, 0)
	fstLine := "HTTP/1.1 " + strconv.Itoa(res.status) + " " + res.message + "\r\n"
	buf = append(buf, []byte(fstLine)...)
	// 2006年1月2日15時4分5秒 フォーマットの例文

	buf = append(buf, []byte("Date: " + time.Now().Format("Mon, 2 Jan 2006 15:04:05 GMT") + "\r\n")...)
	buf = append(buf, []byte("Server: GolangServer\r\n")...)
	buf = append(buf, []byte("Connection: close\r\n")...)
	buf = append(buf, []byte("Content-Type:" + res.contentType + "\r\n")...)
	
	if res.status == 301 && &res.location != nil {
		buf = append(buf, []byte("Location: " + *(res.location) + "\r\n")...)
	}

	buf = append(buf, []byte("\r\n")...)

	buf = append(buf, []byte(res.body)...)

	return buf
}

func getContentType(ext string) string {
	contentTypeMap := map[string]string{
		"html": "text/html",
		"htm": "text/html",
		"txt": "text/plain",
		"css": "text/css",
		"png": "image/png",
		"jpg": "image/jpeg",
		"jpeg": "image/jpeg",
		"gif": "image/gif",
	}

	ret := contentTypeMap[strings.ToLower(ext)]
	if ret == "" {
		return "application/octet-stream" 
	} else {
		return ret
	}
}