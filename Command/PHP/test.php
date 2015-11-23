<?php
  require_once './Command.php';
  require_once './PlayCommand.php';
  require_once './PauseCommand.php';
  require_once './StopCommand.php';
  require_once './Invoker.php';
  require_once './Player.php';

  $invoker = new Invoker();
  $player = new Player();
  $playCommand = new PlayCommand($player);
  $pauseCommand = new PauseCommand($player);
  $invoker->setCommand($playCommand);
  $invoker->run();
  $invoker->undo();
  $invoker->setCommand($pauseCommand);
  $invoker->run();
  $invoker->undo();
?>
