package api_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/OpitiCalvin/novelsTradeIn/pkg/api"
)

type mockUserRepo struct{}

func (m mockUserRepo) CreateUser(request api.NewUserRequest) error {
	if request.Username == "alreadyExists" {
		return errors.New("repository - user already exists in database")
	}

	return nil
}

func TestCreateNewUser(t *testing.T) {
	mockRepo := mockUserRepo{}
	mockUserService := api.NewUserService(&mockRepo)

	tests := []struct {
		name    string
		request api.NewUserRequest
		want    error
	}{
		{
			name: "should create new user successfully",
			request: api.NewUserRequest{
				Username:  "testUser",
				Password:  "testUser",
				Email:     "testuser@gmail.com",
				FirstName: "Test",
				Surname:   "User",
			},
			want: nil,
		}, {
			name: "should return an error because of missing email",
			request: api.NewUserRequest{
				Username:  "testuser",
				Password:  "testUser",
				Email:     "",
				FirstName: "Test",
				Surname:   "User",
			},
			want: errors.New("user service - email required"),
		}, {
			name: "should return an error because of missing username",
			request: api.NewUserRequest{
				Username:  "",
				Password:  "testUser",
				Email:     "testuser@gmail.com",
				FirstName: "Test",
				Surname:   "User",
			},
			want: errors.New("user service - username required"),
		}, {
			name: "should return error from database because user already exists",
			request: api.NewUserRequest{
				Username:  "alreadyExists",
				Password:  "testUser",
				Email:     "testuser@gmail.com",
				FirstName: "Test",
				Surname:   "User",
			},
			want: errors.New("repository - user already exists in database"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := mockUserService.New(test.request)

			if !reflect.DeepEqual(err, test.want) {
				t.Errorf("Test: %v failed. Got: %v, wanted: %v", test.name, err, test.want)
			}
		})
	}
}
