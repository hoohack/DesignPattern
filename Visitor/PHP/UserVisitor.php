<?php
  interface UserVisitor {
    public function addPointForVip($vip_user);
    public function addPointForNormal($normal_user);
  }
