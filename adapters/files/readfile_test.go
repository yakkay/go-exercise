package files

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	// test empty file
	directory := "./../../storage3/"
	hashes, err := ReadFile(directory, "testfile1.txt", 100)
	if err != nil {
		t.Error(err.Error())
	}
	for _, h := range hashes {
		expected := ""
		if h != expected {
			t.Errorf("expected empty hash but got %s", h)
		}
	}
	// test non empty file
	hashes, err = ReadFile(directory, "testfile2.txt", 100)
	if err != nil {
		t.Error(err.Error())
	}
	for _, h := range hashes {
		expected := "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
		if h != expected {
			t.Errorf("expected %s hash but got %s", expected, h)
		}
	}
	// test non empty file with multiple hashes
	hashes, err = ReadFile(directory, "testfile3.txt", 2048)
	if err != nil {
		t.Error(err.Error())
	}
	for i, h := range hashes {
		var expected [4]string
		expected[0] = "afd8fbb0f023a34e79cf7f811cbbd39a63d5779092dd861dde62c2c9768ab413"
		expected[1] = "e1c4b909c433c8a563572b0ad1ee7fb9504854f9f5bad76140d4a50dab81af60"
		expected[2] = "0428d7fb33190e910cf20bba8c54c751ba1d162736d5b8a73145fd42bb8b7585"
		expected[3] = "883683b8707e996e6bb759ec6c884f4605ac1c77ce5067a8868f594b68b4edb4"
		if h != expected[i] {
			t.Errorf("expected %s hash but got %s", expected[i], h)
		}
	}
	// test non empty file with different chunk size
	hashes, err = ReadFile(directory, "testfile3.txt", 4096)
	if err != nil {
		t.Error(err.Error())
	}
	for i, h := range hashes {
		var expected [4]string

		expected[0] = "e164fe182d324d8ff84f337a8baa08fdf1e3888a9d4deec877d1989f28e6bfe1"
		expected[1] = "20e9507887920e4bb2b47a36c083d4a619d269d98429ef6b41b3acdb3a64969b"
		if h != expected[i] {
			t.Errorf("expected %s hash but got %s", expected[i], h)
		}
	}

}
