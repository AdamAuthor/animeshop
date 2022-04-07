package inmemory

import (
	"animeProject/internal/products"
	"animeProject/internal/store"
	"context"
	"fmt"
)

type DB struct {
	data map[int]*products.Manga
}

func NewDB() store.Store {
	return &DB{
		data: make(map[int]*products.Manga),
	}
}

func (D *DB) Create(ctx context.Context, manga *products.Manga) error {
	D.data[manga.ID] = manga
	return nil
}

func (D *DB) All(ctx context.Context) ([]*products.Manga, error) {
	mangas := make([]*products.Manga, 0, len(D.data))
	for _, manga := range D.data {
		mangas = append(mangas, manga)
	}
	return mangas, nil
}

func (D *DB) ByID(ctx context.Context, id int) (*products.Manga, error) {
	laptop, ok := D.data[id]
	if !ok {
		return nil, fmt.Errorf("No laptop with id: %d", id)
	}
	return laptop, nil
}

func (D *DB) Update(ctx context.Context, manga *products.Manga) error {
	D.data[manga.ID] = manga
	return nil
}

func (D *DB) Delete(ctx context.Context, id int) error {
	delete(D.data, id)
	return nil
}
