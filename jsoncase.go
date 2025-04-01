package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	caseFlag := flag.Bool("upper", false, "Convert keys to uppercase")
	flag.Usage = func() {
		fmt.Println("Usage: jsoncase [flags] <input.json>")
		flag.PrintDefaults()
	}
	flag.Parse()

	fmt.Println("Case Flag:", *caseFlag)
	fmt.Println("Arguments:", flag.Args())

	if len(os.Args) < 1 || len(flag.Args()) < 1 {
		fmt.Println("Usage: jsoncase [flags] <input.json>")
		return
	}

	fmt.Println(flag.Arg(0))
	inputFile := flag.Arg(0)
	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	var jsonData map[string]any
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling Json:", err)
		return
	}

	caseJson := convertKeysCase(jsonData, *caseFlag)

	modifiedData, err := json.MarshalIndent(caseJson, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling Json:", err)
		return
	}

	fmt.Println(string(modifiedData))
}

func convertKeysCase(data any, caseFlag bool) any {
	switch v := data.(type) {
	case map[string]any:
		caseMap := make(map[string]any)
		for key, value := range v {
			if caseFlag {
				caseMap[strings.ToUpper(key)] = convertKeysCase(value, caseFlag)
			} else {
				caseMap[strings.ToLower(key)] = convertKeysCase(value, caseFlag)
			}
		}
		return caseMap
	case []any:
		for i, item := range v {
			v[i] = convertKeysCase(item, caseFlag)
		}
		return v
	default:
		return v
	}
}
