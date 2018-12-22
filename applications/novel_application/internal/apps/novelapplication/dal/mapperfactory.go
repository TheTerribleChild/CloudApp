package novelapplicationmapper

import (
	nas "github.com/TheTerribleChild/CloudApp/cloud_appplication_portal/cloud_applications/novel_application/internal/apps/novelapplication/service"
	cs "github.com/TheTerribleChild/CloudApp/cloud_appplication_portal/cloud_applications/novel_application/internal/apps/connectorservice/service"
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