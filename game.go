package main

import (
	"flag"
	"github.com/youfu9527/go-jump1jump/request"
	"github.com/youfu9527/go-jump1jump/server"
	"log"
	"strconv"
)

func main() {
	var (
		score   int64
		session string
		port    int
	)
	flag.Int64Var(&score, "c", 0, "输入您要定制的分数")
	flag.StringVar(&session, "s", "", "输入Sessionid")
	flag.IntVar(&port, "p", -1, "服务器地址,不输入则不启用,如 8888")
	flag.Parse()
	//如果使用-p参数，则不会执行单机刷分操作
	if port != -1 {
		err := server.Serve(strconv.Itoa(port))
		log.Fatalf("[Debug] 出错啦 %s", err.Error())
		return
	}
	//单机执行刷分操作
	if session == "" {
		log.Println(`[Debug] Sessionid 不能为空`)
		return
	}
	if len(session) < 16 {
		log.Println(`[Debug] Sessionid 长度错误`)
		return
	}
	if score < 0 {
		log.Println(`[Debug] Score 不能为负数`)
		return
	}
	request.Gogogo(score, session)
}
