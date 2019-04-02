package akhttp

import (
	"strconv"
	"time"
	"path/filepath"
	// "strings"
	// "io/ioutil"
	// "log"
	// "os"
)

type AKResponse struct {
	httpVersion string
	status int
	message string
	location string
	body []byte
	contentType string
}

func NewResponse() *AKResponse {
	return &AKResponse {}
}

// func NewResponse() *AKResponse {
// 	documentDir, _ := filepath.Abs("./views")
// 	path := requestHandleMap(req.path)
// 	println(path)
// 	absPath, err := filepath.Abs(documentDir + path)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// ディレクトリトラバーサルの検査もする
// 	if !strings.HasPrefix(absPath, documentDir) {
// 		return notFoundError()
// 	}

// 	println(absPath)
// 	fileInfo, err := os.Stat(absPath)
// 	if err != nil {
// 		return internalServerErrorRes()
// 	}
// 	// ディレクトリであれば301
// 	if fileInfo.IsDir() {
// 		body, _ := getHTML(documentDir + "/index.html")
// 		location := "http://localhost:8080/index"
// 		return &AKResponse {
// 			protocol: req.protocol,
// 			status: 301,
// 			message: "Moved Permanently",
// 			location: &location,
// 			body: body,
// 		}
// 	} 
		
// 	// req.pathを元にファイルを読み取る
// 	body, err := getHTML(absPath)
// 	if err != nil {
// 		return notFoundError()
// 	}
// 	// status, messageはファイルの読み取り結果次第
// 	// locationはpath次第
// 	// bodyは読み取ったファイルのstring <- 実はbyteで良いかも
// 	// contentTypeはファイルの拡張子から
// 	return &AKResponse{
// 		protocol: req.protocol,
// 		status: 200,
// 		message: "OK",
// 		location: nil,
// 		body: body,
// 		contentType: getContentType(filepath.Ext(absPath)),
// 	}
// }

func (res *AKResponse) BadRequest() {
	res.contentType = "text/html"
	res.status = 400
	res.message = "bad request"
	res.body = []byte("<h1>bad request</h1>")
}

// rendering html
func (res *AKResponse) Render(body string) {
	// basePath := "../views"
	res.contentType = "text/html"
	res.status = 200
	res.message = "OK"
	res.body = []byte(body)
}

func (res *AKResponse) InternalServerError() {
	res.contentType = "text/html"
	res.status = 500
	res.message = "Internal Server Error"
	res.body = []byte("<h1>Internal Server Error</h1>")
}

func (res *AKResponse) NotFound() {
	res.contentType = "text/html"
	res.status = 404
	res.message = "NotFound"
	res.body = []byte("<h1>Not Found</h1>")
}

// rendering json
func (res *AKResponse) JSON() {
	res.contentType = "application/json"
	res.status = 200
	res.message = "OK"
}

// func (res *AKResponse) SetOK() {
// 	res.status = 200
// 	res.message = "OK"
// }

func (res *AKResponse) SetHttpVersion(v string) {
	res.httpVersion = v
}

func (res *AKResponse) MovedPermanently() {
	res.contentType = "application/json"
	res.status = 301
	res.message = "OK"
	res.location = "http://localhost:8080/messages"
}


func (res *AKResponse) ToByteArray() []byte {
	buf := make([]byte, 0)
	fstLine := res.httpVersion + " " + strconv.Itoa(res.status) + " " + res.message + "\r\n"
	buf = append(buf, []byte(fstLine)...)
	// 2006年1月2日15時4分5秒 フォーマットの例文
	buf = append(buf, []byte("Date: " + time.Now().Format("Mon, 2 Jan 2006 15:04:05 GMT") + "\r\n")...)
	buf = append(buf, []byte("Server: GolangServer\r\n")...)
	buf = append(buf, []byte("Connection: close\r\n")...)
	buf = append(buf, []byte("Content-Type:" + res.contentType + "\r\n")...)
	
	if res.status == 301 && res.location != "" {
		buf = append(buf, []byte("Location: " + res.location + "\r\n")...)
	}

	buf = append(buf, []byte("\r\n")...)
	// println(string(res.body))
	buf = append(buf, res.body...)

	return buf
}

func getExtension(path string) string {
	return filepath.Ext(path)
}


// 返す値をここに定義
// func requestHandleMap(path string) string {
// 	switch path {
// 	case "/index":
// 		html := "/index.html"
// 		return html
// 	case "/form":
// 		html := "/form.html"
// 		return html
// 	default:
// 		return ""
// 	}
// }

// func getContentType(ext string) string {
// 	contentTypeMap := map[string]string{
// 		".html": "text/html",
// 		".htm": "text/html",
// 		".txt": "text/plain",
// 		".css": "text/css",
// 		".png": "image/png",
// 		".jpg": "image/jpeg",
// 		".jpeg": "image/jpeg",
// 		".gif": "image/gif",
// 	}

// 	ret := contentTypeMap[strings.ToLower(ext)]
// 	if ret == "" {
// 		return "application/octet-stream" 
// 	} else {
// 		return ret
// 	}
// }