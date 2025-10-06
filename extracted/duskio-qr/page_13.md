# Page 13

## Text Content

```
created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
shop VARCHAR,
product_variant_id VARCHAR,
synced_at TIMESTAMP WITH TIME ZONE,
shopify_variant_id VARCHAR,
variant_inventory_quantity INTEGER NOT NULL DEFAULT 0,
barcode_data BYTEA,
qr_code_data BYTEA,
code_generated_at TIMESTAMP WITH TIME ZONE,
code_format VARCHAR,
image_width INTEGER,
image_height INTEGER,
inventory_quantity INTEGER,
qr_code_content VARCHAR,
variant_image_url VARCHAR
);
-- Indexes for Variants
CREATE INDEX index_variants_on_code_generated_at ON variants(code_generated_at);
CREATE UNIQUE INDEX index_variants_on_shopify_variant_id ON variants(shopify_variant_id
CREATE UNIQUE INDEX index_variants_on_product_and_variant_sku ON variants(product_id, v
CREATE INDEX index_variants_on_product_id ON variants(product_id);
CREATE INDEX index_variants_on_barcode_and_format ON variants(variant_barcode, code_form

-- =========================================
-- Migration 003: Create Product Images Table
-- =========================================
CREATE TABLE product_images (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
shopify_image_id VARCHAR UNIQUE,
alt_text TEXT,
url VARCHAR NOT NULL,
created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Indexes for Product Images
CREATE INDEX index_product_images_on_product_id ON product_images(product_id);
CREATE UNIQUE INDEX index_product_images_on_shopify_image_id ON product_images(shopify_


```

