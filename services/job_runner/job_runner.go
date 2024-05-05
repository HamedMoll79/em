package job_runner

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/bsm/redislock"
	"github.com/p3ym4n/re"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.snapp.ir/backend/nbo/api-go/bulk_job"
	"gitlab.snapp.ir/backend/nbo/api-go/event"
	"gitlab.snapp.ir/backend/nbo/pkg/broadcast"
	"gitlab.snapp.ir/backend/nbo/tractor/config"
	"gitlab.snapp.ir/backend/nbo/tractor/contract"
	"gitlab.snapp.ir/backend/nbo/tractor/entity"
	"gitlab.snapp.ipackage job_runner

r/backend/nbo/tractor/log"
	"gitlab.snapp.ir/backend/nbo/tractor/metrics"
	"gitlab.snapp.ir/backend/nbo/tractor/pkg/baser"
	"google.golang.org/protobuf/proto"
)

type jobRunner struct {
	serviceName string
	instanceID  string
	config      config.JobRunnerConfig
	consumer    contract.Consumer
	broker      contract.Broker
	queueStore  contract.JobQueueStore
	publisher   *publisher
}
type RPS struct {
	Value uint
	Unit  uint
}

func New(instanceID, serviceName string, consumer contract.Consumer, broker contract.Broker, queueStore contract.JobQueueStore, config config.Config) *jobRunner {
	publisher := newPublisher(config.JobRunner.PublisherMaxWorkers, broker)
	return &jobRunner{
		config: config.JobRunner,
		//TODO: move x-api-key to db and set each bulkjob a unique one
		consumer:    consumer,
		broker:      broker,
		serviceName: serviceName,
		instanceID:  instanceID,
		queueStore:  queueStore,
		publisher:   publisher,
	}
}

