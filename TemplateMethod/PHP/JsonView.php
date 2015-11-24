<?php
  class JsonView extends ViewEngine {
    public function packData($data) {
      return json_encode($data);
    }
  }
