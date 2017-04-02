package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Brands []struct {
	Text  string `json:"text"`
	Value int64  `json:"value"`
}
type Models []struct {
	Text  string `json:"text"`
	Value int64  `json:"value"`
}
type ModelYears []int64
type ModelChasis []struct {
	Text  string `json:"text"`
	Value int64  `json:"value"`
}
type ModelVersions []struct {
	Text  string `json:"text"`
	Value int64  `json:"value"`
}
type Location struct {
	Area     string `json:"area"`
	Location string `json:"location"`
	Picture  string `json:"picture"`
}

func hodo(err error) {
	if err != nil {
		panic(err)
	}
}

func bodyBrands(body []byte) (*Models, error) {
	var s = new(Models)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
func getBrands() {
	res, err := http.Get("https://www.hum.com/bin/core/findPort.json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	brandsresult, err := bodyBrands([]byte(body))

	for i := 0; i < len(*brandsresult); i++ {
		fmt.Println((*brandsresult)[i].Text)
		fmt.Println((*brandsresult)[i].Value)
		getModels((*brandsresult)[i].Value)
	}
}

func bodyModels(body []byte) (*Models, error) {
	var s = new(Models)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
func getModels(brand int64) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	modelsresult, err := bodyModels([]byte(body))

	for i := 0; i < len(*modelsresult); i++ {
		fmt.Println((*modelsresult)[i].Text)
		fmt.Println((*modelsresult)[i].Value)
		getModelYears((*modelsresult)[i].Value)
	}
}

func bodyModelYears(body []byte) (*ModelYears, error) {
	var s = new(ModelYears)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
func getModelYears(brand int64, model int64) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + "." + strconv.FormatInt(model, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	modelyearsresult, err := bodyModelYears([]byte(body))

	for i := 0; i < len(*modelyearsresult); i++ {
		fmt.Println((*modelyearsresult)[i])
		getModelChasis(brand, model, (*modelyearsresult)[i])
	}
}

func bodyModelChasis(body []byte) (*ModelChasis, error) {
	var s = new(ModelChasis)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
func getModelChasis(brand int64, model int64, modelyears int64) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + "." + strconv.FormatInt(model, 10) + "." + strconv.FormatInt(modelyears, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	modelchasisresult, err := bodyModelChasis([]byte(body))

	for i := 0; i < len(*modelchasisresult); i++ {
		fmt.Println((*modelchasisresult)[i].Text)
		fmt.Println((*modelchasisresult)[i].Value)
		getModelVersions(brand, model, modelyears, (*modelchasisresult)[i].Value)
	}
}

func bodyModelVersions(body []byte) (*ModelVersions, error) {
	var s = new(ModelVersions)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
func getModelVersions(brand int64, model int64, modelyears int64, modelchasis int64) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + "." + strconv.FormatInt(model, 10) + "." + strconv.FormatInt(modelyears, 10) + "." + strconv.FormatInt(modelchasis, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	modelversionresult, err := bodyModelVersions([]byte(body))

	for i := 0; i < len(*modelversionresult); i++ {
		fmt.Println((*modelversionresult)[i].Text)
		fmt.Println((*modelversionresult)[i].Value)
		getLocation(brand, model, modelyears, modelchasis, (*modelversionresult)[i].Value)
	}
}

func bodyLocation(body []byte) (*Location, error) {
	var s = new(Location)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}
func getLocation(brand int64, model int64, modelyears int64, modelchasis int64, modelversion int64) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + "." + strconv.FormatInt(model, 10) + "." + strconv.FormatInt(modelyears, 10) + "." + strconv.FormatInt(modelchasis, 10) + "." + strconv.FormatInt(modelversion, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	locationresult, err := bodyLocation([]byte(body))

	fmt.Println(locationresult.Area)
	fmt.Println(locationresult.Location)
	fmt.Println(locationresult.Picture)
}

func main() {
	getBrands()
}
