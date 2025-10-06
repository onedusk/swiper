# Page 32

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

Specific requirements for certain app configurations
Apps are grouped into different categories depending on how they solve problems and meet merchant
needs. If your app is in one of the following categories, then it needs to meet the requirements listed
below. These requirements are in addition to the General requirements for all apps above.
In some cases, an app can have more than one type of configuration. For example, an app could be both
a third-party integration and a dropshipping app.

9. Online store
An online store app modifies a merchant's storefront and theme by using either Shopify's API or other
technical resources.

A. Online store
1. If your app modifies the merchant's theme, then you must implement an automated setup process

using the theme app extensions framework's deep linking using theme app extensions.

To support vintage themes, consider alternative integration methods, such as sharing instructions
with merchants that detail how to add and remove your app features in their theme.
2. If you want to forward requests made to a route on an online store's origin to an external origin to

display data on a store page, then you need to use app proxies.

3. Your app widget must be displayed properly and without any errors in the Theme Editor and Online

Store.

4. Your app doesn't use theme app extensions or blocks to promote your app, promote related apps, or

request reviews.

5. If your app adds a visible element to a merchant’s storefront, then you must allow the merchant to

preview edits before saving and publishing changes to your app’s visual storefront components.

6. If your app includes app blocks, then your app must allow merchants to add, reposition, or remove

app blocks in the theme editor.

7. App blocks must be responsive to the size of the section that they're added to.
8. If your app interacts with a merchant's theme, then you need to ensure that the app also works in the

theme editor environment. If necessary, you can set your app to detect the theme editor so that you
can adjust your app to work in that environment.

9. Your app must have detailed setup instructions on how to use your app embeds and app blocks.

https://shopify.dev/docs/apps/launch/app-requirements-checklist

32/48


```

