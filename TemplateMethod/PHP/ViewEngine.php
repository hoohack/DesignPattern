<?php
  abstract class ViewEngine {
    abstract function packData($data);

    public function getData() {
      return array('code' => 200, 'data' => array('name' => 'tom', 'age' => 22));
    }

    public final function render() {
      $data = $this->getData();
      var_dump($this->packData($data));
    }
  }
?>
