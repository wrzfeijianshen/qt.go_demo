package iobase

import (
	"fmt"
	// "fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	strings "strings"
)

func FullFilename(filespath string) string {
	filespath = strings.Replace(filespath, "\\", "/", -1)
	return filespath
}

func FilenameWithSuffix(filespath string) string {
	filespath = strings.Replace(filespath, "\\", "/", -1)
	filenameWithSuffix := path.Base(filespath) // 得到文件名含后缀
	return filenameWithSuffix
}

func FileSuffix(filespath string) string {
	filespath = strings.Replace(filespath, "\\", "/", -1)

	fileSuffix := path.Ext(filespath) // 得到后缀名
	fileSuffix = strings.Replace(fileSuffix, ".", "", -1)

	return fileSuffix
}
func FilenameOnly(filespath string) string {
	filespath = strings.Replace(filespath, "\\", "/", -1)
	filenameOnly := strings.TrimSuffix(filespath, path.Ext(filespath)) // 得到文件名不含后缀
	return filenameOnly
}

func FilePathOnly(filespath string) string {
	filespath = strings.Replace(filespath, "\\", "/", -1)
	strpath := FilenameWithSuffix(filespath)
	if strpath == filespath {
		return ""
	}

	str := "/"
	str += strpath

	filePathOnly := strings.Replace(filespath, str, "", -1)
	return filePathOnly
}

// AirFilterSlice滤空去除切片内的空切片
func AirFilterSlice(sn []string) []string {
	//fmt.Println("AirFilterSlice 0： ",sn," : ",len(sn))

	//for num,str := range sn {
	//	str = strings.Replace(str, " ", "", -1)
	//	if str == "" {
	//		sn = append(sn[:num], sn[num+1:]...)
	//	}
	//}
	var ss []string
	for _, str := range sn {
		str = strings.Replace(str, " ", "", -1)
		if str != "" {
			ss = append(ss, str)
		}
	}
	//fmt.Println("AirFilterSlice Ok： ",sn," : ",len(ss))
	return ss
}

// 测试接口
func MergeSlice(sn ...[]string) []string {
	//fmt.Print(" =======================", len(sn))

	if len(sn) < 2 {
		return nil
	}

	lenth := 0
	for num, ss := range sn {
		ss = AirFilterSlice(ss)
		if len(ss) <= 0 {
			continue
		}
		lenth += len(sn[num])
	}

	slice := make([]string, lenth)
	//var slice  []string

	for _, ss := range sn {
		//slice = append(slice,ss[num])
		for _, aa := range ss {
			if len(aa) <= 0 {
				continue
			}
			slice = append(slice, aa)
		}
	}

	//for _, v := range files {
	//	urls = MergeSlice(urls, util.ReadFile(v))
	//}
	return slice
}
func IsEmlySlice(sn []string) bool {
	flag := true
	for _, str := range sn {
		//slice = append(slice,ss[num])
		str = strings.Replace(str, " ", "", -1)
		//fmt.Println("IsEmlySlice: [",str,"]",num)
		if str == "" {
			flag = true
			continue
		} else {
			flag = false
		}
	}
	//fmt.Println("IsEmlySliceEND: [",flag,"]")
	//fmt.Println("IsEmlySlice: [",len(sn),"]")
	return flag
}

func RepeatString(str string, num uint64) string {
	var i uint64
	for i = 0; i < num; i++ {
		str += str
	}
	return str
}

//Duplicate removal
func DupRemoveMergeSlice(sn ...[]string) []string {
	//fmt.Print(" =======================", len(sn))

	if len(sn) < 2 {
		return nil
	}

	lenth := 0
	for num, _ := range sn {
		lenth += len(sn[num])
	}

	slice := make([]string, lenth)
	//var slice  []string

	for _, ss := range sn {
		if IsEmlySlice(ss) == true {
			continue
		}
		//slice = append(slice,ss[num])
		for _, str := range ss {
			str = strings.Replace(str, " ", "", -1)
			if str == "" {
				continue
			} else {
				flag := false
				for _, strslice := range slice {
					if str == strslice { // 重复去除
						flag = true
						break
					}
				}
				if flag == false {
					slice = append(slice, str)
				}
			}
		}
	}
	//if IsEmlySlice(slice) == false {
	//	fmt.Println("slice1 : ",slice)
	//}
	//for _, v := range files {
	//	urls = MergeSlice(urls, util.ReadFile(v))
	//}
	slice = AirFilterSlice(slice)
	//if IsEmlySlice(slice) == false {
	//	//fmt.Println("slice2 : ",slice,len(slice))
	//
	//}
	return slice
}

// 判断指定文件夹或文件存在不存在
func IsExists(path string) (bool, int, error) { // bool:文件或文件夹是否存在，int:文件夹为0,文件为1,其他为-1
	fileInfo, err := os.Stat(path)

	if err == nil {
		if fileInfo.IsDir() == true {
			// fmt.Println("文件夹")
			return true, 0, nil
		} else {
			// fmt.Println("文件")
			return true, 1, nil
		}

	}
	if os.IsNotExist(err) {
		return false, -1, nil
	}
	return false, -1, err
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
//files, err := ListDir("E:\\", ".txt")
//fmt.Println(files, err)
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	// PthSep := string(os.PathSeparator)
	suffix = strings.ToLower(suffix) //忽略后缀匹配的大小写
	if suffix == "*.*" {
		suffix = ""
	} else {
		temp := FileSuffix(suffix)
		suffix = "."
		suffix += temp

	}

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToLower(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+fi.Name())
		}
	}
	return files, nil
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。返回值包含路径
func WalkDir(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToLower(suffix)
	if suffix == "*.*" {
		suffix = ""
	} else {
		temp := FileSuffix(suffix)
		suffix = "."
		suffix += temp

	}
	fmt.Println(suffix)
	//忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			fmt.Println(filename)
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。返回名称和后缀
//files, err := ListDir("E:\\", ".txt")
//fmt.Println(files, err)
func ListDirFileName(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	// PthSep := string(os.PathSeparator)
	suffix = strings.ToLower(suffix) //忽略后缀匹配的大小写
	if suffix == "*.*" {
		suffix = ""
	} else {
		temp := FileSuffix(suffix)
		suffix = "."
		suffix += temp

	}

	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		filename := fi.Name()
		if strings.HasSuffix(strings.ToLower(fi.Name()), suffix) { //匹配文件
			filename = FilenameWithSuffix(filename)
			fmt.Println(filename)
			files = append(files, dirPth+fi.Name())
		}
	}
	return files, nil
}

//获取指定目录及所有子目录下的所有文件，可以匹配后缀过滤。
func WalkDirFilename(dirPth, suffix string) (files []string, err error) {
	files = make([]string, 0, 30)
	suffix = strings.ToLower(suffix)
	if suffix == "*.*" {
		suffix = ""
	} else {
		temp := FileSuffix(suffix)
		suffix = "."
		suffix += temp

	}
	fmt.Println(suffix)
	//忽略后缀匹配的大小写
	err = filepath.Walk(dirPth, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		//if err != nil { //忽略错误
		// return err
		//}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		if strings.HasSuffix(strings.ToLower(fi.Name()), suffix) {
			filename = FilenameWithSuffix(filename)
			fmt.Println(filename)
			files = append(files, filename)
		}
		return nil
	})
	return files, err
}