func (jr *jobRunner) Start(ctx context.Context, wg *sync.WaitGroup) {
	cancelFuncs := make(map[uint]context.CancelFunc) // [jobID] : cancelFunc

	defer func() {
		for _, cancelFunc := range cancelFuncs {
			cancelFunc()
		}
		log.Info("shutting down job runner in ", config.Get().JobRunner.GracefulShutdown)
		<-time.After(config.Get().JobRunner.GracefulShutdown)
		wg.Done()
	}()

	timer := time.NewTicker(jr.config.CheckFrequency)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:

			//handles Single Mode jobs
			allJobs, err := jr.queueStore.GetInQueueJobIDs(context.Background())
			if err != nil {
				log.Error(fmt.Sprintf("unable to read job queues : %s", err.Internal()))
				continue
			}
			//looking for new jobs and stopping finished jobs
			for _, jobID := range allJobs {
				isNewJob := true
				for prevJob := range cancelFuncs {
					if jobID == prevJob {
						isNewJob = false
					}
				}
				if isNewJob {
					ctx1, cancelFunc := context.WithCancel(ctx)
					cancelFuncs[jobID] = cancelFunc
					info, err1 := jr.queueStore.GetJobInfo(ctx, jobID)
					if err1 != nil {
						log.Error("unable to get jobInfo", jobID, err1)
						continue
					}
					go jr.handleNewJob(ctx1, jobID, info, wg)
					log.Info(fmt.Sprintf("start consuming jobitems queue of %d", jobID))
				}
			}
			for prevJob, cancelFunc := range cancelFuncs {
				jobFinished := true
				for _, job := range allJobs {
					if job == prevJob {
						jobFinished = false
					}
				}
				if jobFinished {
					cancelFunc()
					delete(cancelFuncs, prevJob)
					log.Info(fmt.Sprintf("stopped consuming job %d queue", prevJob))
				}
			}
		case <-ctx.Done():
			return
		}
	}
}
func (jr *jobRunner) handleNewJob(ctx context.Context, jobID uint, jobInfo map[string]string, wg *sync.WaitGroup) {

	queueName := fmt.Sprintf("job_%d_items_queue", jobID)
	jobRPS, err1 := getJobRPS(jobInfo)
	if err1 != nil {
		log.Error(fmt.Sprintf("unable to convert jobRPS %v", jobInfo[entity.PropertyBulkRPS]))
		return
	}
	jobMaxConcurrentReq, err2 := strconv.Atoi(jobInfo[entity.PropertyBulkMaxConcurrentRequest])
	if err2 != nil {
		log.Error(fmt.Sprintf("unable to convert jobMaxConcurrentReq %v", jobInfo[entity.PropertyBulkMaxConcurrentRequest]))
		return
	}
	bulkJobID, err3 := strconv.Atoi(jobInfo[entity.EntityJobBulkJobID])
	if err3 != nil {
		log.Error(fmt.Sprintf("unable to convert bulkJobID %v", jobInfo[entity.EntityJobBulkJobID]))
	}

	apiKey := jobInfo[entity.EntityBulkJobApiKey]
	if apiKey == "" {
		log.Error("unable to get api key from job info (empty)")
		return
	}

	var lock *redislock.Lock
	timer := time.NewTicker((time.Second / time.Duration(jobRPS.Value)) * time.Duration(jobRPS.Unit))
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			// for rps < 1 we have to set a lock to prevent race
			// as long as running the whole process of job item
			if jobRPS.Unit > 1 {
				length, err := jr.queueStore.GetQueueLength(context.Background(), queueName)
				if err != nil {
					log.Error("unable to get job queue length", err)
					continue
				}

				if length == 0 {
					continue
				}

				var lockErr re.Error
				lockKey := fmt.Sprintf("rps_lock_for_bulk_%d", bulkJobID)
				ttl := time.Second * time.Duration(jobRPS.Unit/jobRPS.Value)
				lock, lockErr = jr.queueStore.ObtainLock(ctx, lockKey, ttl)
				if lockErr != nil || lock == nil {
					continue
				}
			}
			concurrentRequest, err1 := jr.queueStore.GetJobConcurrentReq(ctx, uint(bulkJobID))
			if err1 != nil {
				log.Error("unable to get job concurrent Request", err1)
				continue
			}
			if concurrentRequest >= jobMaxConcurrentReq {
				continue
			}

			currentRPS, err2 := jr.queueStore.GetJobCurrentRPS(ctx, uint(bulkJobID), jobRPS.Unit)
			if err2 != nil {
				log.Error("unable to get job current RPS", err2)
				continue
			}
			if currentRPS >= int64(jobRPS.Value) {
				continue
			}
			message, ok, err := jr.queueStore.FetchMessage(context.Background(), queueName)
			if err != nil {
				log.Error(fmt.Sprintf("job runner reading queue err : %s", err.Internal()))
			} else if ok {
				metrics.Get(metrics.RunnerTotalJobItemsConsumption).(*prometheus.CounterVec).
					WithLabelValues(fmt.Sprintf("%d", bulkJobID), fmt.Sprintf("%d", jobID)).Inc()
				wg.Add(1)
				go func() {
					defer func() {
						if lock != nil {
							err4 := lock.Release(ctx)
							if err4 != nil {
								log.Error(fmt.Sprintf("unable to release lock %v", err4))
							}
						}
						wg.Done()
					}()
					err4 := jr.decodeMessageAndExecute(*message, bulkJobID, apiKey)
					if err4 != nil {
						log.Error("unable to decodeMessageAndExecute %s", err4.Internal())
					}
				}()
			}
		}
	}
}

