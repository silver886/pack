// Package pack provides various single line function for external source
package pack

import (
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr"
	"github.com/sirupsen/logrus"
	llfile "leoliu.io/file"
	"leoliu.io/logger"
)

var (
	intLog    bool
	intLogger *logger.Logger
)

// SetLogger set internal logger for logging
func SetLogger(extLogger *logger.Logger) {
	intLogger = extLogger
	intLog = true
}

// ResetLogger reset internal logger
func ResetLogger() {
	intLogger = nil
	intLog = false
}

// Box add some attributes on packr.Box
type Box struct {
	*packr.Box

	Dest string
}

// Extract extract the file to the destination folder
func (box Box) Extract(file string) (path string, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"box":  box,
				"file": file,
			}),
		).Debugln("Extract file . . .")
	}

	path, err = box.ExtractToDir(box.Dest, file)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"box":            box,
				"path":           path,
				"internal_error": err,
			}),
		).Debugln("Extract file")
	}

	return
}

// ExtractToDir extract the file to the target folder
func (box Box) ExtractToDir(dest string, file string) (path string, err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"box":              box,
				"destination_path": dest,
				"file":             file,
			}),
		).Debugln("Extract file to directory . . .")
	}

	err = box.ExtractTo(dest+"/"+file, file)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot extract file")
		}
		return
	}
	path, _ = filepath.Abs(dest + "/" + file)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"box":            box,
				"path":           path,
				"internal_error": err,
			}),
		).Debugln("Extract file to directory")
	}

	return
}

// ExtractTo extract the file to the target path
func (box Box) ExtractTo(dest string, file string) (err error) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"box":              box,
				"destination_path": dest,
				"file":             file,
			}),
		).Debugln("Extract file to path . . .")
	}

	destPath, _ := filepath.Abs(dest)
	destDirPath := filepath.Dir(destPath)
	if !llfile.Exist(destDirPath) {
		if err := os.MkdirAll(destDirPath, 0755|os.ModeDir); err != nil {
			if intLog {
				intLogger.WithFields(
					logger.DebugInfo(1, logrus.Fields{
						"absolute_destination_directory_path": destDirPath,
						"internal_error":                      err,
					}),
				).Errorln("Cannot create destination directory")
			}
			return err
		}
	}

	fileByte, err := box.MustBytes(file)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot read file")
		}
		return
	}

	_, err = llfile.WriteByte(dest, fileByte)
	if err != nil {
		if intLog {
			intLogger.WithFields(
				logger.DebugInfo(1, logrus.Fields{
					"internal_error": err,
				}),
			).Errorln("Cannot write file")
		}
		return
	}

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"box":              box,
				"destination_path": destPath,
				"internal_error":   err,
			}),
		).Debugln("Extract file to path")
	}

	return
}

// Clear remove a packr.Box and content
func (box Box) Clear() (err error) {
	err = os.RemoveAll(box.Dest)

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"box":            box,
				"internal_error": err,
			}),
		).Debugln("Clear destination directory . . .")
	}

	return
}

// New create a packr.Box
func New(packrBox packr.Box, dest string) (box Box) {
	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"packr_box":        packrBox,
				"destination_path": dest,
			}),
		).Debugln("New box . . .")
	}

	absDest, _ := filepath.Abs(dest)
	if !llfile.Exist(absDest) {
		if err := os.MkdirAll(absDest, 0755|os.ModeDir); err != nil {
			if intLog {
				intLogger.WithFields(
					logger.DebugInfo(1, logrus.Fields{
						"absolute_destination_path": absDest,
						"internal_error":            err,
					}),
				).Errorln("Cannot create destination directory")
			}
			return
		}
	}
	box = Box{
		Box:  &packrBox,
		Dest: absDest,
	}

	if intLog {
		intLogger.WithFields(
			logger.DebugInfo(1, logrus.Fields{
				"box": box,
			}),
		).Debugln("New box")
	}

	return
}
