<?php
  interface LogFileOperateAPI {
    public function readLogFile();
    public function writeLogFile($content);
  }
?>
