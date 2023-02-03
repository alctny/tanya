package util

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Hash64Check(file string, sha string) (bool, error) {
	f, err := os.Open(file)
	if err != nil {
		return false, err
	}

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return false, err
	}

	return sha == fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Download(url string, file string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, os.ModePerm)
}

func HashSumFile(outfile string, file string, sha string) error {
	f, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write([]byte(file + " " + sha + "\n"))
	return err
}
