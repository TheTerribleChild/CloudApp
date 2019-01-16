package model

import (
	fileutil "theterriblechild/CloudApp/tools/utils/file"
)

type FileStat struct {
	FilePath         string
	Size             int64
	LastModifiedTime int
}

type FileRead struct {
	Files []FileStat
}

type FileWrite struct {
	WriteLocation   string
	Files           []FileStat
	Decompress      bool
	ConflictResolve fileutil.MergeMode
}
