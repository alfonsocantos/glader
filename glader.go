package glader

type Glader interface {
	Get(string) interface{}
	List() interface{}
	Add(interface{}) error
	Delete(string) error
}
