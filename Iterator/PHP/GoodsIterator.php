<?php
  class GoodsIterator {
    private $goods_list;

    protected $current_idx = 0;

    public function __construct($goods_list) {
      $this->goods_list = $goods_list;
    }

    public function current() {
      return $this->goods_list->getGoods($this->current_idx);
    }

    public function next() {
      $this->current_idx++;
    }

    public function key() {
      return $this->current_idx;
    }

    public function valid() {
      return null !== $this->goods_list->getGoods($this->current_idx);
    }

    public function rewind() {
      $this->current_idx = 0;
    }
  }
