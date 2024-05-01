package orderservice

import (
	"context"
	"encoding/json"
	"github.com/adjust/redismq"
	"gitlab.sazito.com/sazito/event_publisher/adapter/redisqueue"
	"gitlab.sazito.com/sazito/event_publisher/entity"
	"gitlab.sazito.com/sazito/event_publisher/pkg/errmsg"
	"gitlab.sazito.com/sazito/event_publisher/pkg/richerror"
	"net/http"
	"strings"
	"time"
)

func ReadAndAckFromRedis(consumer *redismq.Consumer) (string, error) {
	pkg, err := consumer.Get()
	if err != nil {
		return "", err
	}

	payload := pkg.Payload

	err = pkg.Ack()
	if err != nil {
		return "", err
	}

	return payload, nil
}

type ReadAndSaveAndSendRequest struct {
	Queue redismq.Queue
}

type ReadAndSaveAndSendResponse struct {
	StatusCode int `json:"status_code"`
}

func (s Service) ReadAndSaveAndSend(ctx context.Context, req ReadAndSaveAndSendRequest) (ReadAndSaveAndSendResponse, error) {
	const op = "orderservice.ReadAndSaveAndSend"

	consumer, err := redisqueue.NewConsumer(req.Queue, "test")
	if err != nil {

		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	payload, err := ReadAndAckFromRedis(consumer)
	if err != nil {

		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	var order entity.Order
	err = json.Unmarshal([]byte(payload), &order)
	if err != nil {
		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgCantUnmarshal.String())
	}

	order, err = s.OrderRepo.Save(ctx, order)
	if err != nil {

		return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err)
	}

	request, err := http.NewRequest("POST", order.WebHookURL, strings.NewReader(order.MetaData))
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
		order, err := s.OrderRepo.ModifyIsPublished(ctx, order, true)
		if err != nil {
			return ReadAndSaveAndSendResponse{}, richerror.New(op).WithErr(err)
		}
	}
}
