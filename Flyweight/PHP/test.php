<?php
  require_once './Flyweight.php';
  require_once './AuthorizationFlyweight.php';
  require_once './FlyweightFactory.php';
  require_once './SecurityMgr.php';

  $mgr = SecurityMgr::getInstance();
  $mgr->login("hhq");
  $mgr->login("tom");
  var_dump($mgr->hasPermit("hhq", "money", "look"));
  var_dump($mgr->hasPermit("tom", "money", "look"));
