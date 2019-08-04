package main

import (
	"flag"

	"./tools"
)

var flagSrc = flag.String("s", "./src", "源代码地址")
var flagDst = flag.String("d", "./dst", "打包地址")
var flagBranch = flag.String("b", "master", "标签所在分支")

func main() {
	flag.Parse()
	//src
	src, err := tools.NewSrc(*flagSrc, *flagBranch)
	if err != nil {
		tools.Out("src加载失败：" + err.Error())
		return
	}
	//dst
	dst, err := tools.NewDst(*flagDst)
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
	//apply it
	err = commit.Apply()
	if err != nil {
		tools.Out("发布失败：" + err.Error())
		return
	}
}
