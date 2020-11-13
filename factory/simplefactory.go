package factory

// 单例支持类型
type ParseType string

const (
	JSON ParseType = "json"
	XML            = "xml"
	HTML           = "html"
)

// 工厂反方应该是一类实现了接口的对象
// 从工厂得到的应该是多态的对象
type Parse interface {
	Serialize()
}

// json 解析
type JsonParser struct {
}

func (jp *JsonParser) Serialize() {
	panic("not implemented") // TODO: Implement
}

// html 解析
type HtmlParser struct {
}

func (h *HtmlParser) Serialize() {
	panic("not implemented") // TODO: Implement
}

// xml 解析
type XmlParser struct {
}

func (x *XmlParser) Serialize() {
	panic("not implemented") // TODO: Implement
}

// 工厂方法
// 传递名字得到不同的Parser
func ParseFactory(t ParseType) Parse {
	switch t {
	case JSON:
		return &JsonParser{}

	case XML:
		return &XmlParser{}

	case HTML:
		return &HtmlParser{}

	default:
		return nil
	}
}
