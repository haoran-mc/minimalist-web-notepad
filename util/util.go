package util

import (
	"io"
	"log/slog"
	"os"

	"math/rand"
)

const letters = "abcdefghijklmnopqrstuvwxyz"

func RandStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

const fireDir = "data"

func ReadFile(fileName string) []byte {
	f, err := os.Open(fireDir + "/" + fileName)
	if err != nil {
		slog.Error("open file", "error", err.Error())
		return []byte{}
	}
	data, err := io.ReadAll(f)
	if err != nil {
		slog.Error("read file", "error", err.Error())
		return []byte{}
	}
	return data
}

func WriteFile(fileName string, fileContent []byte) {
	err := os.WriteFile(fireDir+"/"+fileName, fileContent, 0666)
	if err != nil {
		slog.Error("write file", "error", err.Error())
	}
}

func DeleteFile(fileName string) {
	err := os.Remove(fireDir + "/" + fileName)
	if err != nil {
		slog.Error("delete file", "error", err.Error())
	}
}
