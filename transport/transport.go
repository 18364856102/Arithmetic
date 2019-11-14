package transport

import (
	"ArithmeticOperation/model"
	"ArithmeticOperation/service"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

func MakeArithmeticEndpoint(svc service.Calculate) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {

		req, ok := request.(model.ArithmeticRequest)
		if !ok {
			return nil, errors.New("type transform error")
		}

		var (
			res      int
			calError string
		)

		switch req.RequestType {
		case "add":
			res = svc.Add(req)
		case "subtract":
			res = svc.Subtract(req)
		case "multiply":
			res = svc.Multiply(req)
		case "divide":
			res, calError = svc.Divide(req)
		default:
			return nil, model.ErrInvalidRequestType
		}

		return model.ArithmeticResponse{
			Result: res,
			Error:  calError,
		}, nil

	}
}

// decodeArithmeticRequest decode request params to struct
func DecodeArithmeticRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request model.ArithmeticRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// encodeArithmeticResponse encode response to return
func EncodeArithmeticResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8") //返回数据格式是json
	w.Header().Set("Access-Control-Allow-Origin", "*")               //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")   //header的类型

	return json.NewEncoder(w).Encode(response)
}
