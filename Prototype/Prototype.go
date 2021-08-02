package prototype

type University struct {
	Name  string
	Add   string
	Level string
}

func (u *University) Clone() *University {
	return &University{
		Name:  u.Name,
		Add:   u.Add,
		Level: u.Level,
	}
}
