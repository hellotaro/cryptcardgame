#!/bin/bash

scp -r ubuntu@3jih.l.time4vps.cloud:~/projects/mynuxt/dist .
rm -r static
mv dist static

