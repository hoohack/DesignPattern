<?php
  require_once './Article.php';
  require_once './Author.php';
  require_once './LifeArticle.php';

  $life_article = new LifeArticle();
  $life_article->setTitle("test");
  $author_hector = new Author("hector");
  $life_article->setAuthor($author_hector);

  $article = clone $life_article;
  $article->setTitle("Programmer life");
  var_dump($article);
  var_dump($life_article);

  echo $article->getTitle() . "\n";
  $article->getAuthor()->setName("haha");
  echo $article->getAuthor()->getName() . "\n";
  echo $life_article->getAuthor()->getName() . "\n";
  echo $life_article->getTitle() . "\n";
