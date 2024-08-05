package request

// json buat type application/json, form buat x-www-form-urlencoded/form-data kasih keduanya biar bisa terima json/form-data/x-www-form-urlencoded
// https://docs.gofiber.io/api/ctx#bodyparser
type UserCreateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}
