#!/bin/bash

# Check if an argument is provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <folder_name>"
    exit 1
fi

# Get the argument as the folder name
FOLDER_NAME=$1
FILE_NAME="$FOLDER_NAME/$FOLDER_NAME.go"

# Create the directory
mkdir -p "$FOLDER_NAME"

# Create the Go file inside the directory
touch "$FILE_NAME"

# Add a basic Go template
echo "package main

import \"fmt\"

func main() {
    
}" > "$FILE_NAME"

# Confirmation message
echo "Folder '$FOLDER_NAME' and file '$FILE_NAME' created successfully."