package model

type Example struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

func (e Example) TableName() string {
	return "public.example"
}
