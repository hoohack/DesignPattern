<?php
  abstract class NavigationComponent {
    public function add($navigation_component) {
      echo "Unsupported operation\n";
    }

    public function remove($navigation_component) {
      echo "Unsupported operation\n";
    }

    public function getChild() {
      echo "Unsupported operation\n";
      return null;
    }

    public abstract function getName();
  }
