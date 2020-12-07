package usecase

import(
	"context"
)


//Init Use case interface
type InterfaceUseCase interface {

}

type UseCase struct {
	Context context.Context
}

func NewUseCase(ctx context.Context) InterfaceUseCase {
	return &UseCase{
		Context: ctx,
	}
}