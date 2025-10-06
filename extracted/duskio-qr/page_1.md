# Page 1

## Text Content

```
To ensure your QR code remains linked to the correct product even if the QR code URL or Shopify
GID changes, you can rely on the persistent relationship stored in your database. Here’s how it
works:
1. Persistent Product Identifier:
Each product has a stable internal ID (stored in your products table). The product variants
reference this internal product ID (via the product_id column). This means that even if external
details change, the internal link between a variant and its product remains intact.
2. Decoupled QR Code and GID:
The QR code itself contains only a short URL that references the variant ID. When the QR code is
scanned, your application uses that variant ID to look up the record, which includes both the
persistent product ID (via product_id) and the current Shopify GID. If Shopify updates the GID,
you only need to update that column in your variant record without changing the QR code
content.
3. Dynamic Updates:
Since the QR code is generated on the fly (or its URL is stored and later used to build the
redirect), you can update the mapping in your database whenever there’s a change. This means:
If the QR code content (URL) is updated, your system will generate a new QR code when
needed.
If the Shopify GID changes, you update the variant’s record with the new GID while the QR
code URL remains the same.
The product’s association via the persistent product ID remains constant unless the product
is deleted.
By designing your system this way, you decouple the external identifiers (QR code URL and Shopify
GID) from the internal product identity. This decoupling ensures that, as long as the product exists,
even if external identifiers change, your database mapping (the combination of the variant’s internal
ID, the product’s internal ID, and the current Shopify GID) will always point to the correct product.
In summary, your database stores all three pieces of information:
The QR code URL (which is simple and stable),
The persistent product ID (from the products table),
The Shopify GID (which may change over time).
Even if the QR code URL or the Shopify GID is updated, the product variant record still maintains a
connection to the same product via the persistent product ID, ensuring that the link remains valid
unless the product is deleted.


```

