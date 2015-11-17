<?php
  require_once 'observer.php';
  require_once 'subject.php';
  require_once 'concreteSubject.php';
  require_once 'observerA.php';
  require_once 'observerB.php';

  $subject = new ConcreteSubject();
  $observerA = new ObserverA();
  $observerB = new ObserverB();

  $subject->registerObserver($observerA);
  $subject->registerObserver($observerB);

  $subject->setValue(5);
?>
