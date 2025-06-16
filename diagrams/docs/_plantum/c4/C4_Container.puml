@startuml
title FitLife Container Diagram

top to bottom direction

!includeurl https://raw.githubusercontent.com/RicardoNiepel/C4-PlantUML/master/C4_Component.puml

Person(user, "User", "A user of the fitness club system")
Person(admin, "Administrator", "An administrator managing the system")
System(FitLifeSystem, "FitLife System", "System managing memberships, schedules, and payments")

Container_Boundary(FitLifeSystem, "FitLife System") {
  Container(WebApp, "Web Application", "Java, Spring", "Handles user interactions")
  Container(MobileApp, "Mobile Application", "Kotlin, Swift", "Allows users to manage their membership")
  Container(PaymentService, "Payment Service", "Node.js", "Processes payments")
  Container(Database, "Database", "PostgreSQL", "Stores user data and schedules")
}

System_Ext(api, "Third-Party API", "External API for fitness data integration")
System_Ext(bank, "Bank System", "External bank for processing payments")

Rel(user, WebApp, "Uses the system")
Rel(admin, WebApp,"Manages the system")
Rel(WebApp,PaymentService,"Processes payment requests")
Rel(PaymentService,bank,"Processes payments")
Rel(WebApp,Database,"Reads/Writes user data")
Rel(MobileApp,Database,"Reads/Writes user data")
Rel(WebApp,api,"Fetches fitness data")
@enduml