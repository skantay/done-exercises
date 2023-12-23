#!/bin/bash

tree | tail +2 | head -n -1 | wc -l
