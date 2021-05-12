package api

type NewUserRequest struct {
	Username  string `json:"username"`
	Password  string `'json:"password"`
	Email     string `json:"email"`
	FirstName string `json:"fname"`
	Surname   string `json:"surname"`
}

type UserRequest struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"fname"`
	Surname   string `json:"surname"`
}

// import "time"

// type User struct {
// 	ID            int       `json:"id"`
// 	CreatedAt     time.Time `json:"created_at"`
// 	UpdatedAt     time.Time `json:"updated_at"`
// 	Name          string    `json:"name"`
// 	Age           int       `json:"age"`
// 	Height        int       `json:"height"`
// 	Sex           string    `json:"sex"`
// 	ActivityLevel int       `json:"activity_level"`
// 	WeightGoal    string    `json:"weight_goal"`
// 	Email         string    `json:"email"`

// 	gorm.Model
// 	Email    string `json:"email"`
// 	UserName string `json:"username"`
// 	password []byte
// }

// type UserProfile struct {
// 	ID            int       `json:"id"`
// 	CreatedAt     time.Time `json:"created_at"`
// 	UpdatedAt     time.Time `json:"updated_at"`
// 	Name          string    `json:"name"`
// 	Age           int       `json:"age"`
// 	Height        int       `json:"height"`
// 	Sex           string    `json:"sex"`
// 	ActivityLevel int       `json:"activity_level"`
// 	WeightGoal    string    `json:"weight_goal"`
// 	Email         string    `json:"email"`
// }
