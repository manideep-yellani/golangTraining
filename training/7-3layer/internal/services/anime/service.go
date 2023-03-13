package anime

import "training/7-3layer/internal/stores"

type service struct {
	store stores.Animestorer
}

func New(store stores.Animestorer) *service {
	return &service{store}
}
