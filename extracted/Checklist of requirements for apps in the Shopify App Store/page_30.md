# Page 30

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

3.
4. Your app must have a valid TLS/SSL certificate without any errors.
5. Your app must protect iFrames and prevent domains other than the shop domain from using the app

in an iFrame.

6. Your app must not expose network services unnecessarily.
7. Your app must subscribe to mandatory webhooks.
8. Your app must not expose its shared secret. If your secret is inadvertently exposed, then you must

rotate the secret immediately.

9. If your app uses offline tokens, then your app must not expose a shop's offline access token.
10. Your app must generate secure tokens, including expirations and search indexing protections, where

applicable.

11. Your app must not process payments or orders outside of Shopify's checkout.
12. Your app must not alter or modify Shopify's checkout, except through the APIs and components that

Shopify provides for that purpose.

13. Apps using the Admin APIs to capture payments must subscribe to the GraphQL ORDERS_EDITED

webhook topic, to be notified when an order is edited and a secondary payment needs to be
captured.

14. If your app uses app proxies, then it must verify the authenticity of requests.

7. Data and user privacy
Make sure that your app meets all requirements and best practices for querying, storing, processing, and
deleting Shopify data.

A. Data and user privacy
1. If your app gathers, stores, processes, or shares personal data, then it's your responsibility to make

sure that it complies with privacy laws.

2. You must include a link to a privacy policy in your app listing to communicate how your app uses

data, and to provide transparency and build trust with merchants.

3. If your app handles a significant amount of customer data, then it should have a system in place to

manage that data properly, including secure storage and the ability to erase data at the user's request
where applicable, as per the data rights of individuals.

4. If your app runs marketing or advertising campaigns that require personal data, then it must have a

system for allowing users to provide consent and/or opt-out for marketing promotions where
applicable.

5. All public apps must subscribe to mandatory webhooks so that you can receive any data deletion

requests that are issued by merchants.

6. Customer data collected by your app through a Shopify hosted service using the Online Store/Point

of Sale channels must be synced to the Shopify admin and be made accessible to merchants. More

https://shopify.dev/docs/apps/launch/app-requirements-checklist

30/48


```

