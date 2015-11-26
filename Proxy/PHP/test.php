<?php
  require_once './IDBQuery.php';
  require_once './DBQuery.php';
  require_once './DBQueryProxy.php';

  $idbq = new DBQueryProxy();
  $idbq->request();
