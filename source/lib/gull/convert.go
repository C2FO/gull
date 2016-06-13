package gull

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/c2fo/gull/source/lib/common"
)

type Convert struct {
	DestinationDir        string
	FileNameIsEnvironment bool
	JsonEncode            bool
	Logger                ILogger
}

func NewConvert(destinationDir string, fileNameIsEnvironment bool, jsonEncode bool, logger ILogger) (*Convert, error) {
	err := os.MkdirAll(destinationDir, 0755)
	if err != nil {
		return nil, err
	}
	return &Convert{
		DestinationDir:        destinationDir,
		FileNameIsEnvironment: fileNameIsEnvironment,
		JsonEncode:            jsonEncode,
		Logger:                logger,
	}, nil
}

func (c *Convert) ConvertDirectory(dirPath string) error {
	return filepath.Walk(dirPath, c.convertFileWalk)
}

func (c *Convert) ConvertFile(filePath string) error {
	if strings.Contains(filePath, common.DefaultGullDirectory) {
		return nil
	}
	migration, err := NewMigrationFromConfigFile(filePath, c.FileNameIsEnvironment, c.JsonEncode)
	if err != nil {
		return err
	}
	name := GetMigrationNameFromConfigName(filePath)
	destPath := filepath.Join(c.DestinationDir, name)
	c.Logger.Info("Converting [%v] to [%v]", filePath, destPath)
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
