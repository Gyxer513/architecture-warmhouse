@startuml
title FitLife Web Application Component Diagram

top to bottom direction

!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

Container_Boundary(FitLifeSystem, "FitLife System") {
  Container(WebApp, "Web Application", "Java, Spring", "Handles user interactions")
  Container(Database, "Database", "PostgreSQL", "Stores user data and schedules")
}

Container(WebApp, "Web Application", "Java, Spring") {
  Component(AuthController, "AuthController", "Handles authentication and authorization")
  Component(UserController, "UserController", "Manages user profiles")
  Component(ScheduleController, "ScheduleController", "Manages schedules")
  Component(PaymentController, "PaymentController", "Handles payment processing")
  Component(ServiceLayer, "Service Layer", "Business logic")
  Component(RepositoryLayer, "Repository Layer", "Data access logic")
}

Rel(AuthController,ServiceLayer,"Calls business logic")
Rel(UserController,ServiceLayer,"Calls business logic")
Rel(ScheduleController,ServiceLayer,"Calls business logic")
Rel(PaymentController,ServiceLayer,"Calls business logic")
Rel(ServiceLayer,RepositoryLayer,"Reads/Writes data")
Rel(RepositoryLayer,Database,"Reads/Writes user data")
@enduml