func (jr *jobRunner) decodeMessageAndExecute(message broadcast.Message, bulkJobID int, apiKey string) re.Error {
	const op = re.Op("jobRunner.decodeMessageAndExecute")

	defer func() {
		if r := recover(); r != nil {
			metrics.Get(metrics.RunnerDecodeMessagePanic).(*prometheus.CounterVec).
				WithLabelValues().Inc()
			log.Error("Recovered in jobRunner.decodeMessageAndExecute", r)
		}
	}()

	if message.EntityType != broadcast.EventEntityType {
		return nil
	}

	evtByte, err2 := baser.DecodeStringStd(message.Payload)
	if err2 != nil {
		return re.New(op, err2, re.KindUnexpected)
	}

	evt := new(event.Event)
	err3 := proto.Unmarshal(evtByte, evt)
	if err3 != nil {
		return re.New(op, err3, re.KindUnexpected)
	}

	switch evt.Message.(type) {
	case *event.Event_BulkJobMessage:
		bulkMessage := evt.GetBulkJobMessage()
		switch bulkMessage.Message.(type) {
		case *bulk_job.BulkJobMessage_NewJobItem:
			newJobItem := bulkMessage.Message.(*bulk_job.BulkJobMessage_NewJobItem).NewJobItem
			metrics.Get(metrics.RunnerSingleJobItemsConsumption).(*prometheus.CounterVec).
				WithLabelValues(fmt.Sprintf("%d", bulkJobID), fmt.Sprintf("%d", newJobItem.JobId)).Inc()
			log.Info("New JobItem Single Message : (job item id) ", newJobItem.GetJobItemId())
			jr.executeJobSingle(newJobItem, bulkJobID, apiKey)
		case *bulk_job.BulkJobMessage_NewJobItemList:
			jobItemList := bulkMessage.Message.(*bulk_job.BulkJobMessage_NewJobItemList).NewJobItemList
			//TODO , Refactor api-go package to change name of jobItemlist.job ?:(
			metrics.Get(metrics.RunnerBulkJobItemsConsumption).(*prometheus.CounterVec).
				WithLabelValues(fmt.Sprintf("%d", bulkJobID), fmt.Sprintf("%d", jobItemList.Jobs[0].JobId)).Inc()
			log.Info(fmt.Sprintf("New JobItem Bulk Message : {job_id: %d , len_jobitem_received : %d} ", jobItemList.Jobs[0].JobId, len(jobItemList.Jobs)))
			err4 := jr.executeJobBulk(jobItemList, bulkJobID, apiKey)
			if err4 != nil {
				return err4.Chain(op)
			}

		default:
			return nil
		}
	default:
		return nil
	}

	return nil
}

func (jr *jobRunner) executeJobSingle(newJobItem *bulk_job.NewJobItem, bulkJobID int, apiKey string) {
	defer func() {
		if r := recover(); r != nil {
			metrics.Get(metrics.RunnerSingleExecutePanic).(*prometheus.CounterVec).
				WithLabelValues().Inc()
			log.Error("Recovered in jobRunner.ExecuteJob", r)
		}
	}()

	bulkJob := jr.getBulkJobOfMessage(newJobItem)

	jobItem := &entity.JobItem{
		JobID:   uint(newJobItem.GetJobId()),
		Payload: newJobItem.GetPayload(),
		ID:      uint(newJobItem.GetJobItemId()),
		Job: &entity.Job{
			ID:                 uint(newJobItem.GetJobId()),
			BulkJob:            *bulkJob,
			BulkJobID:          bulkJob.ID,
			JobItemsTotalCount: uint(newJobItem.GetJobTotalCount()),
		},
	}

	metrics.Get(metrics.TotalExecuteWorkflowJobItem).(*prometheus.CounterVec).
		WithLabelValues(fmt.Sprintf("%d", jobItem.JobID)).Inc()

	attempts := 0
	for {
		jobItemStatus, err := jr.callRequestActivitySingle(*bulkJob, *jobItem, &attempts, bulkJobID, apiKey)
		if err != nil {
			attempts++
			log.Error("call request activity failed : ", err)
			continue
		}

		jr.publisher.Channel <- &PublisherMessage{
			JobID:  jobItem.JobID,
			Status: *jobItemStatus,
		}
		break
	}
}

