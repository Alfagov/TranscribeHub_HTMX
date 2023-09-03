package models

type UserFormField struct {
	FormTemplate string
	Id           string
	LabelText    string
	Type         string
	Name         string
	QueryURI     string
	Placeholder  string
	Value        string
	Error        string
}

var RegisterFields = []UserFormField{
	{
		FormTemplate: "formField",
		Id:           "Username",
		LabelText:    "Username",
		Type:         "text",
		Name:         "username",
		QueryURI:     "/validate/username",
		Placeholder:  "Username",
	},
	{
		FormTemplate: "formField",
		Id:           "Email",
		LabelText:    "Your email",
		Type:         "email",
		Name:         "email",
		QueryURI:     "/validate/email",
		Placeholder:  "name@mail.com",
	},
	{
		FormTemplate: "formField",
		Id:           "Password",
		LabelText:    "Your password",
		Type:         "password",
		Name:         "password",
		QueryURI:     "/validate/password",
		Placeholder:  "••••••••",
	},
}

var LoginFields = []UserFormField{
	{
		FormTemplate: "formField",
		Id:           "Email",
		LabelText:    "Your email",
		Type:         "email",
		Name:         "email",
		QueryURI:     "/validate/email",
		Placeholder:  "name@mail.com",
	},
	{
		FormTemplate: "formFieldNoValidation",
		Id:           "Password",
		LabelText:    "Your password",
		Type:         "password",
		Name:         "password",
		QueryURI:     "/validate/password",
		Placeholder:  "••••••••",
	},
}
