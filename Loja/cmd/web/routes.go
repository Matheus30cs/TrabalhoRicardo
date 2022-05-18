package main

import "net/http"

func (app *application) routes() *http.ServeMux {
  mux := http.NewServeMux()
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/Produtos", app.Produtos)
  mux.HandleFunc("/Produtos/Fogao", app.Fogao)
  mux.HandleFunc("/Produtos/Geladeira", app.Geladeira)
  mux.HandleFunc("/Produtos/Lavadora", app.Lavadora)
  mux.HandleFunc("/Produtos/Microondas", app.Microondas)
  mux.HandleFunc("/snippet", app.showSnippet)
  mux.HandleFunc("/snippet/create", app.createSnippet)
  fileServer := http.FileServer(http.Dir("./ui/static/"))
  mux.Handle("/static/",http.StripPrefix("/static", fileServer))

  return mux
}