<?php
  class ServiceFacede {
    protected $store_dao;
    protected $goods_dao;

    public function __construct($store_dao, $goods_dao) {
      $this->store_dao = $store_dao;
      $this->goods_dao = $goods_dao;
    }

    public function goodsDetail() {
      return array(
        'store_info' => $this->store_dao->info(),
        'goods_info' => $this->goods_dao->info()
      );
    }
  }
?>
