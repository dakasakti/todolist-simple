package helper

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ParseParameter(id string) (uint, error) {
	req, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return uint(req), nil
}

func ParseMessageErrors(e validator.ValidationErrors) string {
	for _, val := range e {
		switch val.Tag() {
		case "required":
			if val.Field() == "Title" {
				return "title cannot be null"
			}

			if val.Field() == "ActivityID" {
				return "activity_group_id cannot be null"
			}
		case "email":
			return "email must be username@domain.tld"
		}
	}

	return "email cannot be null"
}

func ParseMessageBinding(e error) string {
	var messages string

	splits := strings.Split(e.Error(), ", ")
	for i := 1; i < 4; i++ {
		value := strings.Split(splits[i], "=")
		messages += fmt.Sprintf("%s=%s", value[0], value[len(value)-1])
		if i != 3 {
			messages += " "
		}
	}

	return messages
}

// func GenerateLogin(id uint, password string) entity.Response {
// 	if id <= 0 || password != "member" {
// 		return entity.Response{
// 			Code:    401,
// 			Message: "failed login",
// 			Errors: echo.Map{
// 				"id":       "id must be greater than 0",
// 				"password": "password must be member",
// 			},
// 		}
// 	}

// 	token, err := middlewares.CreateToken(id)
// 	if err != nil {
// 		return entity.Response{
// 			Code:    401,
// 			Message: "failed login",
// 			Errors:  err.Error(),
// 		}
// 	}

// 	return entity.Response{
// 		Code:    200,
// 		Message: "success login",
// 		Data:    token,
// 	}
// }
