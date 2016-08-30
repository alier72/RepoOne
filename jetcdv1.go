package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
//    "os"
    "log"
)

type Stats struct {
   Action string
   Node struct {
    Key string
    Dir bool
    Nodes []struct {
     Key string
     Value string
     ModifiedIndex int
     CreatedIndex int
   }
  ModifiedIndex int
  CreatedIndex int
}
}

func main() {
file, err := ioutil.ReadFile("etcd.json")
    if err != nil {
        log.Fatal(err)
    }
    var jetcd Stats
    err = json.Unmarshal(file, &jetcd)
    if err != nil {
       fmt.Println("Error reading JSON data:", err)
       return
        }


  err=json.Unmarshal(file, &jetcd)
  fmt.Println(jetcd)
  for s:=range jetcd.Node.Nodes {
  fmt.Print(jetcd.Node.Nodes[s].Key)
  fmt.Println()}
}
