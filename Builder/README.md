# 建造者模式

创建者模式又叫建造者模式，是将一个复杂的对象的构建与它的表示分离，使
得同样的构建过程可以创建不同的表示。创建者模式隐藏了复杂对象的创建过程，它把复杂对象的创建过程加以抽象，通过子类继承或者重载的方式，动态的创建具有复合属性的对象。

## 优点和缺点

### 优点

1.封装性好，构建和表示分离。 2.扩展性好，各个具体的建造者相互独立，有利于系统的解耦。 3.客户端不必知道产品内部组成的细节，建造者可以对创建过程逐步细化，而不对其它模块产生任何影响，便于控制细节风险。

### 缺点

1.产品的组成部分必须相同，这限制了其使用范围。 2.如果产品的内部变化复杂，如果产品内部发生变化，则建造者也要同步修改，后期维护成本较大。

### 应用场景

1.隔离复杂对象的创建和使用，相同的方法，不同执行顺序，产生不同事件结果.  
2.多个部件都可以装配到一个对象中，但产生的运行结果不相同.  
3.产品类非常复杂或者产品类因为调用顺序不同而产生不同作用.  
4.初始化一个对象时，参数过多，或者很多参数具有默认值.  
5.Builder 模式不适合创建差异性很大的产品类产品内部变化复杂，会导致需要定义很多具体建造者类实现变化，增加项目中类的数量，增加系统的理解难度和运行成本.  
6.需要生成的产品对象有复杂的内部结构，这些产品对象具备共性.

## 例子

我们需要建造汽车和摩托车。

### 定义方法

```go
// 建造方法
type Builder interface {
	setTires()
	setCylinder()
	getMotorVehicles() MotorVehicles
}

func newMotorVehicles(mt MotoType) Builder {
	if mt == Car {
		return newCar()
	}

	if mt == Moto {
		return newMoto()
	}
	return nil
}

```

因为非机动车都会有建造轮子和发动机。因此其他非机动车我们都按照这套模式走。

### 汽车

```go
// 汽车工厂
type car struct {
	Tires    int
	Cylinder int
}

func newCar() *car {
	return &car{}
}

func (c *car) setTires() {
	c.Tires = 4
}

func (c *car) setCylinder() {
	c.Cylinder = 4
}

func (c *car) getMotorVehicles() MotorVehicles {
	return MotorVehicles{
		Tires:    c.Tires,
		Cylinder: c.Cylinder,
	}
}
```

### 摩托

```go
// 摩托工厂
type moto struct {
	Tires    int
	Cylinder int
}

func newMoto() *moto {
	return &moto{}
}

func (m *moto) setTires() {
	m.Tires = 2
}

func (m *moto) setCylinder() {
	m.Cylinder = 1
}

func (m *moto) getMotorVehicles() MotorVehicles {
	return MotorVehicles{
		Tires:    m.Tires,
		Cylinder: m.Cylinder,
	}
}
```

### 包工头

```go
type Worker struct {
	b Builder
}

func newWork(w Builder) Worker {
	return Worker{
		b: w,
	}
}

func (w *Worker) setWorker(b Builder) {
	w.b = b
}

func (w *Worker) buildMotorVehicles() MotorVehicles {
	w.b.setCylinder()
	w.b.setTires()
	return w.b.getMotorVehicles()
}
```

### 其他

``` go
const (
	Car  MotoType = "car"
	Moto MotoType = "moto"
)

type MotorVehicles struct {
	Tires    int
	Cylinder int
}
```


### 调用

```go
fun main() {
	car := newMotorVehicles(Car)
	worker := newWork(car)
	worker.buildMotorVehicles()
}
```