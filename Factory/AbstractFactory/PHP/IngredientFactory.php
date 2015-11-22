<?php
  abstract class IngredientFactory {
    abstract function createOS($name);
    abstract function createSpecial($content);
  }
?>
