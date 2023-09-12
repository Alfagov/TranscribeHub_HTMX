package models

type Counter struct {
	Text          string
	Value         int
	Max           int
	ProgressClass string
}

type SubscriptionCounter struct {
	Name     string
	Counters []*Counter
}

func GetDefaultProjectCounter() *Counter {
	return &Counter{
		Text:          "Projects",
		Value:         0,
		Max:           0,
		ProgressClass: "progress-primary",
	}
}

func GetDefaultAssistCounter() *Counter {
	return &Counter{
		Text:          "Assists",
		Value:         0,
		Max:           0,
		ProgressClass: "progress-secondary",
	}
}

func GetDefaultTranscriptionCounter() *Counter {
	return &Counter{
		Text:          "Transcriptions",
		Value:         0,
		Max:           0,
		ProgressClass: "progress-error",
	}
}
