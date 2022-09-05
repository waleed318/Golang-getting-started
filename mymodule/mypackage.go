package mymodule

type BookdataStore struct {
	name   string
	author string
	pbdate string
}

func (e *BookdataStore) SetValues(name_ string, auth string, date string) string {

	e.name = name_
	e.author = auth
	e.pbdate = date
	return ""
}

func (e *BookdataStore) GetName() string {
	return e.name
}
func (e *BookdataStore) Getauthor() string {
	return e.author
}
func (e *BookdataStore) Getpbdate() string {
	return e.pbdate
}
