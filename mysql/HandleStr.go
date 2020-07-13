package mysql


import "strings"

//majorNameSplit 处理专业名称
func MajorNameSplit(str string) (string, string) {
	str = strings.Replace(strings.Replace(strings.Replace(strings.Replace(str, "。", "，", -1), ";", "，", -1), ",", "，", -1), "；", "，", -1)
	var tag, mname string
	if string(str[0]) == "[" {
		strList := strings.Split(str, "]")
		tag = strings.Replace(strList[0], "[", "", 1)
		str = strList[1]
	}
	str = strings.Replace(strings.Replace(str, "{", "(", -1), "}", ")", -1)
	str = strings.Replace(strings.Replace(str, "〔", "(", -1), "〕", ")", -1)
	str = strings.Replace(strings.Replace(strings.Replace(str, "身高〈", "身高<", -1), "〈", "(", -1), "〉", ")", -1)
	splitStr := ""
	for _, s := range str {
		switch string(s) {
		case "【":
			splitStr = "【"
		case "[":
			splitStr = "["
		case "(":
			splitStr = "("
		case "（":
			splitStr = "（"
		}
		if splitStr != "" {
			break
		}
	}
	if splitStr == "" {
		return str, ""
	}
	ss := strings.Split(str, splitStr)
	mname = ss[0]
	if tag != "" {
		tag = tag + "，" + SplitTag(ss[1])
	} else {
		tag = SplitTag(strings.Join(ss[1:], splitStr))
	}
	return mname, tag
}

func SplitTag(str string) string {
	if str == "" {
		return ""
	}
	for {
		if string(str[0]) == "(" {
			str = strings.Replace(str, "(", "", 1)
		} else if string(str[0]) == "[" {
			str = strings.Replace(str, "[", "", 1)
		} else {
			break
		}
	}

	tag := strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(str, ")【(", "，", -1), ")((", "，", -1), ")【", "，", -1), ")(", "，", -1), "）（", "，", -1), ")（", "，", -1), "）(", "，", -1)
	for string([]rune(tag)[len([]rune(tag))-1]) == "】" || string([]rune(tag)[len([]rune(tag))-1]) == ")" || string([]rune(tag)[len([]rune(tag))-1]) == "）" {
		tag = string([]rune(tag)[:len([]rune(tag))-1])
		//beego.Info(tag)
	}
	con1 := strings.Count(tag, "(")
	con2 := strings.Count(tag, ")")
	con3 := strings.Count(tag, "（")
	con4 := strings.Count(tag, "）")
	if con1 > con2 {
		tag = tag + ")"
	}
	if con1 < con2 {
		tag = "(" + tag
	}
	if con3 > con4 {
		tag = tag + "）"
	}
	if con3 < con4 {
		tag = "（" + tag
	}
	return tag
}

var tagKey = []string{"狐臭", "精神病", "肝功能", "文化成绩", "E字表", "cm", "眼疾", "疤痕", "只", "不", "限", "男", "女", "考生", "加试", "矫正", "语种", "度数", "报考", "录取考生", "转氨酶", "嗅觉", "≥", "招生章程", "口试", "笔试", "复试", "口语", "为适应", "要求", "择优", "分制", "录取", "普通话", "相貌", "基础", "单科", "视力", "心", "裸视", "身高", "慎报", "五官", "身材", "体型", "体重", "加试", "近视", "建议", "满分", "招收", "色盲", "色弱", "形象", "身体", "口吃", "四肢", "听力", "晕车", "传染", "疾病", "残疾", "口齿", "辨色", "明显", "纹身", "≥", "高考", "折算", "%", "面试"}
var tagNo = []string{"建档立卡", "宠物", "不同", "有限", "女子学院", "教学计划", "补助", "不动产", "心理", "中心", "学分", "学费", "收费", "大类", "保研", "小动物", "组成", "选拔", "分流", "划分", "修业", "推荐", "企业", "合作", "费用", "人才", "学科", "班", "科学", "确定专业", "所含专业", "阶段", "必修课", "学年", "医学", "培养方式", "统一", "基地", "方向", "不列颠", "符合", "校区", "交流项目"}

func Warning(one string) (string, string) {
	one = strings.Replace(strings.Replace(strings.Replace(strings.Replace(one, ";", "，", -1), "；", "，", -1), "。", "，", -1), ",", "，", -1)
	strList := strings.Split(one, "，")
	warning := make([]string, 0)
	tags := make([]string, 0)
	for _, one := range strList {
		warnList := make([]string, 0)
		tagList := make([]string, 0)
		//for _, sp := range strings.Split(one, "，") {
		if strings.TrimSpace(one) == "" {
			continue
		}
		flag := 0
		for _, keyWord := range tagKey {
			if strings.Contains(one, keyWord) {
				flag1 := 0
				for _, notag := range tagNo {
					if strings.Contains(one, notag) {
						flag1++
						break
					}
				}
				if flag1 != 0 {
					break
				}
				warnList = append(warnList, one)
				flag++
				break
			}
		}
		if flag == 0 {
			tagList = append(tagList, one)
		}
		if len(warnList) > 0 {
			warning = append(warning, strings.Join(warnList, "，"))
		}
		if len(tagList) > 0 {
			tags = append(tags, strings.Join(tagList, "，"))
		}
	}
	warn := strings.Join(warning, "，")
	tag := strings.Join(tags, "，")
	return warn, tag
}

