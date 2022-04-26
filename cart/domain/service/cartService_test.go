package service

// import (
// 	"cartService/domain/model"
// 	"cartService/internal/error"
// 	"reflect"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// import (
// 	"cartService/domain/model"
// 	"cartService/internal/error"
// 	"reflect"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// import (
// 	"cartService/domain/model"
// 	"cartService/internal/error"
// 	"reflect"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// import (
// 	"cartService/domain/model"
// 	"cartService/internal/error"
// 	mocks "cartService/mock"
// 	"reflect"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"
// )

// func TestCartHandler_AddToCart(t *testing.T) {

// 	customer_id := "1"
// 	new_customer_id := "12"

// 	products := model.Product{
// 		ProductId: "1",
// 		Quantity:  1,
// 	}

// 	old_cart := model.Cart{
// 		CustomerId: customer_id,
// 		Products:   []model.Product{products},
// 	}

// 	new_cart := model.Cart{
// 		CustomerId: new_customer_id,
// 		Products:   []model.Product{products},
// 	}

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(cartRepositoryDB *mocks.MockCartRepositoryDB)
// 		checkResponse func(t *testing.T, expected interface{}, actual interface{})
// 	}{
// 		{
// 			name: "SuccessRead&Create",
// 			buildStubs: func(CartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				CartRepositoryDB.EXPECT().
// 					Read(new_customer_id).
// 					Times(1).
// 					Return(nil, nil)

// 				CartRepositoryDB.EXPECT().
// 					Create(new_cart).
// 					Times(1).
// 					Return(nil)
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, true, reflect.ValueOf(actual).IsNil())
// 			},
// 		},
// 		{
// 			name: "FailureCreate",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					Read(new_customer_id).
// 					Times(1).
// 					Return(&new_cart, nil)

// 				cartRepositoryDB.EXPECT().
// 					Create(&new_cart).
// 					Times(0).
// 					Return(error.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, error.NewUnexpectedError(""), actual)
// 			},
// 		},
// 		{
// 			name: "SuccessRead&Update",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					Read(customer_id).
// 					Times(1).
// 					Return(&old_cart, nil)

// 				cartRepositoryDB.EXPECT().
// 					UpdateExisting(old_cart).
// 					Times(1).
// 					Return(error.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, error.NewUnexpectedError(""), actual)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			cartRepositoryDB := mocks.NewMockCartRepositoryDB(ctrl)
// 			tc.buildStubs(cartRepositoryDB)

// 			cartService := CartServiceImpl(cartRepositoryDB)

// 			err := cartService.AddToCart(&old_cart, customer_id)
// 			tc.checkResponse(t, nil, err)
// 		})
// 	}
// }

// func TestCartService_GetCartByCustomerId(t *testing.T) {

// 	customer_id := "1"
// 	invalid_customer_id := "0"

// 	products := model.Product{
// 		ProductId: "1",
// 		Quantity:  1,
// 	}

// 	old_cart := model.Cart{
// 		CustomerId: customer_id,
// 		Products:   []model.Product{products},
// 	}

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(cartRepositoryDB *mocks.MockCartRepositoryDB)
// 		checkResponse func(t *testing.T, expected interface{}, actual interface{})
// 	}{
// 		{
// 			name: "SuccessGet",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					Read(customer_id).
// 					Times(1).
// 					Return(&old_cart, nil)
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, old_cart, actual)
// 			},
// 		},
// 		{
// 			name: "FailureGet",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					Read(invalid_customer_id).
// 					Times(1).
// 					Return(nil, error.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, error.NewUnexpectedError(""), actual)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			cartRepository := mocks.NewMockCartRepositoryDB(ctrl)
// 			tc.buildStubs(cartRepository)

// 			cartService := CartServiceImpl(cartRepository)

// 			cart, err := cartService.GetCart(customer_id)
// 			tc.checkResponse(t, nil, err)
// 			tc.checkResponse(t, old_cart, cart)
// 		})
// 	}
// }

// func TestCartService_UpdateCart(t *testing.T) {

// 	customer_id := "1"
// 	new_customer_id := "12"

// 	products := model.Product{
// 		ProductId: "1",
// 		Quantity:  1,
// 	}

// 	old_cart := model.Cart{
// 		CustomerId: customer_id,
// 		Products:   []model.Product{products},
// 	}

// 	empty_cart := model.Cart{
// 		Products: make([]model.Product, 0),
// 	}

