package exporter

import (
	"encoding/json"
	"fmt"
	"log"
)

func ExportToJson(data []map[string]interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bytes))
}
