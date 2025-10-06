# Page 43

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

3. The donation distribution app must provide proof to a merchant that the funds collected from a

merchant's customers are donated to a registered charitable organization. You can't use a tax receipt
as proof.

4. The donation distribution app must collect funds from a merchant's customers only through the

Shopify checkout.

5. Add a widget to buy the donation product to the product page, cart page or checkout page. This can

be implemented using Theme App Extensions or Checkout UI Extensions.

6. Your app must include instructions on how to hide the add-to-cart button

that is created.

for any donation product

7. The operating cost must be clearly indicated in both the UI and listing.

16. Payments apps
A payments app integrates with the Shopify admin to provide payment processing services.

A. Requirements for third-party payments apps
Third-party payments apps must meet the minimum product requirements in addition to the following
requirements:
1. Revenue share agreement: All Partners are required to have a signed revenue share agreement with

Shopify. You must sign and submit the agreement before Shopify can approve a payments app to
process payments.

2. API Usage: Your payment app must not use any Shopify APIs other than the Payments Apps API.
3. Payment app compatibility: All Partners must submit screencasts of the app's payment flow for all

supported browsers .

4. Cancelling payments: Your app must allow buyers to cancel or abandon the payment and be

redirected back to Shopify’s checkout.

5. Buyer flow redirections (off-site payment apps): You must update your app’s buyer flow on desktop

and mobile devices to redirect from Shopify’s checkout to your app’s payment flow, and then back to
Shopify’s order confirmation page.

6. Off-site payment information: Your payment app must present identical payment information to what

is displayed to the buyer at checkout.

7. Off-site redirects: Must not upsell any product or features in the payment flow
8. Redirection after install: Your payments app must redirect back to the Shopify admin
( https://{shop}.myshopify.com/services/payments_partners/gateways/${api_key}/se
ttings ) after it's installed. After redirecting to that page, the merchant will then immediately be

redirected to the payments app's corresponding page in the Shopify admin.

9. Restricted payment methods: Your payment app must not process payment methods that include,

but aren't limited to, Apple Pay, Google Pay, Shop Pay, PayPal, and Alipay. Shopify has a direct
connection with providers that improves performance and checkout conversion for merchants.

10. No embedding: Shopify prohibits the embedding of payment apps in the Shopify admin.
https://shopify.dev/docs/apps/launch/app-requirements-checklist

43/48


```

