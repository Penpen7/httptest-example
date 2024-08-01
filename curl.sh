#!/bin/bash

curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "age": 13}' http://localhost:8080/api/users/1
