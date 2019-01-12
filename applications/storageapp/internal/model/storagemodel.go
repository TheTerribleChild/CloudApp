package model

type FileStat struct {
	FilePath string
	Size int64
}

type FileRead struct {
	Files []FileStat
}

type FileWrite struct {
	WriteRoot string
	Files []FileStat
	Decompress bool
}