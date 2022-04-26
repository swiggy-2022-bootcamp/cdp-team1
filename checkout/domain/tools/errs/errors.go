package errs

import (
	"net/http"

	"github.com/pkg/errors"
)

//AppError ..
type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

//Error ..
func (e AppError) Error() error {
	return errors.New(e.Message)
}

//AsMessage ..
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

//NewNotFoundError ..
func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

//NewUnexpectedError ..
func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

//NewValidationError ..
func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

//NewBadRequest ..
func NewBadRequest(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusBadRequest,
	}
}

/*
//fmt.Println(paymentModesDTOData)
		//conn, err := grpc.Dial("localhost:9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
		//if err != nil {
		//	log.Fatalf("Connection failed : %v", err)
		//}
		//defer func(conn *grpc.ClientConn) {
		//	err := conn.Close()
		//	if err != nil {
		//
		//	}
		//}(conn)
		//c := protos.NewPaymentClient(conn)
		//paymentRequest := &protos.PaymentRequest{
		//	UserId:  "bb912edc-50d9-42d7-b7a1-9ce66d459abcd",
		//	Amount:  1,
		//	OrderId: "OA-123",
		//	PaymentMode: &protos.PaymentMode{
		//		Mode:       paymentModesDTOData.Mode,
		//		CardNumber: int64(paymentModesDTOData.CardNumber),
		//	},
		//}
		////fmt.Println("hello")
		//result, err := c.CompletePayment(ctx, paymentRequest)
		//if err != nil {
		//	ctx.JSON(http.StatusInternalServerError, err)
		//	log.Printf("Payment failed : %v", err)
		//} else {
		//	//fmt.Println(result.IsPaymentSuccessful)
		//	ctx.JSON(http.StatusAccepted, result)
		//	//log.Printf("%s - Payment Successful", result.GetIsPaymentSuccessful())
		//}
*/
