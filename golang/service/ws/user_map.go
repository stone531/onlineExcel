package ws

import (
	"ark-online-excel/logger"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

var (
	file2ShareObj map[string]*ExcelShareObj
	mu sync.Mutex
)

func init(){
	file2ShareObj = make(map[string]*ExcelShareObj, 0)
}

func AddMap(fileName string,obj *ExcelShareObj) {
	if fileName == "" || obj == nil {
		logger.ZapLogger.Error("AddMap params_bk Error", zap.String("filename",fileName), zap.Any("obj is:",obj))
		return
	}

	mu.Lock()
	file2ShareObj[fileName] = obj
	mu.Unlock()

}

func FindFileObj(fileName string) *ExcelShareObj {
	if fileName == "" {
		logger.ZapLogger.Error("FindFileObj params_bk Error  filename emtpy")
		return nil
	}

	if obj,ok := file2ShareObj[fileName]; ok {
		return obj
	}

	return nil
}

func DelMap(fileName string) {
	mu.Lock()
	delete(file2ShareObj,fileName)
	mu.Unlock()
}

func ClearUnUsingShareRoom() {
	if len(file2ShareObj) == 0{
		return
	}

	for fileName,roomObj := range file2ShareObj {
		if !roomObj.runFlag {
			fmt.Println("ClearUnUsingShareRoom file_id:",fileName)
			DelMap(fileName)
		}
	}
}
