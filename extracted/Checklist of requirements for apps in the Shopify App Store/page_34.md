# Page 34

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

4. If your app uses bulk action links, then they must be complete, functional, and relevant to their

locations in the Shopify admin. You must also make sure that for each bulk action link, the related
action is applied to all items that have been selected.

5. Apps must use the latest version of Shopify App Bridge by adding the app-bridge.js script tag
before any other script tags. We recommend to add it to the <head> of each document of your app

or as the first script element of the body.

6. If your app uses max modal, formerly known as full screen mode, then it must not launch without a

merchant interaction. Max modal can't be launched from the app navigation menu.

7. Max modal, formerly known as full screen mode, is intended to be used for complex editors or other

complex use cases. Max modal should be used to improve user experience when launched.

8. Your app must function in incognito mode in Chrome.
9. Use session tokens to authenticate requests between client and your app's backend.
10. Don't use 3rd party cookies or local storage, because your app might not work on certain browsers,

such as Safari for iOS, or browsers that block third party cookies.

B. Embedding into POS
1. If your app embeds into the Shopify POS app, then any POS actions that it uses (such as cart actions

or checkout actions) must be complete, fully functional, and relevant to the Shopify POS app's
capabilities.

2. If your app embeds into the Shopify POS app, then your app's user interface must be functioning and

accessible from the POS Apps Admin Dashboard.

https://shopify.dev/docs/apps/launch/app-requirements-checklist

34/48


```

## Images

![Image from page 34](images/page_34_img_001.ppm)

![Image from page 34](images/page_34_img_002.ppm)

