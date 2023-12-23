#!/bin/bash
url="https://01.alem.school/assets/superhero/all.json"

curl -s "$url" | jq ".[] | select(.id == $HERO_ID)" | jq '.connections["relatives"]' | tr -d '"'
