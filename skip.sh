#!/usr/bin/env bash
ls -1 | awk '{ if(NR%2==0){print}}'