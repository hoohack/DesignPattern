<?php
  class Logger {
    private $log_id;
    private $ope_user_id;

    public function getLogId() {
      return $this->log_id;
    }

    public function setLogId($log_id) {
      $this->log_id = $log_id;
    }

    public function getOpeUserId() {
      return $this->ope_user_id;
    }

    public function setOpeUserId($ope_user_id) {
      $this->ope_user_id = $ope_user_id;
    }
  }
?>
