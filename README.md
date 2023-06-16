# Moneta
Moneta is a Web-based application to manage our daily cashflow.

## Techstack
| Golang | PostgreSQL | JWT | Gin | Github |
| --- | --- | --- | --- | --- |
| <a href="https://go.dev"><img width="100" src="https://github.com/c23-m4001/.github/raw/master/assets/go.png" /></a> | <a href="https://www.postgresql.org"><img width="100" src="https://github.com/c23-m4001/.github/raw/master/assets/postgresql.png" /></a> | <a href="https://jwt.io"><img width="100" src="https://github.com/c23-m4001/.github/raw/master/assets/jwt.svg" /></a> | <a href="https://gin-gonic.com"><img width="100" src="https://github.com/c23-m4001/.github/raw/master/assets/gin.png" /></a> | <a href="https://github.com"><img width="100" src="https://github.com/c23-m4001/.github/raw/master/assets/github.png" /></a> |

## Deployment

### Deployment Requirement

- `GeoLite2-City.mmdb` from https://dev.maxmind.com/geoip/geolite2-free-geolocation-data
- jwt `private.key` and `public.key` (can be generated from `make jwt-key-gen` command)

<br />

Deployment config please refer to: https://github.com/c23-m4001/deployment

## Command

### Make command
| Command | Usage |
| ---- | --- |
| make http |  Run http server |
| make jwt-key-gen | generate private and public key for JWT |
| make migrate | Migration for database |
| make migrate-fresh | Drop all tables and Migrate all tables |
| make migrate-gen | Generate a new Migration file |
| make seed | Seeder for development |
| make seed-production | Seeder for production |
| make generate-docs | Generate docs using Swaggo to /docs folder |

