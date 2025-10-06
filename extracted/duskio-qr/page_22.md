# Page 22

## Text Content

```
def create_roll_document
Prawn::Document.new(
page_size: [POINTS_PER_INCH, POINTS_PER_INCH],
margin:

DEFAULT_ROLL_MARGIN

)
end
# Lay out one label per page
def generate_roll_labels(document)
# write the QR once to temp file
temp_path = write_image_to_temp_file(qr_code_data) if qr_code_data
@stock_quantity.times do |idx|
document.start_new_page unless idx.zero?
add_roll_label(document, temp_path) if temp_path && File.exist?(temp_path)
end
end
# Place QR and price on a page
def add_roll_label(document, image_path)
dims = calculate_dimensions(document)
# Draw the QR image
document.image image_path,
at: [dims[:x_position], dims[:y_position] + dims[:qr_size]],
width: dims[:qr_size]
# Draw the price text using BaseGenerator.format_price
text = Propel::Pdf::BaseGenerator.format_price(variant_price)
document.font_size 8 do
document.text_box text,
at:

[0, 12],

width: dims[:usable_width],
height: 10,
align: :center
end
end
# Compute sizes & positions
def calculate_dimensions(doc)
w = doc.bounds.width
h = doc.bounds.height
qr = [w, h * 0.8].min


```

