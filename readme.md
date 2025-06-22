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
- **Code Quality**: Clean, maintainable code following Go best practices
- **Documentation**: Comprehensive API documentation with usage examples
- **Open Source Workflow**: Git workflows, branch management, and collaborative development

---

## 🔄 Program Status

**Current Progress**: 2/X Assignments Completed  
**Status**: In Progress - Continuing with advanced API concepts and real-world applications

---

## 📚 Learning Outcomes

Through this fellowship, I've gained valuable experience in:
- Modern API development patterns and best practices
- Open-source collaboration and contribution workflows
- Database design and integration strategies
- Technical documentation and communication skills
- Production-ready code development and testing

---

## 🤝 Connect & Collaborate

Feel free to explore the code, suggest improvements, or connect for discussions about API development and open-source contribution!

**GitHub**: [@utkarshkrsingh](https://github.com/utkarshkrsingh)

---

*This journey continues as I dive deeper into advanced API concepts, testing strategies, and production deployment practices. Stay tuned for more exciting developments!* 🚀
