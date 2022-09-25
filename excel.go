package utility

import (
	"fmt"
	"reflect"
	"time"
	"zjutjh/Join-Us/db/model"

	"github.com/xuri/excelize/v2"
)

func WriteXlsx(sheet string, records interface{}) *excelize.File {
	xlsx := excelize.NewFile()    // new file
	index := xlsx.NewSheet(sheet) // new sheet
	xlsx.SetActiveSheet(index)    // set active (default) sheet
	firstCharacter := 65          // start from 'A' line
	t := reflect.TypeOf(records)
	if t.Kind() != reflect.Slice {
		return xlsx
	}

	s := reflect.ValueOf(records)

	for i := 0; i < s.Len(); i++ {
		elem := s.Index(i).Interface()
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		for j := 0; j < elemType.NumField(); j++ {
			field := elemType.Field(j)
			tag := field.Tag.Get("xlsx")
			name := tag
			column := string(rune(firstCharacter + j))
			if tag == "" {
				continue
			}
			// 设置表头
			if i == 0 {
				xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+1), name)
			}
			// 设置内容
			xlsx.SetCellValue(sheet, fmt.Sprintf("%s%d", column, i+2), elemValue.Field(j).Interface())
		}
	}
	return xlsx
}

type Data struct {
	Name       string    `gorm:"not null" json:"name" xlsx:"姓名"`
	UpdatedAt  time.Time `gorm:"not null" json:"-" xlsx:"更新时间"`
	CreatedAt  time.Time `gorm:"not null" json:"-" xlsx:"创建时间"`
	StuID      string    `gorm:"primaryKey" json:"stu_id" xlsx:"学号"`
	Gender     string    `gorm:"not null" json:"gender" xlsx:"性别"`
	College    string    `gorm:"not null" json:"college" xlsx:"专业"`
	Campus     string    `gorm:"not null" json:"campus" xlsx:"学院"`
	Phone      string    `gorm:"not null" json:"phone" xlsx:"电话号"`
	QQ         string    `gorm:"not null" json:"qq" xlsx:"QQ"`
	Region     string    `gorm:"not null" json:"region" xlsx:"校区"`
	Want1      string    `gorm:"not null" json:"want1" xlsx:"第一志愿"`
	Want2      string    `gorm:"not null" json:"want2" xlsx:"第二志愿"`
	Profile    string    `gorm:"not null" json:"profile" xlsx:"简介"`
	Feedback   string    `gorm:"not null" json:"feedback" xlsx:"反馈"`
	IsModified bool      `gorm:"not null" json:"is_modified" xlsx:"是否修改"`
}

var gender map[string]string
var region [4]string
var want [10]string

func init() {
	gender = map[string]string{
		"0": "男",
		"1": "女",
	}
	region = [4]string{"未选择", "朝晖", "屏峰", "莫干山"}
	want = [10]string{"未选择",
		"办公室",
		"活动部",
		"秘书处",
		"Touch产品部",
		"小弘工作室",
		"编辑工作室",
		"视觉影像部",
		"技术部",
		"易班文化工作站",
	}
}

func GenerateExcel() string {
	forms, _ := model.GetAllNormalForms()
	var data []Data
	for _, form := range forms {
		data = append(data, Data{
			Name:       form.Name,
			UpdatedAt:  form.UpdatedAt,
			CreatedAt:  form.CreatedAt,
			IsModified: form.UpdatedAt != form.CreatedAt,
			StuID:      form.StuID,
			Gender:     gender[form.Gender],
			Campus:     form.Campus,
			College:    form.College,
			Phone:      form.Phone,
			QQ:         form.QQ,
			Region:     region[form.Region],
			Want1:      want[form.Want1],
			Want2:      want[form.Want2],
			Profile:    form.Profile,
			Feedback:   form.Feedback,
		})
	}
	f := WriteXlsx("Sheet1", data)
	filename := "./export/AllForms.xlsx"
	f.SaveAs(filename)
	return filename
}
