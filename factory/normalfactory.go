package factory

//
// 工厂模式
//
// 简单工厂中增加一种解析格式时，工厂本身需要修改
// 不符合开闭原则，所以讲工厂抽象出来形成接口，
// 若添加新工厂是只要实现一个接口就可以，而不用
// 动工厂
type ParseNormFactory interface {
	CreateParser() Parse
}

type XmlParseNormFactory struct {
}

func (xpf *XmlParseNormFactory) CreateParser() Parse {
	return &XmlParser{}
}

type JsonParseNormFactory struct {
}

func (xpf *JsonParseNormFactory) CreateParser() Parse {
	return &JsonParser{}
}

type HtmlParseNormFactory struct {
}

func (xpf *HtmlParseNormFactory) CreateParser() Parse {
	return &HtmlParser{}
}

//
// 工厂方法虽然心田解析方式不用动工厂，但是具体使用哪种工厂被耦合到业务了，
// 因此为了从业务中解耦出来，定义一个简单工厂，它用来处理工厂逻辑，返回
// 工厂
func SimpleFactoryFactory(t ParseType) ParseNormFactory {
	if t == JSON {
		return &XmlParseNormFactory{}
	} else if t == XML {
		return &JsonParseNormFactory{}
	} else {
		return &HtmlParseNormFactory{}
	}
}
