package main

import (
	"context"
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/qiniu/qmgo"
)

// bool, for JSON booleans
// float64, for JSON numbers
// string, for JSON strings
// []interface{}, for JSON arrays
// map[string]interface{}, for JSON objects
// nil for JSON null


type Gene struct{
	Name 				 string
	AnnotatedAlterations []interface{}
	Therapeutic          []interface{}
	Diagnostic			 []interface{}
	Prognostic			 []interface{}
	FdaRecognized		 []interface{}
	TitleInfo			 []interface{}
}




func test1() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	defer func () {
		if err:=client.Close(ctx);err!=nil{
			panic(err)
		}
	}()
	if err != nil {
		panic(err)
	}
	db := client.Database("MyTest")
	oncokb := db.Collection("oncokb")
	// 获取所有基因文件
	folderPath := "./oncokb"
	// 获取文件夹下的所有文件
	fileList,err := getFileList(folderPath)
	if err != nil {
		panic(err)
	}
	// 遍历文件列表
	for _,file := range fileList{
		if filepath.Ext(file)==".json"{
			data,err:=os.ReadFile(file)
			geneNames := strings.Split(file, "-")
			geneName := strings.Split(geneNames[len(geneNames)-1], ".")[0]
			if err != nil {
				panic(err)
			}
			var jsonData map[string]interface{}
			err = json.Unmarshal(data,&jsonData)
			if err != nil {
				panic(err)
			}
			// 添加基因名称用于查询
			jsonData["name"]=geneName
			// 将突变点位提取出来
			therapes_info := jsonData["Therapeutic"]
			therapeArr,ok:=therapes_info.([]interface{})
			if !ok {
				panic("not a slice type")
			}
			if len(therapeArr)!=0{
				therapes_interface := therapeArr[0]
				therapes:=therapes_interface.([]interface{})

				for index,therape_int := range therapes{
					therape:=therape_int.([]interface{})
					url := therape[1].(string)
					mutations := strings.Split(url,"/")
					mutation := mutations[len(mutations)-1]
					therape = append(therape, mutation)
					therapes[index]=therape
				}
			}
			oncokb.InsertOne(ctx,jsonData)
		}

	}

}

func getFileList(folderPath string)([]string,error){
	var fileList []string
	err:=filepath.Walk(folderPath,func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path!=folderPath{
			fileList = append(fileList, path)
		}
		return nil
	})
	if err!=nil{
		return nil,err
	}
	return fileList,nil
}