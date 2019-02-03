package agent

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"

	cldstrg "theterriblechild/CloudApp/applications/storageapp/internal/model"
)

type JobManager struct {
	uploadWorkerChan    chan Job
	downloadWorkerChan  chan Job
	immediateWorkerChan chan Job
	notifyNewJobChan    chan bool
	downloadQueue       []string
	downloadQueueMutex  sync.Mutex
	uploadQueue         []string
	uploadQueueMutex    sync.Mutex

	availableUploadWorker   int
	availableDownloadWorker int

	jobMap map[string]Job
}
type JobType int

const (
	JobType_Immediate JobType = 1
	JobType_Upload    JobType = 2
	JobType_Download  JobType = 3
)

func (instance *JobManager) Initialize() {

	MaxUploadWorker := viper.GetInt("maxUploadWorker")
	MaxDownloadWorker := viper.GetInt("maxDownloadWorker")

	instance.uploadWorkerChan = make(chan Job, MaxUploadWorker)
	instance.downloadWorkerChan = make(chan Job, MaxDownloadWorker)
	instance.notifyNewJobChan = make(chan bool, 100)

	uploadWorkerResultChan := make(chan bool, MaxUploadWorker)
	downloadWorkerResultChan := make(chan bool, MaxDownloadWorker)

	instance.uploadQueue = make([]string, 0)
	instance.downloadQueue = make([]string, 0)

	instance.downloadQueueMutex = sync.Mutex{}
	instance.uploadQueueMutex = sync.Mutex{}

	instance.jobMap = make(map[string]Job)

	go instance.jobDistributor(uploadWorkerResultChan, downloadWorkerResultChan)

	for i := 0; i < MaxUploadWorker; i++ {
		workerName := fmt.Sprintf("UploadWorker-%d", i)
		go doJob(workerName, instance.uploadWorkerChan, uploadWorkerResultChan)
	}
	for i := 0; i < MaxDownloadWorker; i++ {
		workerName := fmt.Sprintf("DownloadWorker-%d", i)
		go doJob(workerName, instance.downloadWorkerChan, downloadWorkerResultChan)
	}
}

func (instance *JobManager) jobDistributor(uploadWorkerResultChan <-chan bool, downloadWorkerResultChan <-chan bool) {
	for {
		select {
		case job, _ := <-instance.immediateWorkerChan:
			go job.f()
		case ready, _ := <-uploadWorkerResultChan:
			if ready {
				instance.availableUploadWorker++
				instance.startUploadDownloadJob()
			} else {
				instance.availableUploadWorker--
			}

		case ready, _ := <-downloadWorkerResultChan:
			if ready {
				instance.availableDownloadWorker++
				instance.startUploadDownloadJob()
			} else {
				instance.availableDownloadWorker--
			}
		case <-instance.notifyNewJobChan:
			instance.startUploadDownloadJob()
		case <-time.After(30 * time.Second): //Check just in case it is missed in other scenarios
			instance.startUploadDownloadJob()
		}
	}
}

func (instance *JobManager) startUploadDownloadJob() {
	log.Printf("Upload worker available: %d  Download worker available: %d", instance.availableUploadWorker, instance.availableDownloadWorker)
	if instance.availableUploadWorker > 0 {
		if taskID := instance.dequeueUploadJob(); len(taskID) > 0 {
			if job, ok := instance.jobMap[taskID]; ok {
				instance.uploadWorkerChan <- job
				delete(instance.jobMap, taskID)
			}
		}
	}
	if instance.availableDownloadWorker > 0 {
		if taskID := instance.dequeueDownloadJob(); len(taskID) > 0 {
			if job, ok := instance.jobMap[taskID]; ok {
				instance.downloadWorkerChan <- job
				delete(instance.jobMap, taskID)
			}
		}
	}
}

func (instance *JobManager) updateTaskProgress(newProgress cldstrg.ProgressUpdate) {
	//instance.jobMap[newProgress.TaskId].progress = newProgress
}

type Job struct {
	taskID    string
	progress  cldstrg.ProgressUpdate
	cancelJob chan bool
	f         func()
}

func (instance *JobManager) createJobFromHandler(handler CommandHandler) Job {
	taskID := handler.agentCommand.GetAgentCommand().TaskID
	cancelJobChan := make(chan bool, 1)
	return Job{taskID: taskID, cancelJob: cancelJobChan, f: handler.handleCommand, progress: cldstrg.ProgressUpdate{State: cldstrg.ProgressUpdate_NotStarted}}
}

func (instance *JobManager) addImmediateJob(handler CommandHandler) {
	job := instance.createJobFromHandler(handler)
	instance.immediateWorkerChan <- job
}

func (instance *JobManager) addUploadJob(handler CommandHandler) {
	taskID := handler.agentCommand.GetAgentCommand().TaskID
	job := instance.createJobFromHandler(handler)
	instance.jobMap[taskID] = job
	instance.uploadQueueMutex.Lock()
	log.Println("Enqueuing upload task: " + taskID)
	instance.uploadQueue = append(instance.uploadQueue, taskID)
	instance.uploadQueueMutex.Unlock()
	instance.notifyNewJobChan <- true
}

func (instance *JobManager) dequeueUploadJob() string {
	var id string
	instance.uploadQueueMutex.Lock()
	if len(instance.uploadQueue) > 0 {
		id = instance.uploadQueue[0]
		instance.uploadQueue = instance.uploadQueue[1:]
	}

	instance.uploadQueueMutex.Unlock()
	return id
}

func (instance *JobManager) addDownloadJob(handler CommandHandler) {
	taskID := handler.agentCommand.GetAgentCommand().TaskID
	job := instance.createJobFromHandler(handler)
	instance.jobMap[taskID] = job
	instance.downloadQueueMutex.Lock()
	log.Println("Enqueuing download task: " + taskID)
	instance.downloadQueue = append(instance.downloadQueue, taskID)
	instance.downloadQueueMutex.Unlock()
	instance.notifyNewJobChan <- true
}

func (instance *JobManager) dequeueDownloadJob() string {
	var id string
	instance.downloadQueueMutex.Lock()
	if len(instance.downloadQueue) > 0 {
		id = instance.downloadQueue[0]
		instance.downloadQueue = instance.downloadQueue[1:]
	}
	instance.downloadQueueMutex.Unlock()
	return id
}

func doJob(workerName string, j <-chan Job, result chan bool) {
	for {
		log.Println(workerName + " waiting for job")
		result <- true
		job := <-j
		result <- false
		log.Println(workerName + " starting job " + job.taskID)
		job.f()
		log.Println(workerName + " finished job " + job.taskID)
	}

}
