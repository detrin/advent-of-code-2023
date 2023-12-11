#!/usr/bin/env bash

runs1=5000
runs2=100
runs3=500

# go_d01_t1=$(./benchmark.sh -l go -d day01 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d01_t2=$(./benchmark.sh -l go -d day01 -t task2 -r $runs1 2>/dev/null | tail -n1)
# go_d02_t1=$(./benchmark.sh -l go -d day02 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d02_t2=$(./benchmark.sh -l go -d day02 -t task2 -r $runs1 2>/dev/null | tail -n1)
# go_d03_t1=$(./benchmark.sh -l go -d day03 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d03_t2=$(./benchmark.sh -l go -d day03 -t task2 -r $runs1 2>/dev/null | tail -n1)
# go_d04_t1=$(./benchmark.sh -l go -d day04 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d04_t2=$(./benchmark.sh -l go -d day04 -t task2 -r $runs1 2>/dev/null | tail -n1)
# go_d05_t1=$(./benchmark.sh -l go -d day05 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d05_t2=$(./benchmark.sh -l go -d day05 -t task2 -r $runs1 2>/dev/null | tail -n1)
# go_d06_t1=$(./benchmark.sh -l go -d day06 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d06_t2=$(./benchmark.sh -l go -d day06 -t task2 -r $runs1 2>/dev/null | tail -n1)
# go_d07_t1=$(./benchmark.sh -l go -d day07 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d07_t2=$(./benchmark.sh -l go -d day07 -t task2 -r $runs1 2>/dev/null | tail -n1)
# go_d08_t1=$(./benchmark.sh -l go -d day08 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d08_t2=$(./benchmark.sh -l go -d day08 -t task2 -r $runs1 2>/dev/null | tail -n1)
go_d09_t1=$(./benchmark.sh -l go -d day09 -t task1 -r $runs1 2>/dev/null | tail -n1)
go_d09_t2=$(./benchmark.sh -l go -d day09 -t task2 -r $runs1 2>/dev/null | tail -n1)
# go_d10_t1=$(./benchmark.sh -l go -d day10 -t task1 -r $runs3 2>/dev/null | tail -n1)
# go_d10_t2=$(./benchmark.sh -l go -d day10 -t task2 -r $runs3 2>/dev/null | tail -n1)
# go_d11_t1=$(./benchmark.sh -l go -d day11 -t task1 -r $runs1 2>/dev/null | tail -n1)
# go_d11_t2=$(./benchmark.sh -l go -d day11 -t task2 -r $runs1 2>/dev/null | tail -n1)

# rust_d01_t1=$(./benchmark.sh -l rust -d day01 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d01_t2=$(./benchmark.sh -l rust -d day01 -t task2 -r $runs1 2>/dev/null | tail -n1)
# rust_d02_t1=$(./benchmark.sh -l rust -d day02 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d02_t2=$(./benchmark.sh -l rust -d day02 -t task2 -r $runs1 2>/dev/null | tail -n1)
# rust_d03_t1=$(./benchmark.sh -l rust -d day03 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d03_t2=$(./benchmark.sh -l rust -d day03 -t task2 -r $runs1 2>/dev/null | tail -n1)
# rust_d04_t1=$(./benchmark.sh -l rust -d day04 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d04_t2=$(./benchmark.sh -l rust -d day04 -t task2 -r $runs1 2>/dev/null | tail -n1)
# rust_d05_t1=$(./benchmark.sh -l rust -d day05 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d05_t2=$(./benchmark.sh -l rust -d day05 -t task2 -r $runs1 2>/dev/null | tail -n1)
# rust_d06_t1=$(./benchmark.sh -l rust -d day06 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d06_t2=$(./benchmark.sh -l rust -d day06 -t task2 -r $runs1 2>/dev/null | tail -n1)
# rust_d07_t1=$(./benchmark.sh -l rust -d day07 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d07_t2=$(./benchmark.sh -l rust -d day07 -t task2 -r $runs1 2>/dev/null | tail -n1)
# rust_d08_t1=$(./benchmark.sh -l rust -d day08 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d08_t2=$(./benchmark.sh -l rust -d day08 -t task2 -r $runs1 2>/dev/null | tail -n1)
rust_d09_t1=$(./benchmark.sh -l rust -d day09 -t task1 -r $runs1 2>/dev/null | tail -n1)
rust_d09_t2=$(./benchmark.sh -l rust -d day09 -t task2 -r $runs1 2>/dev/null | tail -n1)
# rust_d10_t1=$(./benchmark.sh -l rust -d day10 -t task1 -r $runs3 2>/dev/null | tail -n1)
# rust_d10_t2=$(./benchmark.sh -l rust -d day10 -t task2 -r $runs3 2>/dev/null | tail -n1)
# rust_d11_t1=$(./benchmark.sh -l rust -d day11 -t task1 -r $runs1 2>/dev/null | tail -n1)
# rust_d11_t2=$(./benchmark.sh -l rust -d day11 -t task2 -r $runs1 2>/dev/null | tail -n1)

# polars_d01_t1=$(./benchmark.sh -l polars -d day01 -t task1 -r $runs2 2>/dev/null | tail -n1)
# polars_d01_t2=$(./benchmark.sh -l polars -d day01 -t task2 -r $runs2 2>/dev/null | tail -n1)

echo "| Task       | Rust      | Go        | Polars    |"
echo "|------------|-----------|-----------|-----------|"
echo "| Day 1 Task 1 | $rust_d01_t1 | $go_d01_t1 | $polars_d01_t1 |"
echo "| Day 1 Task 2 | $rust_d01_t2 | $go_d01_t2 | $polars_d01_t2 |"
echo "| Day 2 Task 1 | $rust_d02_t1 | $go_d02_t1 |"
echo "| Day 2 Task 2 | $rust_d02_t2 | $go_d02_t2 |"
echo "| Day 3 Task 1 | $rust_d03_t1 | $go_d03_t1 |"
echo "| Day 3 Task 2 | $rust_d03_t2 | $go_d03_t2 |"
echo "| Day 4 Task 1 | $rust_d04_t1 | $go_d04_t1 |"
echo "| Day 4 Task 2 | $rust_d04_t2 | $go_d04_t2 |"
echo "| Day 5 Task 1 | $rust_d05_t1 | $go_d05_t1 |"
echo "| Day 5 Task 2 | $rust_d05_t2 | $go_d05_t2 |"
echo "| Day 6 Task 1 | $rust_d06_t1 | $go_d06_t1 |"
echo "| Day 6 Task 2 | $rust_d06_t2 | $go_d06_t2 |"
echo "| Day 7 Task 1 | $rust_d07_t1 | $go_d07_t1 |"
echo "| Day 7 Task 2 | $rust_d07_t2 | $go_d07_t2 |"
echo "| Day 8 Task 1 | $rust_d08_t1 | $go_d08_t1 |"
echo "| Day 8 Task 2 | $rust_d08_t2 | $go_d08_t2 |"
echo "| Day 9 Task 1 | $rust_d09_t1 | $go_d09_t1 |"
echo "| Day 9 Task 2 | $rust_d09_t2 | $go_d09_t2 |"
echo "| Day 10 Task 1 | $rust_d10_t1 | $go_d10_t1 |"
echo "| Day 10 Task 2 | $rust_d10_t2 | $go_d10_t2 |"
echo "| Day 11 Task 1 | $rust_d11_t1 | $go_d11_t1 |"
echo "| Day 11 Task 2 | $rust_d11_t2 | $go_d11_t2 |"