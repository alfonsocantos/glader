package glader

type Glader interface {
	Get(string) interface{}
	List() []string
	Add(string, interface{})
	Delete(string) error
}
