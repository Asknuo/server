package hotSearch

import "sever/model/other"

type Source interface {
	GetHotSearchData(maxNum int) (HotSearchData other.HotSearchData, err error)
}

func NewSource(sourceStr string) Source {
	switch sourceStr {
	case "baidu":
		return &Baidu{}
	case "kuaishou":
		return &Kuaishou{}
	case "toutiao":
		return &Toutiao{}
	case "zhihu":
		return &Zhihu{}
	default:
		return nil
	}
}
