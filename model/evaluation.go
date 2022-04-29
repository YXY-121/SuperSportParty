package model

type Evaluation struct {
	EvaluationId string `json:"evaluation_id"  gorm:"evaluation_id"`
	EvaluatorId  string `json:"evaluator_id" gorm:"evaluator_id"`
	UserId       string `json:"user_id" gorm:"user_id"`
	Content      string `json:"content" gorm:"content"`
	OrderId      string `json:"order_id" gorm:"order_id"`
	Star         int    `json:"star" gorm:"star"`
	CreateTime   int64  `json:"create_time" gorm:"create_time"`
	UpdateTime   int64  `json:"update_time" gorm:"update_time"`
}

func (Evaluation) TableName() string {
	return "evaluation"
}
