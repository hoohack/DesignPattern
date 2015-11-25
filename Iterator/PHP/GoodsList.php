<?php
  class GoodsList {
    private $goodses;

    public function getGoods($idx) {
      if (isset($this->goodses[$idx])) {
        return $this->goodses[$idx];
      }

      return null;
    }

    public function addGoods($goods) {
      $this->goodses[] = $goods;
      return $this->count();
    }

    public function removeGoods($goods) {
      foreach ($this->goodses as $key) {
        if ($this->goodses[$key]->getName() == $goods->getName()) {
          unset($this->goodses[$key]);
        }
      }

      return $this->count();
    }

    public function count() {
      return count($this->goodses);
    }
  }
