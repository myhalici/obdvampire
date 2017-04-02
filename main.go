package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"
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
		// fmt.Println((*brandsresult)[i].Text)
		// fmt.Println((*brandsresult)[i].Value)
		getModels((*brandsresult)[i].Value, (*brandsresult)[i].Text)
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
func getModels(brand int64, brandname string) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	modelsresult, err := bodyModels([]byte(body))

	for i := 0; i < len(*modelsresult); i++ {
		// fmt.Println((*modelsresult)[i].Text)
		// fmt.Println((*modelsresult)[i].Value)
		getModelYears(brand, (*modelsresult)[i].Value, brandname, (*modelsresult)[i].Text)
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
func getModelYears(brand int64, model int64, brandname string, modelname string) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + "." + strconv.FormatInt(model, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	modelyearsresult, err := bodyModelYears([]byte(body))

	for i := 0; i < len(*modelyearsresult); i++ {
		// fmt.Println((*modelyearsresult)[i])
		getModelChasis(brand, model, (*modelyearsresult)[i], brandname, modelname)
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
func getModelChasis(brand int64, model int64, modelyears int64, brandname string, modelname string) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + "." + strconv.FormatInt(model, 10) + "." + strconv.FormatInt(modelyears, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	modelchasisresult, err := bodyModelChasis([]byte(body))

	for i := 0; i < len(*modelchasisresult); i++ {
		// fmt.Println((*modelchasisresult)[i].Text)
		// fmt.Println((*modelchasisresult)[i].Value)
		getModelVersions(brand, model, modelyears, (*modelchasisresult)[i].Value, brandname, modelname, (*modelchasisresult)[i].Text)
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
func getModelVersions(brand int64, model int64, modelyears int64, modelchasis int64, brandname string, modelname string, modelchasisname string) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + "." + strconv.FormatInt(model, 10) + "." + strconv.FormatInt(modelyears, 10) + "." + strconv.FormatInt(modelchasis, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	modelversionresult, err := bodyModelVersions([]byte(body))

	for i := 0; i < len(*modelversionresult); i++ {
		// fmt.Println((*modelversionresult)[i].Text)
		// fmt.Println((*modelversionresult)[i].Value)
		getLocation(brand, model, modelyears, modelchasis, (*modelversionresult)[i].Value, brandname, modelname, modelchasisname, (*modelversionresult)[i].Text)
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
func getLocation(brand int64, model int64, modelyears int64, modelchasis int64, modelversion int64, brandname string, modelname string, modelchasisname string, modelversionname string) {
	res, err := http.Get("https://www.hum.com/bin/core/findPort." + strconv.FormatInt(brand, 10) + "." + strconv.FormatInt(model, 10) + "." + strconv.FormatInt(modelyears, 10) + "." + strconv.FormatInt(modelchasis, 10) + "." + strconv.FormatInt(modelversion, 10) + ".json")
	hodo(err)
	body, err := ioutil.ReadAll(res.Body)
	hodo(err)
	locationresult, err := bodyLocation([]byte(body))

	carcounter++
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = brandname
	cell = row.AddCell()
	cell.Value = modelname
	cell = row.AddCell()
	cell.Value = strconv.FormatInt(modelyears, 10)
	cell = row.AddCell()
	cell.Value = modelchasisname
	cell = row.AddCell()
	cell.Value = modelversionname
	cell = row.AddCell()
	cell.Value = locationresult.Location
	cell = row.AddCell()
	cell.Value = locationresult.Picture
	cell = row.AddCell()
	cell.Value = strings.Replace(locationresult.Area, "/media/", "https://www.hum.com/content/dam/", -1)
	err = file.Save("obdVampire.xlsx")
	hodo(err)

	fmt.Println("Car count: ", carcounter)
	// fmt.Println(locationresult.Location)
	// fmt.Println(locationresult.Picture)
	// fmt.Println(strings.Replace(locationresult.Area, "/media/", "https://www.hum.com/content/dam/", -1))
}

var file *xlsx.File
var sheet *xlsx.Sheet
var row *xlsx.Row
var cell *xlsx.Cell
var carcounter = 0

func main() {
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("obdVampire")
	hodo(err)
	row = sheet.AddRow()
	getBrands()
}
