# Page 1

## Text Content

```
5/7/25, 7:26 AM

cloud_migration

Below is the proposed outline and mapping for migrating our workloads from Google Cloud Platform (GCP) to Azure. Let me know if you’d like to dive deeper into any of
these sections.

Organization
Cloud Architecture
DevOps
Security & Compliance
Roles & Responsibilities in Azure AD & RBAC
Approval workflows for resource provisioning

Project Scopes for Hardware Needed
On-Prem Connectors (if any)
VPN / ExpressRoute appliances
VM-based Lift-and-Shift
Azure VM sizes mapped to our Compute Engine instances
Container Hosts
AKS node pools sized to match GKE clusters
Networking
Firewalls, load balancers, and peering/Transit Gateway equivalents

Services & Resources
GCP Service

Azure Equivalent(s)

Notes

Cloud Storage

Blob Storage

Similar REST APIs; SDK changes

BigQuery

Synapse Analytics / Azure SQL Data Warehouse

SQL dialect differences; ETL pipeline adjustments

Cloud SQL (MySQL/Postgres)

Azure Database for MySQL/PostgreSQL

Use DMS or pg_dump for schema/data migration

Pub/Sub

Event Grid / Service Bus / Event Hubs

Select based on pub/sub pattern; minor semantic shifts

Cloud Functions

Azure Functions

Rewrite triggers (HTTP, Storage → HTTP, Queue)

App Engine

App Service / Container Apps

app.yaml → ARM/Bicep or Terraform

GKE / Cloud Run

AKS / Azure Container Instances

Helm charts reusable; update image registry

IAM & Roles

Azure RBAC & Azure AD

Parallel role model; revisit policies & blueprints

VPC / Networking

Virtual Network (VNet), NSGs, UDRs

CIDR rules & peering concepts align closely

Secret Manager

Key Vault & (Hashicorp Vault)

API differences; secret-rotation tooling

Total services mapped: 10

Mapping Details
Compute & Containers
GKE → AKS / ACI
Cloud Run → Azure Container Apps
Storage & Databases
Cloud Storage → Blob Storage
Cloud SQL → Azure Database for MySQL/PostgreSQL
Serverless
Functions (GCP) → Azure Functions
Pub/Sub → Event Grid / Service Bus
Networking & Security
VPC → Virtual Network (VNet)
IAM → Azure AD & RBAC
Secret Manager → Key Vault

file:///private/var/folders/bx/xp8_zypj3h1d023vwpc24jm80000gn/T/crossnote202547-21878-1us6kgx.yn4q.html

1/2


```

