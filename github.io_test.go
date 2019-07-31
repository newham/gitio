package gitio

import (
	"fmt"
	"testing"
)

var ia = NewIoAPI(string(key))

func TestIoAPI_FindOwnerIo(t *testing.T) {
	r := ia.FindOwnerIo()
	println(r.Name)
	println(r.Id)
}

func TestIoAPI_List(t *testing.T) {
	fs := ia.List()
	for _, f := range fs {
		fmt.Printf("%s,%s,%d,%s\n", f.Name, f.Sha, f.Size, f.Type)
	}
}

func TestIoAPI_Add(t *testing.T) {
	println(ia.Add("2019-07-30-FFmpeg.md"))
}

func TestIoAPI_Update(t *testing.T) {
	fs := ia.List()
	println(ia.Update(fs[len(fs)-1]))
}

func TestIoAPI_Delete(t *testing.T) {
	fs := ia.List()
	println(ia.Delete(fs[len(fs)-1]))
}
