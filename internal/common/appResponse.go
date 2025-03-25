package common

type successRes struct {
	Data       interface{} `json:"data"`
	Paging     interface{} `json:"paging,omitempty"`
	TotalDocs  interface{} `json:"total_docs,omitempty"`
	TotalPages interface{} `json:"total_pages,omitempty"`
	Filter     interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, totalPages, total, filter interface{}) *successRes {
	return &successRes{Data: data, Paging: paging, TotalPages: totalPages, TotalDocs: total, Filter: filter}
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{Data: data, Paging: nil, Filter: nil, TotalDocs: nil}
}
