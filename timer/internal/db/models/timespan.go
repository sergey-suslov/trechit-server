package models

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	tspb "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/sergey-suslov/trechit-server/api/timer"
	"github.com/sergey-suslov/trechit-server/timer/internal/db"
	"log"
	"time"
)

type TimeSpan struct {
	Id        int32
	UserId    int32
	StartTime *time.Time
	StopTime  *time.Time
	CreatedAt *time.Time
}

func StartTimer(userId int32, startTime *time.Time) (*TimeSpan, error) {
	var timeSpan TimeSpan
	err := db.Pool.QueryRow(context.Background(), `
		insert into timespans values(default, $1, $2, null, default) returning id, userid, starttime, stoptime, created;
	`, userId, *startTime).Scan(&timeSpan.Id, &timeSpan.UserId, &timeSpan.StartTime, &timeSpan.StopTime, &timeSpan.CreatedAt)
	if err != nil {
		log.Println("Error creating time span", err)
		return nil, err
	}

	return &timeSpan, nil
}

func (timeSpan *TimeSpan) ToProtoTimeSpan() (*timer.TimeSpan, error) {
	start, err := ptypes.TimestampProto(*timeSpan.StartTime)
	if err != nil {
		return nil, err
	}

	var end *tspb.Timestamp
	if timeSpan.StopTime != nil {
		end, err = ptypes.TimestampProto(*timeSpan.StopTime)
		if err != nil {
			return nil, err
		}
	}

	var created *tspb.Timestamp
	if timeSpan.StopTime != nil {
		created, err = ptypes.TimestampProto(*timeSpan.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &timer.TimeSpan{
		Id:        timeSpan.Id,
		UserId:    timeSpan.UserId,
		StartTime: start,
		StopTime:  end,
		CreatedAt: created,
	}, nil
}
