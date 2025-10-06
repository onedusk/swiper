# Page 48

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

B. NFT distribution apps requirements

NFT distribution apps include the following types of apps:
NFT minting apps: Enable merchants to create and sell NFTs on Shopify.
NFT gifting apps: Enable merchants to distribute NFTs for free. For example, you might want to offer
free NFTs with a purchase, list an NFT as a product at no cost, or retroactively airdrop to customers.
1. Blockchain apps must identify all NFT variants by automatically populating product metafields.
2. For each fulfilled NFT, blockchain apps must write the blockchain transaction ID to the order's

Fulfillment tracking_number field, and a valid block scanner URL for the NFT fulfillment transaction to
the order's Fulfillment tracking_url field. Optionally, the name of the blockchain, fork, or network can
also be written to the order's Fulfillment tracking_company field, as necessary.

3. App partners must provide a way for customers to acquire a wallet, should they need one. Further,

customers must be able to receive full self-custody of their NFTs without any post-purchase fees,
unless the NFTs will be minted on a permissioned blockchain that prevents buyers from receiving full
self-custody of their NFTs. In such a case, the inability to receive full self-custody must be clearly
disclosed to customers before purchase and no post-purchase fees are permitted.

4. App partners must block stores from using any NFT distribution features while Shopify Payments is

active, including but not limited to minting, gifting, creating or listing NFT products, until the shop is
approved. To determine a shop's approval status, app partners must use the NFT Sales Eligibility API.
For more information, refer to NFT distribution.

C. Tokengating app requirements
Tokengating apps on Shopify enable merchants to gate access to products, promotions, and content
based on the contents of a customer’s Web3 wallet.
1. Any orders that contain line items which are either added or discounted as a result of a buyer

successfully passing a gate-check must be identified using order metafields.

2. Any products that contain one or more gated variants must be identified using product metafields.

Next steps
Prepare your app before submitting - Learn our recommended best practices for preparing and
testing your app before submitting it for review.

https://shopify.dev/docs/apps/launch/app-requirements-checklist

48/48


```

