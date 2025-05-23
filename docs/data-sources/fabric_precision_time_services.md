---
subcategory: "Fabric"
---

# equinix_fabric_precision_time_services (Data Source)

Fabric V4 API compatible data resource that allow user to fetch Equinix Fabric Precision Time Services with pagination details
Additional Documentation:
* API: https://docs.equinix.com/en-us/Content/KnowledgeCenter/Fabric/API-Reference/API-Precision-Time.htm

## Example Usage

```terraform
data "equinix_fabric_precision_time_services" "all" {
  pagination = {
    limit = 2
    offset = 1
  }
  filters = [{
    property = "/type"
    operator = "="
    values = ["PTP"]
  }]
  sort = [{
    direction = "DESC"
    property = "/uuid"
  }]
}


output "ept_service_id" {
  value = data.equinix_fabric_precision_time_services.all.data.0.id
}

output "ept_service_name" {
  value = data.equinix_fabric_precision_time_services.all.data.0.name
}

output "ept_service_state" {
  value = data.equinix_fabric_precision_time_services.all.data.0.state
}

output "ept_service_type" {
  value = data.equinix_fabric_precision_time_services.all.data.0.type
}

output "ept_service_ipv4" {
  value = data.equinix_fabric_precision_time_services.all.data.0.ipv4
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `filters` (Attributes List) List of filters to apply to the stream attachment get request. Maximum of 8. All will be AND'd together with 1 of the 8 being a possible OR group of 3 (see [below for nested schema](#nestedatt--filters))
- `pagination` (Attributes) Pagination details for the returned route aggregations list (see [below for nested schema](#nestedatt--pagination))
- `sort` (Attributes List) Filters for the Data Source Search Request (see [below for nested schema](#nestedatt--sort))

### Read-Only

- `data` (Attributes List) Returned list of route aggregation objects (see [below for nested schema](#nestedatt--data))
- `id` (String) The unique identifier of the resource

<a id="nestedatt--filters"></a>
### Nested Schema for `filters`

Required:

- `operator` (String) Operation applied to the values of the filter
- `property` (String) Property to apply the filter to
- `values` (List of String) List of values to apply the operation to for the specified property

Optional:

- `or` (Boolean) Boolean value to specify if this filter is a part of the OR group. Has a maximum of 3 and only counts for 1 of the 8 possible filters


<a id="nestedatt--pagination"></a>
### Nested Schema for `pagination`

Optional:

- `limit` (Number) Maximum number of search results returned per page. Number must be between 1 and 100, and the default is 20
- `offset` (Number) Index of the first item returned in the response. The default is 0


<a id="nestedatt--sort"></a>
### Nested Schema for `sort`

Optional:

- `direction` (String) The sorting direction. Can be one of: [DESC, ASC], Defaults to DESC
- `property` (String) The property name to use in sorting. One of [/name /uuid /state /type /package/code /changeLog/createdDateTime /changeLog/updatedDateTime] Defaults to /name


<a id="nestedatt--data"></a>
### Nested Schema for `data`

Optional:

- `ntp_advanced_configuration` (Attributes List) NTP Advanced configuration (see [below for nested schema](#nestedatt--data--ntp_advanced_configuration))
- `project` (Attributes) Equinix Project attribute object (see [below for nested schema](#nestedatt--data--project))
- `ptp_advanced_configuration` (Attributes) PTP Advanced Configuration (see [below for nested schema](#nestedatt--data--ptp_advanced_configuration))

Read-Only:

- `account` (Attributes) Equinix User Account associated with Precision Time Service (see [below for nested schema](#nestedatt--data--account))
- `change_log` (Attributes) Details of the last change on the route aggregation resource (see [below for nested schema](#nestedatt--data--change_log))
- `connections` (Attributes List) An array of objects with unique identifiers of connections. (see [below for nested schema](#nestedatt--data--connections))
- `href` (String) Equinix generated Portal link for the created Precision Time Service
- `ipv4` (Attributes) An object that has Network IP Configurations for Timing Master Servers. (see [below for nested schema](#nestedatt--data--ipv4))
- `name` (String) Name of Precision Time Service. Applicable values: Maximum: 24 characters; Allowed characters: alpha-numeric, hyphens ('-') and underscores ('_')
- `order` (Attributes) Precision Time Order (see [below for nested schema](#nestedatt--data--order))
- `package` (Attributes) Precision Time Service Package Details (see [below for nested schema](#nestedatt--data--package))
- `precision_time_price` (Attributes) Precision Time Service Price (see [below for nested schema](#nestedatt--data--precision_time_price))
- `state` (String) Indicator of the state of this Precision Time Service
- `type` (String) Choose type of Precision Time Service
- `uuid` (String) Equinix generated id for the Precision Time Service

<a id="nestedatt--data--ntp_advanced_configuration"></a>
### Nested Schema for `data.ntp_advanced_configuration`

Optional:

- `key` (String) The plaintext authentication key. For ASCII type, the key\
\ must contain printable ASCII characters, range 10-20 characters. For\
\ HEX type, range should be 10-40 characters
- `key_number` (Number) The authentication Key ID
- `type` (String) md5 Authentication type


<a id="nestedatt--data--project"></a>
### Nested Schema for `data.project`

Required:

- `project_id` (String) Equinix Subscriber-assigned project ID


<a id="nestedatt--data--ptp_advanced_configuration"></a>
### Nested Schema for `data.ptp_advanced_configuration`

Optional:

- `domain` (Number) The PTP domain value
- `grant_time` (Number) Unicast Grant Time in seconds. For Multicast and Hybrid transport modes, grant time defaults to 300 seconds. For Unicast mode, grant time can be between 30 to 7200
- `log_announce_interval` (Number) Logarithmic value that controls the rate of PTP Announce packets from the PTP time server. Default is 1 (1 packet every 2 seconds), Unit packets/second
- `log_delay_req_interval` (Number) Logarithmic value that controls the rate of PTP DelayReq packets. Default is -4 (16 packets per second), Unit packets/second..
- `log_sync_interval` (Number) Logarithmic value that controls the rate of PTP Sync packets. Default is -4 (16 packets per second), Unit packets/second..
- `priority1` (Number) The priority1 value determines the best primary clock, Lower value indicates higher priority
- `priority2` (Number) The priority2 value differentiates and prioritizes the primary clock to avoid confusion when priority1-value is the same for different primary clocks in a network
- `time_scale` (String) Time Scale value, ARB denotes Arbitrary and PTP denotes Precision Time Protocol
- `transport_mode` (String) ptp transport mode


<a id="nestedatt--data--account"></a>
### Nested Schema for `data.account`

Read-Only:

- `account_name` (String) Account Name
- `account_number` (Number) Equinix Account Number
- `global_cust_id` (String) Global Customer Id
- `global_org_id` (String) Customer organization naidentifierme
- `global_organization_name` (String) Global organization name
- `org_id` (Number) Customer organization identifier
- `organization_name` (String) Customer organization name
- `reseller_account_name` (String) Reseller account name
- `reseller_account_number` (Number) Reseller account number
- `reseller_org_id` (Number) Reseller customer organization identifier
- `reseller_ucm_id` (String) Reseller account ucmId
- `ucm_id` (String) Global organization name


<a id="nestedatt--data--change_log"></a>
### Nested Schema for `data.change_log`

Read-Only:

- `created_by` (String) User name of creator of the route aggregation resource
- `created_by_email` (String) Email of creator of the route aggregation resource
- `created_by_full_name` (String) Legal name of creator of the route aggregation resource
- `created_date_time` (String) Creation time of the route aggregation resource
- `deleted_by` (String) User name of deleter of the route aggregation resource
- `deleted_by_email` (String) Email of deleter of the route aggregation resource
- `deleted_by_full_name` (String) Legal name of deleter of the route aggregation resource
- `deleted_date_time` (String) Deletion time of the route aggregation resource
- `updated_by` (String) User name of last updater of the route aggregation resource
- `updated_by_email` (String) Email of last updater of the route aggregation resource
- `updated_by_full_name` (String) Legal name of last updater of the route aggregation resource
- `updated_date_time` (String) Last update time of the route aggregation resource


<a id="nestedatt--data--connections"></a>
### Nested Schema for `data.connections`

Required:

- `uuid` (String) Equinix Fabric Connection UUID; Precision Time Service will be connected with it

Read-Only:

- `href` (String) Link to the Equinix Fabric Connection associated with the Precision Time Service
- `type` (String) Type of the Equinix Fabric Connection associated with the Precision Time Service


<a id="nestedatt--data--ipv4"></a>
### Nested Schema for `data.ipv4`

Required:

- `default_gateway` (String) IPv4 address that establishes the Routing Interface where traffic is directed. It serves as the next hop in the Network.
- `network_mask` (String) IPv4 address that defines the range of consecutive subnets in the network.
- `primary` (String) IPv4 address for the Primary Timing Master Server.
- `secondary` (String) IPv4 address for the Secondary Timing Master Server.


<a id="nestedatt--data--order"></a>
### Nested Schema for `data.order`

Read-Only:

- `customer_reference_number` (String) Customer reference number
- `order_number` (String) Order reference number
- `purchase_order_number` (String) Purchase order number


<a id="nestedatt--data--package"></a>
### Nested Schema for `data.package`

Required:

- `code` (String) Time Precision Package Code for the desired billing package

Optional:

- `href` (String) Time Precision Package HREF link to corresponding resource in Equinix Portal


<a id="nestedatt--data--precision_time_price"></a>
### Nested Schema for `data.precision_time_price`

Read-Only:

- `charges` (Attributes List) offering price charge (see [below for nested schema](#nestedatt--data--precision_time_price--charges))
- `currency` (String) Offering price currency

<a id="nestedatt--data--precision_time_price--charges"></a>
### Nested Schema for `data.precision_time_price.charges`

Read-Only:

- `price` (Number) Offering price
- `type` (String) Price charge type; MONTHLY_RECURRING, NON_RECURRING
