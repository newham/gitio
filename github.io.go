package gitio

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	pathNameKey   = "{+path}"
	pathName      = "_posts"
	defaultBranch = "master"
	githubIoName  = ".github.io"
)

type IoAPI struct {
	*API
}

func NewIoAPI(token string) *IoAPI {
	api := NewAPI(token)
	if api == nil {
		return nil
	}
	return &IoAPI{API: api}
}

func (ia *IoAPI) Add(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	form := PushForm{Message: "add file by Gitio", Branch: defaultBranch, Content: Base64Encode(b)}
	return Error(ia.Push(actionAdd, ia.Url(pathName+"/"+f.Name()), form))
}

func Error(code int, b []byte) error {
	if code >= 200 && code < 300 { //ok
		return nil
	}
	return errors.New(string(b))

}

func (ia *IoAPI) Delete(f File) error {
	form := PushForm{Message: "delete file by Gitio", Sha: f.Sha, Branch: defaultBranch}
	return Error(ia.Push(actionDelete, ia.Url(f.Path), form))
}

func (ia *IoAPI) Update(f File) error {
	fl, err := os.Open(f.Name)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(fl)
	if err != nil {
		return err
	}
	form := PushForm{Message: "update file by Gitio", Sha: f.Sha, Branch: defaultBranch, Content: Base64Encode(b)}
	return Error(ia.Push(actionUpdate, ia.Url(f.Path), form))
}

func (ia *IoAPI) Url(path string) string {
	return strings.Replace(ia.FindOwnerIo().ContentsUrl, pathNameKey, path, -1)
}

func (ia *IoAPI) List() []File {
	fs := ia.QueryPath(ia.FindOwnerIo().ContentsUrl, pathName)
	if fs != nil {
		ia.Temp.Posts = fs
		ia.Temp.Save()
	} else {
		return ia.Temp.Posts
	}
	return fs
}

func (ia *IoAPI) PrintList() {
	list := ia.List()
	if list == nil {
		return
	}
	fmt.Printf("%-6s, %s\n", "sha", "name")
	for _, f := range list {
		fmt.Printf("%s, %s\n", f.Sha[:6], f.Name)
	}
}

func (ia *IoAPI) FindOwnerIo() Repository {
	name := ia.Temp.Owner.Login + githubIoName
	return ia.Temp.FindRepo(name)
}

func (ia *IoAPI) FindFileBySha(sha string) *File {
	fs := ia.List()
	for _, f := range fs {
		if strings.HasPrefix(f.Sha, sha) {
			return &f
		}
	}
	return nil
}
