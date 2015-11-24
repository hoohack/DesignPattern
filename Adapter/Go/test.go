package main

import "./adapter"

func main() {
  lg := new(adapter.Logger)
  lg.SetLogId("111")
  lg.SetOpeUserId("user_tom");
  logFileOperate_api := adapter.NewLogFileOperate("");
  db_api := adapter.NewLogAdapter(logFileOperate_api)
  db_api.CreateLog(lg)
}
