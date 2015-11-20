<?php
  class AndroidFactory extends PhoneFactory {
    public function createPhone($type) {
      if ($type == "samsung") {
        return new Samsung();
      } else if ($type == "xiaomi") {
        return new XiaoMi();
      }
    }
  }
?>
