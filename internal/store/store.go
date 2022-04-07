package store

import (
	"animeProject/internal/products"
	"context"
)

type Store interface {
	Create(ctx context.Context, manga *products.Manga) error
	All(ctx context.Context) ([]*products.Manga, error)
	ByID(ctx context.Context, id int) (*products.Manga, error)
	Update(ctx context.Context, manga *products.Manga) error
	Delete(ctx context.Context, id int) error
}

// НА БУДУЩЕЕ
/*type Store interface {
	Clothes() ClothesRepository
	Figurines() FigurinesRepository
	Manga() MangaRepository
}
type ClothesRepository interface {
	Create(ctx context.Context, clothes *products.Clothes) error
	All(ctx context.Context) ([]*products.Clothes, error)
	ByID(ctx context.Context, id int) (*products.Clothes, error)
	Update(ctx context.Context, clothes *products.Clothes) error
	Delete(ctx context.Context, id int) error
}

type FigurinesRepository interface {
	Create(ctx context.Context, figurines *products.Figurines) error
	All(ctx context.Context) ([]*products.Figurines, error)
	ByID(ctx context.Context, id int) (*products.Figurines, error)
	Update(ctx context.Context, figurines *products.Figurines) error
	Delete(ctx context.Context, id int) error
}

type MangaRepository interface {
	Create(ctx context.Context, manga *products.Manga) error
	All(ctx context.Context) ([]*products.Manga, error)
	ByID(ctx context.Context, id int) (*products.Manga, error)
	Update(ctx context.Context, manga *products.Manga) error
	Delete(ctx context.Context, id int) error
}*/
