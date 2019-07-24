package golib

import (
	"bufio"
	"crypto"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)


type Storage struct {
	Schema string
}

var fileStorage = Storage{
	Schema: "file://",
}


//IsExist simple func check file/directory
// 只有不存在时返回false ,存在,无权限都返回true
// func IsExist(path string) bool {
// 	_, err := os.Stat(path)
// 	return !os.IsNotExist(err)
// }

// if stat, err := os.Stat(path); err == nil && stat.IsDir() {
//     // path is a directory
// }


// FileIsExist 检测文件是否存在
// 只有文件不存在时返回  return false
// err == nil | 文件存在 |文件存在但无权限 return true
func FileIsExist(dir string) bool {
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// FileIsExist 效率稍微高一些
// func FileIsExist(file string) bool {
// 	if _, err := os.Stat(file); err != nil {
// 		if os.IsNotExist(err) {
// 			return false
// 		}
// 	}
// 	return true
// }
func Abs(path string) (string, error) {
	fullPath, err := filepath.Abs(path)
	return fileStorage.Schema + ToSlash(fullPath), err
}

func ToSlash(path string) string {
	return filepath.ToSlash(filepath.Clean(path))
}

func MkdirAll(path string) error {
	if FileIsExist(path) != true {
		return os.MkdirAll(path, 0777)
	}
	return nil
}
func GetUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("Get current user got an error: %#v.", err)
	}
	return usr.HomeDir, nil
}
func writeToFile(filePath string, data interface{}) error {
	if strings.HasPrefix(filePath, "~") {
		home, err := GetUserHomeDir()
		if err != nil {
			return err
		}
		if home != "" {
			filePath = strings.Replace(filePath, "~", home, 1)
		}
	}

	os.Remove(filePath)

	var out string
	switch data.(type) {
	case string:
		out = data.(string)
		break
	case nil:
		return nil
	default:
		bs, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return fmt.Errorf("MarshalIndent data %#v got an error: %#v", data, err)
		}
		out = string(bs)
	}

	ioutil.WriteFile(filePath, []byte(out), 422)
	return nil
}

//Md5File 计算md5
func Md5File(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)

	hash := crypto.MD5.New()
	_, err = io.Copy(hash, r)
	if err != nil {
		return "", err
	}
	out := hex.EncodeToString(hash.Sum(nil))
	return out, nil
}
