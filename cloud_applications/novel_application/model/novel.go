package model

type Novel struct {
	Name   string
	Author string
}

type NovelDownload struct {
	ID       string
	Name     string
	Source   string
	Chapters []ChapterDownload
}
