package main

import (
//      "fmt"
        "github.com/antonholmquist/jason"
        "log"
        "io/ioutil"
)

func main() {

//exampleJSON := `{ "name": "Walter White","age": 51, "children": [ "junior","holly"],"other": {"occupation": "chemist","years": 23}}`

file, err := ioutil.ReadFile("data.json")
    if err != nil {
        log.Fatal(err)
    }
  v, _ := jason.NewObjectFromBytes([]byte(file))

  name, _ := v.GetString("name")
  age, _ := v.GetNumber("age")
  occupation, _ := v.GetString("other", "occupation")
  years, _ := v.GetNumber("other", "years")

  log.Println("age:", age)
  log.Println("name:", name)
  log.Println("occupation:", occupation)
  log.Println("years:", years)
}
