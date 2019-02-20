package main

import (
	"log"
	"net"
	"fmt"
	"os"
	"io/ioutil"
	"time"
	"./byteUtil"
	// "bufio"
	"strings"
)

func main() {
	port := ":8080"

	// ポート解放
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("can not listen at", port)
	}
	fmt.Println("listening ", port)

	for {
		// 接続待ち
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("can not established connection")
		}
		println("connection established\n")
		
		// リクエストを読み込む
		reqBuf := make([]byte, 1024)
		_, err = conn.Read(reqBuf)
		if err != nil {
			log.Fatal("can not read request header")
		}
		// fmt.Println(reqBuf)
		// リクエストをbyteからstringに変換
		requestLines := byteUtil.ReadLines(reqBuf)
		methodLine := strings.Split(requestLines[0], " ")

		documentRoot := "./views/"
		var path string

		if methodLine[0] == "GET" && methodLine[1] == "/index" {
			path = "index.html"
		} else if methodLine[0] == "GET" && methodLine[1] == "/" {
			path = "root.html"
		}

		// ファイルを開く
		f, err := os.Open(documentRoot + path)
		if err != nil {
			log.Fatal("can not open ", documentRoot + path)
		}
		defer f.Close()

		// ファイル読み込み
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal("error was occured in reading file")
		}
		

		// buf := make([]byte, 1024)
		// for {
		// 	n, err := f.Read(buf)
		// 	if n == 0 {
		// 		break
		// 	}
		// 	if err != nil {
		// 		log.Fatal("error was occured in reading file")
		// 	}
		// }

		// レスポンスヘッダを返す処理
		// (...)演算子は可変長引数に対し、可変長構造体を与える時につける
		headerBuf := make([]byte, 0)
		headerBuf = append(headerBuf, []byte("HTTP/1.1 200 OK\n")...)
		// 2006年1月2日15時4分5秒 フォーマットの例文
		headerBuf = append(headerBuf, []byte("Date: "+ time.Now().Format("Mon, 2 Jan 2006 15:04:05 GMT\n"))...)
		// headerBuf = append(headerBuf, []byte("Date: "+ time.Now().Format("Tue, 30 Jul 2013 17:47:09 GMT\n"))...)
		headerBuf = append(headerBuf, []byte("Server: GolangServer\n")...)
		headerBuf = append(headerBuf, []byte("Connection: close\n")...)
		headerBuf = append(headerBuf, []byte("Content-Type: text/html\n")...)
		headerBuf = append(headerBuf, []byte("\n")...)
		
		// レスポンスボディを返す
		conn.Write(append(headerBuf, buf...))

		// ヘッダとボディを分けないとブラウザに怒られる
		// どうやらhttp/0.9の仕様らしい
	}
	
	// listen.Close()
}