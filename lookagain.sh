#!/bin/bash

find . -type f -name "*.sh" -exec basename {} \; | sort -r |cut -d. -f1
