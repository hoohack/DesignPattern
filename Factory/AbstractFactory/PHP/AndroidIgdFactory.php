<?php
  class AndroidIgdFactory implements IngredientFactory {
    public function createOS() {
      echo "creating android operation system...\n";
    }

    public function createSoftware() {
      echo "creating special software root\n"
    }
  }
?>
