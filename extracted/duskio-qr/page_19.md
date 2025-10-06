# Page 19

## Text Content

```
# propel/lib/propel/pdf/generators/qr_code.rb
require_relative 'base'
require 'prawn'
require 'rqrcode'
require 'chunky_png'
module Propel
module Pdf
module Generators
# Single‐label QR code PDF generator (1.25″ × 1.25″)
class QrCode < Base
DEFAULT_QR_SIZE = [450, 450]
DEFAULT_MARGIN

= [40, 40, 40, 40]

def initialize(options = {})
super
validate_qr_data!
@qr_size = options[:qr_size] || DEFAULT_QR_SIZE
end
def generate_pdf
Propel.logger.info "Generating QR code PDF - #{qr_context.inspect}"
document = create_document(
page_size: options[:page_size] || :letter,
margin:

options[:margin]

|| DEFAULT_MARGIN

)
if qr_code_data
add_qr_code(document)
add_details(document) if show_details?
else
document.text "No QR Code Available",
align: :center,
size: 16
end
document.render
end
def generate_filename


```

