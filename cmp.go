package xtest

import (
	"bufio"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// CmpIfErr is the macro which'll take an error and the expected cases (old actual and new)
// In case of being errored, it'll cmp actual with old (to check if there are undesired changes)
// Else, it'll cmp actual with new (to check if the changes are correct)
func CmpIfErr(t *testing.T, err error, old, actual, new interface{}, msg string) {
	if err != nil {
		if diff := cmp.Diff(old, actual); diff != "" {
			t.Errorf("%s errored mismatch (-want +got): %s", msg, diff)
		}
		return
	}
	if diff := cmp.Diff(new, actual); diff != "" {
		t.Errorf("%s mismatch (-want +got): %s", msg, diff)
	}
}

// CmpWithGoldenFile performs cmp over the goldenFile and the got bytes given
// Note: avoid testdata/%s when giving the filename
func CmpWithGoldenFile(t *testing.T, got []byte, goldenFilename, msg string) {
	t.Helper()
	want := ReadGoldenFile(t, goldenFilename)
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("%s mismatch (-want +got): %s", msg, diff)
	}
}

func ReaderFromGoldenFile(t *testing.T, filename string) (*bufio.Reader, func() error) {
	t.Helper()
	f, err := os.Open("testdata/" + filename)
	if err != nil {
		t.Fatalf("Error trying to open the goldenFile: %s", err)
	}
	return bufio.NewReader(f), f.Close
}

func ReadGoldenFile(t *testing.T, filename string) []byte {
	t.Helper()
	bytes, err := ioutil.ReadFile("testdata/" + filename)
	if err != nil {
		t.Fatalf("Error trying to read the goldenFile: %s", err)
	}
	return bytes
}
