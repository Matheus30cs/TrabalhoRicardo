package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record Found")

type Snippet struct{
  ID int
  Title string
  Content string
  Created time.Time
  Expires time.Time
}

/*type Produto struct{
  ID int
  Nome string
  CEP string
  Preco float64
}*/