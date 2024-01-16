package api

import (
	"log"

	"github.com/danish45007/hex-architecture-golang/internal/ports"
)

type Application struct {
	arith ArithmeticPort
	db    ports.DBPort
}

func NewApplication(db ports.DBPort, arith ArithmeticPort) *Application {
	return &Application{db: db, arith: arith}
}

func (application Application) GetAddition(a, b int32) (int32, error) {
	result, err := application.arith.Addition(a, b)

	if err != nil {
		return 0, nil
	}
	err = application.db.AddToHistory(result, "addition")
	if err != nil {
		log.Fatalf("Failed to store operation result into database %v", err)
		return 0, err
	}
	return result, nil
}

func (application Application) GetSubtraction(a, b int32) (int32, error) {
	result, err := application.arith.Subtraction(a, b)
	if err != nil {
		return 0, nil
	}
	err = application.db.AddToHistory(result, "subtraction")
	if err != nil {
		log.Fatalf("Failed to store operation result into database %v", err)
		return 0, err
	}
	return result, nil
}

func (application Application) GetMultiplication(a, b int32) (int32, error) {
	result, err := application.arith.Multiplication(a, b)
	if err != nil {
		return 0, nil
	}
	err = application.db.AddToHistory(result, "multiplication")
	if err != nil {
		log.Fatalf("Failed to store operation result into database %v", err)
		return 0, err
	}
	return result, nil
}

func (application Application) GetDivision(a, b int32) (int32, error) {
	result, err := application.arith.Division(a, b)
	if err != nil {
		return 0, nil
	}
	err = application.db.AddToHistory(result, "division")
	if err != nil {
		log.Fatalf("Failed to store operation result into database %v", err)
		return 0, err
	}
	return result, nil
}
