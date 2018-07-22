package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	cs "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/novel_application/connectors/service"
	stomp "github.com/go-stomp/stomp"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func getChapter(ch *cs.ChapterSourceMetadata, c cs.ConnectorClient, ctx context.Context, idx int) {
	start := time.Now()
	chData, err := c.GetChapterData(ctx, &cs.ChapterDownloadRequest{ChapterMetadata: ch})
	elapsed := time.Since(start)
	log.Printf("thread %d  Binomial took %s", idx, elapsed)
	if err != nil {
		fmt.Errorf("could not download chapter: %s", err)
	} else {
		str, _ := json.Marshal(chData)
		fmt.Println(string(str))
	}
}

func testWebService() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := cs.NewConnectorClient(conn)
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	r, err := c.GetNovelData(ctx, &cs.NovelDownloadRequest{SourceId: "69shu.com", NovelId: "29242"})
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
	if err != nil {
		log.Fatalf("could not download novel: %v", err)
	}
	str, _ := json.Marshal(r)
	fmt.Println(string(str))
	i := 0
	for idx, ch := range r.NovelData.Chapters {
		go getChapter(ch, c, ctx, idx)
		i++
	}
	time.Sleep(100 * time.Second)
}

func testMQ() {
	webconn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer webconn.Close()
	c := cs.NewConnectorClient(webconn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	r, err := c.GetNovelData(ctx, &cs.NovelDownloadRequest{SourceId: "69shu.com", NovelId: "29242"})

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := stomp.Dial("tcp", "192.168.1.71:61613")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Publishing")

	for _, ch := range r.NovelData.Chapters {
		msg, _ := proto.Marshal(ch)
		err = conn.Send("/queue/ChapterDownload", "text/plain", msg)
	}

	time.Sleep(1 * time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	//testWebService()
	testMQ()
}
