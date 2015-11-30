<?php
  require_once './Builder.php';
  require_once './TextBuilder.php';
  require_once './XmlBuilder.php';
  require_once './Director.php';

  $header = "title->test";
  $body = "test builder";
  $footer = "end";

  $text_builder = new TextBuilder();
  $text_director = new Director($text_builder);
  $text_director->construct($header, $body, $footer);

  $xml_builder = new XmlBuilder();
  $xml_director = new Director($xml_builder);
  $xml_director->construct($header, $body, $footer);
  echo $text_builder->getResult();
  echo $xml_builder->getResult();
