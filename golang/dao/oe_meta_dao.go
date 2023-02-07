package dao

import (
	"ark-online-excel/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

func AddOnlineExcelFileData(data *models.ExcelMeta) (*models.ExcelMeta, error) {
	o := orm.NewOrm()
	id, err := o.Insert(data)
	if err != nil {
		return nil, err
	}

	data.Id = id
	return data, nil
}


func UpdateFileInfo(fileId ,user string, cell,data,rowData string) (*models.ExcelMeta, error) {

	oldFileMeta, err := GetMetaDataById(fileId)
	if err != nil {
		return nil, err
	}
	if oldFileMeta == nil {
		return nil, errors.New("资源不存在")
	}

	oldFileMeta.Cell = cell
	oldFileMeta.Data = data
	oldFileMeta.Modifier= user
	oldFileMeta.RowData = rowData

	o := orm.NewOrm()
	_, err = o.Update(oldFileMeta,"modifier","row_data", "cell", "data")
	if err != nil {
		return nil, err
	}

	return oldFileMeta, nil
}


func GetMetaDataById(fileId string) (*models.ExcelMeta, error) {
	cond := orm.NewCondition()
	cond = cond.And("file_id", fileId)
	q := orm.NewOrm().QueryTable(&models.ExcelMeta{}).SetCond(cond)
	var modules []*models.ExcelMeta
	_, err := q.All(&modules)
	if err != nil {
		return nil, err
	}
	if modules != nil && len(modules) > 0 {
		return modules[0], nil
	} else {
		return nil, nil
	}
}

func DelFileById(fileId string) error {
	adminOld, err := GetMetaDataById(fileId)
	if err != nil {
		return err
	}
	if adminOld == nil {
		return errors.New("资源不存在")
	}
	o := orm.NewOrm()
	admin := &models.ExcelMeta{
		FileId: fileId,
	}
	_, err = o.Delete(admin, "file_id")
	if err != nil {
		return err
	}
	return nil
}

func QueryMetaByConditions(fileName, author,version string) ([]*models.ExcelMeta, error) {
	if author == "" {
		return nil, fmt.Errorf("author params empty")
	}
	cond := orm.NewCondition()

	cond = cond.And("author", author)

	if fileName != "" {
		cond = cond.And("file_name", fileName)
	}

	if version != "" {
		cond = cond.And("version", version)
	}
	q := orm.NewOrm().QueryTable(&models.ExcelMeta{}).SetCond(cond)

	var modules []*models.ExcelMeta
	_, err := q.All(&modules)
	if err != nil {
		return nil, err
	}

	return modules,nil
}

func GetAllMetaFile(size int64) ([]*models.ExcelMeta, error) {
	q := orm.NewOrm().QueryTable(&models.ExcelMeta{})
	var modules []*models.ExcelMeta
	_, err := q.Limit(size, 0).All(&modules)
	return modules, err
}

func GetMetaFileVersion(fileId string) (int ,error) {
	cond := orm.NewCondition()
	cond = cond.And("file_id", fileId)
	q := orm.NewOrm().QueryTable(&models.ExcelMeta{}).SetCond(cond)
	var modules []*models.ExcelMeta
	_, err := q.All(&modules)
	if err != nil {
		return 0, err
	}
	if modules != nil && len(modules) > 0 {
		return modules[0].Version, nil
	} else {
		return 0, nil
	}
}