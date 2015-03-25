CREATE TABLE "products" (
  "id" integer,
  "name" varchar(255),
  "type" integer,
  "description" varchar(255),
  "curr_quantity" integer,
  "min_quantity" integer , PRIMARY KEY (id)
);

CREATE TABLE "order_products" (
  "order_id" integer,
  "product_id" integer
);

CREATE TABLE "orders" (
  "id" integer,
  "approved" bool , PRIMARY KEY (id)
);
