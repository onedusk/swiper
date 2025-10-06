# Page 2

## Text Content

```
EXAMPLE OF EXACTLY WHAT I NEED TO BE DONE - EXCEPT IT NEEDS TO BE DONE WITHIN THE
RODA APPLICATION
GENERATE THE PROMPTS I NEED TO IMPLEMENT AND ACCOMPLISH THESE TASKS
i use this script to generate and store all the urls for qr code generation => PLEASE ANALYZE FULLY
#!/usr/bin/env ruby
require 'rqrcode'
require 'chunky_png'
require 'sequel'
require 'dotenv/load'

# ensure your DB_CONNECTION_OPTIONS are loaded

require_relative '../config/db_config'
require_relative '../config/environment'
# Connect to your database (adjust DB_CONNECTION_OPTIONS in your .env or config)
DB = Sequel.connect(ENV['DATABASE_URL'] || DB_CONNECTION_OPTIONS)
# Update variants: generate the URL to be encoded in the QR code for variants.
DB[:variants].each do |variant|
qr_url = "https://duskioqr.fly.dev/api/v1/scan/#{variant[:id]}"
DB[:variants].where(id: variant[:id]).update(qr_code_content: qr_url)
puts "Updated product variant #{variant[:id]} with QR code URL: #{qr_url}"
end
# Update items: generate the URL to be encoded in the QR code for products (no variant
DB[:items].each do |item|
qr_url = "https://duskioqr.fly.dev/api/v1/scan/#{item[:id]}"
DB[:items].where(id: item[:id]).update(qr_code_content: qr_url)
puts "Updated item #{item[:id]} with QR code URL: #{qr_url}"
end
puts "QR code URL generation completed for all product variants and items."

now that all are generated, when a user clicks on the print label button / icon via the variants table we
need to match the products id via the variants table and then use the stored qr_content url to encode
the url. the amount of qr codes to generate is completely dependent on that variants ivnentory level /
stock qty.
below i will attach the variants table, the print button(still needs to be integrated into the variants
table, when a user clicks a checkbox or all checkboxes the button needs to appear beside the


```

