package usecase

import "github.com/joseMarciano/intensive-golang/internal/order/entity"

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FInalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: orderRepository}
}

func (useCase *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)

	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()

	if err != nil {
		return nil, err
	}

	err = useCase.OrderRepository.Save(order)

	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FInalPrice: order.FinalPrice,
	}, nil
}
