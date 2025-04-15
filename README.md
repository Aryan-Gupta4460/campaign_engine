# campaign_engine
# Campaign Optimization Engine

This is a Go-based Campaign Optimization Engine designed to manage advertising campaigns across multiple platforms. It simulates real-time bidding, performs predictive analytics, and integrates with Kafka and PostgreSQL for messaging and storage.

## Features

- Real-time campaign simulation
- Predictive analytics for bid optimization
- Kafka producer integration
- PostgreSQL persistence using GORM
- RESTful API with Gin

## Prerequisites

- Go 1.18+
- PostgreSQL (recommended version: 14)
- Kafka (with Zookeeper)

## Installation

1. Clone the repository:
   
Install Go dependencies:

go mod tidy

Set up PostgreSQL:

createdb campaign_db

Run Kafka and Zookeeper (if using locally via Homebrew):

 start zookeeper
 start kafka
Run the service:

go run cmd/main.go
