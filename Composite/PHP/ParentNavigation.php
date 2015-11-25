<?php
  class ParentNavigation extends NavigationComponent {
    private $childs;
    private $name;

    public function __construct($name) {
      $this->name = $name;
    }

    public function add($navigation_component) {
      $this->childs[] = $navigation_component;
    }

    public function remove($navigation_component) {
      foreach ($childs as $key => $child) {
        if ($child->getName() == $navigation_component->getName()) {
          unset($childs[$key]);
          break;
        }
      }
    }

    public function getChild() {
      $result = "";
      foreach ($this->childs as $child) {
        $result .= " " . $child->getName();
      }
      return $result;
    }

    public function getName() {
      return $this->name;
    }
  }
