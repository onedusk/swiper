# Page 14

## Text Content

```
-- =========================================
-- Migration 004: Create Image Annotations Table
-- =========================================
CREATE TABLE image_annotations (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
product_image_id UUID NOT NULL REFERENCES product_images(id) ON DELETE CASCADE,
raw_annotations JSONB DEFAULT '{}',
label_annotations JSONB DEFAULT '{}',
text_annotations JSONB DEFAULT '{}',
landmark_annotations JSONB DEFAULT '{}',
face_annotations JSONB DEFAULT '{}',
created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Indexes for Image Annotations
CREATE INDEX index_image_annotations_on_product_image_id ON image_annotations(product_im
CREATE INDEX index_image_annotations_on_raw_annotations ON image_annotations USING GIN

-- =========================================
-- Migration 005: Create Product Options Table
-- =========================================
CREATE TABLE product_options (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
product_id UUID NOT NULL REFERENCES products(id) ON DELETE CASCADE,
shopify_option_id VARCHAR,
name VARCHAR NOT NULL,
values TEXT[] DEFAULT '{}',
created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- Indexes for Product Options
CREATE INDEX index_product_options_on_product_id ON product_options(product_id);
CREATE INDEX index_product_options_on_name ON product_options(name);

-- =========================================
-- Migration 007: Create Variant Selected Options Table
-- =========================================


```

