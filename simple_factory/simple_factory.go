package simplefactory

// 交通工具接口
type TheTrafficTools interface {
	SetTraffic(name string)
	SetPart(p Part)
}

// 零件结构
type Part struct {
	Name string
	Num  uint64
}

type traffic struct {
	Name  string // 交通工具名称
	Parts []Part // 拥有的零件
}

func (t *traffic) SetTraffic(name string) {
	t.Name = name
}

func (t *traffic) SetPart(p Part) {
	t.Parts = append(t.Parts, p)
}

// 汽车
type Car struct {
	traffic
}

// 自行车
type Bicycle struct {
	traffic
}

func NewCar(name string) TheTrafficTools {
	return &Car{
		traffic: traffic{
			Name:  name,
			Parts: []Part{},
		},
	}
}

func NewBicycle(name string) TheTrafficTools {
	return &Bicycle{
		traffic: traffic{
			Name:  name,
			Parts: []Part{},
		},
	}
}