func (jr *jobRunner) executeJobBulk(newJobItemList *bulk_job.NewJobItemList, bulkJobID int, apiKey string) re.Error {
	if len(newJobItemList.Jobs) == 0 {
		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			metrics.Get(metrics.RunnerBulkExecutePanic).(*prometheus.CounterVec).
				WithLabelValues().Inc()
			log.Error("Recovered in jobRunner.ExecuteJobBulk", r)
		}
	}()

	bulkJob := jr.getBulkJobOfMessage(newJobItemList.Jobs[0]) //parent jobs are the same
	//TODO , change name of var Jobs ? why jobs while means jobItems
	jobItems := make([]entity.JobItem, len(newJobItemList.Jobs))

	jobTotalCount := newJobItemList.Jobs[0].GetJobTotalCount()
	for i, item := range newJobItemList.Jobs {
		jobItems[i] = entity.JobItem{
			ID:      uint(item.GetJobItemId()),
			JobID:   uint(item.GetJobId()),
			Payload: item.GetPayload(),
			Job: &entity.Job{
				ID:                 uint(item.GetJobId()),
				BulkJob:            *bulkJob,
				BulkJobID:          bulkJob.ID,
				JobItemsTotalCount: uint(item.GetJobTotalCount()),
			},
		}
	}

	attempts := 0
	jr.queueStore.IncrConcurrentReq(context.Background(), uint(bulkJobID))
	defer jr.queueStore.DecrConcurrentReq(context.Background(), uint(bulkJobID))
	for {
		err := jr.callRequestActivityBulk(*bulkJob, jobItems, jobTotalCount, &attempts, bulkJobID, apiKey)
		if err != nil {
			attempts++
			log.Error("call request activity failed : ", err)

			errPayload := err.Error()
			trimAfter := 200
			if len(errPayload) > trimAfter {
				errPayload = errPayload[:trimAfter]
			}

			message := broadcast.Message{
				AudienceType: broadcast.SystemAudienceType,
				AudienceID:   0,
				EntityType:   broadcast.EventEntityType,
				EntityID:     jobItems[0].JobID,
				Payload:      errPayload,
			}

			queueName := fmt.Sprintf("job_%d_last_fail_reason", jobItems[0].JobID)

			if err := jr.queueStore.DispatchMessage(context.Background(), queueName, message); err != nil {
				log.Error("call request activity failed : ", err)
			}
			time.Sleep(time.Second * time.Duration(bulkJob.NextItemInterval))
			continue
		}
		break
	}

	return nil
}

func (jr *jobRunner) getBulkJobOfMessage(newJobItem *bulk_job.NewJobItem) *entity.BulkJob {
	bulkJob := newJobItem.GetBulkJob()

	bulkEntity := &entity.BulkJob{
		Name:             bulkJob.GetName(),
		TargetURL:        bulkJob.GetTargetUrl(),
		Method:           bulkJob.GetMethod(),
		TimeOutInterval:  bulkJob.GetTimeOutInterval(),
		RetryCount:       int(bulkJob.GetRetryCount()),
		NextItemInterval: bulkJob.GetNextItemInterval(),
	}

	return bulkEntity
}

func getJobRPS(jobInfo map[string]string) (RPS, error) {
	fRPS, err := strconv.ParseFloat(jobInfo[entity.PropertyBulkRPS], 64)
	if err != nil {
		log.Error("Invalid RPS in ", jobInfo)
		return RPS{}, err
	}
	if fRPS > 1 {
		return RPS{
			Unit:  1,
			Value: uint(fRPS),
		}, nil
	}
	if fRPS >= 0.1 {
		return RPS{
			Unit:  10,
			Value: uint(fRPS * 10),
		}, nil
	}

	if fRPS >= 0.01 {
		return RPS{
			Unit:  100,
			Value: uint(fRPS * 100),
		}, nil
	}
	if fRPS >= 0.001 {
		return RPS{
			Unit:  1000,
			Value: uint(fRPS * 1000),
		}, nil
	}
	return RPS{}, fmt.Errorf("invalid rps %v", fRPS)
}