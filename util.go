package gitio

import (
	"bufio"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"os/user"
	"runtime"
	"time"
)

func NowTime() string {
	return time.Now().Format("2006-01-02 15:04:05 -0700")
}

func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func WriteFile(name string, data []byte) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("create %s failed : %s\n", name, err.Error())
		return
	}
	fmt.Printf("create %s success\n", name)
	_, err = f.Write(data)
	if err != nil {
		fmt.Printf("write data to %s failed : %s\n", name, err.Error())
	}
	f.Close()
}

func ScanLine() string {
	reader := bufio.NewReader(os.Stdin)
	b, _, _ := reader.ReadLine()
	return string(b)
}

func Base64Encode(input []byte) string {
	return base64.StdEncoding.EncodeToString(input)
}

func Base64Decode(encodeString string) []byte {
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		println(err.Error())
	}
	return decodeBytes
}

func Sha256(str string) string {
	//使用sha256哈希函数
	h := sha256.New()
	h.Write([]byte(str))
	sum := h.Sum(nil)

	//由于是十六进制表示，因此需要转换
	s := hex.EncodeToString(sum)
	return s
}

func Home() string {
	home := ""
	user, err := user.Current()
	if nil == err {
		home = user.HomeDir
	}
	return home
}

func PathSeparator() string {
	if runtime.GOOS == "windows" {
		return "\\"
	}
	return "/"
}
