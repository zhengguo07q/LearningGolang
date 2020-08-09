package base

// Class
type Class interface {
	Do()
}

var (
	factoryByName = make(map[string]func() Class)
)

// Register 注册一个类生成工厂
func Register(name string, factory func() Class) {
	factoryByName[name] = factory
}

// Create 根据名称创建对应的类
func Create(name string) Class {
	if f, ok := factoryByName[name]; ok {
		return f()
	} else {
		panic("name not found")
	}
}
