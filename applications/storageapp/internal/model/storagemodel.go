package model

import (
	fileutil "theterriblechild/CloudApp/tools/utils/file"
)

type FileStat struct {
	FilePath         string
	Size             int64
	LastModifiedTime int
	Hash             string
}

type FileHash struct {
	FilePath         string
	Size             int64
	LastModifiedTime int64
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
