package poker_test

import (
	"io"
	poker "learn-go-with-tests"
	"testing"
)

func TestTapeWrite(t *testing.T) {
	file, clean := createTempFile(t, "12345")
	defer clean()

	tape := &poker.Tape{File: file}
	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContent, _ := io.ReadAll(file)

	got := string(newFileContent)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
