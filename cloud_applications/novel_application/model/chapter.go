package model

type Chapter struct {
	Name  string
	Index int
}

type ChapterDownload struct {
	ID       string
	Name     string
	Index    int
	SourceID string
	URL      string
	Content  string
}
