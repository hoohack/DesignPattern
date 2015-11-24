<?php
  require_once './StoreDao.php';
  require_once './GoodsDao.php';
  require_once './ServiceFacade.php';

  $store_dao = new StoreDao();
  $goods_dao = new GoodsDao();

  $goods_service = new ServiceFacede($store_dao, $goods_dao);
  $goods_detail = $goods_service->goodsDetail();
  var_dump($goods_detail);
