# Page 44

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

11. For a complete list of prohibited actions, refer to Prohibited actions

Requirements for payments extensions using checkout extensibility
1. Must not use banners, logos, or graphics in the checkout interface: Any checkout interface

customization needs to support payment completion. Banners, logos, or graphics can’t be used for
error states or as decorative elements.

2. Don’t use scrollable areas within the payment surface: Dropdown menus should display all options.

Don’t create any other scrollable areas once the payment option is selected.

3. Use a single extension with permitted targets: Your extension can only use the following targets:
Checkout::PaymentMethod::Render
Purchase.checkout.payment-option-item.details.render
Purchase.checkout.payment-option-item.hosted-fields.render-after

B. Requirements for testing
When you submit your payments app to the Shopify App Store for review, you need to fill out Part I. App
review instructions on the app listing with the following testing details:
1. A test store with the payments app installed
2. The required credentials to enable installing the payments app for testing (for example, activation

codes and login credentials)

3. Instructions on how to process a test payment and refund
4. A description of specific testing scenarios including installments / deferred payments and 3D Secure

authentication (if applicable)

C. Naming restrictions
To make choosing additional payment methods as straightforward as possible for merchants, you
should adhere to the following rules when naming your payments app:
1. The name of the payments app can't contain marketing text: For example, the name “World's Best

Provider: Get 50 payment methods” isn't allowed. This is because merchants won't see the name of
the payments app until they have chosen the payment method they wish to add to their store.

2. The name of the payment app can't be used by partners to gain a higher listing: There isn't a

general alphabetized directory of payments apps for merchants to navigate. Instead merchants will
discover payments apps using the payment methods they want to add.
You should make sure that the payment methods and locations offered are accurate because this is
the only information that's used to surface the app to merchants. If a name appears to have been
created with the purpose of gaining a higher listing on an alphabetized list, then it will not be allowed.

https://shopify.dev/docs/apps/launch/app-requirements-checklist

44/48


```

