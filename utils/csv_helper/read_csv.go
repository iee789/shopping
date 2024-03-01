package csv_helper

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
)

// 解析csv文件工具类

// 从给定的FileHeader读信息，返回二维切片
func ReadCsv(fileHeader *multipart.FileHeader) ([][]string, error) {
	// 打开文件
	f, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer func(f multipart.File) {
		err := f.Close() // 延迟关闭
		if err != nil {
			log.Print(err)
		}
	}(f)
	// 读取csv文件
	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("err")
		return nil, err
	}
	var result [][]string

	// 通常第一行包含列标题，这个函数只处理数据行
	for _, line := range lines[1:] {
		// 提取前两个字段（索引为0和1），将它们存储在一个新的字符串切片
		data := []string{line[0], line[1]}
		result = append(result, data)
	}

	return result, nil
}
