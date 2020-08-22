package pagination

// 分页结构体
type Pager struct {
	Page  uint `json:"page" form:"page" validate:"empty|number"`
	Limit uint `json:"limit" form:"limit" validate:"empty|number"`
	Total uint `json:"total" form:"total" validate:"empty"`
}

// 设置当前页
func (p *Pager) SetPage(page uint) *Pager {
	p.Page = page
	return p
}

// 设置当前页面大小
func (p *Pager) SetLimit(limit uint) *Pager {
	p.Limit = limit
	return p
}

// 这只总数据量
func (p *Pager) SetTotal(total uint) *Pager {
	p.Total = total
	return p
}

// 获取分页偏移量
func (p *Pager) GetOffset() (offset uint) {
	p.Format()
	offset = (p.Page - 1) * p.Limit
	return
}

// 格式化分页
func (p *Pager) Format() *Pager {
	if p.Page < 1 {
		p.SetPage(1)
	}
	if p.Limit < 1 {
		p.SetLimit(10)
	}
	if p.Total < 0 {
		p.SetTotal(0)
	}
	return p
}
