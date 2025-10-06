# Page 11

## Text Content

```
-- migrations.sql
-- =========================================
-- Enable Necessary PostgreSQL Extensions
-- =========================================
CREATE EXTENSION IF NOT EXISTS pgcrypto;
CREATE EXTENSION IF NOT EXISTS vector;

-- =========================================
-- Migration 001: Create Products Table
-- =========================================
CREATE TABLE products (
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
handle VARCHAR,
title VARCHAR,
body TEXT,
vendor VARCHAR,
product_category VARCHAR,
product_type VARCHAR,
tags TEXT[] DEFAULT '{}',
published BOOLEAN DEFAULT FALSE,
status VARCHAR NOT NULL DEFAULT 'draft',
qr_code VARCHAR,
created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
shop VARCHAR,
product_id VARCHAR,
product_handle VARCHAR,
scans INTEGER DEFAULT 0,
destination VARCHAR DEFAULT 'product',
synced_at TIMESTAMP WITH TIME ZONE,
shopify_product_id VARCHAR,
raw_category_data JSONB DEFAULT '{}',
image_width INTEGER,
image_height INTEGER,
qr_code_content VARCHAR,
online_store_url VARCHAR,
online_store_preview_url VARCHAR,
body_html TEXT
);
-- Indexes for Products


```

