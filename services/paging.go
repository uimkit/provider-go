package services

type Filter map[string]interface{}

type PageQuery struct {
	Filter *Filter `json:"filter,omitempty"`
	Page   int64   `json:"page,omitempty"`
	Limit  int64   `json:"limit,omitempty"`
	Sort   []*Sort `json:"sort,omitempty"`
}

type Sort struct {
	Name      string        `json:"name,omitempty"`
	Direction SortDirection `json:"direction,omitempty"`
}

type SortDirection string

const (
	SortDirectionASC  SortDirection = "ASC"
	SortDirectionDESC SortDirection = "DESC"
)
