package main

import (
	"./service"
	// "./akhttp"
)

func main() {
	port := 8080
	s := service.New()
	s.Start(port)
}

// func main() {
// 	port := ":8080"

// 	// ポート解放
// 	listen, err := net.Listen("tcp", port)
// 	if err != nil {
// 		log.Fatal("can not listen at", port)
// 	}
// 	fmt.Println("listening ", port)

// 	for {
// 		// 接続待ち
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			log.Fatal("can not established connection")
// 		}
// 		go func() {
// 			// コネクションが二回飛んでくるのは, htmlとfaviconを取得しようとするため
// 			// もし, javascript等の読み込みがあるなら, さらに増える
// 			println("connection established\n")

// 			// リクエストを読み込む
// 			reqBuf := make([]byte, 4096)
// 			_, err = conn.Read(reqBuf)
// 			if err != nil {
// 				// 400を返す処理
// 				// fmt.Println(reqBuf)
// 				fmt.Println(err)
// 				log.Fatal("can not read request header")
// 			}
			
// 			println(string(reqBuf))

// 			println(reqBuf)
			
// 			req := akhttp.MakeRequest(reqBuf)
// 			res := akhttp.MakeResponse(req)
// 			resBuf := res.ToByteArray()

// 			// println(string(resBuf))
// 			conn.Write(resBuf)
// 			conn.Close()
// 		} ()
// 	}

// 	// listen.Close()
// }
