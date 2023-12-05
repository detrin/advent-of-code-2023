
import sys
import polars as pl

rows = []
for line in sys.stdin:
    rows.append(line.strip())
df = pl.DataFrame([rows], schema=["input"])

numbers_map = {}
for i, num in enumerate(["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]):
    numbers_map[num] = num + str(i+1) + num

df = df.with_columns(replaced=pl.col("input"))
for num, digit in numbers_map.items():
    df = df.with_columns(replaced=pl.col("replaced").str.replace_all(num, digit, literal=True))

answer = df.with_columns(
    digits=pl.col("replaced").str.extract_all(r"(\d)")
).with_columns(
    first_digit=pl.col("digits").list.first().cast(pl.Int64),
    last_digit=pl.col("digits").list.last().cast(pl.Int64),
).with_columns(
    sum=pl.col("first_digit")*10 + pl.col("last_digit"),
)["sum"].sum()

print(answer)