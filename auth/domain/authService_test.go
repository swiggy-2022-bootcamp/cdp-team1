package domain_test

import (
	"authService/domain"
	"authService/errs"
	mocks2 "authService/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewAuthService(t *testing.T) {

	authSvc := domain.NewAuthService(nil, nil, nil)

	expAuthSvc := &domain.DefaultAuthService{
		AuthDB:  nil,
		AdminDB: nil,
		CustDB:  nil,
	}

	assert.Equal(t, expAuthSvc, authSvc)
}

func TestDefaultAuthService_HealthCheck(t *testing.T) {

	mockAuthRepo := mocks2.AuthRepositoryDB{}
	mockAdminRepo := mocks2.UserRepositoryDB{}
	mockCustRepo := mocks2.UserRepositoryDB{}
	mockErr := errs.NewUnexpectedError("mock error")

	type mockStruct struct {
		authDB  mocks2.AuthRepositoryDB
		adminDB mocks2.UserRepositoryDB
		custDB  mocks2.UserRepositoryDB
	}
	type testStruct struct {
		name        string
		mocks       mockStruct
		mockReturns []*errs.AppError
		want        *errs.AppError
	}

	tests := []testStruct{
		{
			name: "All Services Healthy",
			mocks: mockStruct{
				authDB:  mockAuthRepo,
				adminDB: mockAdminRepo,
				custDB:  mockCustRepo,
			},
			mockReturns: []*errs.AppError{nil, nil, nil},
			want:        nil,
		},
		{
			name: "Auth Repo Unhealthy",
			mocks: mockStruct{
				authDB:  mockAuthRepo,
				adminDB: mockAdminRepo,
				custDB:  mockCustRepo,
			},
			mockReturns: []*errs.AppError{mockErr, nil, nil},
			want:        mockErr,
		},
		{
			name: "Admin Repo Unhealthy",
			mocks: mockStruct{
				authDB:  mockAuthRepo,
				adminDB: mockAdminRepo,
				custDB:  mockCustRepo,
			},
			mockReturns: []*errs.AppError{nil, mockErr, nil},
			want:        mockErr,
		},
		{
			name: "Customer Repo Unhealthy",
			mocks: mockStruct{
				authDB:  mockAuthRepo,
				adminDB: mockAdminRepo,
				custDB:  mockCustRepo,
			},
			mockReturns: []*errs.AppError{nil, nil, mockErr},
			want:        mockErr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authSvc := domain.NewAuthService(&tt.mocks.authDB, &tt.mocks.adminDB, &tt.mocks.custDB)
			tt.mocks.authDB.On("RepoHealthCheck").Return(tt.mockReturns[0])
			tt.mocks.adminDB.On("RepoHealthCheck").Return(tt.mockReturns[1])
			tt.mocks.custDB.On("RepoHealthCheck").Return(tt.mockReturns[2])

			err := authSvc.HealthCheck()

			if tt.want == nil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
				assert.IsType(t, tt.want, err)
			}
		})
	}
}

func TestDefaultAuthService_Login(t *testing.T) {

	t.Run("Successful Admin Login", func(t *testing.T) {

		mockAuthRepo := mocks2.AuthRepositoryDB{}
		mockAdminRepo := mocks2.UserRepositoryDB{}
		mockCustRepo := mocks2.UserRepositoryDB{}
		authSvc := domain.NewAuthService(&mockAuthRepo, &mockAdminRepo, &mockCustRepo)

		username := "ayan59dutta"
		email := ""
		password := "pass@123!"
		mockUser := &domain.User{
			UserID:   "12345",
			Username: username,
			Password: "$2a$14$jB4Ie5PNm5ajFTEEefp/K.ZFyD/pmuUrdLYVCOc7awrFQ/AFi6NHS",
		}
		mockAuthRepo.On("SaveSession", mock.Anything).Return(nil)
		mockAdminRepo.On("FetchByUsername", username).Return(mockUser, nil)

		token, err := authSvc.Login(username, email, password)

		mockAuthRepo.AssertNumberOfCalls(t, "SaveSession", 1)
		mockAdminRepo.AssertNumberOfCalls(t, "FetchByUsername", 1)
		mockAdminRepo.AssertNumberOfCalls(t, "FetchByEmail", 0)
		mockCustRepo.AssertNumberOfCalls(t, "FetchByUsername", 0)
		mockCustRepo.AssertNumberOfCalls(t, "FetchByEmail", 0)
		assert.NotEmpty(t, token)
		assert.Nil(t, err)
	})

	t.Run("Successful User Login", func(t *testing.T) {
		mockAuthRepo := mocks2.AuthRepositoryDB{}
		mockAdminRepo := mocks2.UserRepositoryDB{}
		mockCustRepo := mocks2.UserRepositoryDB{}
		authSvc := domain.NewAuthService(&mockAuthRepo, &mockAdminRepo, &mockCustRepo)

		username := ""
		email := "ayan59dutta@gmail.com"
		password := "pass@123!"
		mockUser := &domain.User{
			UserID:   "12345",
			Email:    email,
			Password: "$2a$14$jB4Ie5PNm5ajFTEEefp/K.ZFyD/pmuUrdLYVCOc7awrFQ/AFi6NHS",
		}
		mockAuthRepo.On("SaveSession", mock.Anything).Return(nil)
		mockCustRepo.On("FetchByEmail", email).Return(mockUser, nil)

		token, err := authSvc.Login(username, email, password)

		mockAuthRepo.AssertNumberOfCalls(t, "SaveSession", 1)
		mockAdminRepo.AssertNumberOfCalls(t, "FetchByUsername", 0)
		mockAdminRepo.AssertNumberOfCalls(t, "FetchByEmail", 0)
		mockCustRepo.AssertNumberOfCalls(t, "FetchByUsername", 0)
		mockCustRepo.AssertNumberOfCalls(t, "FetchByEmail", 1)
		assert.NotEmpty(t, token)
		assert.Nil(t, err)
	})
}

