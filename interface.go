package main

type Inter interface {
	get() error
	gets() interface{}
	post() error
	put() error
	delete() error
}

func getI(i Inter) error {
	return i.get()
}
func postI(i Inter) error {
	return i.post()
}
func getsI(i Inter) interface{} {
	return i.gets()
}
func putI(i Inter) error {
	return i.put()
}
func deleteI(i Inter) error {
	return i.delete()
}
