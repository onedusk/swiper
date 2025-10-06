# Page 40

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

Note
As of March 9th, 2020, all sales channels (with the exception of mobile app builders) must have
the read_only_own_orders scope applied. The read_only_own_orders scope is added
by the review team during the approval process and ensures that a channel can read only the
orders it created. Make sure that your channel is requesting only orders it created for a faster
review and approval.

Shopify supports a variety of ways of building sales channels. The way that you decide to build can
determine who is responsible for payment processing, order fulfillment, and refunds.
Build a sales channel using cart permalinks
Build your sales channel with cart permalinks. These links take customers who want to buy a product
from the sales channel directly to a merchant's store checkout to complete the purchase.
1. Take customers to Shopify's checkout with items pre-loaded in the cart.
2. Use the Billing API. For sales attribution, you can use a storefront access token.

D. Navigation icon
1. The sales channel must include a 16px by 16px navigation icon in SVG format, uploaded through the

Partner Dashboard .

2. The icon must be a single color with a transparent background.
3. The icon's SVG file should be less than 2KB.
4. The icon's SVG file can contain only the following permitted tags: circle , ellipse , g , line ,
path , rect , svg , title .
5. The icon's SVG file can contain only the following permitted attributes: cx , cy , d , height ,
opacity , pathLength , points , r , rx , ry , version , viewBox , width , x1 , x2 , xmlns ,
y1 , y2 , fill-rule , clip-rule .

14. Purchase option apps
A purchase option app provides merchants and customers with various ways to sell and buy products,
beyond the "buy now, pay now, and ship now" experience. For example, merchants can sell a product as
a one-time purchase, a recurring subscription, or a pre-order.

A. Storefront requirements
General principles
https://shopify.dev/docs/apps/launch/app-requirements-checklist

40/48


```

