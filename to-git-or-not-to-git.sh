#!/usr/bin/env bash
curl https://api.github.com/users/MJoshua25 | grep '"id"' | cut -d : -f 2  | cut -d , -f1 | cut -d "" -f2
