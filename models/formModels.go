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

type PricingCard struct {
	Title       string
	Description string
	Price       string
	Elements    []PricingCardElement
}

type PricingCardElement struct {
	Text string
}

var PricingCards = []PricingCard{
	{
		Title:       "Starter",
		Description: "Best option for personal use & for your next project.",
		Price:       "$40",
		Elements: []PricingCardElement{
			{
				Text: "Invalid configuration",
			},
			{
				Text: "No setup, or hidden fees",
			},
			{
				Text: "Team size: 5",
			},
		},
	},
	{
		Title:       "Company",
		Description: "Relevant for multiple users, extended & premium support.",
		Price:       "$99",
		Elements: []PricingCardElement{
			{
				Text: "Invalid configuration",
			},
			{
				Text: "No setup, or hidden fees",
			},
			{
				Text: "Team size: 5",
			},
		},
	},
	{
		Title:       "Enterprise",
		Description: "Best for large scale uses and extended redistribution rights.",
		Price:       "$499",
		Elements: []PricingCardElement{
			{
				Text: "Invalid configuration",
			},
			{
				Text: "No setup, or hidden fees",
			},
			{
				Text: "Team size: 5",
			},
		},
	},
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
