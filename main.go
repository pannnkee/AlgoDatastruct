package main

import (
	"DataStructureGolang/mysql"
	"fmt"
)

type SpiderData struct {
	ID          int    `gorm:"column:id"`
	ProvinceId  string `gorm:"column:provinceId"`
	CollegeId   string `gorm:"column:collegeId"`
	CollegeName string `gorm:"column:collegeName"`
	YearVal     string `gorm:"column:yearVal"`
	MinScoreVal string `gorm:"column:minScoreVal"`
	MaxScoreVal string `gorm:"column:maxScoreVal"`
	AvgScoreVal string `gorm:"column:avgScoreVal"`
	BatchNameVal string `gorm:"column:batchNameVal"`

	Sid      int `gorm:"column:sid"`
	MaxScore int `gorm:"column:max_score"`
	MinScore int `gorm:"column:min_score"`
	AvgScore int `gorm:"column:avg_score"`
	Year     int `gorm:"column:year"`
	ProvID   int `gorm:"column:prov_id"`
	AdmitOrderType int `gorm:"column:admit_order_type"`
	MajorType int `gorm:"column:major_type"`
}

func main() {
	//src.ListInstance()
	//src.DListNodeInstance()
	//stack.ItemStackInterface()
	//stack.ListStackInstance()
	//queue.InstanceQueue()
	//tree.BinaryTreeNodeInstance()
	//tree.BinarySearchTreeInstance()
	//tree.TestTreeNodIenstance()


	config := &mysql.DBConfig{
		Name:     "spider-yzy-init",
		User:     "root",
		Password: "root123",
		Host:     "192.168.0.197",
		Port:     3306,
	}

	config1 := &mysql.DBConfig{
		Name:     "gkzy-school",
		User:     "root",
		Password: "root123",
		Host:     "192.168.0.197",
		Port:     3306,
	}

	config2 := &mysql.DBConfig{
		Name:     "gkzy3",
		User:     "root",
		Password: "root123",
		Host:     "192.168.0.197",
		Port:     3306,
	}

	configSlice := make([]*mysql.DBConfig, 0)
	configSlice = append(configSlice, config)
	configSlice = append(configSlice, config1)
	configSlice = append(configSlice, config2)
	err := mysql.InitDB(configSlice)
	if err != nil {
		fmt.Println("初始化错误:", err.Error())
		return
	}
	

	//s := SpiderData{}
	//data, _ := s.SearchSidError()
	//for _, v := range data {
	//
	//	sid := Gkzy3NameSid{Name: v.CollegeName}
	//	v.Sid = sid.FindSidBySchoolName(sid.Name)
	//	v.Update()
	//
	//	//time.Sleep(time.Second)
	//}

	//
	//s := SpiderData{}
	//data, _ := s.Search()
	//
	//for _,v := range data {
	//
	//	sid := SNameSid{SchoolName: v.CollegeName}
	//
	//	v.Sid = sid.FindSidBySchoolName(sid.SchoolName)
	//	v.MaxScore, _ = strconv.Atoi(v.MaxScoreVal)
	//	v.MinScore, _ = strconv.Atoi(v.MinScoreVal)
	//	v.AvgScore, _ = strconv.Atoi(v.AvgScoreVal)
	//	v.Year, _ = strconv.Atoi(v.YearVal)
	//	v.ProvID = 45
	//
	//	if v.BatchNameVal == "本一批" {
	//		v.AdmitOrderType = 20
	//	} else if v.BatchNameVal == "本二批" {
	//		v.AdmitOrderType = 30
	//	} else {
	//		v.AdmitOrderType = 70
	//	}
	//	v.Update()
	//	//time.Sleep(time.Hour)
	//}

}

func (this *SpiderData) TableName() string {
	return "score-2019-all-45_copy2"
}
func (this *SpiderData) Search() (data []*SpiderData, err error) {
	db := mysql.GetORMByName("spider-yzy-init")
	db = db.Model(this)
	err = db.Find(&data).Error
	return
}

func (this *SpiderData) SearchSidError() (data []*SpiderData, err error) {
	db := mysql.GetORMByName("spider-yzy-init")
	db = db.Model(this)
	err = db.Where("sid=?", -1).Find(&data).Error
	return
}

func (this *SpiderData) Update() {
	db := mysql.GetORMByName("spider-yzy-init")
	db = db.Model(this)
	err := db.Where("id=?", this.ID).Update("sid", this.Sid).Update("max_score", this.MaxScore).Update("min_score", this.MinScore).
		Update("avg_score", this.AvgScore).Update("year", this.Year).Update("prov_id", this.ProvID).
		Update("admit_order_type", this.AdmitOrderType).Error
	if err != nil {
		fmt.Println("update错误:", err.Error())
		return
	}
}


type SNameSid struct {
	SchoolId int `gorm:"column:school_id"`
	SchoolName string `gorm:"column:name"`
}

func (this *SNameSid) TableName() string {
	return "school"
}

func (this *SNameSid)FindSidBySchoolName(schoolName string) (sid int) {
	ss := new(SNameSid)
	db := mysql.GetORMByName("gkzy-school")
	err := db.Model(this).Where("name=?", schoolName).Find(ss).Error
	if err != nil {
		fmt.Println("查询出错:", err.Error())
		return -1
	}
	return ss.SchoolId
}

type Gkzy3NameSid struct {
	Sid int `gorm:"column:sid"`
	Name string `gorm:"column:name"`
}

func (this *Gkzy3NameSid) TableName() string {
	return "tbl_school"
}

func (this *Gkzy3NameSid) FindSidBySchoolName(schoolName string) (sid int) {
	ss := new(Gkzy3NameSid)
	db := mysql.GetORMByName("gkzy3")
	err := db.Model(this).Where("name=?", schoolName).Find(ss).Error
	if err != nil {
		fmt.Println("查询出错:", err.Error())
		return -1
	}
	return ss.Sid
}




