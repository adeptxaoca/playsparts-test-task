# playsparts-test-task

## Description

Пример микросервиса для хранения и обработки абстрактных запчастей.  
Текст тестового задания: [test task description](./docs/test_task.md).  
Приложение тестировалось на _Ubuntu 18.04.4 LTS_ и _Windows 10_.  

## Installation

1. Убедиться, что уже установлены [Docker Engine](https://docs.docker.com/get-docker/) and [Docker
 Compose](https://docs.docker.com/compose/install/);  
2. Создать в директории проекта файл _.env_ и заполнить его согласно _.env.example_;
3. Из директории проекта запустить приложение путем выполнения команды _make deploy-up_.

## Migrations

Файлы миграций базы данных расположены в директории _internal/app/migrations_
. Все миграции будут автоматически выполнены при запуске приложения до последней версии. Текущая версия схемы будет отображена в логах.

## Other

Список дополнительных команд:
- проверка работы приложения путем запуска простого _grpc_ клиента с подготовленным сценарием: _make run-client_;
- генерация _proto_ файлов: _make generate_;
- доступны _docker-compose build, up, stop_ и _down_ командами _make deploy-build, make deploy-up, make deploy-stop_
 и _make deploy-down_ соответственно;
- процент покрытия кода unit-тестами: _make cover_. 