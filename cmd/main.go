package main

import (
	"database/sql"
	"github.com/joseMarciano/intensive-golang/internal/order/infra/database"
	"github.com/joseMarciano/intensive-golang/internal/order/usecase"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := database.NewOrderRepository(db)
	calculateFinalPriceUseCase := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    "99",
		Tax:   12,
		Price: 50.33,
	}

	output, err := calculateFinalPriceUseCase.Execute(input)
	if err != nil {
		panic(err)
	}

	println(output)
}
