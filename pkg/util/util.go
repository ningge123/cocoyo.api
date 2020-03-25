package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"time"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

func FileHash(file *os.File) (string, error) {
	hashOjb := sha1.New()

	_, err := io.Copy(hashOjb, file)

	if err != nil {
		return "", err
	}

	hash := hashOjb.Sum(nil)

	return hex.EncodeToString(hash), nil
}

func GetTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetAppAbsolutePath() string {
	path, _ := os.Getwd()

	return path
}

func MakeDir(path string) error  {
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(path, os.ModePerm)

			if err != nil {
				return err
			}

			return nil
		}
	}

	return nil
}