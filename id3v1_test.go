package audiotag

import (
	"os"
	"testing"
)

func TestReadID3v1Tags(t *testing.T) {
	for _, name := range []string{
		"testdata/id3v1/sample_usascii_v1.mp3",
		"testdata/id3v1/sample_ms932_v1.mp3",
		"testdata/id3v1/sample_utf8_v1.mp3"} {
		doTest(name, 0, 30, t)
	}
	for _, name := range []string{
		"testdata/id3v1/sample_usascii_v1.1.mp3",
		"testdata/id3v1/sample_ms932_v1.1.mp3",
		"testdata/id3v1/sample_utf8_v1.1.mp3"} {
		doTest(name, 1, 28, t)
	}
}

func doTest(name string, track int, length int, t *testing.T) {
	f, _ := os.Open(name)
	metadata, _ := ReadID3v1Tags(f)
	if actual, total := metadata.Track(); actual != track || total != 0 {
		t.Errorf("Track number for %s is (%d, %d) where (%d, 0) is expected.", name, actual, total, track)
	}
	comment := metadata.Raw()["comment"].(string)
	if actual := len(comment); actual != length {
		t.Errorf("Comment length for %s is %d where %d is expected", name, actual, length)
	}
}
