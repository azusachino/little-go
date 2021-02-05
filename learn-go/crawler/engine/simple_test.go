package engine

import (
	"github.com/little-go/learn-go/crawler/types"
	"github.com/little-go/learn-go/crawler/zhenai/parser"
	"testing"
)

func TestSimpleEngine_Run(t *testing.T) {
	SimpleEngine{}.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
