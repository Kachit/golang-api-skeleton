# Golang Rest API skeleton
[![Build Status](https://app.travis-ci.com/Kachit/golang-api-skeleton.svg?branch=master)](https://app.travis-ci.com/Kachit/golang-api-skeleton)
[![Codecov](https://codecov.io/gh/Kachit/golang-api-skeleton/branch/master/graph/badge.svg?token=L1DIXLCL4s)](https://codecov.io/gh/Kachit/golang-api-skeleton)
[![Go Report Card](https://goreportcard.com/badge/github.com/kachit/golang-api-skeleton)](https://goreportcard.com/report/github.com/kachit/golang-api-skeleton)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kachit/golang-api-skeleton/blob/master/LICENSE)

### Uses:
* gin as http framework
* gorm v2 as ORM
* go-fractal as serializer
* go-hashids as numerical ID obfuscation

### Commands list:
| Command | Description | Launch |
| ------ | ------ |------ |
| **develop:test** | Testable command | manually |
| **database:migrations:migrate** | Apply migrations | manually |
| **database:migrations:rollback** | Rollback migrations | manually |
| **database:seeders:seed** | Seed dev data to DB | manually |
| **database:seeders:clear** | Clear dev data | manually |
| **server-api:start** | API WebServer launch | manually |

### Launch
* **./golang-api-skeleton {command}** - simple launch
* **./golang-api-skeleton {command} -config=./config.yml** - launch with args