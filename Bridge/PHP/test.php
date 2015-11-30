<?php
  require_once './AbstractMessage.php';
  require_once './MessageImplementor.php';
  require_once './CommonMessage.php';
  require_once './UrgencyMessage.php';
  require_once './SpecialUrgencyMessage.php';
  require_once './MessageEmail.php';
  require_once './MessageSMS.php';
  require_once './MessageMobile.php';

  $impl = new MessageSMS();
  $m = new CommonMessage($impl);
  $m->sendMessage("hello", "tom");

  $m = new UrgencyMessage($impl);
  $m->sendMessage("hello", "tom");

  $m = new SpecialUrgencyMessage($impl);
  $m->sendMessage("hello", "tom");

  $impl = new MessageMobile();
  $m = new CommonMessage($impl);
  $m->sendMessage("hello", "tom");

  $m = new UrgencyMessage($impl);
  $m->sendMessage("hello", "tom");

  $m = new SpecialUrgencyMessage($impl);
  $m->sendMessage("hello", "tom");
