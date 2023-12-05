
import sys
import polars as pl

rows = []
for line in sys.stdin:
    rows.append(line.strip())
df = pl.DataFrame([rows], schema=["input"])

answer = df.with_columns(
    digits=pl.col("input").str.extract_all(r"(\d)")
).with_columns(
    first_digit=pl.col("digits").list.first().cast(pl.Int64),
    last_digit=pl.col("digits").list.last().cast(pl.Int64),
).with_columns(
    sum=pl.col("first_digit")*10 + pl.col("last_digit"),
)["sum"].sum()

print(answer)