package main

import (
	"./service"
	"./akhttp"
	"./myutil"
)


type Messages []*myutil.Message

func main() {
	port := 8080

	messages := Messages{}
	s := service.New()
	
	s.Get("/messages", func(req *akhttp.AKRequest, res *akhttp.AKResponse) {
			out := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
		</head>
		<body>
		<h1>テスト掲示板</h1>
		<form action="/messages" method="post">
			ハンドル名: <input type="text" name="handle"></br>
			<textarea name="message" cols="60" rows="4"></textarea><br/>
			<input type="submit">
		</form>
		<hr/>
		<!-- messageList ↓ -->
		<ul>
		`

		for _, msg := range messages {
			out += "<li>" + msg.Handle + " : " + msg.Value + "</li>\n"
		}

		out += `
		</ul>
		</body>
		</html>
		`
		res.Render(out)
	})

	s.Post("/messages", func(req *akhttp.AKRequest, res *akhttp.AKResponse) {
		name := req.GetBody()["handle"]
		msg := req.GetBody()["message"]
		if name != "" && msg != "" {
			messages = append(messages, myutil.NewMessage(name, msg))
		}
		
		out := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
		</head>
		<body>
		<h1>テスト掲示板</h1>
		<form action="/messages" method="post">
			ハンドル名: <input type="text" name="handle"></br>
			<textarea name="message" cols="60" rows="4"></textarea><br/>
			<input type="submit">
		</form>
		<hr/>
		<!-- messageList ↓ -->
		<ul>
		`

		for _, msg := range messages {
			out += "<li>" + msg.Handle + " : " + msg.Value + "</li>\n"
		}

		out += `
		</ul>
		</body>
		</html>
		`
		res.Render(out)
	})

	for _, v := range s.GetHandle() {
		v.PrintPM()
	}
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
