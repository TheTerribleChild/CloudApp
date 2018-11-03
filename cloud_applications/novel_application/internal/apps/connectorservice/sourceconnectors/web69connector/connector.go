package main

import (
	"bytes"
	"errors"
	"fmt"
	model "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/novel_application/internal/model"
	"github.com/TheTerribleChild/cloud_appplication_portal/commons/utils/webutil"
	"golang.org/x/net/html"
	"strings"
)

//Export as symbol named "Connector"
var Connector Web69connector

type Web69connector struct {
	sourceCharset string
}

var sourceID = "69shu.com"
var sourceName = "69shu"
var baseURL = "https://www.69shu.com"
var chapterListURLTemplate = "https://www.69shu.com/%s/"
var chapterURLTemplate = "https://www.69shu.com/txt/%s/%s"
var chapterNodeSignature = "root/1/2/3/13/3/"
var titleNodeSignature = "root/1/2/3/7/1/5/0"

var charSet = webutil.GB2312

//GetChapterList Gets the Novel metadata for download
func (connector Web69connector) GetChapterList(requestNovelID string) (model.NovelSourceData, error) {

	var novelDownload model.NovelSourceData

	if len(requestNovelID) == 0 {
		return novelDownload, errors.New("novel ID cannot be null")
	}

	rootNode, err := webutil.GetURLContentAsNode(fmt.Sprintf(chapterListURLTemplate, requestNovelID), string(charSet))

	if err != nil {
		return novelDownload, errors.New("unable to retrieve new chatpers")
	}

	nodeSignatures := webutil.GenerateNodeSignature(rootNode)

	chapterURLPrefix := fmt.Sprintf(chapterURLTemplate, requestNovelID, "")
	var index int32 = 0
	chapters := []*model.ChapterSourceMetadata{}
	var novelName string

	for _, signature := range nodeSignatures {
		if signature.Signature == titleNodeSignature {
			novelName = signature.Node.Data
		}
		if strings.HasPrefix(signature.Signature, chapterNodeSignature) && len(signature.Node.Attr) > 0 {
			if len(novelName) == 0 {
				return novelDownload, fmt.Errorf("invalid novel ID: '%s'. unable to find title of novel", requestNovelID)
			}
			found := false
			for _, attr := range signature.Node.Attr {
				if attr.Key == "href" && strings.HasPrefix(attr.Val, chapterURLPrefix) {
					found = true
					break
				}
			}
			if found {
				chapterInfo, err := getChapterInfoFromNode(signature.Node)
				if err != nil {
					return novelDownload, fmt.Errorf("problem parsing the following into chapter index: %d ", index)
				}
				chapterInfo.Index = index
				chapterInfo.ChapterSourceId = sourceID
				chapters = append(chapters, &chapterInfo)
				index++
			}
		}
	}

	if len(chapters) == 0 || len(novelName) == 0 {
		return novelDownload, fmt.Errorf("invalid novel ID: '%s'. unable to find title of novel", requestNovelID)
	}
	var novelMetadata model.NovelSourceMetadata
	novelMetadata.Id = requestNovelID
	novelMetadata.SourceId = sourceID
	novelDownload.Metadata = &novelMetadata
	novelDownload.Chapters = chapters

	return novelDownload, nil
}

func (connector Web69connector) GetChapterContent(chapterMetadata model.ChapterSourceMetadata) (model.ChapterSourceData, error) {
	var chapterData model.ChapterSourceData

	if chapterMetadata.ChapterSourceId != sourceID {
		return chapterData, fmt.Errorf("'%s' is not accepted by %s connector", chapterMetadata.ChapterSourceId, sourceName)
	}
	if len(chapterMetadata.Url) == 0 || !strings.HasPrefix(chapterMetadata.Url, baseURL) {
		return chapterData, fmt.Errorf("invalid chapter URL: '%s'. url must start with: ", chapterMetadata.Url, baseURL)
	}

	rootNode, err := webutil.GetURLContentAsNode(chapterMetadata.Url, string(charSet))

	if err != nil {
		return chapterData, errors.New("unable to retrieve chatper")
	}

	nodeSignatures := webutil.GenerateNodeSignature(rootNode)

	startFound := false
	var chapterContent bytes.Buffer

	for _, signature := range nodeSignatures {
		if signature.Node.Data == "章节内容开始" {
			startFound = true
		} else if startFound {
			if signature.Node.Data == "script" {
				break
			}
			if signature.Node.Data != "br" {
				chapterContent.WriteString(signature.Node.Data)
			}
		}
	}
	chapterData.Content = chapterContent.String()
	chapterData.Metadata = &chapterMetadata
	return chapterData, nil
}

func (connector Web69connector) GetConnectorID() string {
	fmt.Println("Web69connector GetConnectorID")
	return sourceID
}

func (connector Web69connector) GetConnectorName() string {
	fmt.Println("Web69connector GetConnectorName")
	return sourceName
}

func getChapterInfoFromNode(node *html.Node) (model.ChapterSourceMetadata, error) {
	var chapterDownload model.ChapterSourceMetadata
	var url string
	var name string
	var id string
	if node == nil {
		return chapterDownload, errors.New("getChapterInfoFromNode: parameter cannot be null")
	}
	if len(node.Attr) == 0 {
		return chapterDownload, errors.New("getChapterInfoFromNode: node is missing necessary attribute(s)")
	}
	if node.FirstChild == nil {
		return chapterDownload, errors.New("getChapterInfoFromNode: node is missing mandatory child node")
	}

	for _, attr := range node.Attr {
		if attr.Key == "href" {
			url = attr.Val
			break
		}
	}
	name = node.FirstChild.Data
	idx := strings.Index(name, ".")
	if idx != -1 {
		name = name[idx+1:]
	} else {
		return chapterDownload, fmt.Errorf("getChapterInfoFromNode: unable to extract chapter name. expected '.' in %s", name)
	}

	idx = strings.LastIndex(url, "/")
	if idx != -1 {
		id = url[idx+1:]
	} else {
		return chapterDownload, fmt.Errorf("getChapterInfoFromNode: unable to extract chapter ID. expected '.' in %s", url)
	}
	chapterDownload = model.ChapterSourceMetadata{ Id: id, Url: url}

	return chapterDownload, nil
}
