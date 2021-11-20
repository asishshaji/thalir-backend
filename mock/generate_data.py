import psycopg2
import pandas as pd
import random

conn = psycopg2.connect(
    host="localhost",
    database="postgres",
    user="postgres",
    password="kai1253", port=5433)

df = pd.read_csv("data.csv")
items = df.Item_Name
buy_price = [i + 10 for i in range(items.size)]
sell_price = [i + 30 for i in range(items.size)]

phone = ["23123123", "23232323345", "343123231", "9999999999", "4444444"]
units = [10, 15, 20, 16, 12, 39]
buy_units = [1, 2, 3, 4]
cur = conn.cursor()

orderIds = [i for i in range(11, 20)]
productIds = [i for i in range(1, 10000)]


def insert_products():
    for i in range(items.size):
        sql = """INSERT INTO products(name,type,buy_price,sell_price,units)
                VALUES(%s,%s,%s,%s,%s);"""

        if i % 2 == 0:
            tp = "P"
        else:
            tp = "R"
        if items[i] == 'NaN':
            continue
        print(items[i], tp, buy_price[i],
              sell_price[i], random.choice(units))
        cur.execute(sql, (items[i], tp, buy_price[i],
                          sell_price[i], random.choice(units)))
        conn.commit()


def insert_orders():
    for i in range(10):
        sql = """INSERT INTO orders(phone_number)
                VALUES(%s);"""
        cur.execute(sql, (random.choice(phone),))
        conn.commit()


def insert_order_details():
    for i in range(1, 2000):
        sql = """INSERT INTO order_details(order_id,product_id,units,profit)
                    VALUES(%s,%s,%s,%s);"""
        u = random.choice(buy_units)
        pid = random.choice(productIds)
        cur.execute("""SELECT * FROM products WHERE p_id = %s""", [pid])
        row = cur.fetchone()
        sell_p = row[4]
        buy_p = row[3]
        profit = (sell_p - buy_p) * u

        cur.execute(sql, (random.choice(orderIds), pid, u, profit))
        conn.commit()


# insert_products()
# insert_orders()
# insert_order_details()

cur.close()
