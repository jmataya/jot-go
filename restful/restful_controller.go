package restful

import "net/url"

const (
	LIST    = "LIST"
	SHOW    = "SHOW"
	CREATE  = "CREATE"
	UPDATE  = "UPDATE"
	DESTROY = "DESTROY"
)

type RestfulController interface {
	List(values url.Values, params map[string]string) (int, interface{})
	Show(values url.Values, params map[string]string) (int, interface{})
	Create(values url.Values, params map[string]string) (int, interface{})
	Update(values url.Values, params map[string]string) (int, interface{})
	Destroy(values url.Values, params map[string]string) (int, interface{})
}

type (
	ListNotSupported    struct{}
	ShowNotSupported    struct{}
	CreateNotSupported  struct{}
	UpdateNotSupported  struct{}
	DestroyNotSupported struct{}
)

func (ListNotSupported) List(values url.Values, params map[string]string) (int, interface{}) {
	return 405, nil
}

func (ShowNotSupported) Show(values url.Values, params map[string]string) (int, interface{}) {
	return 405, nil
}

func (CreateNotSupported) Create(values url.Values, params map[string]string) (int, interface{}) {
	return 405, nil
}

func (UpdateNotSupported) Update(values url.Values, params map[string]string) (int, interface{}) {
	return 405, nil
}

func (DestroyNotSupported) Destroy(values url.Values, params map[string]string) (int, interface{}) {
	return 405, nil
}
