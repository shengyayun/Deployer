package tools

import (
	"os"
	"path/filepath"
)

//Dst is the dist folder
type Dst struct {
	folder string
}

//NewDst returns a Dst pointer
func NewDst(folder string) (*Dst, error) {
	abs, err := filepath.Abs(folder)
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(abs, 0777)
	if err != nil {
		return nil, err
	}
	return &Dst{abs}, nil
}
