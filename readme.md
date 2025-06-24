# 🚀 Keploy API Fellowship Journey

Welcome to my coding adventure with **Keploy's API Fellowship Program**! This repository documents my journey through hands-on assignments designed to master API development, open-source contribution, and modern software engineering practices.

---

## 📋 Program Overview

The Keploy API Fellowship is an intensive program focused on:
- **Open Source Contribution** - Building a strong foundation in collaborative development
- **API Development Mastery** - Creating robust, scalable backend systems
- **Real-world Project Experience** - Practical application of modern development practices

---

## ✅ Completed Assignments

### 🎯 Assignment 1: Open Source 101

**Objective**: Establish a strong open-source presence and contribute meaningfully to the community.

#### Task 1: GitHub Profile Enhancement ✨
Created a comprehensive GitHub profile README that showcases:
- **Technical Skills**: Backend development, systems programming, and modern frameworks
- **Project Portfolio**: Highlighting key projects and contributions
- **Professional Presence**: Clean, organized presentation for recruiters and collaborators

🔗 **Profile**: [github.com/utkarshkrsingh](https://github.com/utkarshkrsingh)

#### Task 2: Open Source Contribution 🤝
Successfully contributed to the Keploy ecosystem with a technical improvement:

**Project**: `keploy/samples-go` - Echo MySQL Sample  
**Contribution**: Refactored short link entropy generation using raw hash bytes  
**Impact**: Improved performance by eliminating unnecessary conversions while maintaining security  

🔗 **Pull Request**: [#153 - Refactor: improve short link entropy using raw hash bytes](https://github.com/keploy/samples-go/pull/153)

**Technical Details**:
- Replaced `uint64` conversion with direct `base58` encoding of hash bytes
- Enhanced randomness and reduced collision probability for short codes
- Maintained compact output (~8 characters) while improving efficiency

---

### 🎯 Assignment 2: Master of API

**Objective**: Design and implement a complete API server with database integration and comprehensive documentation.

#### 🎌 Project: Anime Watchlist API
A full-featured REST API for managing personal anime collections with robust CRUD operations.

**Tech Stack**:
- **Backend**: Go (Golang) with Gin framework
- **Database**: MySQL for reliable data persistence
- **Architecture**: RESTful API design with proper HTTP status codes

**Key Features**:
- ✅ **Complete CRUD Operations** - Create, read, update, and delete anime entries
- ✅ **Data Validation** - Comprehensive input validation and error handling  
- ✅ **Status Management** - Track watching progress with multiple status types
- ✅ **Comprehensive Documentation** - Detailed API reference with examples
- ✅ **Production Ready** - Proper error handling and response formatting

**API Endpoints**:
- `POST /watchlist` - Add new anime to watchlist
- `GET /watchlist` - Retrieve all anime entries
- `PUT /watchlist/{id}` - Update existing anime
- `DELETE /watchlist/{id}` - Remove anime from watchlist

🔗 **Repository**: [golang-watchlist](https://github.com/utkarshkrsingh/keploy-api-fellowship/tree/main/golang-watchlist)

---

## 🛠️ Technical Skills Demonstrated

- **API Design**: RESTful architecture with proper HTTP methods and status codes
- **Database Integration**: Efficient data modeling and query optimization
- **Testing Excellence**: Comprehensive testing with 79.1% coverage across unit, integration, and API tests
- **Quality Assurance**: Error handling, edge case validation, and robust testing strategies
- **Code Quality**: Clean, maintainable code following Go best practices
- **Documentation**: Comprehensive API documentation with usage examples
- **Open Source Workflow**: Git workflows, branch management, and collaborative development

---

### 🎯 Assignment 3: Testing Excellence

**Objective**: Implement comprehensive testing strategies for the API server with thorough coverage analysis.

#### 🧪 Testing Implementation
Enhanced the Anime Watchlist API with robust testing suite covering multiple testing layers:

**Testing Strategy**:
- **Unit Tests**: Core business logic validation with 70%+ coverage target
- **Integration Tests**: Database interaction verification for CRUD operations
- **API Tests**: End-to-end endpoint functionality validation

**Coverage Achievement**: **79.1% Total Coverage** 📊

**Detailed Coverage Breakdown**:
- `NewDB`: 73.7% - Database connection and initialization
- `NewRecordRepository`: 100.0% - Repository pattern implementation
- `CreateRecord`: 77.8% - Anime creation logic
- `GetRecords`: 80.0% - Data retrieval operations
- `UpdateRecord`: 75.0% - Anime update functionality
- `DeleteRecord`: 87.5% - Deletion operations
- `HandleRecordRoutes`: 100.0% - Route handler setup
- **API Handlers**:
  - `CreateRecord` endpoint: 69.2%
  - `GetRecords` endpoint: 66.7%
  - `UpdateRecord` endpoint: 76.5%
  - `DeleteRecord` endpoint: 100.0%

**Testing Tools & Frameworks**:
- Go's built-in testing package for unit tests
- Database mocking for isolated testing
- HTTP testing for API endpoint validation
- Coverage analysis with `go tool cover`

**Key Testing Features**:
- ✅ **Comprehensive CRUD Testing** - All database operations thoroughly tested
- ✅ **Error Handling Validation** - Edge cases and error scenarios covered
- ✅ **HTTP Response Testing** - Status codes and response format verification
- ✅ **Database Integration Testing** - Real database interaction validation
- ✅ **Mocking Strategies** - Isolated unit testing with proper mocks

---

## 🔄 Program Status

**Current Progress**: 3/X Assignments Completed  
**Status**: In Progress - Mastering testing methodologies and quality assurance practices

---

## 📚 Learning Outcomes

Through this fellowship, I've gained valuable experience in:
- Modern API development patterns and best practices
- Comprehensive testing strategies and quality assurance methodologies
- Open-source collaboration and contribution workflows
- Database design and integration strategies
- Test-driven development and coverage analysis
- Technical documentation and communication skills
- Production-ready code development with robust error handling

---

## 🤝 Connect & Collaborate

Feel free to explore the code, suggest improvements, or connect for discussions about API development and open-source contribution!

**GitHub**: [@utkarshkrsingh](https://github.com/utkarshkrsingh)

---
