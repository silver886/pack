// Package pack provides various single line function for external source
package pack

import (
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
	llfile "leoliu.io/file"
)

// Box add some attributes on packr.Box
type Box struct {
	*packr.Box

	Dest string
}

// Extract extract the file to the destination folder
func (box Box) Extract(file string) (path string, err error) {
	path, err = box.ExtractToDir(box.Dest, file)
	return
}

// ExtractToDir extract the file to the target folder
func (box Box) ExtractToDir(dest string, file string) (path string, err error) {
	err = box.ExtractTo(dest+"/"+file, file)
	if err != nil {
		return
	}
	path, _ = filepath.Abs(dest + "/" + file)
	return
}

// ExtractTo extract the file to the target path
func (box Box) ExtractTo(dest string, file string) (err error) {
	destPath, _ := filepath.Abs(dest)
	destDirPath := filepath.Dir(destPath)
	if !llfile.Exist(destDirPath) {
		os.MkdirAll(destDirPath, 0600)
	}

	fileByte, err := box.MustBytes(file)
	if err != nil {
		return
	}

	_, err = llfile.WriteByte(dest, fileByte)
	if err != nil {
		return
	}

	return
}

// Clear remove a packr.Box and content
func (box Box) Clear() (err error) {
	err = os.RemoveAll(box.Dest)
	return
}

// New create a packr.Box
func New(src string, dest string) (box Box) {
	absSrc, _ := filepath.Abs(src)
	packrBox := packr.NewBox(absSrc)
	absDest, _ := filepath.Abs(dest)
	if !llfile.Exist(absDest) {
		os.MkdirAll(absDest, 0600)
	}
	box = Box{
		Box:  &packrBox,
		Dest: absDest,
	}
	return
}
