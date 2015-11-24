package adapter

type LogFileOperateAPI interface {
  readLogFile() string
  writeLogFile(content string)
}
