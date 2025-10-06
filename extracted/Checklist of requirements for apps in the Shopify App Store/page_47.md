# Page 47

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

6. Extensions must not add countdown timers to the checkout.
7. Extensions must not collect information, including personally identifiable information, that's already

captured by a standard Shopify checkout form field.

8. Extensions must use only the documented APIs that Shopify provides for customizing checkout.
9. Extensions must either request network access or use a metafield if they need to get data into

checkout that they can't currently get from Shopify.

10. Extensions using network access must not negatively affect the performance of checkout.
11. Extensions using network access must keep response time to under one second. If the extension

requires a response from a network call to render its components, then it must render skeleton
components initially, to avoid blocking checkout rendering.

12. Your app doesn't use extensions to promote your app, promote related apps, or request reviews.
13. Extensions must be feature-complete, and provide novel functionality or content.
14. Apps that implement Chat UI components on checkout pages must use them to provide customer

service using real-time chat as their core feature.

19. Blockchain apps
A blockchain app is defined as any application that exposes merchants to blockchain assets or
functionality, including but not limited to cryptocurrency, NFT distribution, and tokengating.

A. Blockchain app requirements
1. Apps must ensure that no personal data is written or stored on-chain.
2. Apps can't sell, transfer, or modify fungible tokens unless they are a payments partner that's been

approved by the Shopify Payments team.

3. Apps are presently only able to support the primary sales of NFTs on Shopify. All secondary sales

must be completed on a 3rd party platform, and must not be represented by products or hosted in the
Shopify Admin. A gallery display of NFTs on a Shopify store that links out to an external marketplace
is supported.

4. Apps should in no event facilitate the sale or marketing of NFTs that could be classified as one or

more of the following:

Securities or other regulated financial instruments
Activities related to securities or other regulated financial instruments
Having secondary-level or transferrable royalties.
Caution
Royalties should never be dispersed to buyers or recipients of NFTs.

https://shopify.dev/docs/apps/launch/app-requirements-checklist

47/48


```

## Images

![Image from page 47](images/page_47_img_001.ppm)

![Image from page 47](images/page_47_img_002.ppm)

