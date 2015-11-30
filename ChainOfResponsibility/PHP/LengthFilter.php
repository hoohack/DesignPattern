<?php
  class LengthFilter implements Filter {
    public function handleFilter($str){
      return substr($str, 0, 20);
    }
  }
