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
	process func(interface{}, interface{}) // processでresponseを直接いじるような操作を書く.
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

func (s *Service) Get(path string, process func(interface{}, interface{})) {
	h := &Handle{
		path: path,
		method: "GET",
		process: process,
	}
	s.handles = append(s.handles, h)

}

func (s *Service) Post(path string, process func(interface{}, interface{})) {
	h := &Handle{
		path: path,
		method: "POST",
		process: process,
	}
	s.handles = append(s.handles, h)
}

func (s *Service) NotFound() interface{} {
	return nil
}

func (s *Service) InternalServerError() interface{} {
	return nil
}

func (s *Service) handling(req interface{}, res interface{}) interface{} {
	// parse request byte array

	// 
	reqMethod := "GET"
	reqPath := "/index"
	flag := false
	for _, h := range s.handles {
		if h.method == reqMethod && h.path == reqPath {
			h.process(req, res)
			flag := true
		}
	}
	if !flag {
		return s.NotFound()
	}
		return res // res.tobyte
}

func (s *Service) Start(port int) {
	listen, err := net.Listen("tcp", ":" + string(port))
	if err != nil {
		log.Fatal("can not listen at ", port)
	}

	fmt.Println("listening ", port)
	for {
		// success 3way hand shake.
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal("cant not established connection.")
		}
		go func() {
			println("connection established\n")

			reqBuf := make([]byte, 1024 * 100) // 100KB
			_, err := conn.Read(reqBuf)
			if err != nil {
				s.InternalServerError()
			}
			// ここでrequestとresponseとなるインスタンスを生成．

			s.handling(req, res)
			conn.Close()
		} ()
	}

	// listen.Close()
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