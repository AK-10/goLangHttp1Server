package main

import (
	"log"
	"net"
	"fmt"
	"os"
	"io/ioutil"
	"time"
	// "strings"
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

		documentRoot := "./views/"
		index := "index.html"

		// ファイルを開く
		f, err := os.Open(documentRoot + index)
		if err != nil {
			log.Fatal("can not open ", documentRoot + index)
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
		headerBuf := make([]byte, 1024)
		headerBuf = append(headerBuf, []byte("HTTP/1.1 200 OK\n")...)
		headerBuf = append(headerBuf, []byte("Date: "+ time.Now().Format("Tue, 30 Jul 2013 17:47:09 GMT\n"))...)
		headerBuf = append(headerBuf, []byte("Server: GolangServer\n")...)
		headerBuf = append(headerBuf, []byte("Connection: close\n")...)
		headerBuf = append(headerBuf, []byte("Content-Type: text/html\n")...)
		headerBuf = append(headerBuf, []byte("\n")...)
		
		// レスポンスボディを返す
		conn.Write(append(headerBuf, buf...))
	}
	
	// listen.Close()
}