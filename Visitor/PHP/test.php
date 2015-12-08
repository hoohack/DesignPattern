<?php
  require_once './User.php';
  require_once './UserVisitor.php';
  require_once './AddPointVisitor.php';
  require_once './NormalUser.php';
  require_once './VipUser.php';
  require_once './ObjectStructure.php';

  $add_point_visitor = new AddPointVisitor();

  $os = new ObjectStructure();
  $os->addUser(new VipUser("vip_hector"));
  $os->addUser(new NormalUser("continue"));
  $user_tom = new NormalUser("tom");
  $os->addUser($user_tom);
  $os->addUser($user_tom);

  $os->handleVisitor($add_point_visitor);
