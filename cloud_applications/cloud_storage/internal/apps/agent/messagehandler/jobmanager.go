package agentmessagehandler

import(
	cldstrg "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/cloud_storage/internal/model"
	"log"
	"sync"
	"time"
	"fmt"
)

type JobManager struct {
	uploadWorkerChan chan Job;
	uploadWorkerResultChan int;
	downloadWorkerChan chan Job;
	immediateWorkerChan chan Job;
	downloadQueue []string;
	downloadQueueMutex sync.Mutex;
	uploadQueue []string;
	uploadQueueMutex sync.Mutex;

	availableUploadWorker int;
	availableDownloadWorker int;
	
	jobMap map[string] Job;
}

func(instance *JobManager) Initialize() {

	MAX_UPLOAD_WORKER := 1
	MAX_DOWNLOAD_WORKER := 2
	
	instance.uploadWorkerChan = make(chan Job , MAX_UPLOAD_WORKER)
	instance.downloadWorkerChan = make(chan Job, MAX_DOWNLOAD_WORKER)
	
	uploadWorkerResultChan := make(chan bool, MAX_UPLOAD_WORKER)
	downloadWorkerResultChan := make(chan bool, MAX_DOWNLOAD_WORKER)
	
	instance.uploadQueue = make([]string, 0)
	instance.downloadQueue = make([]string, 0)

	instance.downloadQueueMutex = sync.Mutex{}
	instance.uploadQueueMutex = sync.Mutex{}
	
	instance.jobMap = make(map[string] Job)
	
	go instance.jobDistributor(uploadWorkerResultChan, downloadWorkerResultChan)
	
	for i := 0; i < MAX_UPLOAD_WORKER; i++ {
		workerName := fmt.Sprintf("UploadWorker-%d", i)
		go doJob(workerName, instance.uploadWorkerChan, uploadWorkerResultChan)
	}
	for i := 0; i < MAX_DOWNLOAD_WORKER; i++ {
		workerName := fmt.Sprintf("DownloadWorker-%d", i)
		go doJob(workerName, instance.downloadWorkerChan, downloadWorkerResultChan)
	}
	
}

func(instance *JobManager) jobDistributor(uploadWorkerResultChan <- chan bool, downloadWorkerResultChan <- chan bool){
	for {
		select{
		case job, _ := <-instance.immediateWorkerChan:
			go job.f()
		case ready, _ := <- uploadWorkerResultChan:
			if ready {
				instance.availableUploadWorker++
				instance.startUploadDownloadJob()
			} else{
				instance.availableUploadWorker--
			}
			
		case ready,_ := <- downloadWorkerResultChan:
			if ready {
				instance.availableUploadWorker++
				instance.startUploadDownloadJob()
			} else{
				instance.availableDownloadWorker--
			}
		case <-time.After(1 * time.Second):
			instance.startUploadDownloadJob()
		}
	}
}

func(instance *JobManager) startUploadDownloadJob(){
	if instance.availableUploadWorker > 0 {
		if taskId := instance.dequeueUploadJob(); len(taskId) > 0{
			if job, ok := instance.jobMap[taskId]; ok {
				instance.uploadWorkerChan <- job
				delete(instance.jobMap, taskId)
			}
		}
	}
	if instance.availableDownloadWorker > 0 {
		if taskId := instance.dequeueDownloadJob(); len(taskId) > 0{
			if job, ok := instance.jobMap[taskId]; ok {
				instance.downloadWorkerChan <- job
				delete(instance.jobMap, taskId)
			}
		}
	}
}

func(instance *JobManager) updateTaskProgress(newProgress cldstrg.ProgressUpdate){
	//instance.jobMap[newProgress.TaskId].progress = newProgress
}

type Job struct{
	taskId string;
	progress cldstrg.ProgressUpdate
	cancelJob chan bool;
	f func();
}

func(instance *JobManager) AddJobForHandler(handler MessageHandlerWrapper) {
	cancelJobChan := make(chan bool, 1)
	taskId := handler.taskId
	job := Job{taskId: handler.taskId, cancelJob: cancelJobChan, f: handler.HandleMessage, progress : cldstrg.ProgressUpdate{State:cldstrg.ProgressUpdate_NotStarted}}
	instance.jobMap[taskId] = job
	switch handler.messageHandler.(type) {
	case UploadFileHandler:
		instance.addUploadJob(taskId)
		break
	case DownloadFileHandler:
		instance.addDownloadJob(taskId)
		break
	default:
		instance.immediateWorkerChan <- job
		break
	}
}

func(instance *JobManager) addUploadJob(id string) {
	instance.uploadQueueMutex.Lock()
	log.Println("Enqueuing upload task: " + id)
	instance.uploadQueue = append(instance.uploadQueue, id)
	instance.uploadQueueMutex.Unlock()
}

func(instance *JobManager) dequeueUploadJob() string {
	var id string
	instance.uploadQueueMutex.Lock()
	if len(instance.uploadQueue) > 0 {
		id = instance.uploadQueue[0]
		instance.uploadQueue = instance.uploadQueue[1:]
	}
	
	instance.uploadQueueMutex.Unlock()
	return id
}

func(instance *JobManager) addDownloadJob(id string) {
	instance.downloadQueueMutex.Lock()
	log.Println("Enqueuing download task: " + id)
	instance.downloadQueue = append(instance.downloadQueue, id)
	instance.downloadQueueMutex.Unlock()
}

func(instance *JobManager) dequeueDownloadJob() string {
	var id string
	instance.downloadQueueMutex.Lock()
	if len(instance.downloadQueue) > 0 {
		id = instance.downloadQueue[0]
		instance.downloadQueue = instance.downloadQueue[1:]
	}
	instance.downloadQueueMutex.Unlock()
	return id
}

func doJob(workerName string, j <- chan Job, result chan bool) {
	for {
		log.Println(workerName + " waiting for job")
		result <- true
		job := <- j
		result <- false
		log.Println(workerName + " starting job " + job.taskId)
		job.f()
		log.Println(workerName + " finished job " + job.taskId)
	}
	
}