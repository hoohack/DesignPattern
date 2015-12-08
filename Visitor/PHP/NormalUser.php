<?php
  class NormalUser extends User {

    public function accept($visitor) {
      $visitor->addPointForNormal($this);
    }
  }
