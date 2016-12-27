package utils

import (
	"os"
	"strings"
	"path/filepath"
)

//遍历application下文件，取出所有文件名
func GetAppsBywalkdir(dirPth, suffix string) (files []string, err error) {
    files = make([]string, 0, 30)
    suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

    err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
    //if err != nil { //忽略错误
   // return err
   //}

    if fi.IsDir() { // 忽略目录
     	return nil
     }

   if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) {
  	 ss := strings.Split(filename, "/")

    	files = append(files, ss[len(ss)-1])
    }

    	return nil
  })

 return files, err
}
