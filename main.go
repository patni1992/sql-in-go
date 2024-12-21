package main

import (
	"database/sql"
	"log"
	"sql-in-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbConnection, err := sql.Open("sqlite3", "./shop.db")

	if err != nil {
		log.Fatal(err)
	}

	defer dbConnection.Close()

	orderRepository := &database.OrderRepository{Db: dbConnection}

	err = orderRepository.CreateTable()
	if err != nil {
		log.Fatal("Error creating orders table:", err)
	}

	err = orderRepository.Insert(database.Order{Product: "Laptop", Amount: 10})
	if err != nil {
		log.Fatal("Error inserting order", err)
	}

	err = orderRepository.Insert(database.Order{Product: "Keyboard", Amount: 50})
	if err != nil {
		log.Fatal("Error inserting order", err)
	}

	orders, err := orderRepository.GetAll()
	if err != nil {
		log.Fatal(">Error getting orders", err)
	}

	log.Println(orders)

	order, err := orderRepository.GetById(orders[0].Id)
	if err != nil {
		log.Fatal("Error getting order", err)
	}

	order.Amount = 1500
	err = orderRepository.Update(order)
	if err != nil {
		log.Fatal("Error updating order", err)
	}

	orders, err = orderRepository.GetAll()
	if err != nil {
		log.Fatal(">Error getting orders", err)
	}

	log.Println(orders)

	err = orderRepository.Delete(order.Id)
	if err != nil {
		log.Fatal("Error deleting order", err)
	}

	orders, err = orderRepository.GetAll()
	if err != nil {
		log.Fatal(">Error getting orders", err)
	}

	log.Println(orders)
}
