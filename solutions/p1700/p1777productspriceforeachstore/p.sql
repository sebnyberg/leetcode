select * from products
pivot (
    sum(price)
          for store in (store1, store2, store3)
    ) as pivot_table
