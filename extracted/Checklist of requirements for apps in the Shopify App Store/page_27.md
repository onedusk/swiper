# Page 27

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

You can specify which merchants can install your app by setting the install requirements in the app
submission form. By adding install requirements in the app submission form, you can reduce the number
of uninstalls and negative reviews related to merchant eligibility for your app.
For example, when a merchant installs an app that they can't use, such as a free shipping app that
doesn't work in their country, they will uninstall your app shortly after installing it. They may also be
frustrated about the experience and leave a negative review. Both uninstalls and negative reviews affect
your ranking in the Shopify App Store.
1. Sales channel requirements
If your app requires the Shopify Online Store or Shopify POS sales channels in order to work, then you
only want merchants who use either, or both, to install your app. For example, if a merchant doesn't have
an online store, then you want to prevent them from installing your app.
If your app interacts with the merchant's online store, such as using theme app extensions or editing
theme assets, then select Shopify Online Store. If your app embeds features in Shopify Point of Sale, then
select Shopify POS.
Shopify is responsible for the final determination about whether your app has specific sales channel
requirements, and may update this setting before approving your app.
2. Geography requirements
Set the geography requirements to make your app available only to merchants who meet specific
geographic criteria. For example, if your app is a tax app that helps merchants in Germany file their taxes,
then you should specify that only merchants with a business address in Germany can install your app.
You can restrict the installation of your app to merchants who:
have a business address in a specific country or countries
ship to a specific country or countries
accept a specific currency or currencies.
For each requirement, you can specify a list of countries or currencies that meet the requirement. For
example, if your app works for stores who accept any of USD, CAD or GBP, then you can specify all three
acceptable currencies.
Note
If you specify multiple geographic requirements, then only merchants who meet all of the
requirements can install the app.

What if a merchant changes their store settings after installation?
Within your app, use endpoints and webhooks to check if a merchant changes their store settings after
installation. If a merchant does change their settings, then you can notify them within the app or by email.

https://shopify.dev/docs/apps/launch/app-requirements-checklist

27/48


```

