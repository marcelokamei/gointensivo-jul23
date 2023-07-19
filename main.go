package main

import (
	"database/sql"
	"fmt"

	"github.com/marcelokamei/gointensivo-jul23/internal/infra/database"
	"github.com/marcelokamei/gointensivo-jul23/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

// type Car struct {
// 	Model string
// 	Color string
// }

// // metodo
// func (c Car) Start() {
// 	println(c.Model + " has been started")
// }

// func (c Car) ChangeColor(color string) {
// 	c.Color = color
// 	println("New color: " + c.Color)
// }

// // funcao
// func soma(x, y int) int {
// 	return x + y
// }

func main() {

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "12345",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	//println(output)
	fmt.Println(output)

	// order, err := entity.NewOrder("1", 10, 1)
	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(order.ID)
	// }
	// println(order.ID)

	/* car := Car{
		Model: "Ferrari",
		Color: "Red",
	}

	println(car.Model)

	car.Model = "Fusca"

	println(car.Model)

	car.Start()
	car.ChangeColor("Blue")

	println(car.Color) */

}