// 	new_cart := model.Cart{
// 		CustomerId: new_customer_id,
// 		Products:   []model.Product{products},
// 	}

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(cartRepositoryDB *mocks.MockCartRepositoryDB)
// 		checkResponse func(t *testing.T, expected interface{}, actual interface{})
// 	}{
// 		{
// 			name: "SuccessUpdate",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					UpdateExisting(old_cart).
// 					Times(1).
// 					Return(nil)
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, true, reflect.ValueOf(actual).IsNil())
// 			},
// 		},
// 		{
// 			name: "FailureUpdate",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					UpdateExisting(new_cart).
// 					Times(1).
// 					Return(error.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, error.NewUnexpectedError(""), actual)
// 			},
// 		},
// 		{
// 			name: "FailuerNoCustomerId",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					UpdateExisting(empty_cart).
// 					Times(1).
// 					Return(error.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, error.NewUnexpectedError(""), actual)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			cartRepository := mocks.NewMockCartRepositoryDB(ctrl)
// 			tc.buildStubs(cartRepository)

// 			cartService := CartServiceImpl(cartRepository)

// 			err := cartService.UpdateCart(&new_cart, customer_id)
// 			tc.checkResponse(t, nil, err)
// 		})
// 	}
// }

// func TestCartService_DeleteCartByCustomerId(t *testing.T) {

// 	customer_id := "1"
// 	invalid_customer_id := "0"

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(cartRepositoryDB *mocks.MockCartRepositoryDB)
// 		checkResponse func(t *testing.T, expected interface{}, actual interface{})
// 	}{
// 		{
// 			name: "SuccessDelete",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					Delete(customer_id).
// 					Times(1).
// 					Return(nil)
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, true, reflect.ValueOf(actual).IsNil())
// 			},
// 		},
// 		{
// 			name: "FailureDelete",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					Delete(invalid_customer_id).
// 					Times(1).
// 					Return(error.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, error.NewUnexpectedError(""), actual)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			cartRepository := mocks.NewMockCartRepositoryDB(ctrl)
// 			tc.buildStubs(cartRepository)

// 			cartService := CartServiceImpl(cartRepository)

// 			err := cartService.DeleteCart(customer_id)
// 			tc.checkResponse(t, nil, err)
// 		})
// 	}
// }

// func TestCartService_DeleteCartItem(t *testing.T) {

// 	customer_id := "1"
// 	invalid_customer_id := "0"
// 	product_id := "2"

// 	products := make([]model.Product, 0)
// 	products = append(products, model.Product{
// 		ProductId: "1",
// 		Quantity:  1,
// 	})
// 	products = append(products, model.Product{
// 		ProductId: "2",
// 		Quantity:  1,
// 	})

// 	old_cart := model.Cart{
// 		CustomerId: customer_id,
// 		Products:   products,
// 	}

// 	new_products := make([]model.Product, 0)
// 	new_products = append(new_products, model.Product{
// 		ProductId: "1",
// 		Quantity:  1,
// 	})

// 	new_cart := model.Cart{
// 		CustomerId: customer_id,
// 		Products:   new_products,
// 	}

// 	invalid_cart := model.Cart{
// 		CustomerId: invalid_customer_id,
// 		Products:   new_products,
// 	}

// 	testCases := []struct {
// 		name          string
// 		buildStubs    func(cartRepositoryDB *mocks.MockCartRepositoryDB)
// 		checkResponse func(t *testing.T, expected interface{}, actual interface{})
// 	}{
// 		{
// 			name: "SuccessDelete",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					Read(customer_id).
// 					Times(1).
// 					Return(&old_cart)

// 				cartRepositoryDB.EXPECT().
// 					UpdateExisting(new_cart).
// 					Times(1).
// 					Return(nil)
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, true, reflect.ValueOf(actual).IsNil())
// 			},
// 		},
// 		{
// 			name: "FailureDelete",
// 			buildStubs: func(cartRepositoryDB *mocks.MockCartRepositoryDB) {
// 				cartRepositoryDB.EXPECT().
// 					Read(customer_id).
// 					Times(1).
// 					Return(&old_cart)

// 				cartRepositoryDB.EXPECT().
// 					UpdateExisting(invalid_cart).
// 					Times(1).
// 					Return(error.NewUnexpectedError(""))
// 			},
// 			checkResponse: func(t *testing.T, expected interface{}, actual interface{}) {
// 				require.Equal(t, error.NewUnexpectedError(""), actual)
// 			},
// 		},
// 	}

// 	for i := range testCases {
// 		tc := testCases[i]

// 		t.Run(tc.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			cartRepository := mocks.NewMockCartRepositoryDB(ctrl)
// 			tc.buildStubs(cartRepository)

// 			cartService := CartServiceImpl(cartRepository)

// 			err := cartService.DeleteCartItem(customer_id, product_id)
// 			tc.checkResponse(t, nil, err)
// 		})
// 	}
// }
