<?php
  class AddPointVisitor implements UserVisitor {
    public function addPointForVip($vip_user) {
      $vip_user->addPoint(10);
      echo "给VIP{$vip_user->getName()}加10分，总分{$vip_user->getPoint()}分\n";
    }

    public function addPointForNormal($normal_user) {
      $normal_user->addPoint(5);
      echo "给普通用户{$normal_user->getName()}加5分，总分{$normal_user->getPoint()}分\n";
    }
  }
