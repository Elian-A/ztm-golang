package library

import (
	"errors"
	"fmt"
	"time"
)

type Member struct {
	Name       string
	BookChecks []BookCheck
}

func (m *Member) CheckoutBook(b string, l Library) error {
	for _, book := range l.Books {
		if book.Title == b {
			newCheck := BookCheck{
				Book:         book,
				CheckoutTime: time.Now(),
			}
			fmt.Printf("Checkin?: %v", newCheck.CheckinTime)

			m.BookChecks = append(m.BookChecks, newCheck)
			return nil
		}
	}
	return errors.New("Book not found")
}

func (m *Member) CheckinBook(title string) error {
	for idx, checks := range m.BookChecks {
		if checks.Book.Title == title {
			m.BookChecks[idx].CheckinTime = time.Now()
		}
	}
	return errors.New("Book hasn't been checkout")
}
