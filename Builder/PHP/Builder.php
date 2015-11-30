<?php
  interface Builder {
    public function buildHeader($header);

    public function buildBody($body);

    public function buildFooter($footer);
  }
