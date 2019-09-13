package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	err := Init()
	if err != nil{
		log.Printf("init  read idl.yaml failed [%v]", err)
		os.Exit(1)
	}
	err =  os.MkdirAll(IDLConfig.Generate.OutPath,0777)
	if err != nil{
		log.Printf("create out directory failed [%v]", err)
		os.Exit(1)
	}
	importPath := []string{IDLConfig.IDLPath}
	importPath = append(importPath, IDLConfig.IncludePath...)
	descSource, err := DescriptorSourceFromProtoFiles(importPath, IDLConfig.Depends...)
	if err != nil {
		log.Println(err, "Failed to process proto source files.")
		os.Exit(1)
	}
	generateGo(descSource)
}



func generateGo(descSource DescriptorSource){
	MMap := IDLConfig.Generate.GoOptions.ExtraModifiers
	exMap := IDLConfig.ExcludesMap
	idlPath := IDLConfig.IDLPath
	outPath := IDLConfig.Generate.OutPath
	//log.Printf("%v", outPath)
	var M string
	for name, fileDesc := range descSource.Files() {
		//fmt.Println("name: ", name)
		if _, ok := exMap[name]; ok {
			continue
		}
		path := filepath.Join(idlPath, name)
		//log.Println("path:", path)
		fs := fileDesc.GetDependencies()
		var ms []string
		for _, fd := range fs {
			dependName := fd.GetName()
			//log.Println("depend:", fd.GetName())
			_, ok := exMap[dependName]
			if ok {
				continue
			}
			opt := fd.GetFileOptions().GoPackage
			if opt == nil {
				fmt.Println("go_package not define")
				os.Exit(1)
			}
			//log.Printf("mmap %v", MMap)

			if m, ok := MMap[dependName];ok{
				ms =append(ms, "M"+dependName+"="+m)
				//log.Printf("ms:%v", ms)
			}
		}

		M = strings.Join(ms, ",")
		//log.Println("\tM:", M)

		var args []string
		args = []string{"-I" + idlPath}
		for _, itr := range IDLConfig.IncludePath {
			args = append(args, "-I" + itr)
		}
		if len(fileDesc.GetServices()) == 0 {
				args = append(args, "--go_out=paths=source_relative,"+M+":"+outPath)
		} else {
				args = append(args, "--go_out=paths=source_relative,plugins=grpc,"+M+":"+outPath)
		}
		args = append(args, path)
		log.Printf("%v", args)
		cmd := exec.Command("protoc", args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(cmd.Args)
			log.Println("Error:", err)
			log.Println(string(out))
		}

	}
}
