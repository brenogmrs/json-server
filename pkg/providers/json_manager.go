package providers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"json-server/pkg/utils"
	"os"
)

const (
	INSERT = "INSERT"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

func GetDBData() map[string]interface{} {
	workDir, _ := os.Getwd()

	resourcesFilePath := workDir + "/" + "db.json"

	file, err := os.Open(resourcesFilePath)
	utils.ErrorChecker(err)

	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	utils.ErrorChecker(err)

	var data map[string]any
	json.Unmarshal([]byte(byteValue), &data)

	return data
}

func GetResources(data map[string]any) []string {

	resourcesNames := make([]string, 0, len(data))

	for k := range data {
		resourcesNames = append(resourcesNames, k)
	}

	return resourcesNames

}

func SetResources(body map[string]any, resource string, id any, action string) map[string]interface{} {

	data := GetDBData()
	dbKeys := GetResources(data)
	chosenResource := data[resource].([]interface{})

	fullDataMap := make(map[string]interface{})

	switch action {
	case INSERT:

		body["id"] = len(chosenResource) + 1

		fullDataMap[resource] = append(chosenResource, body)

	case UPDATE:

		intId, ok := id.(int)

		if !ok {
			fmt.Println("Error converting ID to int at UPDATE")
			return nil
		}

		chosenResource[intId-1] = body

		fullDataMap[resource] = chosenResource

	case DELETE:

		intId, ok := id.(int)

		if !ok {
			fmt.Println("Error converting ID to int at DELETE")
			return nil
		}

		dataWithoutIndex := utils.RemoveIndex(chosenResource, intId-1)
		fullDataMap[resource] = dataWithoutIndex

	default:
		fmt.Println("UNKNOWN ACTION")
	}

	for i := 0; i < len(dbKeys); i++ {

		key := dbKeys[i]

		if key != resource {
			fullDataMap[key] = data[key]
		}
	}

	return fullDataMap
}

func SetJSON(data map[string]any) error {
	jsonByteArray, err := json.MarshalIndent(data, "", "\t")

	if err != nil {
		return err
	}

	workDir, _ := os.Getwd()

	resourcesFilePath := workDir + "/" + "db.json"

	err = ioutil.WriteFile(resourcesFilePath, jsonByteArray, 0644)

	if err != nil {
		return err
	}

	return nil
}
