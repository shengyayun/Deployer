package tools

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//Src is a git project
type Src struct {
	folder string
	branch string
}

//NewSrc returns a Src pointer
func NewSrc(folder string, branch string) (*Src, error) {
	abs, err := filepath.Abs(folder)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(abs, 0777)
	if err != nil {
		return nil, err
	}
	return &Src{abs, branch}, nil
}

//Pull is git pull
func (src *Src) Pull() (string, error) {
	_, err := src.execute(fmt.Sprintf("git checkout %s", src.branch))
	if err != nil {
		return "", err
	}
	return src.execute("git pull")
}

//Latest is the lastest commit id of the branch
func (src *Src) Latest() (string, error) {
	return src.execute(fmt.Sprintf("git log %s --pretty=format:\"%%h\" -1", src.branch))
}

//Build is to build
func (src *Src) Build() error {
	_, err := src.execute("npm run build")
	return err
}

//run a shell
func (src *Src) execute(s string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", s)
	cmd.Dir = src.folder
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return strings.Trim(out.String(), "\r\n"), err
}
