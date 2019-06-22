package main

import (
    "fmt"
    "io"
    "net/http"
    "bytes"
    "strings"
    "io/ioutil"
    "math/rand"
  "time"
    "github.com/julienschmidt/httprouter"
)



func uploadFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  rand.Seed(time.Now().UnixNano())
  chars := []rune("ABCĆDEĘFGHIJKLMNOPQRSTUVWXYZÅĄÄÖabcćdeęfghijklmnopqrstuvwxyzåąäö0123456789")
  length := 3
  var b strings.Builder
  for i := 0; i < length; i++ {
    b.WriteRune(chars[rand.Intn(len(chars))])
  }
  var Buf bytes.Buffer
  file, header, err := r.FormFile("file")
  if err != nil {
    fmt.Printf("Invalid file")
    fmt.Fprintf(w, "Invalid file, :\\")
  }
  defer file.Close()
  name := strings.Split(header.Filename, ".")
  io.Copy(&Buf, file)
  contents := Buf.String()
  fname := b.String()
  err_ := ioutil.WriteFile("img/" + fname + "." + name[1], []byte(contents), 0644)
  if err != nil {
    panic(err_)
  }
  Buf.Reset()
  fmt.Fprintf(w, "{\"file\":\"https://clyx.ml/i/"+ fname+ "." + name[1] + "\"}")
  return
}
func ImageServe(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
  http.ServeFile(w,r,"./img/" +ps.ByName("clap"))
}
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
  http.ServeFile(w,r,"./public/index.html")
}
func Style(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
  http.ServeFile(w,r,"./public/style.css")
}
func setupRoutes() {
  router := httprouter.New()
  
  router.POST("/upload", uploadFile)
  router.GET("/", Index)
  router.GET("/style.css", Style)
  router.GET("/i/:clap", ImageServe)
  
  http.ListenAndServe(":8080", router)
}

func main() {
    fmt.Println("Server started")
    setupRoutes()
}