package adapter

import "fmt"

type LogFileOperate struct {
  LogFileOperateAPI
  file_name string
}

func NewLogFileOperate(file_name string) *LogFileOperate{
  return &LogFileOperate{file_name: file_name}
}

func (this *LogFileOperate) ReadLogFile() string {
  content := "test content from " + this.file_name
  return content
}

func (this *LogFileOperate) WriteLogFile(content string) {
  fmt.Printf("Writing %s\n", content)
}
