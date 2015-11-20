<?php
  require_once './PhoneFactory.php';
  require_once './Phone.php';
  require_once './IPhoneFactory.php';
  require_once './IPhone.php';
  require_once './Android.php';
  require_once './AndroidFactory.php';
  require_once './OtherFactory.php';
  require_once './Other.php';

  $iphoneFactory = new IphoneFactory();
  $androidFactory = new AndroidFactory();
  $otherFactory = new OtherFactory();

  $iphone = $iphoneFactory->createPhone();
  echo "Get phone " . $iphone->getName() . "\n";

  $android = $androidFactory->createPhone();
  echo "Get phone " . $android->getName() . "\n";

  $other = $otherFactory->createPhone();
  echo "Get phone " . $other->getName() . "\n";
?>
