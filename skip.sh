#!/usr/bin/env bash
ls -l | awk '{ if(NR%2==1){print}}'