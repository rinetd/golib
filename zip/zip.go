package zip

import (
	"io"
	"log"
	"os"

	"github.com/yeka/zip"
)

func NewZipWriter(name string) *zip.Writer {
	zipname, err := os.Create(name)
	if err != nil {
		log.Fatalln(err)
	}
	return zip.NewWriter(zipname)
}

func zippass(dst string, src ...string) {
	fzip, err := os.Create(`D:/待测试数据.zip`)
	if err != nil {
		log.Fatalln(err)
	}
	zipw := zip.NewWriter(fzip)
	defer zipw.Close()
	for _, n := range src {
		w, err := zipw.Encrypt(n, `hangruan2017`, zip.AES256Encryption)
		if err != nil {
			log.Fatal(err)
		}
		f, err := os.Open(n)
		if err != nil {
			return
		}

		_, err = io.Copy(w, f)
		if err != nil {
			log.Fatal(err)
		}
	}
	zipw.Flush()
}
