package gitio

import (
	"fmt"
	"os"
	"strings"
)

type Client interface {
	Run()
	Help(string)
	Args()
}

type CmdClient struct {
}

func (c *CmdClient) Run() {
	c.Args()
}

var cmd = map[string]string{
	"help": "Gitio is a tool to creat template of markdown file that can be push to github.io.\n\n" +
		"Usage:\n\n" +
		"\tgitio <command> [arguments]\n\n" +
		"The commands are:\n",
	"create": "create a new template, you can input the title et. of this file",
	"login":  "sign in your github account by token",
	"logout": "logout your github account",
	"list":   "list all of your markdown file in github",
	"add":    "add local markdown to remote",
	"delete": "delete remote markdown by sha",
	"update": "update remote markdown by sha, also check the file named if exist",
}

func (c *CmdClient) Help(key string) {
	key = strings.ToLower(key)
	if key == "" || key == "help" {
		println(cmd["help"])
		for k, v := range cmd {
			if k == "help" {
				continue
			}
			fmt.Printf("\t%-11s%s\n", k, v)
		}
	} else {
		fmt.Printf("\t%-11s%s\n", key, cmd[key])
	}
}

func (c *CmdClient) Args() {
	if len(os.Args) < 2 {
		c.Help("help")
		os.Exit(0)
	}
	switch os.Args[1] {
	case "create":
		println("Start to create a new template of markdown file of github.io.\n" +
			"Please input the content by step")
		a := NewDefaultArticle()
		total := len(a.HeadKeys())
		for i, key := range a.HeadKeys() {
			fmt.Printf("%d/%d %s:\n",
				i+1, total, key)
			input := ScanLine()
			if input != "" {
				a.SetHead(key, input)
			} else {
				fmt.Printf("Use \"%s\" as default\n", a.GetHead(key))
			}
		}
		a.Save()
		break
	case "login":
		println("Start to sign in github.com")
		token := ""
		for {
			println("input Token:")
			token = ScanLine()
			if token != "" {
				break
			}
		}
		api := NewIoAPI(token)
		if api == nil {
			println("login failed")
		} else {
			println("login success")
		}
		break
	case "logout":
		err := os.Remove(defaultTmpName)
		if err != nil {
			println(err.Error())
			println("logout failed")
		} else {
			println("logout success")
		}
		break
	case "list":
		api := NewIoAPI("")
		if api == nil {
			break
		}
		fs := api.List()
		println("name, sha, size(Byte), type")
		for _, f := range fs {
			fmt.Printf("%s, %s, %d, %s\n", f.Name, f.Sha, f.Size, f.Type)
		}
		break
	case "add":
		fileName := ""
		if len(os.Args) == 3 {
			fileName = os.Args[2]
		} else {
			for {
				println("input file's name you want to push:\nfileName:")
				fileName = ScanLine()
				if fileName != "" {
					break
				}
			}
		}
		api := NewIoAPI("")
		if api == nil {
			break
		}
		err := api.Add(fileName)
		if err != nil {
			println("add", fileName, "failed")
		} else {
			println("add", fileName, "success")
		}

		break
	case "delete":
		sha := ""
		if len(os.Args) == 3 {
			sha = os.Args[2]
		} else {
			for {
				println("input file's sha you want to delete:\nsha:")
				sha = ScanLine()
				if sha != "" {
					break
				}
			}
		}
		api := NewIoAPI("")
		if api == nil {
			break
		}
		f := api.FindFileBySha(sha)
		if f == nil {
			println("file's sha", sha, "not exist")
			break
		}
		err := api.Delete(*f)
		if err != nil {
			println("delete", f.Sha, "failed")
		} else {
			println("delete", f.Sha, "success")
		}
	case "update":
		sha := ""
		if len(os.Args) == 3 {
			sha = os.Args[2]
		} else {
			for {
				println("input file's sha you want to update:\nsha:")
				sha = ScanLine()
				if sha != "" {
					break
				}
			}
		}
		api := NewIoAPI("")
		if api == nil {
			break
		}
		f := api.FindFileBySha(sha)
		if f == nil {
			println("file's sha ", sha, " not exist")
			break
		}
		//check if file named exist
		if _, err := os.Stat(f.Name); err != nil {
			println("file named ", f.Name, " not exist")
			break
		}
		err := api.Update(*f)
		if err != nil {
			println("update ", f.Sha, " failed")
		} else {
			println("update ", f.Sha, " success")
		}
	default:
		c.Help("help")
		break
	}
}
