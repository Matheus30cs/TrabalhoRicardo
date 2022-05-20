package mysql

import (
	"database/sql"
	"main/pkg/models"
)

type ProdutoModel struct{
  DB *sql.DB
}

func(m *ProdutoModel)Insert(Nome string, CEP string, Preco float64) (int, error){
  stmt := `INSERT INTO produtos (Nome, CEP, Preco) VALUES (?, ?, ?)`

  result, err := m.DB.Exec(stmt, Nome, CEP, Preco)
  if err != nil{
    return 0, err
  }

  ID, err := result.LastInsertId()
  
  if err != nil{
    return 0, err
  }

  return int(ID), nil
}

func(m *ProdutoModel) Get(ID int) (*models.Produto, error){

  stmt := `SELECT ID, Nome, CEP, Preco FROM produtos
  WHERE Preco > 0 AND ID = ?`
  row := m.DB.QueryRow(stmt, ID)

  s := &models.Produto{}

  err := row.Scan(&s.ID, &s.Nome, &s.CEP, &s.Preco)
  
  if err == sql.ErrNoRows{
    return nil, models.ErrNoRecord
  } else if err != nil{
    return nil, err
  }
  
  return s, nil
}
func(m *ProdutoModel) Latest()([]*models.Snippet, error){
  stmt := `SELECT ID, Nome, CEP, Preco FROM produtos
   WHERE Preco > 0 ORDER BY created DESC LIMIT 10`

  rows, err := m.DB.Query(stmt)
  if err != nil{
    return nil, err
  }
  defer rows.Close()

  produtos := []*models.Snippet{}
  for rows.Next(){
    s := &models.Snippet{}
    err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil{
      return nil, err
    }
    produtos = append(produtos, s)
  }
  err = rows.Err()
  if err != nil{
    return nil, err
  }
  return produtos, nil
}