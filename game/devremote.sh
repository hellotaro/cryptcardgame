#!/bin/bash

scp -r ubuntu@3jih.l.time4vps.cloud:~/project/gamecenter/dist .
rm -r static
mv dist static

