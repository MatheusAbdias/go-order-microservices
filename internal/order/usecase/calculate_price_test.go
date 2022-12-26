package usecase

import (
	"database/sql"
	"testing"

	"github.com/MatheusAbdias/microservices/internal/order/domain"
	"github.com/MatheusAbdias/microservices/internal/order/infra/database"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type CalculatePriceUseCaseTestSuite struct {
	suite.Suite
	OrderRepository database.OrderRepository
	Db              *sql.DB
}

func (suite *CalculatePriceUseCaseTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	db.Exec("CREATE TABLE orders (id VARCHAR(255) NOT NULL, price FLOAT NOT NULL, tax FLOAT NOT NULL, final_price FLOAT NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
	suite.OrderRepository = *database.NewOrderRepository(db)
}

func (suite *CalculatePriceUseCaseTestSuite) TearDownTest() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(CalculatePriceUseCaseTestSuite))
}

func (suite *CalculatePriceUseCaseTestSuite) TestCalculatePrice() {
	order, err := domain.NewOrder("1", 10, 2)
	suite.NoError(err)
	order.CalculatePrice()

	calculatePrinceInput :=
		OrderInputDTO{
			ID:    order.ID,
			Price: order.Price,
			Tax:   order.Tax,
		}
	calculatePriceUseCase := NewCaculatePriceUseCase(suite.OrderRepository)
	output, err := calculatePriceUseCase.Execute(calculatePrinceInput)

	suite.NoError(err)
	suite.Equal(order.ID, output.ID)
	suite.Equal(order.Price, output.Price)
	suite.Equal(order.Tax, output.Tax)
	suite.Equal(order.FinalPrice, output.FinalPrice)

}
