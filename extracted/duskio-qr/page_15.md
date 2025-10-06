# Page 15

## Text Content

```
CREATE TABLE variant_selected_options (
id UUID PRIMARY KEY DEFAULT gen_random_uuid()
variant_id UUID NOT NULL REFERENCES variants(
name VARCHAR NOT NULL,
value VARCHAR NOT NULL,
created_at TIMESTAMP WITH TIME ZONE NOT NULL
updated_at TIMESTAMP WITH TIME ZONE NOT NULL
);
-- Indexes for Variant Selected Options
CREATE INDEX index_variant_selected_options_on_variant_id ON variant_selected_options(v
CREATE INDEX index_variant_selected_options_on_name_and_value ON variant_selected_optio
-- =========================================
-- Migration 008: Create Inventories Table
-- =========================================
CREATE TABLE inventories (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
shopify_product_id VARCHAR NOT NULL,
handle VARCHAR NOT NULL,
title VARCHAR NOT NULL,
option1_name VARCHAR,
option1_value VARCHAR,
option2_name VARCHAR,
option2_value VARCHAR,
option3_name VARCHAR,
option3_value VARCHAR,
sku VARCHAR NOT NULL,
hs_code VARCHAR,
coo VARCHAR,
location VARCHAR NOT NULL,
incoming INTEGER DEFAULT 0,
unavailable INTEGER DEFAULT 0,
committed INTEGER DEFAULT 0,
available INTEGER DEFAULT 0,
on_hand INTEGER DEFAULT 0
);
-- Indexes for Inventories
CREATE INDEX index_inventories_on_handle ON inventories(handle);
CREATE UNIQUE INDEX index_inventories_on_sku_and_location ON inventories(sku, location)
CREATE INDEX index_inventories_on_product_id ON inventories(product_id);


```

