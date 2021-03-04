package model

type Example struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}

func (Example) TableName() string {
	return "public.example"
}
