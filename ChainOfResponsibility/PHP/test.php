<?php
  require_once './Filter.php';
  require_once './HTMLFilter.php';
  require_once './LengthFilter.php';
  require_once './FilterChain.php';
  require_once './MsgProcessor.php';

  $msg = "<p>hello, i am tom.how are you?</p>";
  $filter_chain = new FilterChain();
  $filter_chain->addFilter(new HTMLFilter())->addFilter(new LengthFilter());
  $msg_processor = new MsgProcessor($msg, $filter_chain);
  echo $msg_processor->process() . "\n";
