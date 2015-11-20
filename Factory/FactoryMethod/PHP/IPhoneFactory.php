<?php
  class IPhoneFactory implements PhoneFactory {
    function createPhone() {
      return new IPhone();
    }
  }
?>
