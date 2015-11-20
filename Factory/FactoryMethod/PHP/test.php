<?php
  require_once './PhoneFactory.php';
  require_once './Phone.php';
  require_once './IPhoneFactory.php';
  require_once './IPhone6.php';
  require_once './IPhone6s.php';
  require_once './AndroidFactory.php';
  require_once './BlackBerry.php';
  require_once './Nokia.php';
  require_once './OtherFactory.php';
  require_once './Samsung.php';
  require_once './XiaoMi.php';

  $iphoneFactory = new IphoneFactory();
  $androidFactory = new AndroidFactory();
  $otherFactory = new OtherFactory();

  $iphone = $iphoneFactory->orderPhone("6");
  echo "Get phone " . $iphone->getName() . "\n";

  $android = $androidFactory->orderPhone("samsung");
  echo "Get phone " . $android->getName() . "\n";

  $other = $otherFactory->orderPhone("blackberry");
  echo "Get phone " . $other->getName() . "\n";
?>
