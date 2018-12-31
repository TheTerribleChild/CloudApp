package connector

import (
	"time"
	// "encoding/json"
	"fmt"

	cs "github.com/TheTerribleChild/CloudApp/cloud_appplication_portal/cloud_applications/novel_application/internal/app/connectorservice/service"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	// "golang.org/x/net/netutil"
	"log"
	"net"

	"github.com/go-stomp/stomp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	// "time"
)

type connectorServer struct{}

func (s *connectorServer) GetNovelData(ctx context.Context, request *cs.NovelDownloadRequest) (*cs.NovelDownloadReply, error) {
	var reply cs.NovelDownloadReply
	requestedSourceId := request.SourceId
	sourceIds := GetSourceConnectorIDs()
	found := false
	for _, sourceId := range sourceIds {
		if sourceId == requestedSourceId {
			found = true
		}
	}

	if !found {
		return &reply, grpc.Errorf(codes.Unavailable, "cannot find source connector with ID: %s", requestedSourceId)
	}

	sourceConnector, err := GetSourceConnector(requestedSourceId)
	if err != nil {
		return &reply, grpc.Errorf(codes.Internal, "unable to retrieve source connector '%s'", requestedSourceId)
	}

	novelData, err := sourceConnector.GetChapterList(request.NovelId)

	if err != nil {
		return &reply, grpc.Errorf(codes.Internal, "unable to retrieve novel data for novel ID '%s' from source '%s'", request.NovelId, requestedSourceId)
	}

	return &cs.NovelDownloadReply{NovelData: &novelData}, nil
}

func (s *connectorServer) GetChapterData(ctx context.Context, request *cs.ChapterDownloadRequest) (*cs.ChapterDownloadReply, error) {
	var reply cs.ChapterDownloadReply
	chapterData, err := getChapterData(*request.ChapterMetadata)
	if err != nil {
		return &reply, err
	}
	return &cs.ChapterDownloadReply{ChapterData: &chapterData}, nil
}

func getChapterData(request cs.ChapterSourceMetadata) (cs.ChapterSourceData, error) {
	var chapterSourceData cs.ChapterSourceData
	requestedSourceId := request.SourceId
	sourceIds := GetSourceConnectorIDs()
	found := false
	for _, sourceId := range sourceIds {
		if sourceId == request.SourceId {
			found = true
		}
	}

	if !found {
		return chapterSourceData, grpc.Errorf(codes.Unavailable, "cannot find source connector with ID: %s", requestedSourceId)
	}

	sourceConnector, err := GetSourceConnector(requestedSourceId)
	if err != nil {
		return chapterSourceData, grpc.Errorf(codes.Internal, "unable to retrieve source connector '%s'", requestedSourceId)
	}

	chapterData, err := sourceConnector.GetChapterContent(request)

	if err != nil {
		return chapterSourceData, grpc.Errorf(codes.Internal, "unable to retrieve chapter data from URL '%s' for source '%s'", request.Url, requestedSourceId)
	}
	return chapterData, nil
}

func InitializeController() {
	go initializeChapterDownloadQueueListener()
	initializeWebService()
}

func initializeWebService() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.MaxConcurrentStreams(5))
	cs.RegisterConnectorServer(s, &connectorServer{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func initializeChapterDownloadQueueListener() {

	for true {
		conn, err := stomp.Dial("tcp", "192.168.1.71:61613", stomp.ConnOpt.HeartBeat(0*time.Second, 0*time.Second))
		//conn, err := stomp.Dial("tcp", "192.168.1.71:61613")

		if err != nil {
			fmt.Println(err)
		}
		sub, err := conn.Subscribe("/queue/ChapterDownload", stomp.AckClient)
		// stomp.SubscribeOpt.
		if err != nil {
			fmt.Errorf("Sub error: %s", err)
			break
		}
		for {
			// msg := <-sub.C
			msg, err := sub.Read()
			if err != nil {
				fmt.Errorf("Error in receiving message: %s ", err)
			}
			if msg == nil {
				break
			}
			request := &cs.ChapterSourceMetadata{}
			proto.Unmarshal(msg.Body, request)
			getChapterData(*request)
			log.Println("Ack message " + request.Url)
			conn.Ack(msg)
		}
		//sub.Unsubscribe()
		defer conn.Disconnect()
	}

}
