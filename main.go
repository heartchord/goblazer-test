package main

import (
	"fmt"
	"io/ioutil"
	"os"

	gmencoder "github.com/heartchord/jxonline/gameencoder"
	"github.com/henrylee2cn/mahonia"
)

func main() {
	f, err := os.Open("D:/温柔的づ杀戮↓.bak")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	encoder := new(gmencoder.RoleBakEncoder)
	encoder.Decode(data)

	mdecoder := mahonia.NewDecoder("GBK")

	rolename := string(encoder.RoleData.BaseData.RoleName[:])
	rolename = mdecoder.ConvertString(rolename)
	fmt.Println(rolename)

	account := string(encoder.RoleData.BaseData.Account[:])
	account = mdecoder.ConvertString(account)
	fmt.Println(account)

	for i := int16(0); i < encoder.RoleData.FightSkillCount; i++ {
		fmt.Printf("SkillData.SkillID = %d\n", encoder.SkillData[i].SkillID)
		fmt.Printf("SkillData.SkillLv = %d\n", encoder.SkillData[i].SkillLv)
		fmt.Printf("SkillData.SkillExp = %d\n", encoder.SkillData[i].SkillExp)
	}
}
