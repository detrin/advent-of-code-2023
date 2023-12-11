#!/bin/bash

# Initialize variables
lang=""
day=""
task=""

# Parse command-line options
while getopts ":l:d:t:a:h" opt; do
  case $opt in
    l) lang="$OPTARG";;
    d) day="$OPTARG";;
    t) task="$OPTARG";;
    a) all="$OPTARG";;
    h) 
      echo "Benchmarking script usage:"
      echo "-l: Language of the task ('rust' or 'go')."
      echo "-d: Day of the task."
      echo "-t: Task number."
      echo "-a: Run all tasks."
      echo "-h: Print help."
      exit 0
      ;;
    \?) echo "Invalid option -$OPTARG" >&2; exit 1;;
    :) echo "Option -$OPTARG requires an argument." >&2; exit 1;;
  esac
done

task1_answers=(
    0 
    55172       # Day01
    2679        # Day02
    560670      # Day03
    24542       # Day04
    309796150   # Day05
    1313763     # Day06
    24845353130 # Day07
    18157       # Day08
    1806615041  # Day09
    7005        # Day10
)
task2_answers=(
    0 
    54925           # Day01
    77607           # Day02
    91622824        # Day03
    8736438         # Day04
    50716416        # Day05
    3412343720      # Day06
    24845353130     # Day07
    142997638331814 # Day08
    1211            # Day09
    417             # Day10
    702152204842    # Day11
)

if [[ "$all" == "true" ]]; then
    for day_i in $(seq 1 4); do
        for lang in rust go; do
            day_index=$(printf "%02d" $day_i)
            day="day$day_index"
            for task in task1 task2; do
                echo "Validating $lang $day $task"
                result=$(/bin/bash run.sh -l "$lang" -d "$day" -t "$task" | tail -n1)
                echo "Result: $result"

                if [ "$task" == "task1" ]; then
                    expected=${task1_answers["$day_index"]}
                else
                    expected=${task2_answers["$day_index"]}
                fi

                echo "Expected: $expected"

                if [ "$result" == "$expected" ]; then
                    echo "Success!"
                else
                    echo "Failure!"
                    exit 1
                fi
            done
        done
    done
    exit 0
else
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

    day_index=${day#day}  
    echo "Validating $lang $day $task"
    result=$(/bin/bash run.sh -l "$lang" -d "$day" -t "$task" | tail -n1)
    echo "Result: $result"

    if [ "$task" == "task1" ]; then
        expected=${task1_answers["$day_index"]}
    else
        expected=${task2_answers["$day_index"]}
    fi

    echo "Expected: $expected"

    if [ "$result" == "$expected" ]; then
        echo "Success!"
    else
        echo "Failure!"
        exit 1
    fi
fi


