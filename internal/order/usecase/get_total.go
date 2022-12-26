package usecase

import "github.com/MatheusAbdias/microservices/internal/order/domain"

type GetTotalOutPutDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository domain.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository domain.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutPutDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}

	return &GetTotalOutPutDTO{Total: total}, nil
}
