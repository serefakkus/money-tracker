package interfaces

type ISendSms interface {
	GetPhone() string
	GetCode() string
}
