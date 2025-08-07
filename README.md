# Best Logging Packages in Go: Exploration & Comparison

Structured logging is a crucial practice for modern Go applications. Selecting the right logging package impacts your app’s performance, observability, and maintainability.

This repository provides an in-depth review, comparisons, benchmarks, and practical examples of the leading structured logging libraries available for Go developers.

## Table of Contents

- [Overview](#overview)
- [Supported Libraries](#supported-libraries)
- [Why Structured Logging?](#why-structured-logging)
- [Benchmark Summary](#benchmark-summary)
- [How To Use This Repo](#how-to-use-this-repo)
- [Advanced Production Setup](#advanced-production-setup)
- [Contributing](#contributing)
- [License](#license)

## Overview

Structured logging adds powerful context to your logs, enabling better searchability and integration with monitoring and alerting tools. This repo dives into:

- Native Go structured logging via `log/slog`
- High-performance logging with Zerolog and Uber’s Zap
- Legacy-friendly and stable logging with Logrus
- Benchmarks comparing speed, memory allocation, and throughput
- Production-ready configurations, including rotation and formatting

## Supported Libraries

| Library       | Description                                  | Recommended For                              |
|---------------|----------------------------------------------|---------------------------------------------|
| **log/slog**  | Go standard library structured logging       | Dependency-free native Go logging            |
| **Zerolog**   | Zero-allocation, blazing-fast JSON logger    | High-performance, microservices, cloud-native |
| **Zap**       | Uber’s performant, configurable logger       | Latency-sensitive & high-throughput systems |
| **Logrus**    | Veteran, easy-to-use logger                   | Legacy codebases, ease of migration          |

## Why Structured Logging?

- Enables **machine-readable** logs (JSON and key-value pairs)
- Eases **searching**, **filtering**, and **tracing** in log aggregation platforms
- Adds **context** to events for better observability (user IDs, request IDs, etc.)
- Facilitates integration with modern **monitoring and tracing** tools

## Benchmark Summary

You’ll find detailed benchmark results inside to compare:

| Library        | Speed (ns/op) | Memory Allocations | Performance Notes                   |
|----------------|---------------|--------------------|------------------------------------|
| Zerolog        | Fastest       | Lowest             | Best for high-volume microservices |
| Zap (Core)     | Very Fast     | Low                | Configurable and highly performant |
| Zap (Sugared)  | Fast          | Slightly higher     | Ergonomic API with minimal cost    |
| Logrus         | Slowest       | Highest            | Stable but resource-heavy           |

## How To Use This Repo

- Browse the sample code in each library’s folder or the examples directory
- Review configuration and setup guides to get started quickly
- Use benchmark results to select the correct logger for your needs
- Try the production-ready Zerolog + Lumberjack setup for log rotation and formatting

## Advanced Production Setup

We recommend combining **Zerolog** with the **Lumberjack** package for log rotation and multi-format logging. This setup ensures:

- **Human-readable console logs** during development
- **Structured JSON logs** in rotated files for production ingestion
- Automated **log size management**, backup, and compression
- Smooth integration with containers, Kubernetes, serverless, and observability pipelines

Example configurations and code are provided in the `zerolog-lumberjack` directory.

## Contributing

Contributions are welcome! Please feel free to:

- Add or update example code
- Add benchmarks for newer versions or other logging packages
- Suggest improvements to documentation or code quality

Open issues or pull requests to participate in this community-driven repository.

*Happy logging, and may your applications be ever observable!* 🚀
