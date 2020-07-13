package main

import (
	"DataStructureGolang/asyncgoroutine"
	"DataStructureGolang/mysql"
)

type SpiderData struct {
	MainID         int    `gorm:"column:mainid"`
	ProfessionName string `gorm:"column:professionName"`
	Sid            int    `gorm:"column:sid"`
	Mid            int    `gorm:"column:mid"`
	MName          string `gorm:"column:mname"`
	MTag           string `gorm:"column:mtag"`
	MWaring        string `gorm:"column:mwaring"`
	Batch          int    `gorm:"column:batch"`
	Aot            int    `gorm:"column:aot"`
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

	asyncgoroutine.WorkerInstance()

	//config := &mysql.DBConfig{
	//	Name:     "spider-yzy-init",
	//	User:     "root",
	//	Password: "root123",
	//	Host:     "192.168.0.197",
	//	Port:     3306,
	//}
	//
	//config1 := &mysql.DBConfig{
	//	Name:     "gkzy-school",
	//	User:     "root",
	//	Password: "root123",
	//	Host:     "192.168.0.197",
	//	Port:     3306,
	//}
	//
	//config2 := &mysql.DBConfig{
	//	Name:     "gkzy3",
	//	User:     "root",
	//	Password: "root123",
	//	Host:     "192.168.0.197",
	//	Port:     3306,
	//}
	//
	//config3 := &mysql.DBConfig{
	//	Name:     "gkzy-major",
	//	User:     "root",
	//	Password: "root123",
	//	Host:     "192.168.0.197",
	//	Port:     3306,
	//}
	//
	//configSlice := make([]*mysql.DBConfig, 0)
	//configSlice = append(configSlice, config)
	//configSlice = append(configSlice, config1)
	//configSlice = append(configSlice, config2)
	//configSlice = append(configSlice, config3)
	//err := mysql.InitDB(configSlice)
	//if err != nil {
	//	fmt.Println("初始化错误:", err.Error())
	//	return
	//}
	//
	//s := SpiderData{}
	//data, _ := s.SearchData()
	//for k, v := range data {
	//	//mname, tag := mysql.MajorNameSplit(v.ProfessionName)
	//	//mwarn, mtag := mysql.Warning(tag)
	//	//
	//	//v.MName = mname
	//	//v.MTag = mtag
	//	//v.MWaring = mwarn
	//	//
	//	//sd:= SpiderData{}
	//	//_ = sd.Update(v.MainID, v.MName, v.MTag, v.MWaring)
	//
	//	//ms := MajorStruct{}
	//	//mid, _ := ms.FindMid(v.MName, v.Batch)
	//	//
	//	//sd := SpiderData{}
	//	//_ = sd.Update(v.MainID, mid, v.MName, v.MTag, v.MWaring)
	//	//fmt.Println("正在处理:", k)
	//
	//	if v.Batch == 1 {
	//		v.Aot = 20
	//	}
	//	if v.Batch == 2 {
	//		v.Aot = 30
	//	}
	//	if v.Batch == 3 {
	//		v.Aot = 70
	//	}
	//	sd := SpiderData{}
	//	_ = sd.UpdateAot(v.MainID, v.Aot)
	//	//time.Sleep(time.Second * 3)
	//	fmt.Println("正在处理:", k)
	//}

}

func (this *SpiderData) TableName() string {
	return "major_score-2019-all-45_copy2"
}

func (this *SpiderData) SearchData() (data []*SpiderData, err error) {
	db := mysql.GetORMByName("spider-yzy-init")
	db = db.Model(this)
	err = db.Find(&data).Error
	return
}

func (this *SpiderData) Update(mainid, mid int, mname, mtag, mwaring string) (err error) {
	db := mysql.GetORMByName("spider-yzy-init")
	db = db.Model(this)
	err = db.Where("mainid=?", mainid).Update("mname", mname).Update("mtag", mtag).Update("mwaring", mwaring).Update("mid", mid).Error
	return
}

func (this *SpiderData) UpdateAot(mainid, aot int) (err error) {
	db := mysql.GetORMByName("spider-yzy-init")
	db = db.Model(this)
	err = db.Where("mainid=?", mainid).Update("aot", aot).Error
	return
}

type MajorStruct struct {
	MajorID   int    `gorm:"major_id"`
	MajorName string `gorm:"major_name"`
}

func (this *MajorStruct) TableName() string {
	return "majors"
}

func (this *MajorStruct) FindMid(majorName string, batch int) (majorID int, err error) {
	ms := MajorStruct{}

	db := mysql.GetORMByName("gkzy-major")
	db = db.Model(this)

	if batch == 3 {
		err = db.Where("major_name=?", majorName).Where("deep>?", 1).Where("major_type=?", 2).Find(&ms).Error
	}
	if batch == 1 || batch == 2 {
		err = db.Where("major_name=?", majorName).Where("deep>?", 1).Where("major_type=?", 1).Find(&ms).Error
	}

	if err == nil {
		majorID = ms.MajorID
	} else {
		majorID = -1
	}
	return
}
