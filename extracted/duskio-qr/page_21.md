# Page 21

## Text Content

```
def variant_sku;

options[:variant_sku];

end

def variant_price;

options[:variant_price];

end

def qr_code_data;

options[:qr_code_data];

end

def qr_context
{
variant_identifier: variant_identifier,
variant_sku:

variant_sku,

has_qr:

!qr_code_data.nil?,

show_details:

show_details?,

qr_size:

@qr_size,

page_size:

options[:page_size] || :letter

}
end
end
# Continuous‐roll PDF generator: one 1.25″×1.25″ page per label
class QrCodeRoll < QrCode
DEFAULT_ROLL_MARGIN = [2, 2, 2, 2]

# 2-point edges

def initialize(options = {})
super
@stock_quantity = options[:stock_quantity].to_i
end
protected
# Overridden entry point
def generate_pdf
Propel.logger.info "Generating roll QR codes - #{roll_context.inspect}"
document = create_roll_document
generate_roll_labels(document)
document.render
end
def generate_filename
"qr_roll_#{variant_identifier || Time.now.to_i}.pdf"
end
private
# Create a 1.25″×1.25″ Prawn document for roll labels


```

