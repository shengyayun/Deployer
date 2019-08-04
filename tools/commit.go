package tools

import "os"

//Commit is a commit's id in git project
type Commit struct {
	id  string
	src *Src
	dst *Dst
}

//NewCommit return a Commit Pointer
func NewCommit(id string, src *Src, dst *Dst) *Commit {
	return &Commit{id, src, dst}
}

//Exist returns if this commit is done
func (commit *Commit) Exist() (bool, error) {
	_, err := os.Stat(commit.dst.folder + "/" + commit.id)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//Apply is move dist from src to dst, and create a symlink
func (commit *Commit) Apply() error {
	err := commit.src.Build()
	if err != nil {
		return err
	}
	err = os.Rename(commit.src.folder+"/dist", commit.dst.folder+"/"+commit.id)
	if err != nil {
		return err
	}
	return os.Symlink(commit.dst.folder+"/"+commit.id, commit.dst.folder+"/html")
}
