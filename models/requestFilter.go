package models

type RequestFilter struct {
	Filter     map[string]interface{} `json:"filter"`
	Projection map[string]int         `json:"projection"`
	Sort       map[string]int         `json:"sort"`
	Skip       int64                  `json:"skip"`
	Limit      int64                  `json:"limit"`
}
