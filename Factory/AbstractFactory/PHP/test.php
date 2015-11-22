<?php
  require_once './IngredientFactory.php';
  require_once './IPhoneFactory.php';
  require_once './AndroidFactory.php';
  require_once './PhoneFactory.php';
  require_once './OS.php';
  require_once './IPhoneOS.php';
  require_once './AndroidOS.php';
  require_once './Special.php';
  require_once './IPhoneSpecial.php';
  require_once './AndroidSpecial.php';


  $iPhoneFactory = new IPhoneFactory();
  $iphone_os = $iPhoneFactory->createOS("ios");
  echo $iphone_os->create();
  $iphone_special = $iPhoneFactory->createSpecial("siri");
  echo $iphone_special->create();

  $androidFactory = new AndroidFactory();
  $android_os = $androidFactory->createOS("android");
  echo $android_os->create();
  $android_special = $iPhoneFactory->createSpecial("nfc");
  echo $android_special->create();
?>
