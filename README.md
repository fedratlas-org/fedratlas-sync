## fedratlas-sync

A distributed synchronization service for the **Fedratlas** platform, designed to enable efficient sharing and consistency of geospatial data across multiple nodes.

---

## 📌 Overview

`fedratlas-sync` is part of a larger system focused on managing and synchronizing geospatial datasets in a distributed environment. The service ensures that spatial data remains consistent across different peers using modern API standards and efficient data handling techniques.

---

## 🎯 Objectives

* Enable **data synchronization** between distributed nodes
* Support **geospatial data handling** using PostGIS
* Implement **OGC API standards** for interoperability
* Ensure **scalability and reliability** in distributed environments

---

## 🏗️ Architecture (High-Level)

The project follows a modular architecture:

* **API Layer** – Exposes endpoints (OGC API-based where applicable)
* **Service Layer** – Handles synchronization logic
* **Data Layer** – Manages PostgreSQL/PostGIS interactions
* **Peer Management** – Maintains registry of connected nodes

---

## 🛠️ Tech Stack

* **Language:** Go
* **Database:** PostgreSQL + PostGIS
* **ORM:** GORM
* **API Standards:** OGC API (Features, Records – planned usage)
* **Logging:** Zap (structured logging)

---

## 📁 Project Structure

```
fedratlas-sync/
│
├── cmd/
│   └── server/        # Application entry point
│
├── internal/
│   ├── api/           # API handlers & routing
│   ├── service/       # Business logic
│   ├── data/          # Database access layer
│   └── peer/          # Peer registry & sync logic
│
├── configs/           # Configuration files
├── scripts/           # Utility scripts
└── docs/              # Documentation
```

---

## 🚀 Current Progress

* ✅ Initial project structure setup
* ✅ Basic architecture design
* ✅ Database setup (PostgreSQL + PostGIS)
* ⏳ API implementation (in progress)
* ⏳ Synchronization logic (planned)

---

## 🔄 Synchronization Concept

The system is designed to:

* Track changes in spatial datasets
* Propagate updates across peer nodes
* Maintain eventual consistency

---

## 📡 OGC API Integration

Planned integration includes:

* OGC API - Features for spatial data access
* Standardized endpoints for interoperability
* Compatibility with external geospatial systems

---

## ⚙️ Getting Started

### Prerequisites

* Go (1.20+ recommended)
* PostgreSQL with PostGIS extension
* Git

### Setup

```bash
git clone https://github.com/your-username/fedratlas-sync.git
cd fedratlas-sync

go mod tidy
go run cmd/server/main.go
```

---

## 🧪 Future Improvements

* Full OGC API compliance
* Conflict resolution strategies
* Event-driven synchronization
* Authentication & authorization
* Performance optimizations

---

## 🤝 Contributing

This project is currently part of a capstone effort. Contributions and suggestions are welcome in future phases.

---

## 📄 License

To be decided.

---

## 👤 Author

Developed as part of a **Capstone Project** focusing on distributed geospatial systems and data synchronization.
