#!/usr/bin/env bash

runs1=1000

go_d01_t1=$(./benchmark.sh -l go -d day01 -t task1 -r $runs1 2>/dev/null | tail -n1)
go_d01_t2=$(./benchmark.sh -l go -d day01 -t task2 -r $runs1 2>/dev/null | tail -n1)
go_d02_t1=$(./benchmark.sh -l go -d day02 -t task1 -r $runs1 2>/dev/null | tail -n1)
go_d02_t2=$(./benchmark.sh -l go -d day02 -t task2 -r $runs1 2>/dev/null | tail -n1)
go_d03_t1=$(./benchmark.sh -l go -d day03 -t task1 -r $runs1 2>/dev/null | tail -n1)
go_d03_t2=$(./benchmark.sh -l go -d day03 -t task2 -r $runs1 2>/dev/null | tail -n1)
go_d04_t1=$(./benchmark.sh -l go -d day04 -t task1 -r $runs1 2>/dev/null | tail -n1)
go_d04_t2=$(./benchmark.sh -l go -d day04 -t task2 -r $runs1 2>/dev/null | tail -n1)

rust_d01_t1=$(./benchmark.sh -l rust -d day01 -t task1 -r $runs1 2>/dev/null | tail -n1)
rust_d01_t2=$(./benchmark.sh -l rust -d day01 -t task2 -r $runs1 2>/dev/null | tail -n1)
rust_d02_t1=$(./benchmark.sh -l rust -d day02 -t task1 -r $runs1 2>/dev/null | tail -n1)
rust_d02_t2=$(./benchmark.sh -l rust -d day02 -t task2 -r $runs1 2>/dev/null | tail -n1)
rust_d03_t1=$(./benchmark.sh -l rust -d day03 -t task1 -r $runs1 2>/dev/null | tail -n1)
rust_d03_t2=$(./benchmark.sh -l rust -d day03 -t task2 -r $runs1 2>/dev/null | tail -n1)
rust_d04_t1=$(./benchmark.sh -l rust -d day04 -t task1 -r $runs1 2>/dev/null | tail -n1)
rust_d04_t2=$(./benchmark.sh -l rust -d day04 -t task2 -r $runs1 2>/dev/null | tail -n1)

echo "| Task       | Rust      | Go        |"
echo "|------------|-----------|-----------|"
echo "| Day 1 Task 1 | $rust_d01_t1 | $go_d01_t1 |"
echo "| Day 1 Task 2 | $rust_d01_t2 | $go_d01_t2 |"
echo "| Day 2 Task 1 | $rust_d02_t1 | $go_d02_t1 |"
echo "| Day 2 Task 2 | $rust_d02_t2 | $go_d02_t2 |"
echo "| Day 3 Task 1 | $rust_d03_t1 | $go_d03_t1 |"
echo "| Day 3 Task 2 | $rust_d03_t2 | $go_d03_t2 |"
echo "| Day 4 Task 1 | $rust_d04_t1 | $go_d04_t1 |"
echo "| Day 4 Task 2 | $rust_d04_t2 | $go_d04_t2 |"
