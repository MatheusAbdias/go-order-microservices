package usecase

import (
	"github.com/MatheusAbdias/microservices/internal/order/domain"
	"github.com/MatheusAbdias/microservices/internal/order/infra/database"
)

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculatePriceUseCase struct {
	OrderRepository domain.OrderRepositoryInterface
}

func NewCaculatePriceUseCase(orderRepository database.OrderRepository) *CalculatePriceUseCase {
	return &CalculatePriceUseCase{
		OrderRepository: &orderRepository,
	}
}

func (c *CalculatePriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := domain.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculatePrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
