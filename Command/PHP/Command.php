<?php
  interface Command {
    public function excute();
    public function undo();
  }
