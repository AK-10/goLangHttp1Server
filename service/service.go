package service

import (
	"../myutil"
	"log"
	"net"
	"fmt"
	"errors"
	"../akhttp"
	"path/filepath"
)

type Handle struct {
	path string
	method string
	process func(interface{}, interface{}) // processでresponseを直接いじるような操作を書く.
}

type Service struct {
	handles []*Handle
}


func New() *Service {
	return &Service{
		handles: []*Handle{},
	}
}

func (s *Service) Get(path string, process func(req *akhttp.AKRequest, res *akhttp.AKRequest)) {
	h := &Handle{
		path: path,
		method: "GET",
		process: process,
	}
	s.handles = append(s.handles, h)
}

func (s *Service) Post(path string, process func(req *akhttp.AKRequest, res *akhttp.AKRequest)) {
	h := &Handle{
		path: path,
		method: "POST",
		process: process,
	}
	s.handles = append(s.handles, h)
}

func (s *Service) NotFound() *akhttp.AKResponse {
	return nil
}

func (s *Service) InternalServerError() *akhttp.AKResponse {
	return nil
}


func (s *Service) handling(req *akhttp.AKRequest, res *akhttp.AKRequest) error {
	flag := false
	for _, h := range s.handles {
		if h.method == req.method && h.path == req.path {
			h.process(req, res)
			flag := true
		}
	}
	if !flag {
		return errors.New("no handle")
	}
	return nil
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

			reqBuf := make([]byte, 1024 * 100) // 100KBのrequestBuffer
			_, err := conn.Read(reqBuf)
			req := akhttp.NewRequestFromBytes(reqBuf)
			res := akhttp.NewResponse()

			res.SetHttpVersion(req.GetHttpVersion())

			if err != nil {
				res = s.InternalServerError()
			} else {
				err = s.handling(&req, &res) // ここでのerrはnot foundを意味する
				if err != nil {
					res = s.NotFound()
				}


			// directory traversal検査 (http1serverより上にアクセスしているか検査)
			// repo := filepath.Abs("../views")
			// path := req.GetPath()
			// absPath, err := filepath.Abs(repo + path)
			// if err != nil {
			// 	res.InternalServerError()
			// }
			// if strings.HasPrefix(absPath, repo) {
			// 	res.BadRequest()
			// }

			// // 301検査
			// fileInfo, err := os.Stat(absPath)
			// if err != nil {
			// 	res.InternalServerError()
			// }
			// if fileInfo.IsDir() {
			// 	res.MovedParmanently()
			// }

			conn.Write(res.ToByteArray())
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