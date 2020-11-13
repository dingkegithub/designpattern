package factory

var parseInstance map[ParseType]Parse

//
// 结合单例，不用每次初始化
//
func init() {
	parseInstance[JSON] = &JsonParser{}
	parseInstance[XML] = &XmlParser{}
	parseInstance[HTML] = &HtmlParser{}
}

// 工厂方法
// 传递名字得到不同的Parser
func ParseSingletonFactory(t ParseType) Parse {
	if p, ok := parseInstance[t]; ok {
		return p
	}

	return nil
}
