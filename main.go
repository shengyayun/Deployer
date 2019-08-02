package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	for {
		do()
		time.Sleep(time.Second * 5)
	}
}

func do() {
	//工作目录
	wd, err := os.Getwd()
	if err != nil {
		out("工作目录获取失败：" + err.Error())
		return
	}
	//重置分支
	_, err = shell(wd+"/src", "/usr/bin/git checkout master")
	if err != nil {
		out("重置分支失败：" + err.Error())
		return
	}
	//拉取代码
	_, err = shell(wd+"/src", "/usr/bin/git pull")
	if err != nil {
		out("代码拉取失败：" + err.Error())
		return
	}
	//获取最新tag
	tag, err := shell(wd+"/src", "/usr/bin/git describe --abbrev=0 --tags")
	if err != nil {
		out("最新标签获取失败：" + err.Error())
		return
	}
	tag = strings.Trim(tag, "\r\n")
	//判断dist/${tag}是否存在
	_, err = os.Stat(wd + "/dist/" + tag)
	if err == nil {
		//存在则跳过处理
		return
	}
	out("获取到最新标签：" + tag)
	//切换标签
	_, err = shell(wd+"/src", fmt.Sprintf("/usr/bin/git checkout %s", tag))
	if err != nil {
		out("切换标签失败：" + err.Error())
		return
	}
	//项目打包
	_, err = shell(wd+"/src", fmt.Sprintf("/usr/local/bin/npm run build"))
	if err != nil {
		out("项目打包失败：" + err.Error())
		return
	}
	//dist移动
	err = os.Rename(wd+"/src/dist", wd+"/dist/"+tag)
	if err != nil {
		out("Dist移动失败：" + err.Error())
		return
	}
	//重新生成软链接
	shell(wd, fmt.Sprintf("/bin/ln -snf %s ./html", wd+"/dist/"+tag))
	out("打包成功：" + tag)
}

func out(msg string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " " + msg)
}

func shell(p string, s string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", s)
	cmd.Dir = p
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}
