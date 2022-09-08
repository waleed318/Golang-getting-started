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

type TempBook struct {
	name   string
	author string
	pbdate string
}
type BookList struct {
	Items []TempBook
}

func (box *BookList) AddItem(item TempBook) []TempBook {
	box.Items = append(box.Items, item)
	return box.Items
}
