package bridge

type MsgSender interface {
	Send(msg string) error
}

type EmailSender struct {
	emails []string
}

type SMSSender struct {
	phones []string
}

func (s *EmailSender) Send(msg string) error {
	// 省略逻辑
	return nil
}

func (s *SMSSender) Send(msg string) error {
	// 省略逻辑
	return nil
}

type NotificationService interface {
	Notify(msg string) error
}

type ErrorEmailNotify struct {
	sender EmailSender
}

func NewErrorNotify(sender EmailSender) *ErrorEmailNotify {
	return &ErrorEmailNotify{sender: sender}
}

func (n *ErrorEmailNotify) Notify(msg string) error {
	return n.sender.Send(msg)
}
