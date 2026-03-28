<?php

use Symfony\Component\Dotenv\Dotenv;

require dirname(__DIR__).'/vendor/autoload.php';

(new Dotenv())->bootEnv(dirname(__DIR__).'/.env');

$_SERVER += $_ENV;

$_SERVER['APP_ENV'] = $_ENV['APP_ENV'] ?? 'dev';

$_SERVER['APP_DEBUG'] = $_SERVER['APP_DEBUG'] ?? $_ENV['APP_DEBUG'] ?? 'prod' !== $_SERVER['APP_ENV'];

$_SERVER['APP_DEBUG'] = $_SERVER['APP_DEBUG'] ? '1' : '0';