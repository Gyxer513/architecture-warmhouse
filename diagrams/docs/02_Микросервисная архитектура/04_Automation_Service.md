# Диаграмма компонентов Automation Service
**Automation Service** реализует пользовательские сценарии, расписания и обработку 
событий по автоматизации управления устройствами (например, "если температура ниже X — включить отопление").  
## Компоненты сервиса и их функциональность
- Automation API — Контроллер, принимающий REST/gRPC-запросы на управление сценариями и расписаниями.
- Rule Engine — Ядро бизнес-логики: обрабатывает сценарии, триггеры, условия, действия.
- Scheduler — Планировщик: хранит расписания, срабатывает в нужное время.
- Scenario Repository — Доступ к сценариям, условиям, действиям и метаданным из БД (CRUD).
- Event Listener — Подписчик на события из других сервисов по message bus (например, новые данные с датчика, событие пользователя).
- Action Dispatcher — Отвечает за отправку команд в Control Service или другие сервисы (для выполнения действий по сценарию).
- Event Publisher — Публикует события о выполнении сценариев (например, для логгирования или уведомлений).
- Automation DB — База данных для сценариев, расписаний и состояний.

```puml
@startuml
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

' Внешние акторы и сервисы
Person(user, "Пользователь")
Person(admin, "Администратор")
Container(message_bus, "Message Bus", "RabbitMQ/Kafka/NATS", "Event-driven интеграция")
Container(control_service, "Control Service", "Go", "Выполнение команд управления устройствами")

' Automation Service как контейнер
Container_Boundary(automation_srv, "Automation Service") {
  Component(automation_api, "Automation API", "REST/gRPC Controller", "Внешний HTTP/gRPC-интерфейс для управления сценариями")
  Component(rule_engine, "Rule Engine", "Service", "Обработка сценариев, условий, бизнес-логика автоматизации")
  Component(scheduler, "Scheduler", "Service", "Планировщик срабатывания сценариев по времени или событиям")
  Component(scenario_repository, "Scenario Repository", "Data Access", "CRUD для сценариев, условий, расписаний")
  Component(event_listener, "Event Listener", "Integration", "Подписка на события из других сервисов через message bus")
  Component(action_dispatcher, "Action Dispatcher", "Integration", "Отправка команд для выполнения действий")
  Component(event_publisher, "Event Publisher", "Integration", "Публикация событий выполнения сценариев")
  ComponentDb(automation_db, "Automation DB", "MongoDB/Redis", "БД сценариев, расписаний и состояний автоматизации")
}

' Внешние взаимодействия
Rel(user, automation_api, "Управление сценариями и расписаниями", "REST/gRPC")
Rel(admin, automation_api, "Администрирование автоматизации", "REST/gRPC")

' Внутренние связи компонентов
Rel(automation_api, rule_engine, "Передача команд и запросов")
Rel(rule_engine, scheduler, "Планирование задач")
Rel(rule_engine, scenario_repository, "Загрузка/сохранение сценариев")
Rel(rule_engine, event_listener, "Получение внешних событий (срабатывания)")
Rel(rule_engine, action_dispatcher, "Отправка команд на выполнение действий")
Rel(scheduler, action_dispatcher, "Запуск действий по расписанию")
Rel(action_dispatcher, control_service, "Команды для устройств", "REST/gRPC/event")
Rel(action_dispatcher, event_publisher, "Генерация событий об исполнении сценариев")
Rel(event_publisher, message_bus, "Публикация событий", "event")
Rel(event_listener, message_bus, "Подписка на события (например, телеметрия, действия пользователей)", "event")
Rel(scenario_repository, automation_db, "Чтение/запись сценариев и расписаний", "NoSQL/Redis")

@enduml
```