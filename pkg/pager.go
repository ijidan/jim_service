package pkg

type Pager struct {
	Page       uint64      `json:"page,omitempty;query:page"`
	PageSize   uint64      `json:"page_size,omitempty;query:page_size"`
	TotalRows  uint64      `json:"total_rows"`
	TotalPages uint64      `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

func (p *Pager) GetPage() uint64 {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pager) GetPageSize() uint64 {
	if p.PageSize == 0 {
		p.PageSize = 10
	}
	return p.PageSize
}

func (p *Pager) GetTotalRows() uint64 {
	return p.TotalRows
}

func (p *Pager) GetTotalPages() uint64  {
	return p.TotalPages
}

func (p *Pager) GetRows() interface{}  {
	return p.Rows
}