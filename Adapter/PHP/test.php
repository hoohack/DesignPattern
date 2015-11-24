<?php
  require_once './Logger.php';
  require_once './LogFileOperateAPI.php';
  require_once './LogFileOperate.php';
  require_once './LogDbOperateAPI.php';
  require_once './LogAdapter.php';

  $logger = new Logger();
  $logger->setLogId("111");
  $logger->setOpeUserId("user_tom");
  $logFileOperate_api = new LogFileOperate("");
  $db_api = new LogAdapter($logFileOperate_api);
  $db_api->createLog($logger);
?>
