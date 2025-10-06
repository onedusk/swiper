# Page 29

## Text Content

```
5/6/25, 4:40 PM

Checklist of requirements for apps in the Shopify App Store

In this section, you can provide instructions on how to test your app during your app review. You also
need to include your app's performance ratio. To calculate your app's performance ratio, refer to B.
Testing methodology in section 4. App performance of this document.
Including instructions on how your app should be used lets us give you valuable feedback if we encounter
issues while testing.
Login credentials must be provided if your app integrates with third-party platforms. For example, if your
app requires account access to a marketplace, then you must provide credentials for an active test
account for that specific marketplace. Failure to provide a test account will result in the rejection of your
app submission.
If your app requires login credentials, then the credentials you provide for review must be valid, and grant
full access to the app's complete feature set. Double-check any credentials before submission to avoid
issues during review.
Include a screencast when you submit your app for review. Follow these guidelines:
Create a complete demo: Include the setup process and all functionality and features as detailed in
the app’s description.
Add step-by-step instructions: Provide a walkthrough on how to configure and use the app. Demo
the expected outcome for each test case.
Show external setup: Add any additional steps for coding or external configurations.
Make it accessible: Make the screencast in English or include English subtitles.

6. Security and merchant risk
Security is a critical part of any web-based business because online apps can be exposed or
compromised in many different ways. Before you submit your app, you need to make sure that it's secure
so that the merchants who use it won't be at risk.

A. Security
Caution
By January 31, 2024, embedded apps need to load in the new admin.shopify.com domain.
Refer to our changelog for details. To resolve this issue, reference our Setting up Iframe
protection document and ensure your app is using App bridge 2.0 or later (App bridge 3.0 is
recommended).

1. Your app must not collect Shopify user credentials. As explained in Shopify API Authentication, public

apps must use OAuth and public embedded apps must use session tokens and OAuth.

2. If your app stores its own credentials, then it must only store salted password hashes instead of

actual passwords, as described on the Open Web Application Security Project

website.

Your app must be protected against common web security vulnerabilities.
https://shopify.dev/docs/apps/launch/app-requirements-checklist

29/48


```

