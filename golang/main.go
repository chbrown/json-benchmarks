package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "io"
  "os"
  // "errors"
)

type Tweet struct {
  Delete   interface{} `json:"delete"`
  Entities struct {
    Hashtags     []interface{} `json:"hashtags"`
    Symbols      []interface{} `json:"symbols"`
    Urls         []interface{} `json:"urls"`
    UserMentions []interface{} `json:"user_mentions"`
  } `json:"entities"`
  // Hashtags
  // }
  // Value map[string]interface{}
  // Entites
  // Inserted          pq.NullTime `db:"-"`
}

// Nested (with anonymous struct)

// type Frame struct {
//     Type    string
//     Value struct {
//         Imagedata string `json:"image_data"`
//     }
// }
// Seperate structs

// type Frame struct {
//     Type    string
//     Value   value
// }

// type value struct {
//     Imagedata string `json:"image_data"`
// }

func main() {
  file, err := os.Open("/Users/chbrown/Desktop/sample-all-2014-03-16.json")
  defer file.Close()

  if err != nil {
    panic(err)
  }

  ndeletes := 0
  nentities := 0
  colors := map[string]int{}

  reader := bufio.NewReader(file)
  for {
    line, _, err := reader.ReadLine()
    if err == io.EOF {
      break
    }
    // fmt.Sprintf("%s: %s", key, val)
    // fmt.Printf("%s\n", line)
    var doc Tweet
    json.Unmarshal(line, &doc)
    // if _, ok := doc["delete"]; ok {
    //   ndeletes++
    // }
    // fmt.Println(doc)
    if doc.Delete != nil {
      ndeletes++
      // fmt.Println("--delete--")
    } else {
      // entities, _ := doc["entities"]
      // entities := doc.Entities
      // fmt.Println("entities", )
      // fmt.Println()

      nentities += len(doc.Entities.Hashtags) +
        len(doc.Entities.Symbols) +
        len(doc.Entities.Urls) +
        len(doc.Entities.UserMentions)
    }
  }

  fmt.Printf("%d deletes\n", ndeletes)
  fmt.Printf("%d entities\n", nentities)
  fmt.Println("colors: %s", colors)
}

// scanner := bufio.NewScanner(file)
// for scanner.Scan() {
//   fmt.Println(scanner.Text())
// }
// if err := scanner.Err(); err != nil {
//   panic(err)
// }
