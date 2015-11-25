<?php
  require_once './NavigationComponent.php';
  require_once './ParentNavigation.php';
  require_once './ChildNavigation.php';

  $store_navigation = new ParentNavigation('store');

  $nike = new ChildNavigation('nike');
  $adidas = new ChildNavigation('adidas');
  $new_balance = new ChildNavigation('new balance');
  $assics = new ChildNavigation('assics');

  $store_navigation->add($nike);
  $store_navigation->add($adidas);
  $store_navigation->add($new_balance);
  $store_navigation->add($assics);

  var_dump($store_navigation->getChild());

  $food = new ChildNavigation('food');
  $food->add(new ChildNavigation('kfc'));
