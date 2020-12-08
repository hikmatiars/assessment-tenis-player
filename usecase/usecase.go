package usecase

import (
	"assessment-tennis-player/http/response"
	"context"
	"github.com/labstack/gommon/log"
	"math/rand"
)

//Init Use case interface
type InterfaceUseCase interface {
  PutBallToContainerUseCase() []*response.DataContainer
}

type UseCase struct {
	Context context.Context
	Data []*response.DataContainer
	FlagReady int
}

func NewUseCase(ctx context.Context, data []*response.DataContainer, flag int) InterfaceUseCase {
	return &UseCase{
		Context: ctx,
		Data: data,
		FlagReady: flag,
	}
}

func (uc *UseCase) PutBallToContainerUseCase() []*response.DataContainer {
	var (
		setContainer = 5
		ballReady = 5
	)

	//Using Flag if player ready to play
	if uc.FlagReady == 1 {
		return uc.Data
	}

	if len(uc.Data) == 0 {
		//Insert data if container empty
		for i := 1; i <= setContainer; i ++   {
			uc.Data = append(uc.Data, &response.DataContainer{
				ContainerId: i,
				Ball: 0,
				Status: "not ready",
			})
		}

	}else {
		//Random put the ball to container
		n := Random(len(uc.Data))

		ball := uc.Data[n].Ball + 1
		log.Print(ball)
		if ball == ballReady {
			//set to ready
			uc.Data[n].Ball = ball
			uc.Data[n].Status = "ready"
			uc.FlagReady = 1
		}else {
			//put the ball
			uc.Data[n].Ball = ball
		}
	}

	return uc.Data
}

func Random( count int ) int {
	return rand.Int() % count
}