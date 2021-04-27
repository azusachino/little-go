#!/usr/bin/bash
files = $(ls)
for f in files 
  do
    sed -i 's/main/init/g' f
  done
