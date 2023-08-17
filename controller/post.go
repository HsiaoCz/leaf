package controller

// 这里主要是放post请求传递过来的数据

// UserRegisterPost 用户注册请求的数据
type UserRegisterPost struct {
	PhoneNumber string `json:"phone_number" binding:"omitempry"`
	Email       string `json:"email" binding:"email,omitempty"`
	Password    string `json:"password" binding:"required"`
}

// UserLoginPost  用户登录的时候需要的数据
type UserLoginPost struct {
	PhoneNumber string `json:"phone_number" binding:"omitempty"`
	Email       string `json:"email" binding:"email,omitempty"`
}
