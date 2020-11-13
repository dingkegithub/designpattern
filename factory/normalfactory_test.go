package factory

import (
	"strings"
	"testing"
)

func TestNormalFactory(t *testing.T) {
	file := "xx.html"

	suffix := strings.Split(file, ".")[1]

	var normalFactory ParseNormFactory

	// 业务中耦合了工厂逻辑判断
	// 新增工厂，业务接着改
	if suffix == "html" {
		normalFactory = &HtmlParseNormFactory{}
	} else if suffix == "json" {
		normalFactory = &JsonParseNormFactory{}
	} else {
		normalFactory = &XmlParseNormFactory{}
	}

	// 多态获得不同工厂的表现形式
	parse := normalFactory.CreateParser()
	parse.Serialize()
}

func TestNormalFactory2(t *testing.T) {
	file := "xx.html"

	suffix := strings.Split(file, ".")[1]

	var normalFactory ParseNormFactory

	// 简单工厂代理后，工厂模式修改不再:
	normalFactory = SimpleFactoryFactory(ParseType(suffix))
	parse := normalFactory.CreateParser()
	parse.Serialize()
}
