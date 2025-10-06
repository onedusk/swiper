# Page 23

## Text Content

```
{
usable_width:

w,

usable_height: h,
qr_size:

qr,

x_position:

(w - qr) / 2,

y_position:

h - 2 - qr

}
end
def roll_context
qr_context.merge(
stock_quantity: @stock_quantity,
roll_dimensions: [POINTS_PER_INCH, POINTS_PER_INCH]
)
end
end
end
end
end


```

