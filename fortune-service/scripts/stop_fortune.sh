#!/bin/bash
echo "finding fortune process..."
if pgrep -f fortune-linux &> /dev/null 2>&1; then
    echo "stopping fortune process..."
    killall fortune-linux
fi