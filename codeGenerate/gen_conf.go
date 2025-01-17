package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

type GenConf struct {
}

const (
	genConfName       = "genconfname"
	genConfDir        = "conf"
	G_TestConfName    = "test"
	G_ProductConfName = "product"
)

func init() {
	g := &GenConf{}
	err := GenMgrInstance.RegisterGen(genConfName, g)
	if err != nil {
		log.Printf("genConf init error,err:%+v\n", err)
		return
	}
}

func (g *GenConf) Run(opt *GenOption, metaData *protoMetaData) error {
	testConf := path.Join(opt.OutputPath, genConfDir, G_TestConfName, fmt.Sprintf("%s.yaml", G_TestConfName))
	err := g.genConfUseTemple(testConf, opt, metaData)
	if err != nil {
		log.Printf("genConf Run genConfUseTemple test error,err:%+v\n", err)
		return err
	}
	productConf := path.Join(opt.OutputPath, genConfDir, G_ProductConfName, fmt.Sprintf("%s.yaml", G_ProductConfName))
	err = g.genConfUseTemple(productConf, opt, metaData)
	if err != nil {
		log.Printf("genConf Run genConfUseTemple product error,err:%+v\n", err)
		return err
	}
	return nil
}

func (g *GenConf) genConfUseTemple(filename string, opt *GenOption, metaData *protoMetaData) error {
	//打开文件
	filef, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Printf("genConf genConfUseTemple open file error,err:%+v\n", err)
		return err
	}
	defer filef.Close()
	//解析
	err = parseTemple(filef, confTemp, nil)
	if err != nil {
		log.Printf("genConf genConfUseTemple parseTemple error,err:%+v\n", err)
		return err
	}
	return nil
}
