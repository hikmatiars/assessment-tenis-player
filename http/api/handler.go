package api

import(
	"core-project/usecase"
)

type Handler struct {
	Usecase usecase.InterfaceUseCase
}