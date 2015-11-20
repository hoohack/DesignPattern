<?php
  class AndroidFactory implements PhoneFactory {
    public function createPhone() {
      return new Android();
    }
  }
?>
