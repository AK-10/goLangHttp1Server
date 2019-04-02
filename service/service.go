package service

import (
	// "../myutil"
	"log"
	"net"
	"fmt"
	// "errors"
	"../akhttp"
	"path/filepath"
	"strings"
	"os"
	"strconv"
)

type Handle struct {
	path string
	method string
	process func(*akhttp.AKRequest, *akhttp.AKResponse) // processでresponseを直接いじるような操作を書く.
}

func (h *Handle) PrintPM() {
	println(h.path)
	println(h.method)
}

type Service struct {
	handles []*Handle
}


func New() *Service {
	return &Service{
		handles: []*Handle{},
	}
}

func (s *Service) Get(path string, process func(req *akhttp.AKRequest, res *akhttp.AKResponse)) {
	h := &Handle{
		path: path,
		method: "GET",
		process: process,
	}
	s.handles = append(s.handles, h)
}

func (s *Service) GetHandle() []*Handle {
	return s.handles
}

func (s *Service) Post(path string, process func(req *akhttp.AKRequest, res *akhttp.AKResponse)) {
	h := &Handle{
		path: path,
		method: "POST",
		process: process,
	}
	s.handles = append(s.handles, h)
}


func (s *Service) handling(req *akhttp.AKRequest, res *akhttp.AKResponse) {
	// flag := false
	for _, h := range s.handles {
		if req.EqualMethodAndPath(h.method, h.path) {
			h.process(req, res)
			// flag = true
			return
		}
	}
	res.NotFound()
	return 
	// if !flag {
	// 	res.NotFound()
	// 	return
	// }
	// directory traversal検査 (http1serverより上にアクセスしているか検査)
	repo, err := filepath.Abs("../views")
	path := req.GetPath()
	absPath, err := filepath.Abs(repo + path)
	if err != nil {
		res.InternalServerError()
		return
	}
	if strings.HasPrefix(absPath, repo) {
		res.BadRequest()
		return
	}

	// 301検査
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		res.InternalServerError()
		return
	}
	if fileInfo.IsDir() {
		res.MovedPermanently()
		return
	}
}

func (s *Service) Start(port int) {
	portString := ":" + strconv.Itoa(port)
	listen, err := net.Listen("tcp", portString)
	if err != nil {
		println("can not listen at ", port, err)
		log.Fatal(err)
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
			req.Print()
			res := akhttp.NewResponse()

			res.SetHttpVersion(req.GetHTTPVersion())

			if err != nil {
				res.InternalServerError()
			} else {
				s.handling(req, res) // ここでのerrはnot foundを意味する
			}
			conn.Write(res.ToByteArray())
			conn.Close()
		} ()
	}
	// listen.Close()
}
