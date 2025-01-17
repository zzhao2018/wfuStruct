package main

import (
	"log"
	"os"
	"path"
)

/*
  包路径生成器
*/
var DirList []string = []string{
	"service/client",
	"service/server",
	"scripts",
	"conf/product",
	"conf/test",
	"router",
	"model",
	"generate",
	"controller",
	"service/client/clientTool",
}


type PackGenerator struct {
}

//注册模块
func init(){
	packGen:=&PackGenerator{}
	GenMgrInstance.RegisterGen(packGen.Name(),packGen)
}

func(p *PackGenerator)Name()string{
	return DirGenName
}


func(p *PackGenerator) Run(opt *GenOption,metaData *protoMetaData)error{
	//生成pack
	for _,ele:=range DirList {
		packStr:=path.Join(opt.OutputPath,ele)
		err:=os.MkdirAll(packStr,0777)
		if err!=nil {
			log.Printf("pack Run generator dir %s error,err:%+v\n",packStr,err)
			return err
		}
	}
	return nil
}
