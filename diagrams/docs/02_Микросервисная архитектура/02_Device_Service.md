# Диаграмма компонентов Device Service

**Device Service** — отвечает за управление, регистрацию и настройку устройств, 
а также за внутреннее представление и хранение состояний.

## Компоненты сервиса и их функциональность
- Device API REST/gRPC контроллеры, принимающие внешние запросы (например, добавить/удалить устройство, обновить настройки, просмотреть статус).
- Command Handler Слой бизнес-логики, обрабатывающий команды с API и события от других сервисов (например, обрабатывать изменение статусов, валидацию, формирование событий для message bus).
- Device State Manager Хранит текущее состояние устройств, обеспечивает чтение-запись из базы данных, синхронизирует состояние с командами и событиями.
- Event Publisher Генерирует события об изменениях устройств и публикует их в message bus для других микросервисов.
- Device Repository Слой доступа к базе данных со всеми методами CRUD для устройств (может быть реализован через ORM или на прямых SQL/NoSQL-запросах).
- Integration Adapter (optional) Модуль для интеграции с внешними/партнёрскими устройствами или API.


```puml
@startuml
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml


' Внешние акторы и контексты
Person(admin, "Администратор")
Person(user, "Пользователь")
Container(message_bus, "Message Bus", "RabbitMQ/Kafka/NATS", "Event-driven интеграция")

' Микросервис как контейнер
Container_Boundary(device_srv, "Device Service") {
  Component(device_api, "Device API", "REST/gRPC Controller", "Внешний HTTP/gRPC интерфейс, приём команд от пользователей и админов")
  Component(command_handler, "Command Handler", "Service", "Бизнес-логика, обработка команд и валидация действий")
  Component(state_manager, "Device State Manager", "Service", "Работа с текущим состоянием устройств, синхронизация с БД")
  Component(repository, "Device Repository", "Data Access", "CRUD-доступ к БД устройств")
  Component(event_publisher, "Event Publisher", "Integration", "Публикация событий об изменениях устройств в шину событий")
  Component(ext_adapter, "Integration Adapter", "Integration", "Адаптер для внешних API/протоколов устройств (опционально)")
  ' БД как boundary
  ComponentDb(device_db, "Devices DB", "PostgreSQL", "Хранение устройств и их параметров")
}

' Взаимодействие внешних пользователей
Rel(user, device_api, "Взаимодействует через Web/API", "REST/gRPC")
Rel(admin, device_api, "Управляет устройствами", "REST/gRPC")

' Внутренние связи компонентов
Rel(device_api, command_handler, "Передаёт команды/запросы")
Rel(command_handler, state_manager, "Обрабатывает изменения состояния")
Rel(state_manager, repository, "Читает и пишет", "SQL")
Rel(state_manager, event_publisher, "Сообщает об изменениях")
Rel(event_publisher, message_bus, "Публикует события", "event")
Rel(command_handler, ext_adapter, "Интеграция (если нужно)")
Rel(repository, device_db, "Хранение и получение данных", "SQL")

@enduml
```