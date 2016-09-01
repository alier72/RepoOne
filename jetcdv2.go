package main

import (
//      "fmt"
        "github.com/antonholmquist/jason"
        "log"
        "io/ioutil"
)

func main() {

//exampleJSON := `{ "name": "Walter White","age": 51, "children": [ "junior","holly"],"other": {"occupation": "chemist","years": 23}}`

file, err := ioutil.ReadFile("etcd.json")
    if err != nil {
        log.Fatal(err)
    }
  v, _ := jason.NewObjectFromBytes([]byte(file))

  action, _ := v.GetString("action")
  //mindex, _ := v.GetNumber("node","nodes","modifiedIndex")
  //nkey, _ := v.GetString("node","nodes", "key")
  dirv, _ := v.GetBoolean("node", "dir")

  log.Println("action:", action)
  narray, _ := v.GetObjectArray("node","nodes")
  for _, narray := range narray {
  mindex, _ := narray.GetNumber("modifiedIndex")
  log.Println("Modified Index:", mindex)
  nkey, _ := narray.GetString("key")
  log.Println("Node Key:", nkey)
  }
  log.Println("Dir:", dirv)
}
