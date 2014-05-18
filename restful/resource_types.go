package restful

import "net/url"

type Resource interface {
	Get(values url.Values) (int, interface{})
	Post(values url.Values) (int, interface{})
	Put(values url.Values) (int, interface{})
	Delete(values url.Values) (int, interface{})
}

type GetNotSupported struct{}

func (GetNotSupported) Get(values url.Values) (int, interface{}) {
	return 405, nil
}

type PostNotSupported struct{}

func (PostNotSupported) Post(values url.Values) (int, interface{}) {
	return 405, nil
}

type PutNotSupported struct{}

func (PutNotSupported) Put(values url.Values) (int, interface{}) {
	return 405, nil
}

type DeleteNotSupported struct{}

func (DeleteNotSupported) Delete(values url.Values) (int, interface{}) {
	return 405, nil
}
