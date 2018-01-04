package server

import (
	"fmt"
	"github.com/youfu9527/go-jump1jump/request"
	"net/http"
	"os"
	"strconv"
)

func Serve(port string) (err error) {
	fmt.Println(os.Getwd())
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		http.ServeFile(resp, req, "./template/index.html")
	})
	http.HandleFunc("/score", Index)
	err = http.ListenAndServe(":"+port, nil)
	return
}

func Index(resp http.ResponseWriter, req *http.Request) {
	score, err := strconv.Atoi(req.PostFormValue("score"))
	if err != nil || score < 0 {
		resp.Write([]byte("分数错误"))
		return
	}

	session := req.PostFormValue("session")
	if session == "" || len(session) < 16 {
		resp.Write([]byte("sessionID错误"))
		return
	}
	err = request.Gogogo(int64(score), session)
	if err != nil {
		resp.Write([]byte(err.Error()))
	}
	resp.Write([]byte("y"))
}
