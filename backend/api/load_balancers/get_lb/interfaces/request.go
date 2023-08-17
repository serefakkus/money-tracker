package interfaces

type IRequestClient interface {
	FromJson(data *[]byte) bool
}

type IRequestServer interface {
	ToJson() []byte
}
