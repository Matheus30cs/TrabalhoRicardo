package main

//go run cmd/web/*
import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

  "main/pkg/models"
)

func(app *application) home (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/"{
    app.notFound(rw)
    return
  }

  snippets, err := app.snippets.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/index.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, snippets)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Produtos (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos"{
    app.notFound(rw)
    return
  }

  snippets, err := app.snippets.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/produtos.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, snippets)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Fogao (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Fogao"{
    app.notFound(rw)
    return
  }

  snippets, err := app.snippets.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Fogao.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, snippets)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Geladeira (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Geladeira"{
    app.notFound(rw)
    return
  }

  snippets, err := app.snippets.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Geladeira.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, snippets)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Lavadora (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Lavadora"{
    app.notFound(rw)
    return
  }

  snippets, err := app.snippets.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Lavadora.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, snippets)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

func(app *application) Microondas (rw http.ResponseWriter, r *http.Request){
  if r.URL.Path != "/Produtos/Microondas"{
    app.notFound(rw)
    return
  }

  snippets, err := app.snippets.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

  files := []string{
    "./ui/html/Microondas.html",
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, snippets)
  if err != nil{
    app.serverError(rw, err)
    return
  }
}

//http://localhost:4000/snippet?id=123
func(app *application) showSnippet(rw http.ResponseWriter, r *http.Request){
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(rw)
    return
  }

  s, err := app.snippets.Get(id)
  if err == models.ErrNoRecord {
    app.notFound(rw)
    return
  } else if err != nil{
    app.serverError(rw, err)
    return
  }
  
  files := []string{
    "./ui/html/show.page.tmpl.html",
    "./ui/html/base.layout.tmpl.html",
    "./ui/html/footer.partial.tmpl.html",   
  }
  ts, err := template.ParseFiles(files...)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  err = ts.Execute(rw, s)
  if err != nil{
    app.serverError(rw, err)
    return
  }
  
}

func(app *application) createSnippet(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    return
  }

  title := "Aula de hoje"
  content := "Tentando lidar com o Banco de Dados"
  expires := "7"

  id, err := app.snippets.Insert(title, content, expires)

  if err != nil{
    app.serverError(rw, err)
    return
  }

  http.Redirect(rw, r, fmt.Sprintf("/snippet?id=%d", id),http.StatusSeeOther)
  //rw.Write([]byte("Criar novo Snippet"))
}
//curl -igp -X POST http://localhost:4000/snippet/create
//curl -i -X GET http://localhost:4000/snippet/create