package memfs

import (
	"io"
	"io/fs"
	"time"
)

type FS interface {
	Open(string) (File, error)
}

type MemFS struct{}

var _ FS = (*MemFS)(nil)

// Open opens the named file.
//
// When Open returns an error, it should be of type *PathError
// with the Op field set to "open", the Path field set to name,
// and the Err field describing the problem.
//
// Open should reject attempts to open names that do not satisfy
// ValidPath(name), returning a *PathError with Err set to
// ErrInvalid or ErrNotExist.
func (mem *MemFS) Open(name string) (File, error) {
	panic("not implemented") // TODO: Implement
}

type File interface {
	io.Reader
	io.Writer
	io.Seeker
	io.ReaderAt
	io.WriterAt
	io.Closer
	fs.File
}

type MemFile struct{}

var _ File = (*MemFile)(nil)

func (mFile *MemFile) Stat() (fs.FileInfo, error) {
	panic("not implemented") // TODO: Implement
}

func (mFile *MemFile) Read(_ []byte) (int, error) {
	panic("not implemented") // TODO: Implement
}

func (mFile *MemFile) Close() error {
	panic("not implemented") // TODO: Implement
}

func (mFile *MemFile) Write(p []byte) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

func (mFile *MemFile) Seek(offset int64, whence int) (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (mFile *MemFile) ReadAt(p []byte, off int64) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

func (mFile *MemFile) WriteAt(p []byte, off int64) (n int, err error) {
	panic("not implemented") // TODO: Implement
}

type MemFileInfo struct{}

func (mFileInfo *MemFileInfo) Name() string {
	panic("not implemented") // TODO: Implement
}

func (mFileInfo *MemFileInfo) Size() int64 {
	panic("not implemented") // TODO: Implement
}

func (mFileInfo *MemFileInfo) Mode() fs.FileMode {
	panic("not implemented") // TODO: Implement
}

func (mFileInfo *MemFileInfo) ModTime() time.Time {
	panic("not implemented") // TODO: Implement
}

func (mFileInfo *MemFileInfo) IsDir() bool {
	panic("not implemented") // TODO: Implement
}

func (mFileInfo *MemFileInfo) Sys() interface{} {
	panic("not implemented") // TODO: Implement
}

func New() *MemFS {
	return &MemFS{}
}
