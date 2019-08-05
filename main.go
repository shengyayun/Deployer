package main

import (
	"flag"
	"time"

	"./tools"
)

var flagSrc = *flag.String("s", "./src", "源代码路径")
var flagDst = *flag.String("d", "./dst", "打包路径")
var flagBranch = *flag.String("b", "master", "项目分支")
var flagInterval = *flag.Int64("i", 5, "更新频率")

func main() {
	for {
		pack()
		time.Sleep(time.Second * time.Duration(flagInterval))
	}
}

func pack() {
	flag.Parse()
	//src
	src, err := tools.NewSrc(flagSrc, flagBranch)
	if err != nil {
		tools.Out("src加载失败：" + err.Error())
		return
	}
	//dst
	dst, err := tools.NewDst(flagDst)
	if err != nil {
		tools.Out("dst加载失败：" + err.Error())
		return
	}
	//pull code
	_, err = src.Pull()
	if err != nil {
		tools.Out("src拉取失败：" + err.Error())
		return
	}
	//the lastest commit id
	id, err := src.Latest()
	if err != nil {
		tools.Out("commit获取失败：" + err.Error())
		return
	}
	commit := tools.NewCommit(id, src, dst)
	//if this exist is done
	exist, err := commit.Exist()
	if err != nil {
		tools.Out("commit加载失败：" + err.Error())
		return
	}
	if exist {
		return
	}
	tools.Out("正在打包...")
	//apply it
	err = commit.Apply()
	if err != nil {
		tools.Out("发布失败：" + err.Error())
		return
	}
	tools.Out("打包成功：" + id)
}
