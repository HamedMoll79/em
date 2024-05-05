package event_service

import (
	"context"
	"encoding/json"
	"gitlab.sazito.com/sazito/event_publisher/entity"
	"gitlab.sazito.com/sazito/event_publisher/pkg/richerror"
	"net/http"
	"strings"
	"time"
)

//func ReadAndAckFromRedis(consumer *redismq.Consumer) (string, error) {
//	pkg, err := consumer.Get()
//	if err != nil {
//		return "", err
//	}
//
//	payload := pkg.Payload
//
//	err = pkg.Ack()
//	if err != nil {
//		return "", err
//	}
//
//	return payload, nil
//}

type ReadAndSaveAndSendRequest struct {
}

type ReadAndSaveAndSendResponse struct {
	StatusCode int `json:"status_code"`
}

func (s Service) ReadAndSaveAndSend(ctx context.Context, req ReadAndSaveAndSendRequest) (ReadAndSaveAndSendResponse, error) {
	const op = "event_service.ReadAndSaveAndSend"

	message, err := s.RedisDB.FetchMessage(ctx, s.RedisQueueName)

	var event entity.Event
	err = json.Unmarshal([]byte(message), event)
	if err != nil {

		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	event, err = s.EventRepo.Save(ctx, event)
	if err != nil {

		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err)
	}

	jsonEvent, err := json.Marshal(event)
	if err != nil {

		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	webhook, err := s.WebhookRepo.GetByStoreID(ctx, event.StoreID)
	if err != nil {

		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err)
	}

	request, err := http.NewRequest("POST", webhook.Url, strings.NewReader(string(jsonEvent)))
	if err != nil {

		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	request.Header.Add("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Do(request)
	if err != nil {
		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindBadRequest)
	}

	var response ReadAndSaveAndSendResponse
	response.StatusCode = resp.StatusCode

	if response.StatusCode == 200 {
		event, err = s.EventRepo.ModifyIsPublished(ctx, event, true)
		if err != nil {
			return response, richerror.New(op).WithErr(err)
		}
	}

	return response, nil
}
