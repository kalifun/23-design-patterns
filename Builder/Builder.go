package builder

type MotoType string

const (
	Car  MotoType = "car"
	Moto MotoType = "moto"
)

type MotorVehicles struct {
	Tires    int
	Cylinder int
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

// 建造方法
type Builder interface {
	setTires()
	setCylinder()
	getMotorVehicles() MotorVehicles
}

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
