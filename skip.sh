#!/usr/bin/env bash
ls -l | awk '{ if(NR%2==0 or NR==1){print}}'