package main

import (
	"flag"
	"time"

	"./tools"
)

var (
	flagSrc      string
	flagDst      string
	flagBranch   string
	flagInterval int64
)

func init() {
	flag.StringVar(&flagSrc, "s", "./src", "源代码路径")
	flag.StringVar(&flagDst, "d", "./dst", "打包路径")
	flag.StringVar(&flagBranch, "b", "master", "项目分支")
	flag.Int64Var(&flagInterval, "i", 5, "更新频率")
}

func main() {
	flag.Parse()
	for {
		pack()
		time.Sleep(time.Second * time.Duration(flagInterval))
	}
}

func pack() {
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
