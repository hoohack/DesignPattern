<?php
    require_once './Memento.php';
    require_once './Workout.php';
    require_once './CareTaker.php';
    $workout = new Workout();
    $workout->running();
    $care_taker = new CareTaker();
    $memento = $workout->createMemento();
    $care_taker->saveMemento($memento);

    $workout->bicycle();
    $workout->setMemento($care_taker->retriveMemento());
    $workout->abdominal_ripper_x();
