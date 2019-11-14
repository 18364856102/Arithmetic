package service

import (
	"ArithmeticOperation/model"
	"github.com/go-kit/kit/log"
)

// Service Define a service interface
type Calculate interface {
	Add(req model.ArithmeticRequest) int
	Subtract(req model.ArithmeticRequest) int
	Multiply(req model.ArithmeticRequest) int
	Divide(req model.ArithmeticRequest) (result int, err string)
}

//ArithmeticService implement Service interface
type ArithmeticService struct {
	Logger log.Logger
}

func New(logger log.Logger) Calculate {
	return &ArithmeticService{
		Logger: logger,
	}
}

// Add implement Add method
func (s *ArithmeticService) Add(req model.ArithmeticRequest) int {
	return req.A + req.B
}

// Subtract implement Subtract method
func (s *ArithmeticService) Subtract(req model.ArithmeticRequest) int {
	return req.A - req.B
}

// Multiply implement Multiply method
func (s *ArithmeticService) Multiply(req model.ArithmeticRequest) int {
	return req.A * req.B
}

// Divide implement Divide method
func (s *ArithmeticService) Divide(req model.ArithmeticRequest) (result int, err string) {
	if req.B == 0 {
		_ = s.Logger.Log("[Divide]", "the dividend can not be zero!")
		return 0, "the dividend can not be zero！"
	}
	return req.A / req.B, ""
}
