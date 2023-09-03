package interfaces

type IResponseServer interface {
	FromJson(data *[]byte) bool
}
