package diff

import (
	"testing"
)

func TestCompareFiles(t *testing.T) {
	path := "./../../storage3/"
	testfile1 := "testfile1.txt"
	testfile2 := "testfile2.txt"
	testfile3 := "testfile3.txt"
	testfile4 := "testfile4.txt"

	// test equal files
	equal, err := CompareFiles(path, testfile2, path, testfile2)
	if err != nil {
		t.Error(err.Error())
	}
	if !equal {
		expected := true
		t.Errorf("expected equeal = %t but got %t", expected, equal)
	}
	// test different files
	equal, err = CompareFiles(path, testfile2, path, testfile3)
	if err != nil {
		t.Error(err.Error())
	}
	if equal {
		expected := false
		t.Errorf("expected equeal = %t but got %t", expected, equal)
	}
	// test different with empty file
	equal, err = CompareFiles(path, testfile1, path, testfile2)
	if err != nil {
		t.Error(err.Error())
	}
	if equal {
		expected := false
		t.Errorf("expected equeal = %t but got %t", expected, equal)
	}
	// test equal empty files
	equal, err = CompareFiles(path, testfile1, path, testfile1)
	if err != nil {
		t.Error(err.Error())
	}
	if !equal {
		expected := true
		t.Errorf("expected equeal = %t but got %t", expected, equal)
	}
	// test chunk removed
	equal, err = CompareFiles(path, testfile3, path, testfile4)
	if err != nil {
		t.Error(err.Error())
	}
	if equal {
		expected := false
		t.Errorf("expected equeal = %t but got %t", expected, equal)
	}
}
