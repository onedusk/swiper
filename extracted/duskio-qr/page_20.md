# Page 20

## Text Content

```
"qr_code_#{variant_identifier || Time.now.to_i}.pdf"
end
private
# Inserts the QR image
def add_qr_code(document)
temp_path = write_image_to_temp_file(qr_code_data)
document.image temp_path, fit: @qr_size, position: options[:position] || :cen
end
# Inserts SKU, price, and optional extra text
def add_details(document)
document.move_down 10
if variant_sku
document.text "SKU: #{variant_sku}", size: 12, align: :center
end
if variant_price
text = Propel::Pdf::BaseGenerator.format_price(variant_price)
document.text text, size: 16, align: :center
end
if options[:additional_text]
document.move_down 5
document.text options[:additional_text], size: 10, align: :center
end
end
def show_details?
options[:show_details] != false &&
(variant_sku || variant_price || options[:additional_text])
end
# Required data
def validate_qr_data!
unless qr_code_data || options[:skip_validation]
raise Propel::Errors::Error, "QR code data must be provided"
end
end
def variant_identifier; options[:variant_identifier]; end


```

