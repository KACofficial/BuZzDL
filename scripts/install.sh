#!/bin/bash

# Install script 

go mod tidy

sudo go build -ldflags="-w -s" -o /bin/buzzdl

sudo chmod +x /bin/buzzdl


echo "To use BuZzDL, type 'buzzdl' in your terminal."

echo "Done!"