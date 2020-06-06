package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
)

func filesystemGetCurrentPath() string{
	//获取当前文件绝对路径
	_, fPath, _, _  := runtime.Caller(1)
	//获取路径，去掉文件部分
	dir := path.Dir(fPath)
	fmt.Println("current dir:", dir)
	return dir
}

func filesystemGetParentDirPath(dir string) string{
	path := filepath.Dir(dir)
	fmt.Println("获取当前目录的上级目录路径", path)
	return path
}

func filesystemGetCurrentDirOrFilename(fileOrDirPath string) string{
	//last element
	path := filepath.Base(fileOrDirPath)
	fmt.Println("获取指定目录或文件的当前目录名或文件名",path)
	return path
}

func filesystemGetListFolderContents(dir string){
	rd, err := ioutil.ReadDir(dir)
	if err != nil{
		panic(err)
	}
	//fmt.Println(rd)
	contentPrefix := fmt.Sprintf("获取指定目录(%s)下的文件或目录列表", dir)
	for _, rdi := range rd{
		nowFilePath := fmt.Sprintf("%s/%s", dir, rdi.Name())
		nowFileOrDir := "这是一个文件"
		if rdi.IsDir(){
			nowFileOrDir = "这是一个目录"
		}
		content := fmt.Sprintf("%s: %s: %s", contentPrefix, nowFilePath, nowFileOrDir)
		fmt.Println(content)
	}
}

func main(){
	currentDir := filesystemGetCurrentPath()
	newFilePath := fmt.Sprintf("%s\\%s", currentDir, "demo_file_dir.go")
	filesystemGetParentDirPath(currentDir)
	filesystemGetCurrentDirOrFilename(currentDir)
	filesystemGetCurrentDirOrFilename(newFilePath)
	filesystemGetListFolderContents(currentDir)
}

