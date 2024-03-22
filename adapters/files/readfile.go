package files

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func ReadFile(directory, fileName string, bufferSize int) (hashes []string, err error) {
	fileSys := os.DirFS(directory)
	file, err := fileSys.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	buff := make([]byte, bufferSize)
	i := 0
	for {
		i++
		var chunk bytes.Buffer
		n, err := file.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			return hashes, err
		}
		chunk.Write(buff[:n])
		hash := sha256.Sum256(chunk.Bytes())
		hexHash := hex.EncodeToString(hash[:])
		hashes = append(hashes, hexHash)
	}

	return
}
