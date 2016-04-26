package gull

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/c2fo/gull/source/lib/common"
)

type Convert struct {
	DestinationDir string
}

func NewConvert(destinationDir string) (*Convert, error) {
	err := os.MkdirAll(destinationDir, 0755)
	if err != nil {
		return nil, err
	}
	return &Convert{
		DestinationDir: destinationDir,
	}, nil
}

func (c *Convert) ConvertDirectory(dirPath string) error {
	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		return err
	}
	return filepath.Walk(absPath, c.convertFileWalk)
}

func (c *Convert) ConvertFile(filePath string) error {
	if strings.Contains(filePath, common.DefaultGullDirectory) {
		return nil
	}
	migration, err := NewMigrationFromConfigFile(filePath)
	if err != nil {
		return err
	}
	name := GetMigrationNameFromConfigName(filePath)
	destPath := filepath.Join(c.DestinationDir, name)
	err = migration.WriteToFile(destPath)
	return err
}

func (c *Convert) convertFileWalk(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if !f.IsDir() {
		return c.ConvertFile(path)
	}
	return nil
}
