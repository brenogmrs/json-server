package service

import (
	"errors"
	"fmt"
	"json-server/pkg/providers"
	"json-server/pkg/utils"
	"reflect"
)

type IService interface {
	Create(body map[string]interface{}) map[string]interface{}
	GetAll() []interface{}
	GetById(id int) (interface{}, error)
	Update(id int, body map[string]interface{}) interface{}
	Delete(id int) error
}

type service struct {
	resource string
}

func New(resource string) IService {
	return &service{
		resource: resource,
	}
}

func (s *service) Create(body map[string]interface{}) map[string]interface{} {

	data := providers.SetResources(body, s.resource, nil, "INSERT")

	err := providers.SetJSON(data)
	utils.ErrorChecker(err)

	return body
}

func (s *service) GetAll() []interface{} {

	data := providers.GetDBData()

	return data[s.resource].([]interface{})

}

func (s *service) GetById(id int) (interface{}, error) {

	data := providers.GetDBData()

	chosenResource := data[s.resource].([]interface{})

	var indexes []float64
	for _, value := range chosenResource {
		v, ok := value.(map[string]interface{})

		if !ok {
			fmt.Println("error")
		}

		indexes = append(indexes, v["id"].(float64))
	}

	foundIndex := utils.Search(id, indexes)

	if foundIndex < 0 {
		return nil, errors.New("Not Found")
	}

	return chosenResource[foundIndex], nil
}

func (s *service) Update(id int, body map[string]interface{}) interface{} {

	foundData, err := s.GetById(id)

	if err != nil {
		return err
	}

	keys := reflect.ValueOf(foundData).MapKeys()

	updatedRecord := make(map[string]interface{})

	parseFoundData := foundData.(map[string]interface{})

	for _, key := range keys {
		stringKey := key.String()

		if body[stringKey] != parseFoundData[stringKey] {
			updatedRecord[stringKey] = body[stringKey]
		}

		if body[stringKey] == nil {
			updatedRecord[stringKey] = parseFoundData[stringKey]
		}
	}

	fullDataMap := providers.SetResources(updatedRecord, s.resource, id, "UPDATE")

	err = providers.SetJSON(fullDataMap)
	utils.ErrorChecker(err)

	return updatedRecord
}

func (s *service) Delete(id int) error {

	if _, err := s.GetById(id); err != nil {
		return err
	}

	fullDataMap := providers.SetResources(nil, s.resource, id, "DELETE")

	err := providers.SetJSON(fullDataMap)
	utils.ErrorChecker(err)

	return nil
}
