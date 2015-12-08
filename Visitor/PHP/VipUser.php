<?php
  class VipUser extends User {

    public function accept($visitor) {
      $visitor->addPointForVip($this);
    }
  }
