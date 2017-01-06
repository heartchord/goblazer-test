package main

import (
	"fmt"
	"io/ioutil"
	"os"

	gmencoder "github.com/heartchord/jxonline/gameencoder"
	"github.com/henrylee2cn/mahonia"
)

func main() {
	f, err := os.Open("D:/道骨仙风.bak")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	encoder := gmencoder.NewRoleBakEncoder()
	encoder.Decode(data)

	mdecoder := mahonia.NewDecoder("GBK")

	rolename := string(encoder.RoleBaseData.RoleName[:])
	rolename = mdecoder.ConvertString(rolename)
	fmt.Println(rolename)

	account := string(encoder.RoleBaseData.Account[:])
	account = mdecoder.ConvertString(account)
	fmt.Println(account)

	encoder.RoleExtData.PrintEquipComposeData()
	//encoder.PrintAllSkillData()
	//encoder.PrintAllTaskData()
	encoder.PrintAllItemData()
}
