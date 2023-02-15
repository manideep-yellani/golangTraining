package service

import (
	"github.com/awesomeProject/internal/models"
	"github.com/awesomeProject/internal/store"
)

type service struct {
	svc store.Storer
}

func New(storer store.Storer) service {
	return service{}
}

func Post(st models.Student) (models.Student, string) {
	if l := len(st.Phone); l != 10 {
		return models.Student{}, "WRONG PHONE GIVEN"

	}
	st = store.Post(st)
	return st, ""

}
