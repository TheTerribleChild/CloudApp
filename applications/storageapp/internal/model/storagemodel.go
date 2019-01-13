package model

type MergeMode uint8

const (
	OVERWRITE MergeMode = 0
	RENAME    MergeMode = 1
	SKIP      MergeMode = 2
)

type FileStat struct {
	FilePath string
	Size     int64
}

type FileRead struct {
	Files []FileStat
}

type FileWrite struct {
	WriteLocation   string
	Files           []FileStat
	Decompress      bool
	ConflictResolve MergeMode
}
