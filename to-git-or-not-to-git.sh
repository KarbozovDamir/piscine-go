#!/bin/bash
curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"damirkap89\"}}){id}}"}' | grep -Eo '[0-9]{1,4}'
