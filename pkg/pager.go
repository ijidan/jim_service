package pkg

type Pager struct {
	Code         int32         `json:"code"`
	BusinessCode int32         `json:"business_code"`
	Message      string        `json:"message"`
	List         []interface{} `json:"list"`
	Total        uint32        `json:"total"`
	Page         uint32        `json:"page"`
	PageSize     uint32        `json:"page_size"`
	PageTotal    int32         `json:"page_total"`
}
