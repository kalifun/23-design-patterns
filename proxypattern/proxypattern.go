package proxypattern

type IUser interface {
	Login(username, password string) error
}

type User struct{}

func (u *User) Login(username, password string) error {
	// 省略
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxy(user *User) *UserProxy {
	return &UserProxy{
		user: user,
	}
}

func (up *UserProxy) Login(username, password string) error {
	// 省略
	if err := up.user.Login(username, password); err != nil {
		return err
	}
	return nil
}
