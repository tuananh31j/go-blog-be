package common

type successRes struct {
	Data      interface{} `json:"data"`
	Paging    interface{} `json:"paging,omitempty"`
	TotalDocs interface{} `json:"total_docs,omitempty"`
	Filter    interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, total, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, TotalDocs: total, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data, Paging: nil, Filter: nil}
}
