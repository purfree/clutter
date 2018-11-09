#### ES分组

{
  "size": 0,
  "aggs": {
​    "group_by_state": {
​      "terms": {
​        "field": "source_date",
​        "size": 100
​      }
​    }
  }
}

类似于mysql语句

select source_date as group_by_state from user_hard_income_before_check group by source_date

