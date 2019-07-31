package gitio

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

var key, _ = ioutil.ReadFile("token.key")

var api = NewAPI(string(key))

func TestAPI_QueryUrl(t *testing.T) {
	api.QueryUrl()
}

func TestAPI_QueryOwnerRepositories(t *testing.T) {
	r := api.QueryOwnerRepositories()
	println(r[0].Description)
	println(r[0].Name)
	println(r[0].FullName)
}

func TestAPI_QueryOwner(t *testing.T) {
	o := api.QueryOwner()
	println(o.Login)
	println(o.Email)
}

func TestNewAPI(t *testing.T) {
	api := NewAPI("")
	println(api.Temp.Owner.Login)
	println(api.Temp.Owner.Email)
}

func TestTemp_Read(t *testing.T) {
	println(api.Temp.Owner.Login)
	println(api.Temp.Owner.Email)
	for k, v := range api.Temp.UrlMap {
		println(k, v)
	}
}

func TestAPI_Push(t *testing.T) {
	file := "2019-07-30-FFmpeg操作.md"
	repo := api.Temp.FindRepo("newham.github.io")
	if &repo == nil {
		panic("repo not exist")
	}
	contentsUrl := strings.Replace(repo.ContentsUrl, "{+path}", "_posts", -1)
	url := fmt.Sprintf("%s/%s", contentsUrl, file)
	println(url)
	//api.Push(url, "master", "add new article", file)
}

func TestAPI_QueryPath(t *testing.T) {
	fs := api.QueryPath(api.Temp.FindRepo("newham.github.io").ContentsUrl, "_posts")
	for _, f := range fs {
		fmt.Printf("%s,%s,%d,%s\n", f.Name, f.Sha, f.Size, f.Type)
	}
}
