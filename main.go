package main

import "github.com/codegangsta/martini"

func HandleIndex() string {
    return "Hello Amsterdam Web Developers"
}

// our main handler

func main() {
  m := martini.Classic()
  m.Get("/", HandleIndex)
  m.Run()
}
