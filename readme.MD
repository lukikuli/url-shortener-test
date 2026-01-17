# Distributed URL Shortener

## 1. Architecture Overview

### High-Level Design

This service implements a **Distributed URL Shortener** using **Clean Architecture with Domain-Driven Design (DDD)** principles.

The architecture enforces strict separation of concerns to ensure:

- Testability
- Replaceable infrastructure (in-memory → managed cloud storage)
- Deterministic business logic

### Logical Layers

┌──────────────────────────┐
│ HTTP Layer │ (Gin Handlers, Middleware)
│ - Request validation │
│ - Response formatting │
│ - Observability headers │
└───────────┬──────────────┘
│
┌───────────▼──────────────┐
│ Use Case Layer │
│ - ShortenURL │
│ - RedirectURL │
│ - Stats │
│ - Orchestration logic │
└───────────┬──────────────┘
│
┌───────────▼──────────────┐
│ Domain Layer │
│ Entities: UrlShorten │
│ Value Objects: LongURL │
│ Domain Services │
│ Repository Interfaces │
└───────────┬──────────────┘
│
┌───────────▼──────────────┐
│ Infrastructure Layer │
│ In-Memory Repository │
│ (Replaceable with any DB)│
└──────────────────────────┘

---

### Data Model

**UrlShorten Entity**

| Field              | Description                     |
| ------------------ | ------------------------------- |
| `short_code`       | Generated human-safe short code |
| `long_url`         | Validated HTTP/HTTPS URL        |
| `created_at`       | Creation timestamp              |
| `expires_at`       | TTL-based expiration            |
| `click_count`      | Atomic redirect counter         |
| `last_accessed_at` | Last redirect timestamp         |

---

## 2. Gap Analysis — In-Memory vs Cloud Storage

### Why In-Memory Storage Is Not Suitable for Production

The current implementation uses an In-Memory storage for simplicity and testability.  
However, this approach is unsuitable for a stateless cloud environment because:

- Volatility
- No scalability
- No durability ~ No crash recovery

### Proposed Managed Storage

**Recommended Production Stack**

| Component       | Choice                       | Reason                                                                          |
| --------------- | ---------------------------- | ------------------------------------------------------------------------------- |
| Primary storage | MySQL / PostgreSQL           | ACID, Strong consistency, Transactions support, Good Performance using Indexing |
| Cache           | Redis / ElastiCache / Valkey | Hot-path redirects                                                              |
| TTL cleanup     | Native TTL feature           | No manual cron required                                                         |

**How this solves the gap:**

- Strong consistency for writes
- Automatic TTL expiration
- Atomic counters for click tracking

---

## 3. Capacity Planning

### Storage Estimation (12 Months)

Assumptions:

- 100 million new URLs per month
- 12 months retention
- Average record size ≈ **300 bytes**
- 100,000,000 × 12 = 1.2 billion records
- 1.2B × 300 bytes ≈ 360 GB

**Estimated storage: 360~460 GB**

### Redirect Endpoint Scaling (10,000 RPS)

**Scaling Strategy**

1. **Read-through cache (Redis)**
   - Key: `short_code`
   - Value: `long_url`
   - TTL aligned with expiration

2. **Write-behind counters**
   - Batch click count updates
   - Reduce DB write amplification

3. **Stateless compute**
   - Horizontal scaling behind Load Balancer
   - No session affinity

## 4. Service Level Management

### Service Level Indicators (SLIs)

| SLI              | Description                            |
| ---------------- | -------------------------------------- |
| Redirect Latency | Time to complete `GET /s/{code}`       |
| Availability     | Successful redirects vs total requests |

---

### Service Level Objectives (SLOs)

| SLI                    | Target          |
| ---------------------- | --------------- |
| Redirect latency (p99) | ≤ 50 ms         |
| Availability           | ≥ 99.9% monthly |

---

## 5. Future Enhancements

Given more time and resources, the following enhancements would be prioritized to improve reliability, scalability, and observability:

1. **Caching Improvements**
   - Introduce **TTL jitter** for Redis cache entries to prevent cache stampedes during high traffic.
   - Implement **hybrid caching strategies**, e.g., warm-up popular URLs, to further reduce database load.

2. **Managed Storage and Persistence**
   - Migrate from in-memory storage to a persistent managed database (MySQL/PostgreSQL) for durability, backup, and analytics.
   - Use Redis ElastiCache or Valkey as a caching layer to accelerate frequent lookups.

3. **Infrastructure as Code**
   - Implement Terraform templates for provisioning cloud resources in a reproducible manner.
   - Note: Limited experience with Terraform currently; additional time would be required to define serverless compute, managed storage, and IAM roles.

4. **Observability and Monitoring**
   - Extend metrics, tracing, and logging for better production visibility.
   - Include alerting for high error rates, expired URL hits, or performance anomalies.

5. **Scalability and Reliability**
   - Explore **horizontal scaling** with load balancers to support high QPS (10,000+ requests/sec).
   - Implement background jobs or hybrid TTL expiration strategies to efficiently clean up expired URLs.
   - Introduce rate limiting and abuse detection for URL creation and redirection endpoints.

---
