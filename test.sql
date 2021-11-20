
CREATE TABLE "products" ("p_id" BIGSERIAL NOT NULL,
                                          "name" VARCHAR, "buy_price" REAL, "sell_price" REAL, "units" REAL, "type" VARCHAR, PRIMARY KEY ("p_id"));


CREATE TABLE "order_history" ("id" BIGSERIAL NOT NULL,
                                             "profit" REAL, "product_id" BIGINT, "order_id" BIGINT, "units" REAL, PRIMARY KEY ("id"));


CREATE TABLE "orders" ("order_id" BIGSERIAL NOT NULL,
                                            "created_at" TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
                                                                                      PRIMARY KEY ("order_id"));