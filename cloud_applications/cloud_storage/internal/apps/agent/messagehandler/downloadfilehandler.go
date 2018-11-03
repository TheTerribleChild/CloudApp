package agentmessagehandler

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
)

type DownloadFileHandler struct {
	asc cldstrg.AgentServiceClient;
	ssc cldstrg.StorageServiceClient;

	message  *cldstrg.AgentMessage;
}

func (handler DownloadFileHandler) HandleMessage(){

}