<?php
  class IPhoneFactory extends IngredientFactory {
    public function createOS($name) {
      return new IPhoneOS($name);
    }

    public function createSpecial($content) {
      return new IPhoneSpecial($content);
    }
  }
?>
