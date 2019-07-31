package gitio

import (
	"errors"
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
	if code > 300 { //error
		println(code, string(b))
		return errors.New(string(b))
	}
	return nil
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
	return fs
}

func (ia *IoAPI) FindOwnerIo() Repository {
	name := ia.Temp.Owner.Login + githubIoName
	return ia.Temp.FindRepo(name)
}

func (ia *IoAPI) FindFileBySha(sha string) *File {
	fs := ia.List()
	for _, f := range fs {
		if f.Sha == sha {
			return &f
		}
	}
	return nil
}
