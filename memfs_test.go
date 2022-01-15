package memfs

import "testing"

func TestMemFS(t *testing.T) {
	mem := New()

	f, err := mem.Open("/path/to/file")
	if err != nil {
		t.Fatal(err)
	}

	f.Read(nil)
	f.Write([]byte("hello"))

	f.Close()
}
