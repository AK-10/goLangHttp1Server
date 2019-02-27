package main

import (
	"log"
	"net"
	"fmt"
	"strings"
	"./response"
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
		go func() {
			println("connection established\n")
			
			// リクエストを読み込む
			reqBuf := make([]byte, 1024)
			_, err = conn.Read(reqBuf)
			if err != nil {
				log.Fatal("can not read request header")
			}
			// fmt.Println(reqBuf)
			// リクエストをbyteからstringに変換
			requestLines := readLines(reqBuf)
			method, reqPath := strings.Split(requestLines[0], " ")[0], strings.Split(requestLines[0], " ")[1]

			buf, status, msg, loc := getResponseItem(method, reqPath)
			if status == 301 {
				buf = nil
			}
			// レスポンスヘッダを返す処理
			// (...)演算子は可変長引数に対し、可変長構造体を与える時につける

			header := response.MakeResponseHeader(status, msg, getUTCTime(), loc)
			// headerBuf := make([]byte, 0)
			// fstLine := "HTTP/1.1 " + strconv.Itoa(status) + " " + msg + "\r\n"
			// headerBuf = append(headerBuf, []byte(fstLine)...)
			// // 2006年1月2日15時4分5秒 フォーマットの例文
			// headerBuf = append(headerBuf, []byte("Date: "+ getUTCTime())...)
			// // headerBuf = append(headerBuf, []byte("Date: "+ time.Now().Format("Tue, 30 Jul 2013 17:47:09 GMT\n"))...)
			// headerBuf = append(headerBuf, []byte("Server: GolangServer\r\n")...)
			// headerBuf = append(headerBuf, []byte("Connection: close\r\n")...)
			// headerBuf = append(headerBuf, []byte("Content-Type: text/html\r\n")...)
			// headerBuf = append(headerBuf, []byte("\r\n")...)
			
			// レスポンスボディを返す
			conn.Write(append(header, buf...))
			conn.Close()
			// ヘッダとボディを分けないとブラウザに怒られる
			// どうやらhttp/0.9の仕様らしい
		} ()
	}
	
	// listen.Close()
}