package restful

import "net/url"

const (
	LIST    = "LIST"
	SHOW    = "SHOW"
	CREATE  = "CREATE"
	UPDATE  = "UPDATE"
	DESTROY = "DESTROY"
)

type RestfulType interface {
	List(values url.Values) (int, interface{})
	Show(values url.Values) (int, interface{})
	Create(values url.Values) (int, interface{})
	Update(values url.Values) (int, interface{})
	Destroy(values url.Values) (int, interface{})
	SetBasePath(path string)
	GetBasePath() string
	ActionMatch(path string, method string) (bool, string)
}

type (
	ListNotSupported    struct{}
	ShowNotSupported    struct{}
	CreateNotSupported  struct{}
	UpdateNotSupported  struct{}
	DestroyNotSupported struct{}
)

func (ListNotSupported) List(values url.Values) (int, interface{}) {
	return 405, nil
}

func (ShowNotSupported) Show(values url.Values) (int, interface{}) {
	return 405, nil
}

func (CreateNotSupported) Create(values url.Values) (int, interface{}) {
	return 405, nil
}

func (UpdateNotSupported) Update(values url.Values) (int, interface{}) {
	return 405, nil
}

func (DestroyNotSupported) Destroy(values url.Values) (int, interface{}) {
	return 405, nil
}
