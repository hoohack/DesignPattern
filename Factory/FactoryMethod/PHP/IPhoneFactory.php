<?php
  class IPhoneFactory extends PhoneFactory {
    function createPhone($type) {
      if ($type == "6") {
        return new IPhone6();
      } else if ($type == "6s") {
        return new IPhone6s();
      }
    }
  }
?>
