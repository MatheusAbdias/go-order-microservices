# go-order-microservices

## Project context

- My first project was with golang, so this project simulates microservice for orders. Imagine you have an e-commerce and receive multiples payments orders, these requests need passing for multiple businesses so for improving these processes you need to build two services.
  The first process is a publisher, only pulled orders request in queue for another service (Consumer) make the all business logic for making this processes complete. In consumer, imagine if you need to calculate a final price for each order based on delivery tax this business logic is calculated by the consumer.

## Tecnologies

- Golang: programming language
- Rabbitmq: Queue manager
- Grafana: Dashboard statistic
- Prometheus: For Monitoring the rabbitmq
- sqlite3: Database

## How running this project

- First need <a href="https://go.dev/dl/"> golong </a> installed, afeter dowload goland need rung

```cmd
go mod tidy
```

- After this, you need <a href= "https://docs.docker.com/get-docker/"> docker </a> to run the docker compose for rabbitmq, grafana, and prometheus.

```
docker compose up -d
```

- Must need to create the orders table in the database so for this run:

```
sqlite3
CREATE TABLE orders (id VARCHAR(255) NOT NULL, price FLOAT NOT NULL, tax FLOAT NOT NULL, final_price FLOAT NOT NULL, PRIMARY KEY (id));
```

- Finally for run the consumer and producer:

```
go run ./cmd/consumer/main.go
go run ./cmd/producer/main.go
```
