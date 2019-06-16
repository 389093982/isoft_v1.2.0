package fileutil

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 写文件,文件不存在时会自动创建
func WriteFile(filename string, data []byte, append bool) error {
	if append {
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		n, err := f.Write(data)
		if err == nil && n < len(data) {
			err = io.ErrShortWrite
		}
		if err1 := f.Close(); err == nil {
			err = err1
		}
		return err
	} else {
		return ioutil.WriteFile(filename, data, 0666)
	}
}

// golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断:
// 如果返回的错误为nil,说明文件或文件夹存在
// 如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
// 如果返回的错误为其它类型,则不确定是否在存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 拷贝文件夹
func CopyDir(sourceDir string, destDir string) error {
	// 遍历文件夹
	err := filepath.Walk(sourceDir, func(path string, f os.FileInfo, err error) error {
		if f != nil {
			// 是目录同时是子目录则进行目录拷贝
			if f.IsDir() {
				if sourceDir != path {
					err = CopyDir(path, strings.Replace(ChangeToLinuxSeparator(path), ChangeToLinuxSeparator(sourceDir), ChangeToLinuxSeparator(destDir), -1))
					if err != nil {
						return err
					}
				}
			} else {
				// 是文件则进行文件拷贝
				err = CopyFile(path, strings.Replace(ChangeToLinuxSeparator(path), ChangeToLinuxSeparator(sourceDir), ChangeToLinuxSeparator(destDir), -1))
			}
		}
		return nil
	})
	return err
}

// 拷贝文件,要拷贝的文件路径,拷贝到哪里
func CopyFile(source, dest string) error {
	if source == "" || dest == "" {
		return errors.New("source file path or dest file path is empty.")
	}
	// 判断源文件是否存在
	if exist, _ := PathExists(source); exist == false {
		return errors.New("source file is exist.")
	}
	// 判断目标文件是否存在,不存在则创建
	if exist, _ := PathExists(dest); exist == false {
		os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	}
	//打开文件资源
	source_open, err := os.Open(source)
	//养成好习惯.操作文件时候记得添加 defer 关闭文件资源代码
	if err != nil {
		return err
	}
	defer source_open.Close()
	//只写模式打开文件,如果文件不存在进行创建 并赋予 644的权限.详情查看linux 权限解释
	dest_open, err := os.Create(dest)
	if err != nil {
		return err
	}
	//养成好习惯.操作文件时候记得添加 defer 关闭文件资源代码
	defer dest_open.Close()
	//进行数据拷贝
	_, copy_err := io.Copy(dest_open, source_open)
	if copy_err != nil {
		return copy_err
	}
	return nil
}

func GetAllSubFile(dirpath string) (filepaths []string, files []os.FileInfo, err error) {
	// 读取目录 dirpath 中的所有目录和文件,不包括子目录
	files, err = ioutil.ReadDir(dirpath)
	if err == nil {
		for _, file := range files {
			filepaths = append(filepaths, filepath.Join(dirpath, file.Name()))
		}
	}
	return filepaths, files, err
}

// 递归获取目录下的所有文件
func GetAllFile(dirpath string, recursion bool) (filepaths []string, files []os.FileInfo, err error) {
	if recursion == false {
		// 非递归场景
		return GetAllSubFile(dirpath)
	}
	// 递归场景
	err = filepath.Walk(dirpath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if f != nil {
			filepaths = append(filepaths, path)
			files = append(files, f)
		}
		return nil
	})
	return filepaths, files, err
}

func IsFile(path string) bool {
	return !IsDir(path)
}

func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err == nil {
		return f.IsDir()
	}
	return false
}

func ChangeToLinuxSeparator(path string) string {
	return strings.Replace(path, "\\", "/", -1)
}

func RemoveFileOrDirectory(filepath string) {
	if exist, err := PathExists(filepath); exist == true {
		if err = os.RemoveAll(filepath); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("clean dir %s\n", filepath)
		}
	}
}
