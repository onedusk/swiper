# Page 5

## Text Content

```
GET /api/cart - View current cart
POST /api/cart_items - Add item to cart

More endpoints detailed in the API documentation

Email Integration
Think Events uses the Resend email service for sending transactional emails. The integration includes
comprehensive email templates for:
Welcome emails
Order confirmations
Event reminders
Event cancellation notices
Password reset emails
See RESEND_INTEGRATION.md for detailed configuration instructions.

Testing
The application includes a comprehensive test suite. Run the tests with:
bin/rails test

For system tests that require a browser:
bin/rails test:system

Deployment
Think Events is designed to be deployed to Fly.io, but can be deployed to any platform that supports
Rails applications.
See DEPLOYMENT.md for detailed deployment instructions.

Environment Variables for Production
Required environment variables:


```

