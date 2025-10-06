# Page 42

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

Tip
The Shopify Theme Store requires themes to display the selling plan name in the cart. Be sure
to check whether the selling_plan.name is already present in the theme's cart.liquid
file before attempting to insert it.

Post-purchase
Apps that offer subscriptions must include navigation to a customer portal, both on the order status page
and through a post-purchase email to a merchant's customers so that they're able to manage their
subscription.
Customer portal
1. The customer portal must give each customer a single login to access subscriptions and their order
history.
2. The customer portal must display to each customer all of their purchased subscriptions. Details must

include the associated products, delivery frequency, price, and order schedule.

3. The customer portal must include an option for a merchant's customers to cancel their subscription.

The subscription app must allow the merchant to clearly communicate conditions of purchase on
their storefront's product page and customer portal.

4. The customer portal must provide a merchant's customers with the option to modify the payment

method associated with their subscriptions.

B. Shopify admin and in-app requirements
1. The purchase option app must use the app extension on the product page. Changes that are made to

purchase options from the Shopify admin must be reflected in the app.

2. Merchants need to be able to create

and manage purchase options in the Shopify admin using
the app extension. This includes letting merchants remove products from a selling plan.

3. Apps that offer subscriptions must include a direct link to orders and customers in the Shopify admin

from the purchase option.

4. Links to the subscription app from the orders and customers pages in the Shopify admin must go to

the correct subscription resource.

15. Donation distribution apps
A donation distribution app collects and distributes funds to a charity on behalf of a merchant.
1. The donation distribution app must use the Billing API or a PCI-compliant third-party gateway when

collecting donation funds from merchants through the app’s user interface.

2. If the donation distribution app allows merchants to collect charity donation funds from their

customers, then you must provide proof of charitable status in the app's user interface.

https://shopify.dev/docs/apps/launch/app-requirements-checklist

42/48


```

