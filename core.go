package gitio

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

type ArticleHead struct {
	Layout      string
	Title       string
	Date        string
	Categories  string
	Description string
}

func (h ArticleHead) ToBytes() []byte {
	format := "%s : %s\n"
	t := reflect.TypeOf(h)
	v := reflect.ValueOf(h)
	buff := bytes.NewBuffer(nil)
	for i := 0; i < v.NumField(); i++ {
		name := strings.ToLower(t.Field(i).Name)
		value := v.Field(i).Interface()
		a := ""
		switch value.(type) {
		case []string:
			a = strings.Join(value.([]string), " ")
		case string:
			a = strings.ToLower(strings.Trim(value.(string), " "))
		}
		buff.WriteString(fmt.Sprintf(format, name, a))
	}
	return buff.Bytes()
}

func (a Article) HeadKeys() []string {
	var heads []string
	t := reflect.TypeOf(a.Head)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "Date" {
			continue
		}
		heads = append(heads, t.Field(i).Name)
	}
	return heads
}

func (a *Article) SetHead(key, value string) {
	h := &a.Head
	v := reflect.ValueOf(h)
	f := v.Elem().FieldByName(key)
	if f.CanSet() {
		f.SetString(value)
	}
}

func (a *Article) GetHead(key string) string {
	h := &a.Head
	v := reflect.ValueOf(h)
	return v.Elem().FieldByName(key).Interface().(string)
}

type Article struct {
	Head    ArticleHead
	Content string
}

func NewArticle(
	layout string,
	title string,
	categories string,
	description string, content string) *Article {
	return &Article{Head: ArticleHead{layout, title, NowTime(), categories, description}, Content: content}
}

const (
	defaultLayout      = "post"
	defaultTitle       = "null"
	defaultCategories  = "article"
	defaultDescription = "null"
	defaultContent     = "null"
)

func NewDefaultArticle() *Article {
	return &Article{Head: ArticleHead{defaultLayout, defaultTitle, NowTime(), defaultCategories, defaultDescription}, Content: defaultContent}
}

func (a Article) ToMD() []byte {
	buff := bytes.NewBuffer(nil)
	buff.WriteString("---\n")
	buff.Write(a.Head.ToBytes())
	buff.WriteString("---\n\n")
	buff.WriteString(a.Content)
	return buff.Bytes()
}

func (a Article) FileName() string {
	data := a.Head.Date[0:10]
	title := strings.Replace(strings.Trim(a.Head.Title, " "), " ", "-", -1)
	return strings.ToLower(fmt.Sprintf("%s-%s.md", data, title))
}

func (a Article) Save() {
	WriteFile(a.FileName(), a.ToMD())
}
