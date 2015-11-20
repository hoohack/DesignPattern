<?php
  abstract class PhoneFactory {
    public function orderPhone($type) {
      $phone = $this->createPhone($type);
      return $phone;
    }

    abstract function createPhone($type);
  }
?>
