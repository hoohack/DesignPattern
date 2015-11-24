<?php
  class LogAdapter implements LogDbOperateAPI {
    private $logFileOperate;

    public function LogAdapter($logFileOperate) {
      $this->logFileOperate = $logFileOperate;
    }

    public function createLog($logger) {
      $content = $this->logFileOperate->readLogFile();
      $this->logFileOperate->writeLogFile("id: " . $logger->getLogId() . ";ope_user_id: " . $logger->getOpeUserId() . ";content: " . $content);
      echo "write into db\n";
    }
  }