func TestDefaultAuthService_Logout(t *testing.T) {

	t.Run("Successful Logout", func(t *testing.T) {

		mockAuthRepo := mocks2.AuthRepositoryDB{}
		mockAdminRepo := mocks2.UserRepositoryDB{}
		mockCustRepo := mocks2.UserRepositoryDB{}
		authSvc := domain.NewAuthService(&mockAuthRepo, &mockAdminRepo, &mockCustRepo)
		token, _ := authSvc.GenerateJWT("mock-id", "mock-role")

		mockAuthRepo.On("UpdateSessionStatusByTokenSign", mock.Anything).Return(nil)

		err := authSvc.Logout(token)

		mockAuthRepo.AssertNumberOfCalls(t, "UpdateSessionStatusByTokenSign", 1)
		assert.Nil(t, err)
	})
}

func TestDefaultAuthService_CreateSession(t *testing.T) {
	type fields struct {
		AuthDB  domain.AuthRepositoryDB
		AdminDB domain.UserRepositoryDB
		CustDB  domain.UserRepositoryDB
	}
	type args struct {
		tokenSign string
		userID    string
		role      string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *errs.AppError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := domain.DefaultAuthService{
				AuthDB:  tt.fields.AuthDB,
				AdminDB: tt.fields.AdminDB,
				CustDB:  tt.fields.CustDB,
			}
			assert.Equalf(t, tt.want, d.CreateSession(tt.args.tokenSign, tt.args.userID, tt.args.role), "CreateSession(%v, %v, %v)", tt.args.tokenSign, tt.args.userID, tt.args.role)
		})
	}
}

func TestDefaultAuthService_GenerateJWT(t *testing.T) {
	type fields struct {
		AuthDB  domain.AuthRepositoryDB
		AdminDB domain.UserRepositoryDB
		CustDB  domain.UserRepositoryDB
	}
	type args struct {
		userID string
		role   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  *errs.AppError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := domain.DefaultAuthService{
				AuthDB:  tt.fields.AuthDB,
				AdminDB: tt.fields.AdminDB,
				CustDB:  tt.fields.CustDB,
			}
			got, got1 := d.GenerateJWT(tt.args.userID, tt.args.role)
			assert.Equalf(t, tt.want, got, "GenerateJWT(%v, %v)", tt.args.userID, tt.args.role)
			assert.Equalf(t, tt.want1, got1, "GenerateJWT(%v, %v)", tt.args.userID, tt.args.role)
		})
	}
}

func TestDefaultAuthService_ExtractTokenSign(t *testing.T) {
	type fields struct {
		AuthDB  domain.AuthRepositoryDB
		AdminDB domain.UserRepositoryDB
		CustDB  domain.UserRepositoryDB
	}
	type args struct {
		token string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := domain.DefaultAuthService{
				AuthDB:  tt.fields.AuthDB,
				AdminDB: tt.fields.AdminDB,
				CustDB:  tt.fields.CustDB,
			}
			assert.Equalf(t, tt.want, d.ExtractTokenSign(tt.args.token), "ExtractTokenSign(%v)", tt.args.token)
		})
	}
}

