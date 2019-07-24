package zip

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/yeka/zip"
)

func TestNewZipWriter(t *testing.T) {
	zw := NewZipWriter("test.zip")
	defer zw.Close()
	{
		//  zip之后的文件名
		w, err := zw.Encrypt("www", "passwd", zip.AES256Encryption)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Open("zip.go")
		if err != nil {
			return
		}

		_, err = io.Copy(w, f)
		if err != nil {
			log.Fatal(err)
		}
	}

	zw.Flush()

}
