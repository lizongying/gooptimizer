package gooptimizer

type lang uint8

const (
	EN lang = iota
	CN
)

type i18n struct {
	lang    lang
	mapping map[lang]map[string]string
}

var DefaultI18n *i18n

func init() {
	DefaultI18n = &i18n{
		lang: EN,
		mapping: map[lang]map[string]string{
			CN: {
				"Field alignment arrangement before": "字段对齐排列前",
				"Field alignment arrangement after":  "字段对齐顺序排列后",
				"Field":                              "字段",
				"Type":                               "类型",
				"Align":                              "对齐",
				"Size":                               "大小",
				"Expect Size":                        "期望大小",
				"Actual Size":                        "实际大小",
				"Bytes":                              "字节",
				"You should optimize the structure; there's potential to save": "您应该优化该结构体, 可以节约",
			},
		},
	}
}

func (i *i18n) Get(name string) (value string) {
	mapping, ok := i.mapping[i.lang]
	if !ok {
		value = name
		return
	}
	if value, ok = mapping[name]; !ok {
		value = name
	}
	return
}
