package novelapplicationmapper

import (
	cs "theterriblechild/CloudApp/cloud_appplication_portal/cloud_applications/novel_application/internal/app/connectorservice/service"
	nas "theterriblechild/CloudApp/cloud_appplication_portal/cloud_applications/novel_application/internal/app/novelapplication/service"
)

type Mapper interface {
	AddNovel(nas.Novel)
	AddChapter(nas.Chapter)
	AddNovelSource(cs.NovelSourceMetadata)
	AddChapterSource(cs.ChapterSourceMetadata)
}

func GetMapper() string {
	return "hello"
}
