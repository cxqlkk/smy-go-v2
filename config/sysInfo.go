package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
var UploadPath="/home/cxq/upload"
type sysInfo struct {
	UploadPath string `yaml:"uploadPath"`
}
var SysInfo *sysInfo=&sysInfo{}
func drop()  {
	fmt.Println("sysInfo ----- reading")
	bts,e:=ioutil.ReadFile(GetAppPath()+"/sysInfo.yml")
	if e!=nil{
		log.Println(e)
	}
	yaml.Unmarshal(bts,SysInfo)
}


func GetAppPath() string{
	file,_:=exec.LookPath(os.Args[0])
	fmt.Println(file)
	path,_:=filepath.Abs(file)

	fmt.Println(path)
	index:=strings.LastIndex(path,string(os.PathSeparator))
	return path[:index]

}