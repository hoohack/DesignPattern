<?php
  class Invoker {
    protected $command;

    public function setCommand($cmd) {
      $this->command = $cmd;
    }

    public function run() {
      $this->command->excute();
    }

    public function undo() {
      $this->command->undo();
    }
  }
?>
