#!/bin/bash

date

if [ $? -eq 0 ]; then
    # Run the compiled program
    ../../cmd/nids_monitor
else
    echo "Compilation failed."
fi