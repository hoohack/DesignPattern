<?php
  require_once 'Beverage.php';
  require_once 'CondimentDecorator.php';
  require_once 'DarkRoast.php';
  require_once 'Decaf.php';
  require_once 'Espresso.php';
  require_once 'HouseBlend.php';
  require_once 'Mocha.php';
  require_once 'Soy.php';
  require_once 'Whip.php';

  $dark_roast = new DarkRoast();
  $espresso = new Espresso();
  $decaf = new Decaf();
  $house_blend = new HouseBlend();

  $mocha = new Mocha($dark_roast);
  $double_mocha = new Mocha($mocha);

  $espresso_soy = new Soy($espresso);

  $espresso_soy_mocha = new Mocha($espresso_soy);

  echo "description " . $double_mocha->getDescription() . "===price : " . $double_mocha->cost() . "$\n";
  echo "description " . $espresso_soy->getDescription() . "===price : " . $espresso_soy->cost() . "$\n";
  echo "description " . $espresso_soy_mocha->getDescription() . "===price : " . $espresso_soy_mocha->cost() . "$\n";
?>
