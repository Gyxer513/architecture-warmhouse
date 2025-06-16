# Диаграмма компонентов User Service

## Компоненты сервиса и их функциональность
- User API — REST/gRPC-контроллер, принимающий и отдающий данные клиентским приложениям и проверенным сервисам.
- Authentication Handler — Принимает запросы на вход, обработку токенов и паролей, двухфакторную аутентификацию, если требуется.
- Manager — Проверяет права доступа, привилегии, формирует или валидирует скоупы (roles/permissions).
- Profile Manager — Обработка изменений профиля, предоставление информации о пользователе, смена пароля/адреса и т.д.
- User Repository — Взаимодействие с БД: хранение, обновление, поиск учетных записей.
- Event Publisher — Публикует события (например, "Пользователь зарегистрирован", "Пользователь изменил e-mail") для оповещения других сервисов через message bus.

```puml
@startuml
!includeurl https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml

' Внешние акторы и контексты
Person(user, "Пользователь")
Person(admin, "Администратор")
Container(message_bus, "Message Bus", "RabbitMQ/Kafka/NATS", "Event-driven интеграция")

' Микросервис как контейнер
Container_Boundary(user_srv, "User Service") {
  Component(user_api, "User API", "REST/gRPC Controller", "Внешний HTTP/gRPC интерфейс для работы с пользователями")
  Component(auth_handler, "Authentication Handler", "Service", "Аутентификация, проверка паролей, генерация access/refresh токенов")
  Component(authz_manager, "Authorization Manager", "Service", "Проверка прав, ролей, scopes")
  Component(profile_manager, "Profile Manager", "Service", "Работа с профилями, изменение персональных данных")
  Component(user_repository, "User Repository", "Data Access", "Доступ к данным о пользователях в БД")
  Component(event_publisher, "Event Publisher", "Integration", "Публикация событий во внешние сервисы (например, регистрация, смена пароля)")
  ComponentDb(user_db, "Users DB", "PostgreSQL", "Данные об учетных записях, ролях, правах")
}

' Внешние связи
Rel(user, user_api, "Авторизация и управление профилем", "REST/gRPC")
Rel(admin, user_api, "Администрирование пользователей", "REST/gRPC")

' Внутренние связи компонентов
Rel(user_api, auth_handler, "Передача запросов на вход/выход")
Rel(user_api, authz_manager, "Проверка прав пользователя")
Rel(user_api, profile_manager, "Просмотр/изменение профиля")
Rel(profile_manager, user_repository, "Читает и пишет профиль пользователя")
Rel(auth_handler, user_repository, "Проверка паролей, аутентификация")
Rel(authz_manager, user_repository, "Проверка прав, ролей")
Rel(profile_manager, event_publisher, "Публикация событий профиля")
Rel(auth_handler, event_publisher, "Публикация событий авторизации")
Rel(event_publisher, message_bus, "Публикация событий", "event")
Rel(user_repository, user_db, "Чтение/запись данных", "SQL")

@enduml
```