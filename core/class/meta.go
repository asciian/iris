package class

type meta struct {
	name string
}

func (m *meta) ToString() string {
	return m.name
}

func (*meta) Parents() []Class {
	return []Class{Object}
}

func (*meta) Class() Class {
	return BuiltInClass
}

func (m *meta) Value() Value {
	return m
}

func (m *meta) New(value ...Value) Instance {
	return &builtin{[]Class{value[0].(Class)}, value[1].(string)}
}

func (m *meta) IsInstanceOf(class Class) bool {
	if m.Class() == class {
		return true
	}
	for _, p := range m.Class().Parents() {
		if test(p, class) {
			return true
		}
	}
	return false
}

var BuiltInClass = &meta{"<built-in-class>"}
var StandardClass = &meta{"<standard-class>"}