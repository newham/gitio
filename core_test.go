package gitio

import (
	"fmt"
	"testing"
)

var a = NewArticle("default", " Generating a new SSH key", "ssh linux", "How to Generating a new SSH", "How to Generating a new SSH")

func TestArticle_ToMD(t *testing.T) {

	fmt.Printf("%s\n", a.ToMD())
}

func TestArticle_FileName(t *testing.T) {
	fmt.Printf("filename:%s\n", a.FileName())
}

func TestArticle_Save(t *testing.T) {
	a.Save()
}

func TestArticleHead_Keys(t *testing.T) {
	fmt.Printf("%s", a.HeadKeys())
}

func TestArticleHead_Set(t *testing.T) {
	println(a.Head.Layout)
	a.SetHead("Layout", "dark")
	println(a.Head.Layout)
}
