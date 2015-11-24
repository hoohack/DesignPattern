<?php
  class LogFileOperate implements LogFileOperateAPI {
    private $file_name;

    public function __construct($file_name) {
      $this->file_name = $file_name;
    }

    public function readLogFile() {
      $content = "test content from " . $this->file_name;
      return $content;
    }

    public function writeLogFile($content) {
      echo "writing $content\n";
    }
  }
?>
