package main

import (
	"encoding/json"
	"fmt"
	"testing"
)
func TestMain(t *testing.T){
	// 示例JSON数据
	jsonData := map[string]interface{}{
		"field1": []interface{}{
			[]int{1, 2, 3},
			5,
			6,
		},
	}

	// 修改元素
	if arr, ok := jsonData["field1"].([]interface{}); ok {
		if len(arr) > 0 {
			if innerArr, ok := arr[0].([]int); ok {
				if len(innerArr) > 1 {
					innerArr[1] = 20
				}
			}
		}
	}

	// 将修改后的JSON数据转换为字符串
	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输出修改后的JSON字符串
	fmt.Println(string(jsonBytes))
}

	
