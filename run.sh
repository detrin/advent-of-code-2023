#!/bin/bash

# Initialize variables
lang=""
day=""
task=""

# Parse command-line options
while getopts ":l:d:t:r:h" opt; do
  case $opt in
    l) lang="$OPTARG";;
    d) day="$OPTARG";;
    t) task="$OPTARG";;
    r) runs="$OPTARG";;
    h) 
      echo "Benchmarking script usage:"
      echo "-l: Language of the task ('rust' or 'go')."
      echo "-d: Day of the task."
      echo "-t: Task number."
      echo "-r: Number of runs."
      echo "-h: Print help."
      exit 0
      ;;
    \?) echo "Invalid option -$OPTARG" >&2; exit 1;;
    :) echo "Option -$OPTARG requires an argument." >&2; exit 1;;
  esac
done

if [[ "$lang" == "" ]]; then
    echo "Language not specified."
    exit 1
fi

if [[ "$day" == "" ]]; then
    echo "Day not specified."
    exit 1
fi

if [[ "$task" == "" ]]; then
    echo "Task not specified."
    exit 1
fi

if [[ "$runs" == "" ]]; then
    runs=500
fi

# Run the appropriate command
if [[ "$lang" == "rust" ]]; then
    cd ./rust/$day/$task
    cat ../../../data/$day/input.txt | cargo run
    cd - > /dev/null
elif [[ "$lang" == "go" ]]; then
    cd ./go/$day
    cat ../../data/$day/input.txt | go run ${task}.go
    cd - > /dev/null
else
    echo "Unknown language: $lang"
    exit 1
fi
