# Page 36

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

When reviewing your app, we test both the mobile app builder and the apps it makes to verify that all
requirements are met.
1. The app builder must be converted into a Sales Channel from the App Settings area of the Partner

Dashboard . This lets mobile apps that it builds create a checkout.

2. Your app must use Shopify App Bridge version 2.0. If your app is currently using the Embedded App

SDK (EASDK), then you need to migrate to use Shopify App Bridge.

3. The app builder must have either a customizable theme builder or include preset themes for

merchants to choose from.

4. The app builder must provide detailed instructions on how to create a developer account for either

the Apple App Store or the Google Play store.

5. The app builder must include information about the app marketplace submission process for either

the Apple App Store or the Google Play store to inform the merchant of wait times and app
requirements.

6. Apps made by the app builder must not make any requests to the authenticated GraphQL Admin API.

The app's client secret and API access token must be stored on a secure web server and not on the
mobile device.

7. Apps made by the app builder must not include the OAuth access token. All calls to the Shopify

Admin API must be made through a secure web server.

13. Sales channels
A sales channel app lets merchants publish their products from their Shopify admin to your platform,
whether they're selling online, on mobile apps, or through social media.
Caution
Ensure your app meets Shopify's definition of a Sales Channel. If it does, turn your app into a
Sales Channel and ensure compliance with all Sales Channel requirements before you submit
your app for review.

Overview
For a diagram that shows the lifecycle of a sales channel from the merchant's perspective, refer to
Building Shopify channels.
The key features of a sales channel app are as follows:
Building your sales channel: Your sales channel app must use Polaris components

and style guide.

Onboarding and account connection: Get permission from merchants to install your app, and then
connect them to your channel.
Product publishing: Import products into your channel, manage product errors, and stay in sync with
merchants' product catalogs.
https://shopify.dev/docs/apps/launch/app-requirements-checklist

36/48


```

