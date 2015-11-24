package adapter

import "fmt"

type LogAdapter struct {
  LogFileOperateAPI
  *LogFileOperate
}

func NewLogAdapter(logFileOperate *LogFileOperate) *LogAdapter {
  return &LogAdapter{LogFileOperate: logFileOperate}
}

func (this *LogAdapter) CreateLog(lg *Logger) {
  content := this.LogFileOperate.ReadLogFile()
  this.LogFileOperate.WriteLogFile("id: " + lg.GetLogId() + ";ope_user_id: " + lg.GetOpeUserId() + ";content: " + content);
  fmt.Printf("Write into db\n");
}