func TestDefaultAuthService_ParseAuthToken(t *testing.T) {

	mockAuthRepo := mocks2.AuthRepositoryDB{}
	mockAdminRepo := mocks2.UserRepositoryDB{}
	mockCustRepo := mocks2.UserRepositoryDB{}
	authSvc := domain.NewAuthService(&mockAuthRepo, &mockAdminRepo, &mockCustRepo)

	t.Run("Valid Token", func(t *testing.T) {

		mockUserID := "mock-id"
		mockRole := "mock-role"
		token, _ := authSvc.GenerateJWT(mockUserID, mockRole)

		userId, role, err := authSvc.ParseAuthToken(token)

		assert.Equal(t, mockUserID, userId)
		assert.Equal(t, mockRole, role)
		assert.Nil(t, err)
	})
}

func TestDefaultAuthService_VerifyAuthToken(t *testing.T) {

	t.Run("Valid Token", func(t *testing.T) {

		mockAuthRepo := mocks2.AuthRepositoryDB{}
		mockAdminRepo := mocks2.UserRepositoryDB{}
		mockCustRepo := mocks2.UserRepositoryDB{}
		authSvc := domain.NewAuthService(&mockAuthRepo, &mockAdminRepo, &mockCustRepo)

		mockUserID := "mock-id"
		mockRole := "mock-role"
		token, _ := authSvc.GenerateJWT(mockUserID, mockRole)
		mockSession := &domain.Session{
			IsActive: true,
		}

		mockAuthRepo.On("FetchSessionByTokenSign", mock.Anything).Return(mockSession, nil)
		mockAuthRepo.On("UpdateSessionStatusByTokenSign", mock.Anything).Return(nil)

		isValid, err := authSvc.VerifyAuthToken(token, mockUserID, mockRole)

		mockAuthRepo.AssertNumberOfCalls(t, "FetchSessionByTokenSign", 1)
		mockAuthRepo.AssertNumberOfCalls(t, "UpdateSessionStatusByTokenSign", 0)
		assert.True(t, isValid)
		assert.Nil(t, err)
	})

	t.Run("Inactive Session", func(t *testing.T) {

		mockAuthRepo := mocks2.AuthRepositoryDB{}
		mockAdminRepo := mocks2.UserRepositoryDB{}
		mockCustRepo := mocks2.UserRepositoryDB{}
		authSvc := domain.NewAuthService(&mockAuthRepo, &mockAdminRepo, &mockCustRepo)
		wantErr := errs.NewAuthenticationError("inactive session")

		mockUserID := "mock-id"
		mockRole := "mock-role"
		token, _ := authSvc.GenerateJWT(mockUserID, mockRole)
		mockSession := &domain.Session{
			IsActive: false,
		}

		mockAuthRepo.On("FetchSessionByTokenSign", mock.Anything).Return(mockSession, nil)
		mockAuthRepo.On("UpdateSessionStatusByTokenSign", mock.Anything).Return(nil)

		isValid, err := authSvc.VerifyAuthToken(token, mockUserID, mockRole)

		mockAuthRepo.AssertNumberOfCalls(t, "FetchSessionByTokenSign", 1)
		mockAuthRepo.AssertNumberOfCalls(t, "UpdateSessionStatusByTokenSign", 0)
		assert.False(t, isValid)
		assert.EqualValues(t, wantErr, err)
	})

	//t.Run("Expired Token", func(t *testing.T) {
	//
	//	mockAuthRepo := mocks.AuthRepositoryDB{}
	//	mockAdminRepo := mocks.UserRepositoryDB{}
	//	mockCustRepo := mocks.UserRepositoryDB{}
	//	authSvc := domain.NewAuthService(&mockAuthRepo, &mockAdminRepo, &mockCustRepo)
	//	wantErr := errs.NewAuthenticationError("expired session")
	//
	//	mockUserID := "mock-id"
	//	mockRole := "mock-role"
	//	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoibW9jay1pZCIsInJvbGUiOiJtb2NrLXJvbGUiLCJleHAiOjEyNTc4OTQwNjB9.w0C93xE8rfqYsAeiotoCZq7MQqULQa1_BqGGxzHvl3c"
	//	mockSession := &domain.Session{
	//		IsActive: true,
	//	}
	//
	//	mockAuthRepo.On("FetchSessionByTokenSign", mock.Anything).Return(mockSession, nil)
	//	mockAuthRepo.On("UpdateSessionStatusByTokenSign", mock.Anything).Return(nil)
	//
	//	isValid, err := authSvc.VerifyAuthToken(token, mockUserID, mockRole)
	//
	//	mockAuthRepo.AssertNumberOfCalls(t, "FetchSessionByTokenSign", 1)
	//	mockAuthRepo.AssertNumberOfCalls(t, "UpdateSessionStatusByTokenSign", 1)
	//	assert.False(t, isValid)
	//	assert.EqualValues(t, wantErr, err)
	//})
}
