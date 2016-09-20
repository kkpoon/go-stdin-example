package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type JsonData struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func main() {
	s := bufio.NewScanner(os.Stdin)

	var data JsonData
	for s.Scan() {
		err := json.Unmarshal([]byte(s.Text()), &data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s,%d\n", data.Name, data.Age)
	}
}
