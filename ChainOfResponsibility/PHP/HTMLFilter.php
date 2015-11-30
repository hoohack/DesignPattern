<?php
  class HTMLFilter implements Filter {
    public function handleFilter($str) {
      $str = str_replace("<", "&lt;", $str);
      $str = str_replace(">", "&gt;", $str);
      return $str;
    }
  }
