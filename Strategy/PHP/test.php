<?php
  require_once 'travelStrategy.php';
  require_once 'airplaneStrategy.php';
  require_once 'bicycleStrategy.php';
  require_once 'trainStrategy.php';
  require_once 'person.php';

  $person = new Person();

  $trainStrategy = new TrainStrategy();
  $person->setStrategy($trainStrategy);
  $person->travel();

  $bicycleStrategy = new BicycleStrategy();
  $person->setStrategy($bicycleStrategy);
  $person->travel();

  $airplaneStrategy = new AirplaneStrategy();
  $person->setStrategy($airplaneStrategy);
  $person->travel();

?>
