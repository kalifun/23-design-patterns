package chain

import "fmt"

// department 
type department interface {
	execute(*patient)
	setNext(department)
}

// patient 
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

// doctor 
type doctor struct {
	next department
}

// execute 
//  @receiver d 
//  @param p 
func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor check already done")
		d.next.execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

// setNext 
//  @receiver d 
//  @param next 
func (d *doctor) setNext(next department) {
	d.next = next
}

// cashier 
type cashier struct {
	next department
}

// execute 
//  @receiver c 
//  @param p 
func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

// setNext 
//  @receiver c 
//  @param next 
func (c *cashier) setNext(next department) {
	c.next = next
}

// medical 
type medical struct {
	next department
}

// execute 
//  @receiver m 
//  @param p 
func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.execute(p)
}

// setNext 
//  @receiver m 
//  @param next 
func (m *medical) setNext(next department) {
	m.next = next
}

// reception 
type reception struct {
    next department
}

// execute 
//  @receiver r 
//  @param p 
func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.registrationDone = true
	r.next.execute(p)
}

// setNext 
//  @receiver r 
//  @param next 
func (r *reception) setNext(next department) {
	r.next = next
}
