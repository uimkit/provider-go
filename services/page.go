package services

type PageExtra struct {
	Page  int32 `json:"page,omitempty"`  // 页码，从1开始
	Limit int32 `json:"limit,omitempty"` // 翻页大小
	Total int64 `json:"total,omitempty"` // 数据总量
}

type Page[T any] struct {
	Extra *PageExtra `json:"extra,omitempty"` // 查询信息
	Items []*T       `json:"items,omitempty"` // 查询结果数据
}

type ListExtra struct {
	HasMore    bool   `json:"has_more,omitempty"`    // 是否有更多数据
	Limit      int32  `json:"limit,omitempty"`       // 翻页大小
	NextCursor string `json:"next_cursor,omitempty"` // 查询游标
}

type List[T any] struct {
	Extra *ListExtra `json:"extra,omitempty"` // 查询信息
	Items []*T       `json:"items,omitempty"` // 查询结果数据
}
