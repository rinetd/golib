package golib

import (
	"os"
	"testing"
)

func Test_IsNotExist(t *testing.T) {
	if _, err := os.Stat("file_test.go"); os.IsNotExist(err) {
		t.Logf("file not exits") // file not found
	} else if err == nil {
		t.Logf("no errror") // file not found
	} else {
		t.Log("file exist")
	}

	// 不存在
	if _, err := os.Stat("file_noexist.go"); os.IsNotExist(err) {
		t.Logf("file not exits") // file not found
	} else if err == nil {
		t.Logf("no errror") // file not found
	} else {
		t.Log("file exist")
	}

	if _, err := os.Stat("file_test.go"); os.IsExist(err) {
		t.Logf("file is exits") // file not found
	} else if err == nil {
		t.Logf("no errror") // file not found
	} else {
		t.Log("file not exist")
	}

}
func TestFS_IsExits(t *testing.T) {
	path := "../src"
	if IsExist(path) {
		t.Logf("path : [%s] exits", path)
	} else {
		t.Fatalf("path [%s] not exits", path)
	}
}

func TestFS_Abs(t *testing.T) {
	path := "../src"
	fullpath, err := Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(fullpath)
}

func TestFS_MkdirAll(t *testing.T) {
	path := "./tmp/123/测试"
	err := MkdirAll(path)
	if err != nil {
		t.Fatal(err)
	}
}
