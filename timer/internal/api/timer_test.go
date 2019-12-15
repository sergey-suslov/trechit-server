package api

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/sergey-suslov/trechit-server/api/timer"
	"github.com/sergey-suslov/trechit-server/timer/internal/db"
	"github.com/sergey-suslov/trechit-server/timer/internal/db/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStartTimer(t *testing.T) {
	_, err := db.Init()
	if err != nil {
		t.Error(err)
	}
	now := time.Now()
	var userId int32 = 1
	nowTimestamp, err := ptypes.TimestampProto(now)
	if err != nil {
		t.Error(err)
	}
	request := &timer.StartTimerRequest{
		UserId:               userId,
		StartTime:            nowTimestamp,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}
	timerService := &TimerService{}
	response, err := timerService.StartTimer(context.Background(), request)
	if err != nil {
		t.Error(err)
	}
	timeSpan, err := models.GetTimeSpanById(response.GetTimeSpan().GetId())
	if err != nil {
		t.Error(err)
	}
	responseTimeSpan := response.GetTimeSpan()
	responseStartTime, err := ptypes.Timestamp(responseTimeSpan.StartTime)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, timeSpan.Id, responseTimeSpan.Id, "Start time Id is not equal to the given one")
	assert.Equal(t, timeSpan.UserId, responseTimeSpan.UserId, "Start time UserId is not equal to the given one")
	assert.Equal(t, timeSpan.StartTime.Unix(), responseStartTime.Unix(), "Start time is not equal to given one")
}

func TestStopTimer(t *testing.T) {
	_, err := db.Init()
	if err != nil {
		t.Error(err)
	}
	now := time.Now()
	var userId int32 = 1
	nowTimestamp, err := ptypes.TimestampProto(now)
	createdTimeSpan, err := models.StartTimer(userId, &now)
	if err != nil {
		t.Error(err)
	}
	request := &timer.StopTimerRequest{
		UserId:     userId,
		StopTime:   nowTimestamp,
		TimeSpanId: createdTimeSpan.Id,
	}
	timerService := &TimerService{}
	response, err := timerService.StopTimer(context.Background(), request)
	if err != nil {
		t.Error(err)
	}
	timeSpan, err := models.GetTimeSpanById(response.GetTimeSpanId())
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, timeSpan.Id, response.GetTimeSpanId(), "StopTime returned id does not match the given one")
	assert.Equal(t, int32(timeSpan.StopTime.Nanosecond()), nowTimestamp.GetNanos(), "StopTime returned id does not match the given one")
}
