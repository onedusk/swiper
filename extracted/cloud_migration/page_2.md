# Page 2

## Text Content

```
5/7/25, 7:26 AM

cloud_migration

Infrastructure as Code & Automation
Terraform
# before
"google" ">provider "google" { … }
# after
"azurerm" ">provider "azurerm" { … }
resource "google_storage_bucket" → resource "azurerm_storage_account"

ARM / Bicep
Rewrite any hand-crafted Deployment Manager templates into ARM or Bicep modules
CI/CD Pipelines
Update build/release definitions for Azure DevOps or GitHub Actions
Validate end-to-end deployments in a sandbox subscription

Considerations & Challenges
API Differences
SDK and REST API semantics vary
Deployment Pipeline Compatibility
Triggers, secrets, service connections
Learning Curve
Azure CLI / Portal / SDKs; cost-management tools
Cost & Licensing
Reserved instances vs. committed-use discounts
Monitoring & Logging
Migrate Stackdriver → Azure Monitor, Log Analytics, Application Insights
Governance & Compliance
Rewrite GCP Org Policies → Azure Policies & Blueprints

file:///private/var/folders/bx/xp8_zypj3h1d023vwpc24jm80000gn/T/crossnote202547-21878-1us6kgx.yn4q.html

2/2


```

