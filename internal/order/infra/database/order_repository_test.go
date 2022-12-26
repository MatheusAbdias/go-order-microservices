package database

import (
	"database/sql"
	"testing"

	"github.com/MatheusAbdias/microservices/internal/order/domain"
	"github.com/stretchr/testify/suite"

	//sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)

	db.Exec("CREATE TABLE orders (id VARCHAR(255) NOT NULL, price FLOAT NOT NULL, tax FLOAT NOT NULL, final_price FLOAT NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	order, err := domain.NewOrder("123", 100, 10)
	suite.NoError(err)
	suite.NoError(order.CalculatePrice())

	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult domain.Order
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
