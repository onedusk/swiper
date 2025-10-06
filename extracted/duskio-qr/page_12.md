# Page 12

## Text Content

```
CREATE UNIQUE INDEX index_products_on_handle ON products(handle);
CREATE INDEX index_products_on_product_category ON products(product_category);
CREATE INDEX index_products_on_product_type ON products(product_type);
CREATE INDEX index_products_on_raw_category_data ON products USING GIN (raw_category_da
CREATE UNIQUE INDEX index_products_on_shopify_product_id ON products(shopify_product_id
CREATE INDEX index_products_on_tags ON products USING GIN (tags);
CREATE UNIQUE INDEX index_products_on_uuid ON products(id);
CREATE INDEX index_products_on_vendor ON products(vendor);

-- =========================================
-- Migration 002: Create Variants Table
-- =========================================
CREATE TABLE variants (
uuid UUID PRIMARY KEY DEFAULT gen_random_uuid(),
product_id UUID NOT NULL,
variant_sku VARCHAR,
variant_grams NUMERIC,
variant_inventory_tracker VARCHAR,
variant_inventory_policy VARCHAR,
variant_fulfillment_service VARCHAR,
variant_price NUMERIC,
variant_compare_at_price NUMERIC,
variant_requires_shipping BOOLEAN,
variant_taxable BOOLEAN,
variant_barcode VARCHAR,
image_src VARCHAR,
image_position INTEGER,
image_alt_text VARCHAR,
gift_card BOOLEAN,
seo_title VARCHAR,
seo_description VARCHAR,
variant_image VARCHAR,
variant_weight_unit VARCHAR,
variant_tax_code VARCHAR,
cost_per_product NUMERIC,
included_us BOOLEAN,
price_us NUMERIC,
compare_at_price_us NUMERIC,
included_international BOOLEAN,
price_international NUMERIC,
compare_at_price_international NUMERIC,
product_collection VARCHAR,


```

