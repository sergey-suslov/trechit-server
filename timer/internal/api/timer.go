package api

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/sergey-suslov/trechit-server/api/timer"
	"github.com/sergey-suslov/trechit-server/timer/internal/db/models"
	"log"
)

type TimerService struct{}

func (TimerService) StartTimer(context context.Context, request *timer.StartTimerRequest) (*timer.StartTimerResponse, error) {
	log.Println("Starting timer")
	startTime, err := ptypes.Timestamp(request.GetStartTime())
	if err != nil {
		return nil, err
	}
	var timeSpan *models.TimeSpan
	timeSpan, err = models.StartTimer(request.GetUserId(), &startTime)
	if err != nil {
		return nil, err
	}
	var response *timer.TimeSpan
	response, err = timeSpan.ToProtoTimeSpan()
	return &timer.StartTimerResponse{
		TimeSpan: response,
	}, nil
}

func (TimerService) StopTimer(context context.Context, request *timer.StopTimerRequest) (*timer.StopTimerResponse, error) {
	stopTime, err := ptypes.Timestamp(request.GetStopTime())
	if err != nil {
		return nil, err
	}
	err = models.StopTimer(request.GetTimeSpanId(), &stopTime)
	if err != nil {
		return nil, err
	}
	return &timer.StopTimerResponse{
		TimeSpanId: request.GetTimeSpanId(),
	}, nil
}
