#!/bin/zsh

# Check if the correct number of arguments is provided
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <image_file>"
    exit 1
fi

# Assign the provided image file path to a variable
image_file="$1"

# Specify the path to your Python script
python_script="python/production_run.py"

# Check if the Python script exists
if [ ! -f "$python_script" ]; then
    echo "Error: Python script not found at $python_script"
    exit 1
fi

# Execute the Python script with the specified image parameter
python "$python_script" --image "$image_file"
