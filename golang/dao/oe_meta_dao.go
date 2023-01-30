package dao

import (
	"ark-online-excel/models"
	"errors"
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


func UpdateFileInfo(fileId string, cell,data,rowData string) (*models.ExcelMeta, error) {

	metaInfo := &models.ExcelMeta{
		FileId:fileId,
	}

	adminOld, err := GetMetaDataById(fileId)
	if err != nil {
		return nil, err
	}
	if adminOld == nil {
		return nil, errors.New("资源不存在")
	}

	newMeta := &models.ExcelMeta{
		Cell:cell,
		Data:  data,

		RowData:rowData,
	}
	o := orm.NewOrm()
	_, err = o.Update(newMeta,"rowData", "cell", "data")
	if err != nil {
		return nil, err
	}
	return metaInfo, nil
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
	cond := orm.NewCondition()
	cond = cond.And("file_name", fileName)

	if author != "" {
		cond = cond.And("author", author)
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
