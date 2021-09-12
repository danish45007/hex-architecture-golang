package api

import (
	"log"

	"github.com/danish45007/hex-architecture-golang/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
	db    ports.DBPort
}

func NewAdapter(arith ports.ArithmeticPort, db ports.DBPort) *Adapter {
	return &Adapter{arith: arith, db: db}
}

func (apiAd Adapter) GetAddition(a, b int32) (int32, error) {
	result, err := apiAd.arith.Addition(a, b)

	if err != nil {
		return 0, nil
	}
	err = apiAd.db.AddToHistory(result, "addition")
	if err != nil {
		log.Fatalf("Failed to store operation result into database %v", err)
		return 0, err
	}
	return result, nil
}

func (apiAd Adapter) GetSubtraction(a, b int32) (int32, error) {
	result, err := apiAd.arith.Subtraction(a, b)
	if err != nil {
		return 0, nil
	}
	err = apiAd.db.AddToHistory(result, "subtraction")
	if err != nil {
		log.Fatalf("Failed to store operation result into database %v", err)
		return 0, err
	}
	return result, nil
}

func (apiAd Adapter) GetMultiplication(a, b int32) (int32, error) {
	result, err := apiAd.arith.Multiplication(a, b)
	if err != nil {
		return 0, nil
	}
	err = apiAd.db.AddToHistory(result, "multiplication")
	if err != nil {
		log.Fatalf("Failed to store operation result into database %v", err)
		return 0, err
	}
	return result, nil
}

func (apiAd Adapter) GetDivision(a, b int32) (int32, error) {
	result, err := apiAd.arith.Division(a, b)
	if err != nil {
		return 0, nil
	}
	err = apiAd.db.AddToHistory(result, "division")
	if err != nil {
		log.Fatalf("Failed to store operation result into database %v", err)
		return 0, err
	}
	return result, nil
}
