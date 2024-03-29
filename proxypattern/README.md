# 代理模式

> 为其他对象提供一种代理以控制对这个对象的访问。

## 优点和缺点

### 优点

1.职责清晰；高扩展性；智能化

### 缺点

1.由于在客户端和真实主题之间增加了代理对象，因此有些类型的代理模式可能会造成请求的处理速度变慢。
2.实现代理模式需要额外的工作，有些代理模式的实现非常复杂。

## Demo
```go
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

```