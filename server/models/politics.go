package models

import (
	"fmt"
	orm "graduation-project/server/database"
	"time"
)

type Politics struct {
	PoliticsId    int       `gorm:"column:id" json:"politicsId"`
	Title         string    `gorm:"column:title" json:"title"`
	Abstract      string    `gorm:"column:abstract" json:"abstract"`
	Content       string    `gorm:"column:content" json:"content"`
	EnclosurePath string    `gorm:"column:enclosure_path" json:"enclosurePath"`
	CreateTime    time.Time `gorm:"column:create_time" json:"createTime"`
	PoliticsType  string    `gorm:"column:politics_type" json:"politicsType"` // guide、report、notify
	ReadingSigns  string    `gorm:"column:reading_signs" json:"readingSigns"`
}

func (p *Politics) TableName() string { return "politics" }

func (p *Politics) QueryPoliticsNotify(userId int, politicsType string) (politicsList []Politics, err error) {
	err = orm.DB.Raw("select * from politics where politics_type = ? and reading_signs not like ?", politicsType,
		fmt.Sprintf("%%[%d]%%", userId)).Find(&politicsList).Error
	for index, _ := range politicsList {
		newCol := politicsList[index].ReadingSigns + fmt.Sprintf(",[%d]", userId)
		err = orm.DB.Model(&politicsList[index]).Update("reading_signs", newCol).Error
	}

	return
}

// 删除
// 删除单条操作日志
func (p *Politics) DeletePolitics(politicsId int) error {
	return orm.DB.Where("id = ?", politicsId).Delete(p).Error
}

// 添加
func (p *Politics) AddPolitics(title string, abstract string, content string, politicsType string) (Politics, error) {
	var pp = Politics{
		PoliticsType: politicsType,
		Title:        title,
		Abstract:     abstract,
		Content:      content,
		CreateTime:   time.Now(),
	}

	err := orm.DB.Create(&pp).Error
	return pp, err
}

// 查询
func (p *Politics) QueryPoliticsToPage(politicsType string, title string, current int, size int) ([]Politics, int) {
	politicsDB := orm.DB.Model(p).Where("politics_type = ? and title like ?", politicsType,
		fmt.Sprintf("%%%s%%", title)).Order("create_time desc")
	var total int
	politicsDB.Count(&total)

	var list []Politics
	politicsDB.Offset((current - 1) * size).Limit(size).Find(&list)

	return list, total
}
