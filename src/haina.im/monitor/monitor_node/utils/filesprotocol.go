package utils

import (
	"io"
	"os"
	"path/filepath"
	"strings"
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

/**********************************************************************************************************/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func Writefile(path, data string) {
	var f *os.File
	//var err1 error

	if checkFileIsExist(path) { //如果文件存在
		f, _ = os.OpenFile(path, os.O_APPEND, 0666) //打开文件
	} else {
		f, _ = os.Create(path) //创建文件
	}

	defer f.Close()

	//check(err1)
	io.WriteString(f, data) //写入文件(字符串)
	//check(err1)
}
