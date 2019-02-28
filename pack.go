package pack

import (
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
	sfile "github.com/silver886/file"
)

// Box add some attributes on packr.Box
type Box struct {
	*packr.Box

	Dest string
}

// Extract extract the file to the destination folder
func (box *Box) Extract(file string) (string, error) {
	return box.ExtractToDir(box.Dest, file)
}

// ExtractToDir extract the file to the target folder
func (box *Box) ExtractToDir(dest string, file string) (string, error) {
	if err := box.ExtractTo(dest+"/"+file, file); err != nil {
		return "", err
	} else if path, err := filepath.Abs(dest + "/" + file); err != nil {
		return "", err
	} else {
		return path, nil
	}
}

// ExtractTo extract the file to the target path
func (box *Box) ExtractTo(dest string, file string) error {
	if destPath, err := filepath.Abs(dest); err != nil {
		return err
	} else if !sfile.Exist(filepath.Dir(destPath)) {
		if err := os.MkdirAll(filepath.Dir(destPath), 0755|os.ModeDir); err != nil {
			return err
		}
	}

	if fileByte, err := box.Find(file); err != nil {
		return err
	} else if _, err = sfile.WriteByte(dest, fileByte); err != nil {
		return err
	}

	return nil
}

// Clear remove a packr.Box and content
func (box *Box) Clear() error {
	return os.RemoveAll(box.Dest)
}

// New create a packr.Box
func New(packrBox packr.Box, dest string) *Box {
	absDest, _ := filepath.Abs(dest)
	if !sfile.Exist(absDest) {
		if err := os.MkdirAll(absDest, 0755|os.ModeDir); err != nil {
			return nil
		}
	}
	return &Box{
		Box:  &packrBox,
		Dest: absDest,
	}
}
