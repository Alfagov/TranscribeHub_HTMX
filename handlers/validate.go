package handlers

import (
	"TranscribeHub_HTMX/database"
	"TranscribeHub_HTMX/models"
	"github.com/gin-gonic/gin"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[\w-]+(\.[\w-]+)*@([\w-]+\.)+[a-zA-Z]{2,7}$`)

func ValidateEmailHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fieldId := c.Param("fieldId")
		formField := &models.UserFormField{
			Id:          fieldId,
			LabelText:   "Your email",
			Type:        "email",
			Name:        "email",
			QueryURI:    "/validate/email",
			Placeholder: "name@mail.com",
		}

		email := c.PostForm("email")
		if email == "" {
			formField.Error = "Email is required"
			c.HTML(200, "errorFormField", formField)
			return
		}

		if !emailRegex.MatchString(email) {
			formField.Value = email
			formField.Error = "Invalid Email"
			c.HTML(200, "errorFormField", formField)
			return
		}

		formField.Value = email
		c.HTML(200, "validFormField", formField)
	}
}

func ValidatePasswordHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fieldId := c.Param("fieldId")
		formField := &models.UserFormField{
			Id:          fieldId,
			LabelText:   "Your password",
			Type:        "password",
			Name:        "password",
			QueryURI:    "/validate/password",
			Placeholder: "••••••••",
		}

		password := c.PostForm("password")
		if password == "" {
			formField.Error = "Password is required"
			c.HTML(200, "errorFormField", formField)
			return
		}

		lengthCheck := len(password) >= 8

		letterCheck, err := regexp.MatchString(`[a-zA-Z]`, password)
		if err != nil {
			formField.Error = "Internal server error"
			formField.Value = password
			c.HTML(200, "errorFormField", formField)
			return
		}

		numberCheck, err := regexp.MatchString(`[0-9]`, password)
		if err != nil {
			formField.Error = "Internal server error"
			formField.Value = password
			c.HTML(200, "errorFormField", formField)
			return
		}

		if !lengthCheck || !letterCheck || !numberCheck {
			formField.Value = password
			formField.Error = "Password must be at least 8 characters long and contain at least one letter and one number"
			c.HTML(200, "errorFormField", formField)
			return
		}

		formField.Value = password
		c.HTML(200, "validFormField", formField)

	}
}

func ValidateUsernameHandler(db database.Dao) gin.HandlerFunc {
	return func(c *gin.Context) {
		fieldId := c.Param("fieldId")
		formField := &models.UserFormField{
			Id:          fieldId,
			LabelText:   "Username",
			Type:        "text",
			Name:        "username",
			QueryURI:    "/validate/username",
			Placeholder: "Username",
		}

		username := c.PostForm("username")
		if username == "" {
			formField.Error = "Username is required"
			c.HTML(200, "errorFormField", formField)
			return
		}

		if len(username) < 3 {
			formField.Value = username
			formField.Error = "Username must be at least 3 characters long"
			c.HTML(200, "errorFormField", formField)
			return
		}

		if len(username) > 20 {
			formField.Value = username
			formField.Error = "Username must be at most 20 characters long"
			c.HTML(200, "errorFormField", formField)
			return
		}

		taken, err := db.UsernameExists(username)
		if err != nil {
			formField.Error = "Internal server error"
			formField.Value = username
			c.HTML(200, "errorFormField", formField)
			return
		}

		if taken {
			formField.Value = username
			formField.Error = "Username is already taken"
			c.HTML(200, "errorFormField", formField)
			return
		}

		formField.Value = username
		c.HTML(200, "validFormField", formField)
	}
}
