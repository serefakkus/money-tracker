package interfaces

type IResponseClient interface {
	ToJson() []byte
}

type IResponseServer interface {
	FromJson(data *[]byte) bool
}
