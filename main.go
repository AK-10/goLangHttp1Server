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
			conn.Write(append(header, buf...))
			conn.Close()
			// ヘッダとボディを分けないとブラウザに怒られる
			// どうやらhttp/0.9の仕様らしい
		} ()
	}
	
	// listen.Close()
}