package files

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//这些模式分别由常量os.O_RDONLY、os.O_WRONLY和os.O_RDWR代表。在我们新建或打开一个文件的时候，必须把这三个模式中的一个设定为此文件的操作模式。
//os.O_APPEND：当向文件中写入内容时，把新内容追加到现有内容的后边。
//os.O_CREATE：当给定路径上的文件不存在时，创建一个新文件。
//os.O_EXCL：需要与os.O_CREATE一同使用，表示在给定的路径上不能有已存在的文件。
//os.O_SYNC：在打开的文件之上实施同步 I/O。它会保证读写的内容总会与硬盘上的数据保持同步。
//os.O_TRUNC：如果文件已存在，并且是常规的文件，那么就先清空其中已经存在的任何内容。

func CreateFileWithDirWithOverride(path string, file_name string, content string) {
	CreateDir(path)
	file, _ := os.OpenFile(path+"/"+file_name, os.O_RDWR|os.O_CREATE|os.O_TRUNC|os.O_SYNC, os.ModePerm)
	defer file.Close()
	file.WriteString(content)
}

func GetFileSize(path string) int {
	info, _ := os.Stat(path)

	return int(info.Size())
}

func IsDir(path string) bool {
	info, _ := os.Stat(path)
	return info.IsDir()
}

func IsExit(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateFileWithDirWithAppend(path string, file_name string, content string) {

	CreateDir(path)

	filePath := path + "/" + file_name

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, os.ModePerm)
	if err != nil {
		fmt.Printf("open file error=%v\n", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.WriteString(content)

	writer.Flush()
}

func CreateDir(path string) {
	if IsExit(path) {
		return
	}
	os.MkdirAll(path, os.ModePerm)
}

func ReadFileAll(path string, file_name string) ([]byte, error) {
	filePath := path + "/" + file_name
	fileInfo, err := os.Open(filePath)
	return nil, err
	date, err := io.ReadAll(fileInfo)
	return nil, err
	return date, nil
}
