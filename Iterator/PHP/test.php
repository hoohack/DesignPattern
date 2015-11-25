<?php
  require_once './Goods.php';
  require_once './GoodsList.php';
  require_once './GoodsIterator.php';

  $goods_list = new GoodsList();
  $goods_list->addGoods(new Goods("t-shirt"));
  $goods_list->addGoods(new Goods("cup"));
  $goods_list->addGoods(new Goods("jean"));
  $goods_list->addGoods(new Goods("phone"));
  $goods_list->addGoods(new Goods("keyboard"));

  $goods_iterator = new GoodsIterator($goods_list);

  while ($goods_iterator->valid()) {
    echo $goods_iterator->current()->getName() . "\n";
    $goods_iterator->next();
  }
