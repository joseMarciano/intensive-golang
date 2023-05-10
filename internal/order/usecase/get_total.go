package usecase

import "github.com/joseMarciano/intensive-golang/internal/order/entity"

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	GetTotalRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{GetTotalRepository: orderRepository}
}

func (useCase *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {

	total, err := useCase.GetTotalRepository.GetTotal()
	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDTO{
		Total: total,
	}, nil
}
