# Page 35

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

11. Product sourcing
A product sourcing app lets merchants find and sell a wide range of products by providing product
discovery and sales features directly in the app.

A. Product sourcing
1. Product sourcing apps are exempt from using the Billing API for the sale of goods to their merchants,

and can instead use a PCI compliant gateway. However, any other costs associated with the app must
be charged using the Billing API.

2. If your app fulfills product orders on behalf of a merchant, then it must not automatically fulfill orders

that are in the pending payment state.

3. Your app must add the cost of goods to the Cost field on the product page of the merchant's Shopify

admin.

4. Your app must not sell high-risk products. Products that violate Shopify's Acceptable Use Policy

and the Terms of Service for Payment Providers are prohibited. Products like cannabis, alcohol,
pharmaceutical drugs, weapons and items listed as prohibited businesses are included in this
restriction.

5. Your app allows merchants to request fulfillment. Use the
fulfillmentOrderSubmitFulfillmentRequest mutation to allow merchants to request fulfillment

from the dropshipping app when an order is created.

12. Mobile app builders
A mobile app builder lets merchants create a mobile app based on their online store.

A. Mobile app builders
https://shopify.dev/docs/apps/launch/app-requirements-checklist

35/48


```

## Images

![Image from page 35](images/page_35_img_001.ppm)

![Image from page 35](images/page_35_img_002.ppm)

