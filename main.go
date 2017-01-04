package main

import (
	"fmt"
	"io/ioutil"
	"os"

	gmencoder "github.com/heartchord/jxonline/gameencoder"
	"github.com/henrylee2cn/mahonia"
)

func main() {
	f, err := os.Open("D:/蓝水晶.bak")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	encoder := new(gmencoder.RoleBakEncoder)
	encoder.Decode(data)

	mdecoder := mahonia.NewDecoder("GBK")

	rolename := string(encoder.RoleData.RoleName[:])
	rolename = mdecoder.ConvertString(rolename)
	fmt.Println(rolename)

	account := string(encoder.RoleData.Account[:])
	account = mdecoder.ConvertString(account)
	fmt.Println(account)

	//encoder.PrintAllSkillData()
	//encoder.PrintAllTaskData()
	encoder.PrintAllItemData()
}
