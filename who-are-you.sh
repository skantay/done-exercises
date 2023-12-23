
#!bin/bash

curl -s "https://01.alem.school/assets/superhero/all.json" | jq -r '.[] | select(.id==70) | "\"\(.name)\""'
