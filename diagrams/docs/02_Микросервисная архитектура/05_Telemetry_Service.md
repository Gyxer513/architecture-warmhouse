
# Диаграмма компонентов Telemetry Service

**Telemetry Service** отвечает за управление сценариями и расписаниями, позволяя пользователю создавать правила, 
принимает внешние запросы, реагирует на события из системы, 
планирует выполнение по времени, а также запускает и логирует действия, 
такие как отправка команд устройствам или уведомления пользователям.

## Компоненты сервиса и их функциональность
- Telemetry API REST/gRPC-контроллер, предоставляющий данные по телеметрии по запросу пользователей, других сервисов и админов.
- Telemetry Ingestor Получает, валидирует и агрегирует новые данные от устройств (например, PUSH через HTTP, MQTT-сообщения, Polling).
- Event Listener Подписчик на внешние события, если интеграция происходит через шину событий (например, интеграция с внешними IoT-платформами).
- Telemetry Repository Слой работы с БД: запись новых строк телеметрии и получение агрегированных данных (TimeSeries DB или аналог).
- Analytics Engine Выполняет on-the-fly расчёты, агрегации, генерацию метрик и, возможно, простейшие алерты (по запросу или в реальном времени).
- Event Publisher Публикует важные события в message bus (например, "вышли за порог температуры", "сработал алерт").
- Telemetry DB Таймсерийная база данных — хранение истории показаний.

```puml
@startuml
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

' Внешние акторы и системы
Person(user, "Пользователь")
Person(admin, "Администратор")
System_Ext(iot_device, "IoT-устройство", "Датчики")
Container(message_bus, "Message Bus", "RabbitMQ/Kafka/NATS", "Event-driven интеграция")

' Контейнер микросервиса Telemetry Service
Container_Boundary(telemetry_srv, "Telemetry Service") {
  Component(telemetry_api, "Telemetry API", "REST/gRPC Controller", "Запрос исторических данных и метрик пользователями и сервисами")
  Component(telemetry_ingestor, "Telemetry Ingestor", "Service", "Приём новых данных телеметрии c устройств (HTTP/MQTT/PUSH)")
  Component(event_listener, "Event Listener", "Integration", "Обработка телеметрии через событийную шину")
  Component(telemetry_repository, "Telemetry Repository", "Data Access", "Работа с TimeSeries DB: запись/чтение данных")
  Component(analytics_engine, "Analytics Engine", "Service", "Агрегация, вычисление метрик, оповещения")
  Component(event_publisher, "Event Publisher", "Integration", "Публикация событий и алертов в шину")
  ComponentDb(telemetry_db, "Telemetry DB", "InfluxDB/Clickhouse", "TimeSeries-хранилище истории телеметрии")
}

' Внешние взаимодействия
Rel(iot_device, telemetry_ingestor, "Передаёт данные", "MQTT/HTTP")
Rel(user, telemetry_api, "Запрашивает статистику и историю", "REST/gRPC")
Rel(admin, telemetry_api, "Анализ состояния", "REST/gRPC")

' Внутренние связи компонентов
Rel(telemetry_api, telemetry_repository, "Читает/агрегирует данные по запросу")
Rel(telemetry_ingestor, telemetry_repository, "Пишет новые значения")
Rel(event_listener, telemetry_repository, "Альтернативно может писать значения (через message bus)")
Rel(telemetry_repository, analytics_engine, "Передаёт данные для анализа")
Rel(analytics_engine, event_publisher, "Генерирует события, алерты")
Rel(event_publisher, message_bus, "Публикация событий и алертов", "event")
Rel(event_listener, message_bus, "Подписка на телеметрию/команды", "event")
Rel(telemetry_repository, telemetry_db, "Чтение/запись временных рядов", "TSDB")

@enduml
```