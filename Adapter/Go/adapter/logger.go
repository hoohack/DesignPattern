package adapter

type Logger struct {
  log_id string
  ope_user_id string
}

func (lg *Logger) GetLogId() string {
  return lg.log_id
}

func (lg *Logger) SetLogId(log_id string) {
  lg.log_id = log_id
}

func (lg *Logger) GetOpeUserId() string {
  return lg.ope_user_id
}

func (lg *Logger) SetOpeUserId(ope_user_id string) {
  lg.ope_user_id = ope_user_id
}
