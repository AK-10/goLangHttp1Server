package akhttp

import (
	"strconv"
	"time"
	"strings"
	"path/filepath"
	"io/ioutil"
	"log"
	"os"
)

type AKResponse struct {
	protocol string
	status int
	message string
	location *string
	body []byte
	contentType string
}


func NewResponse() *AKResponse {
	documentDir, _ := filepath.Abs("./views")
	path := requestHandleMap(req.path)
	println(path)
	absPath, err := filepath.Abs(documentDir + path)
	if err != nil {
		log.Fatal(err)
	}
	// ディレクトリトラバーサルの検査もする
	if !strings.HasPrefix(absPath, documentDir) {
		return notFoundError()
	}

	println(absPath)
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return internalServerErrorRes()
	}
	// ディレクトリであれば301
	if fileInfo.IsDir() {
		body, _ := getHTML(documentDir + "/index.html")
		location := "http://localhost:8080/index"
		return Response {
			protocol: req.protocol,
			status: 301,
			message: "Moved Permanently",
			location: &location,
			body: body,
		}
	} 
		
	// req.pathを元にファイルを読み取る
	body, err := getHTML(absPath)
	if err != nil {
		return notFoundError()
	}
	// status, messageはファイルの読み取り結果次第
	// locationはpath次第
	// bodyは読み取ったファイルのstring <- 実はbyteで良いかも
	// contentTypeはファイルの拡張子から
	return &Response{
		protocol: req.protocol,
		status: 200,
		message: "OK",
		location: nil,
		body: body,
		contentType: getContentType(filepath.Ext(absPath)),
	}
}

func (res *Response) ToByteArray() []byte {
	buf := make([]byte, 0)
	fstLine := res.protocol + " " + strconv.Itoa(res.status) + " " + res.message + "\r\n"
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
	// println(string(res.body))
	buf = append(buf, res.body...)

	return buf
}

func getExtension(path string) string {
	return filepath.Ext(path)
}


// 返す値をここに定義
func requestHandleMap(path string) string {
	switch path {
	case "/index":
		html := "/index.html"
		return html
	case "/form":
		html := "/form.html"
		return html
	default:
		return ""
	}
}

func getContentType(ext string) string {
	contentTypeMap := map[string]string{
		".html": "text/html",
		".htm": "text/html",
		".txt": "text/plain",
		".css": "text/css",
		".png": "image/png",
		".jpg": "image/jpeg",
		".jpeg": "image/jpeg",
		".gif": "image/gif",
	}

	ret := contentTypeMap[strings.ToLower(ext)]
	if ret == "" {
		return "application/octet-stream" 
	} else {
		return ret
	}
}


func getHTML(path string) ([]byte, error) {
	f, err := os.Open(path)
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	return buf, err
}


func notFoundError() Response {
	nfFilePath, _ := filepath.Abs("./views/errors/404.html")
	
	buf, _ := getHTML(nfFilePath)

	return Response {
		protocol: "HTTP/1.1",
		status: 404,
		message: "Not Found",
		location: nil,
		body: buf,
		contentType: getContentType(filepath.Ext(nfFilePath)),
	}	
}

func internalServerErrorRes() Response {
	iseFilePath, _ := filepath.Abs("./views/errors/500.html")
	
	buf, _ := getHTML(iseFilePath)

	return Response {
		protocol: "HTTP/1.1",
		status: 500,
		message: "Internal Server Error",
		location: nil,
		body: buf,
		contentType: getContentType(filepath.Ext(iseFilePath)),
	}
}