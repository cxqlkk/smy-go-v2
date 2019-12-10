package vo

type Page struct {
	PageNo   int
	PageSize int //大写
}

type PageOption func(p *Page)

func PageBuild(op ...PageOption) {
	pg := &Page{}

	for _, o := range op {
		o(pg)
	}

}
// pageSize
func PageSize(pageSize int) PageOption {
	return func(p *Page) {
		p.PageSize = pageSize
	}
}

//PageNo
func PageNo(pageNo int)PageOption{
	return func(p *Page) {
		p.PageNo=pageNo
	}
}
