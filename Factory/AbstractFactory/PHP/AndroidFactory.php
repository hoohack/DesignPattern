<?php
  class AndroidFactory extends IngredientFactory {
    public function createOS($name) {
      return new AndroidOS($name);
    }

    function createSpecial($content) {
      return new AndroidSpecial($content);
    }
  }
?>
