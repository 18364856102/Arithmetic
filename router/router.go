package router

import (
	"ArithmeticOperation/service"
	"ArithmeticOperation/transport"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func HandlerArithmeticOperation(r *mux.Router, svc service.Calculate) {
	getArithmeticOperation := kithttp.NewServer(
		transport.MakeArithmeticEndpoint(svc),
		transport.DecodeArithmeticRequest,
		transport.EncodeArithmeticResponse,
	)

	r.Handle("/calculate", getArithmeticOperation)
}
