# Page 37

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

Payments and order management: Generate orders for merchants by taking customers to Shopify's
checkout with items pre-loaded in the cart using cart permalinks.

A. Onboarding and account connection
Merchant onboarding
1. Merchants must install sales channel apps using Shopify managed installation or during authorization

code grant, and sales channels must embed in the Shopify admin using Shopify App Bridge.

2. After the merchant installs the sales channel app via OAuth, they must be redirected to the account

section's account connection component. Connecting to the sales channel account must be done
in a modal window in the app's UI and occur outside of Shopify. This process returns the merchant to
the channel upon completion.

3. If the sales channel has any qualifying steps, eligibility requirements, or additional onboarding

requirements, then these must be included in the account connection form.

Account section
1. The sales channel must have an account section where the account connection
always visible (labelled with your channel name, such as "Sample channel").

component is

2. The account section for the sales channel must let merchants disconnect their account.
3. If there is an approval process for creating an account for the sales channel, then this must be

communicated to merchants using the banner component. The app must stay in the pending state
while the merchant awaits approval from the channel.

https://shopify.dev/docs/apps/launch/app-requirements-checklist

37/48


```

