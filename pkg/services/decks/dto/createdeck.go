package dto

type CreateDeckParams struct {
	UserID    int32
	Name      string
	IsPremade bool
}
