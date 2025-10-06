# Page 3

## Text Content

```
bin/dev

5. Visit the application
Open your browser and navigate to http://localhost:4000/

Project Structure
app/
├── controllers/

# Application controllers

│

├── admin/

# Admin interface controllers

│

├── api/

# API controllers

├── models/

# Data models

├── views/

# UI templates

│

# Admin interface views

├── admin/

├── mailers/

# Email templates

├── javascript/

# JS with Stimulus controllers

├── assets/

# Static assets and Tailwind config

db/
├── migrations/

# Database migrations

config/

# Application configuration

docs/
├── api/

# API documentation

test/

# Test suite

Key Components
Models
Event: Core entity representing events with details like name, date, location, and price
Customer: Represents ticket purchasers with contact information
Order: Records ticket purchases linking customers to events
Cart & CartItem: Manages the shopping cart system
PaymentMethod: Handles different payment options
QRCode: Generates QR codes for tickets
User: Admin users who can manage events (via Devise)


```

