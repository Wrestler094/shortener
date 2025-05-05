package postgres

import "errors"

// ErrURLAlreadyExists возвращается при попытке сохранить URL,
// который уже существует в хранилище
var ErrURLAlreadyExists = errors.New("url already exists")
