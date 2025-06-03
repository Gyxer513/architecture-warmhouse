# Диаграмма компонентов Control Service

**Control Service** ― микросервиса, отвечающего за доставку команд управления на устройства и подтверждение их исполнения.

## Компоненты сервиса и их функциональность
- Control API REST/gRPC Controller — принимает запросы на управление устройствами (например, "включить отопление", "изменить режим").
- Command Processor Бизнес-логика: валидация команд, проверка ограничений, расписания и прав, формирование финального payload для устройства.
- Device Communicator Работа с физическими устройствами по нужным протоколам (MQTT, HTTP, CoAP и др.), отправка команд, ожидание отклика/квитирования.
- Command Status Tracker Отслеживание состояния отправленных команд: успех, ошибка, повторная попытка, таймаут.
- Event Publisher Публикация статусов/событий команд (успешно исполнено, неудача, подтверждение устройства и пр.) в message bus для других сервисов (например, уведомления, аудит).
- Event Listener Получение команд для устройств от других микросервисов по событийному message bus (например, от Automation Service).
- Control Repository Хранение истории команд, результатов исполнения, диагностических метрик, логов.
- Control DB База данных всех операций управления и логов.

```puml
@startuml
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

' Внешние сервисы и устройства
Person(user, "Пользователь")
Person(admin, "Администратор")
Container(message_bus, "Message Bus", "RabbitMQ/Kafka/NATS", "Event-driven интеграция")
System_Ext(iot_device, "IoT-устройство", "Датчики/исполнители (контроллеры, реле и др.)")

' Control Service как контейнер
Container_Boundary(control_srv, "Control Service") {
  Component(control_api, "Control API", "REST/gRPC Controller", "Внешний интерфейс для команд управления")
  Component(command_processor, "Command Processor", "Service", "Бизнес-логика, валидация команд")
  Component(device_communicator, "Device Communicator", "Integration", "Взаимодействие с устройствами (MQTT, HTTP и др.)")
  Component(status_tracker, "Command Status Tracker", "Service", "Мониторинг состояния команд, обработка результатов и ошибок")
  Component(event_publisher, "Event Publisher", "Integration", "Публикация событий исполнения/статусов команд")
  Component(event_listener, "Event Listener", "Integration", "Приём команд из message bus (например, от Automation Service)")
  Component(control_repository, "Control Repository", "Data Access", "CRUD для истории управления, статусов, логов")
  ComponentDb(control_db, "Control DB", "PostgreSQL/NoSQL", "База управления и логов")
}

' Внешние связи
Rel(user, control_api, "Запрос управления", "REST/gRPC")
Rel(admin, control_api, "Администрирование устройств", "REST/gRPC")
Rel(control_api, command_processor, "Передача команд")
Rel(command_processor, device_communicator, "Команды на устройство")
Rel(command_processor, status_tracker, "Мониторинг выполнения")
Rel(device_communicator, status_tracker, "Отправка статуса ответа")
Rel(status_tracker, control_repository, "Сохранение результата", "SQL/NoSQL")
Rel(control_repository, control_db, "Чтение/запись логов", "SQL/NoSQL")
Rel(status_tracker, event_publisher, "Передача статусов и событий")
Rel(event_publisher, message_bus, "Публикация статусов", "event")
Rel(event_listener, message_bus, "Получение команд от других сервисов", "event")
Rel(event_listener, command_processor, "Создание команд по событиям")
Rel(device_communicator, iot_device, "Отправка управляющих команд, получение откликов", "MQTT/HTTP")

@enduml
```