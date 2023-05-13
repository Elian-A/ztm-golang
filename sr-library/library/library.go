package library

import "time"

type Library struct {
	Members      []Member
	Books        []Book
	booksChecked []BookCheck
}

type BookCheck struct {
	Member       Member
	Book         Book
	CheckinTime  time.Time
	CheckoutTime time.Time
}

func (l *Library) AddBook(b Book) {
	l.Books = append(l.Books, b)
}

func (l *Library) CheckedBooks(b Book) {
	l.Books = append(l.Books, b)
}

func (l *Library) AddMember(n string) {
	var member Member
	member.Name = n
}
