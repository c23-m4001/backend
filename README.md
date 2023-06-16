# Moneta
Moneta is a Web-based application to manage our daily cashflow.

| Golang | PostgreSQL | JWT | Gin | Github |
| --- | --- | --- | --- | --- |
|  <img width="100" src="https://raw.githubusercontent.com/get-icon/geticon/fc0f660daee147afb4a56c64e12bde6486b73e39/icons/go.svg" /> | <img width="100" src="https://raw.githubusercontent.com/get-icon/geticon/fc0f660daee147afb4a56c64e12bde6486b73e39/icons/postgresql.svg" /> | <img width="100" src="https://cdn.worldvectorlogo.com/logos/jwt-3.svg" /> | <img width="100" src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" /> | <img width="100" src="https://raw.githubusercontent.com/get-icon/geticon/fc0f660daee147afb4a56c64e12bde6486b73e39/icons/github.svg" /> |

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

