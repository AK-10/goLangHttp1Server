package service

import (
	"../myutil"
	"log"
	"net"
	"fmt"
)

type Messages []*myutil.Message

type Handle struct {
	path string
	method string
	process func()
}

type Service struct {
	msgs Messages
	handles []*Handle
}


func NewService() *Service {
	return &Service{
		msgs: Messages{},
		handles: []*Handle{},
	}
}

func (s *Service) Get(path string, process func()) {
	h := &Handle{
		path: path,
		method: "get",
		process: process,
	}
	s.handles = append(s.handles, h)

}

func (s *Service) Post(path string, process func()) {
	h := &Handle{
		path: path,
		method: "post",
		process: process,
	}
	s.handles = append(s.handles, h)
}

func (s *Service) handling([]byte) {
	// parse request byte array

	// 
	reqMethod := "get"
	reqPath := "/post"
	for _, h := range s.handles {
		if h.method == reqMethod && h.path == reqPath {
			h.process()
		}
	}
}

func (s *Service) Start(port int) {
	listen, err := net.Listen("tcp", ":" + string(port))
	if err != nil {
		log.Fatal("can not listen at ", port)
	}
	for {
		// success 3way hand shake.
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("cant not established connection.")
		}
		fmt.Println("listening ", port)
		go func() {
			println("connection established\n")

			reqBuf := make([]byte, 8196)
			_, err := conn.Read(reqBuf)
			if err != nil {
				s.handling(reqBuf)
			}
		} ()
	}

	listen.Close()
}

func (s *Service) DoGet() {
	// something
	// contenttype := "text/html;charset=UTF-8"
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
	<form action="/message" method="post">
		ハンドル名: <input type="text" name="handle"></br>
		<textarea name="message" cols="60" rows="4"></textarea><br/>
		<input type="submit">
	</form>
	<hr/>
	<!-- messageList ↓ -->
	<ul>
	`

	for _, msg := range s.msgs {
		out += "<li>" + msg.Handle + " : " + msg.Value + "</li>\n"
	}

	out += `
	</ul>
	</body>
	</html>
	`
}

// func (s *Service) DoPost(req AKRequest, res AKResponse) {
// 	// something
// 	req.setCharacterEncoding("UTF-8")
// 	handle := req.getParameter("handle")
// 	m := req.getParameter("message")
// 	message := myutil.NewMessage(handle, m)
// 	append(s.msgs, message)
// }