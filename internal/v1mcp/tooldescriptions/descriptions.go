package tooldescriptions

var NextLink = "The link used to retrieve the next page of results"

var DefaultTop = "The number of records to display per page. One of 10, 50, 100, 200, 500, 1000"

var CAMListGCPProjectsFilterDescription = `
string <= 254 characters
Examples:

    state eq 'managed' or state eq 'outdated' - List Google Cloud projects with statuses of 'managed' or 'outdated'.
    state eq 'managed' and (contains(name, 'lab') or contains(id, '123')) - List managed Google Cloud projects with names containing 'lab' or IDs containing '123'.

The filter for retrieving a list of a subset of connected Google Cloud projects.

Supported fields:
Field 	Description 	Possible values
id 	The Google Cloud project number used as the ID for managing the connected Google Cloud project in Cloud Accounts. 	Any value
name 	The name of the Google Cloud project. 	Any value
state 	The status of the Google Cloud project. 	managed, outdated, failed
workloadIdentityPoolId 	The workload identity pool ID of the Google Cloud project. 	Any value
oidcProviderId 	The OIDC provider ID of the Google Cloud project. 	Any value
serviceAccountId 	The service account ID of the Google Cloud project. 	Any value
featureId 	The features enabled for the Google Cloud project. 	cloud-sentry
gcpRegion 	The region where Cloud Accounts is deployed. 	Any supported region value
Supported operators: 		
Operator 	Description 	
--------- 	--------- 	
eq 	Operator 'equal to' 	
and 	Operator 'and' 	
or 	Operator 'or' 	
not 	Operator 'not' 	
( ) 	Symbols for grouping operands with their correct operator. 	
contains 	Operator that allows you to search for a specified string in a field 	

Note: Include this parameter in every request that generates paginated output.
`

var FilterAWSAccounts = `
string <= 254 characters
Examples:

    state eq 'managed' or state eq 'outdated' - List AWS accounts with states of 'managed' or 'outdated'.
    state eq 'managed' and (contains(name, 'lab') or contains(id, '123')) - List managed AWS accounts with names containing 'lab' or IDs containing '123'.

The filter for retrieving a list of a subset of connected AWS accounts.

Supported fields:
Field 	Description 	Supported values
name 	The name of the AWS account. 	Any value
state 	The state of the AWS account. 	managed, outdated, failed
featureId 	The features enabled for the AWS account. 	container-security, cloud-response
id 	The ID of the AWS account. 	Any value

Supported operators:
Operator 	Description
eq 	Operator 'equal to'
and 	Operator 'and'
or 	Operator 'or'
not 	Operator 'not'
( ) 	Symbols for grouping operands with their correct operator.
contains 	Operator that allows you to search for a specified string in a field

Note: Include this parameter in every request that generates paginated output.
`
var FilterAlibabaAccounts = `
string <= 254 characters
Examples:

    state eq 'managed' or state eq 'outdated' - List Alibaba Cloud accounts with states of 'managed' or 'outdated'.
    state eq 'managed' and (contains(name, 'lab') or contains(id, '123')) - List managed Alibaba Cloud accounts with names containing 'lab' or IDs containing '123'.

The filter for retrieving a list of a subset of connected Alibaba Cloud accounts.

Supported fields:
Field 	Description 	Supported values
id 	The ID of the Alibaba Cloud account. 	Any value
name 	The name of the Alibaba Cloud account. 	Any value
state 	The state of the Alibaba Cloud account. 	managed, outdated, failed
Supported operators: 		
Operator 	Description 	
--------- 	--------- 	
eq 	Operator 'equal to' 	
and 	Operator 'and' 	
or 	Operator 'or' 	
not 	Operator 'not' 	
( ) 	Symbols for grouping operands with their correct operator. 	
contains 	Operator that allows you to search for a specified string in a field 	

Note: Include this parameter in every request that generates paginated output.
`

var FilterAttackSurfaceDevices = `
string <= 1024 characters
Examples:

    (latestRiskScore ge 70) and (installedAgents eq 'Trend Vision One Agent') -
    hassubset(installedAgents, ['Trend Micro Deep Security']) -
    startswith(lastUser,'john') -
    hassubset(discoveredBy, ['Trend Micro Deep Security', 'Trend Vision One Agent']) - 

Filter for retrieving a subset of the device information list.

Supported fields:
Field 	Description 	Supported values
deviceName 	Device name 	Any value
id 	The ID of the device on the Trend Vision One platform. 	Any value
ip 	The IP addresses of the device 	Any value
deviceType 	Whether a device can be assessed. You can only include one device type per query 	Can be assessed, Cannot be assessed, With managed agents, Unmanaged device
latestRiskScore 	The most recent Risk Score of the device 	Any value
criticality 	The criticality of the device 	high, medium, low
osPlatform 	Operating system of the device 	Android, Linux, macOS, Windows, Other
lastUser 	The last user who signed in to the device 	Any value
installedAgents 	The agents installed on the device 	The values in the "installedAgents" field when the request is successful.
discoveredBy 	The data sources that discovered the device 	The values in the "discoveredBy" field when the request is successful.
assetCustomTagIds 	The tag ID of each asset in assetCustomTags 	Any value in field.

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	Not applicable to discoveredBy
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands 	-
gt 	Operator 'greater than' 	Only applicable to 'latestRiskScore'
ge 	Operator 'greater than or equal' 	Only applicable to 'latestRiskScore'
le 	Operator 'less than or equal' 	Only applicable to 'latestRiskScore'
lt 	Operator 'less than' 	Only applicable to 'latestRiskScore'

Additional functions:
Function 	Description 	Notes
startswith() 	Determines if the specified string begins with the specified characters 	Only applicable to deviceName and lastUser
hassubset() 	Checks if the array contains a subset 	Applicable to discoveredBy, installedAgents, assetCustomTagIds and ip only
`

var FilterDomainAccounts = `
string <= 1024 characters
Examples:

    (latestRiskScore ge 70) and (userType eq 'member') -
    hassubset(discoveredBy, ['Trend Micro Deep Security', 'Trend Vision One Agent']) - 

The filter for retrieving a subset of the domain accounts list.

Supported fields:
Field 	Description 	Supported values
name 	The name of the domain account 	Any value
id 	The ID of the asset on the Trend Vision One platform. 	Any value
type 	Whether a account can be assessed. You can only include one account type per query 	Any value
latestRiskScore 	The most recent Risk Score of the account 	Any value
criticality 	The criticality of the account 	high, medium, low
location 	The location of the account 	Any value
jobTitle 	The job title of the user 	Any value
discoveredBy 	The data sources that discovered the account 	The values in the discoveredBy field when the request is successful.

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	Not applicable to discoveredBy
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands 	-
gt 	Operator 'greater than' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
ge 	Operator 'greater than or equal' 	Only applicable to: 'latestRiskScore''firstSeenDateTime','lastDetectedDateTime'
le 	Operator 'less than or equal' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
lt 	Operator 'less than' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'

Additional functions:
Function 	Description 	Notes
startswith() 	Determines if the specified string begins with the specified characters 	Only applicable to name
hassubset() 	Checks if the array contains a subset 	Only applicable to discoveredBy
`

var FilterFQDNS = `
string <= 1024 characters
Examples:

    (latestRiskScore ge 70) and (provider eq 'Trend Vision One Agent') -
    hassubset(discoveredBy, ['Trend Micro Deep Security', 'Trend Vision One Agent']) - 

Filter for retrieving a subset of the internet-facing domains list.

Supported fields:
Field 	Description 	Supported values
rootDomain 	The root domain 	Any value
id 	The ID of the domain on the Trend Vision One platform. 	Any value
provider 	The domain provider. You can only include one provider per query 	Any value
latestRiskScore 	The most recent Risk Score of the domain 	Any value
criticality 	The criticality of the domain 	high, medium, low
discoveredBy 	The data sources that discovered the domain 	The values in the "discoveredBy" field when the request is successful.

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	Not applicable to discoveredBy
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands 	-
gt 	Operator 'greater than' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
ge 	Operator 'greater than or equal' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
le 	Operator 'less than or equal' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
lt 	Operator 'less than' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'

Additional functions:
Function 	Description 	Notes
startswith() 	Determines if the specified string begins with the specified characters 	Only applicable to rootDomain
hassubset() 	Checks if the array contains a subset 	Applicable to discoveredBy and ipAddresses only
`

var FilterIps = `
string <= 1024 characters
Examples:

    (latestRiskScore ge 70) and (provider eq 'Amazon') -
    hassubset(provider, ['Amazon']) -
    hassubset(discoveredBy, ['Trend Micro Deep Security', 'Trend Vision One Agent']) - 

Filter for retrieving a subset of the public IP addresses list.

Supported fields:
Field 	Description 	Supported values
ipAddress 	The public IP address 	Any value
id 	The ID of the IP address on the Trend Vision One platform. 	Any value
provider 	The provider of the asset. You can only include one provider per query 	Any value
latestRiskScore 	The most recent Risk Score of the IP address 	Any value
criticality 	The criticality of the IP address 	high, medium, low
discoveredBy 	The data sources that discovered the IP address 	The values in the "discoveredBy" field when the request is successful.

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	Not applicable to discoveredBy
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands 	-
gt 	Operator 'greater than' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
ge 	Operator 'greater than or equal' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
le 	Operator 'less than or equal' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
lt 	Operator 'less than' 	Only applicable to: 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'

Additional functions:
Function 	Description 	Notes
startswith() 	Determines if the specified string begins with the specified characters 	Only applicable to ipAddress
hassubset() 	Checks if the array contains a subset 	Applicable to discoveredByonly
`

var FilterCloudAssets = `
string <= 1024 characters
Example: assetType eq 'EKS Cluster'

The filter for retrieving a subset of the cloud asset information list.

Supported fields:
Field 	Description 	Supported values
id 	The ID of the cloud asset on the Trend Vision One platform 	Any value
latestRiskScore 	The most recent Risk Score of the cloud asset 	Any value
assetName 	The name of the cloud asset 	Any value
assetType 	The type of the cloud asset 	Any value
assetCategory 	The category of the cloud asset 	Any value
criticality 	The criticality of the cloud asset 	high, medium, low
provider 	The provider of the cloud asset 	Any value
service 	The cloud service related to the cloud asset 	Any value
location 	The geographical location of the cloud asset 	Any value
region 	The cloud region where the asset is located 	Any value
cloudAccountName 	The name of the cloud account associated with the asset 	Any value
protectionStatus 	Indicates if the cloud asset is protected by Container Security 	enabled, not enabled, unknown
assetCustomTagIds 	The tag ID of each asset in assetCustomTags 	Any value in field

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	-
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands with their correct operator 	-
gt 	Operator 'greater than' 	Only applicable to latestRiskScore
ge 	Operator 'greater than or equal' 	Only applicable to latestRiskScore
le 	Operator 'less than or equal' 	Only applicable to latestRiskScore
lt 	Operator 'less than' 	Only applicable to latestRiskScore

Additional functions:
Function 	Description 	Notes
hassubset() 	Checks if the array contains a subset 	Applicable to assetCustomTagIds only
`

var FilterHighRiskUsers = `
string <= 1024 characters
Example: (userPrincipalName eq 'demo_account@visionone.trendmicro.com') or (userName eq 'demo_account')

Filter for retrieving a subset of the at-risk users list.

Supported fields and operators:

    'id' - The ID of a user on the Trend Vision One platform.
    'riskScore' - Risk score of a user.
    'userPrincipalName' - String that identifies an account.
    'userName' - User name.

Supported fields:
Field 	Description
id 	The ID of a user on the Trend Vision One platform.
riskScore 	The risk score of a user.
userPrincipalName 	String that identifies an account.
userName 	User name

Supported operators:
Operator 	Description
eq 	Operator 'equal to'.
and 	Operator 'and'.
or 	Operator 'or'.
not 	Operator 'not'.
( ) 	Symbols for grouping operands.
gt 	Operator 'greater than'.
ge 	Operator 'greater than or equal'.
le 	Operator 'less than or equal'.
lt 	Operator 'less than'.
`

var FilterApiKeys = `
string <= 1024 characters
Example: role eq 'Master Administrator'

Filter for retrieving a subset of the API keys list.

Supported fields:
Field 	Description
id 	The unique identifier of the API key
name 	The unique name of an API key
role 	The user role assigned to the API key
status 	The status of an API key

Supported operators:
Operator 	Description
eq 	Operator 'equal to'
and 	Operator 'and'
or 	Operator 'or'
not 	Operator 'not'
() 	Symbols for grouping operands with their correct operator.

Note: Include this parameter in every request that generates paginated output.
`

var FilterCloudPostureChecks = `
string <= 1783 characters
Examples:

    accountId eq '3c8f0d33-65f0-4802-97f3-4475bb70e43e' or accountId eq 'be08d97c-55c4-4709-976c-24955ff59c8d' - List the checks with accountId is '3c8f0d33-65f0-4802-97f3-4475bb70e43e' or 'be08d97c-55c4-4709-976c-24955ff59c8d'.
    accountId eq '3c8f0d33-65f0-4802-97f3-4475bb70e43e' and ruleId eq 'EC2-001' - List the checks with accountId is '3c8f0d33-65f0-4802-97f3-4475bb70e43e' and ruleId is 'EC2-001'

The filter for retrieving a subset of the Cloud Risk Management checks.

Supported fields:
Field	Description	Supported values
accountId	The Cloud Risk Management IDs	

    Single account ID:accountId eq '3c8f0d33-65f0-4802-97f3-4475bb70e43e'

    multiple account IDs:(accountId eq '9d4kd94l-d839-3932-88d0-x9839d85jfp0' or accountId eq '38d92k21-3821-d829-s823-38d920s84kfi')

region	The region of the account	For more information about regions, see
service	The cloud service of the check to filter on.	For more information about regions, see
categories	A list of categories of the check.	

    security
    cost-optimisation
    reliability
    performance-efficiency
    operational-excellence
    sustainability


Example:['sustainability', 'performance-efficiency']
riskLevel	The risk level of the check	

    LOW
    MEDIUM
    HIGH
    VERY_HIGH
    EXTREME

status	The status of the check	SUCCESS,FAILURE
ruleId	The rule IDs of checks to be returned	Any value
resource	The resource ID	Any value
description	The check description	Any value
suppressed	

Whether the check is suppressed

Default: All checks
	true,false
tags	The tags associated with a cloud resource	Any value
compliances	A list of supported standard or framework IDs	

    AWAF
    AZUREWAF-2024
    GCPWAF
    CISAWSF-1_5_0
    CISAWSF-2_0
    CISAWSF-3_0
    CISAWSF-4_0_1
    CISAZUREF-2_0
    CISAZUREF-2_1
    CISGCPF-1_3_0
    CISGCPF-2_0
    CISGCPF-3_0
    CISABCF-1_0
    CIS-V8
    NIST4
    NIST5
    SOC2
    NIST-CSF
    NIST-CSF-2_0
    ISO27001
    ISO27001-2022
    AGISM
    AGISM-2024
    HIPAA
    HITRUST
    ASAE-3150
    PCI
    PCI-V4
    APRA
    FEDRAMP
    MAS
    GDPR
    ENISA
    NIS-2
    FISC-V9
    LGPD

Example:['AWAF', 'PCI']
Operator 	Description 	Supported fields 	Example
eq 	Operator 'equal to' 	All 	service eq 'EC2'
and 	Operator 'and' 	All 	service eq 'EC2' and riskLevels eq 'HIGH'
or 	Operator 'or' that can only be used in ( ) 	All 	(service eq 'EC2' or service eq 'S3') and riskLevels eq 'HIGH'
not 	Operator 'not' 	All 	service eq 'EC2' and not riskLevels eq 'HIGH'
( ) 	Operator to group filters for precedence. 	All 	(service eq 'EC2' or service eq 'S3')
contains 	Operator to check if a string contains the specified string. Supports partial matching. 	service, description, ruleId, resource 	contains(description, 'Role IAM') and contains(ruleId, 'S3')
hassubset 	Operator to check if an array contains the specified elements. Supports partial matching. 	categories, compliances 	hassubset(compliances, ['GDPR', 'PCI'])
any 	Operator to filter tags. You can use the eq and contains operators to filter tags. 	tags 	tags/any(tag: tag eq 'Environment::development'), tags/any(tag: contains(tag, 'Service'))

    Important

    The and operator is not supported inside parentheses.
`

var FilterWorkbenchAlerts = `
string <= 5000 characters
Examples:

    investigationStatus eq 'New' and contains(impactScopeEntityValue,'nimda') - Filters the list by alert status (exact match) and impacted entity (partial match).
    impactScopeEntityValue eq 'nimda' - Filters the list by impacted entity (exact match)
    indicatorValue eq '8.8.8.8' - Filters the list by detected indicator (exact match)

Filter for retrieving a subset of the alert list.

Supported fields:
Field 	Description 	Possible values
id 	The unique identifier of an alert 	Any value
investigationStatus (Deprecated) 	The current status of the Workbench alert or investigation 	New, In Progress, True Positive, False Positive, Benign True Positive, Closed
status 	The status of the case or investigation 	Open, In Progress, Closed
investigationResult 	The findings of the case or investigation 	No Findings, Noteworthy, True Positive, False Positive, Benign True Positive, Other Findings
alertProvider 	Source of a Workbench alert 	SAE, TI
modelId 	ID of the detection model that triggered the alert 	Any value
model 	The detection model that triggered the alert 	Any value
modelType 	The type of detection model that triggered the alert 	preset, custom
severity 	The severity assigned to a model that triggered the alert 	critical, high, medium, low
impactScopeEntityValue 	Entities affected within the company network 	Any value
indicatorValue 	Objects found using root cause analysis or sweeping 	Any value
incidentId 	The unique identifier of an incident 	Any value

Supported operators:
Operator 	Description
eq 	Operator 'equal to'
and 	Operator 'and'
or 	Operator 'or'
not 	Operator 'not'
( ) 	Symbols for grouping operands with their correct operator.
contains 	Operator that allows you to search for a specified string in a field

Note: Include this parameter in every request that generates paginated output.
`

var WorkbenchOrderBy = `
string <= 200 characters
Default: "createdDateTime desc"
Example: orderBy=score desc,severity desc

Specifies the field by which the results are sorted.

Records are returned in descending order by default. To return records in ascending order, add the phrase asc after the parameter name.

You can specify multiple fields separated by commas.

Available values:

    id
    caseId
    name
    investigationStatus (deprecated)
    status
    investigationResult
    modelId
    model
    score
    severity
    createdDateTime
    updatedDateTime
    firstInvestigatedDateTime
`

var ObservedAttackFilter = `
string <= 4000 characters
Example: (riskLevel eq 'high') and (endpointName eq 'my-computer')

Filter for retrieving a subset of the collected Observed Attack Techniques events. Include this parameter in every request that generates paginated output.

Important: The name of the containerName field might change depending on the products you purchase and the supported products in your region.

Supported fields:
Field 	Description
uuid 	The ID of an Observed Attack Techniques event.
riskLevel 	The severity of a detection. Possible values: undefined, info, low, medium, high, critical
filterName 	The detection filter that triggered the event
filterMitreTacticId 	The ID of the MITRE ATT&CK tactic associated with an event.
filterMitreTechniqueId 	The ID of the MITRE ATT&CK technique or sub-technique associated with an event.
endpointName 	The name of an endpoint
agentGuid 	The ID of the installed agent
endpointIp 	The IP address of the endpoint
productCode 	Product that generated the alert
containerName 	The name of the container.

Supported operators:
Operator 	Description
eq 	Operator 'equal to'
and 	Operator 'and'
or 	Operator 'or'
not 	Operator 'not'
() 	Symbols for grouping operands with their correct operator
`

var FilterUserAccounts = `
string <= 256 characters
Example: status eq 'enabled' and authType eq 'local'

Filter for retrieving a subset of the retrieved user list.

Supported fields:
Field 	Description
id 	The unique identifier of a user.
email 	The email address of a user
authType 	The type of the user account. Available values: local, saml, samlGroup
status 	The status of an account. Available values: enabled, disabled, invited

Supported operators:
Operator 	Description
eq 	Operator 'equal to'
and 	Operator 'and'
or 	Operator 'or'
not 	Operator 'not'
`

var FilterServiceAccounts = `
string <= 1024 characters
Examples:

    (latestRiskScore ge 70) and (type eq 'application') -
    hassubset(discoveredBy, ['Trend Micro Deep Security', 'Trend Vision One Agent']) - 

Filter for retrieving a subset of the service accounts list.

Supported fields:
Field 	Description 	Supported values
name 	The name of the service Account 	Any value
id 	The ID of the asset on the Trend Vision One platform. 	Any value
type 	Whether a service account can be assessed. You can only include one service account type per query 	Any value
latestRiskScore 	The most recent Risk Score of the service account 	Any value
criticality 	The criticality of the account 	high, medium, low
source 	The source of the service account 	Any value
status 	The status of the service account 	Disabled, Enabled
discoveredBy 	The data sources that discovered the account 	The values in the "discoveredBy" field when the request is successful.

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	Not applicable to discoveredBy
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands 	-
gt 	Operator 'greater than' 	Only applicable to 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
ge 	Operator 'greater than or equal' 	Only applicable to 'latestRiskScore''firstSeenDateTime','lastDetectedDateTime'
le 	Operator 'less than or equal' 	Only applicable to 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
lt 	Operator 'less than' 	Only applicable to 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'

Additional functions:
Function 	Description 	Notes
startswith() 	Determines if the specified string begins with the specified characters 	Only applicable to name
hassubset() 	Checks if the array contains a subset 	Applicable to discoveredBy only
`

var FilterCloudAssetRiskIndicators = `
string <= 1024 characters
Example: riskLevel eq 'high'

The filter for retrieving a subset of the cloud asset information list.

Supported fields:
Field 	Description 	Supported values
id 	The ID of a risk event 	Any value
riskLevel 	The risk level of the risk event 	high, medium, low
riskFactor 	The risk factor of the risk event 	Threat detection, Security configuration, System configuration, Vulnerability detection, Anomaly detection, Account compromise, Cloud app activity, XDR detection
status 	The status of the risk event 	new, inProgress, remediated, dismissed, accepted, mitigated
detectedDateTime 	The time the event was detected 	Any value

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	-
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands with their correct operator 	-
gt 	Operator 'greater than' 	Only applicable to detectedDateTime
ge 	Operator 'greater than or equal' 	Only applicable to detectedDateTime
le 	Operator 'less than or equal' 	Only applicable to detectedDateTime
lt 	Operator 'less than' 	Only applicable to detectedDateTime
`

var FilterLocalApps = `
string <= 1024 characters
Example: operatingSystem eq 'Linux'

The filter for retrieving a subset of the local application information list.

Supported fields:
Field 	Description 	Supported values
id 	The ID of the local application on the Trend Vision One platform 	Any value
name 	The name of the local application 	Any value
osPlatform 	The operating system of the local application 	Windows, Linux, iOS, Android, macOS, Other
latestRiskScore 	The most recent Risk Score of the local application 	Any value
vendor 	The vendor of the local application 	Any value
permissionStatus 	The permission status of the local application 	allowed, blocked
firstSeenDateTime 	The first time Attack Surface Discovery detected the local application 	Any value
lastDetectedDateTime 	The last time a highly-exploitable CVE was detected on the local application 	Any value

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	-
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands with their correct operator 	-
gt 	Operator 'greater than' 	Only applicable to 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
ge 	Operator 'greater than or equal' 	Only applicable to 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
le 	Operator 'less than or equal' 	Only applicable to 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'
lt 	Operator 'less than' 	Only applicable to 'latestRiskScore','firstSeenDateTime','lastDetectedDateTime'

Additional functions:
Function 	Description 	Notes
contains() 	Checks if the string contains the specified value 	Applicable to vendor only
`

var FilterLocalAppRiskIndicators = `
string <= 1024 characters
Example: riskLevel eq 'high'

The filter for retrieving a subset of the local application information list.

Supported fields:
Field 	Description 	Supported values
id 	The ID of a risk event 	Any value
riskLevel 	The risk level of the risk event 	high, medium, low
riskFactor 	The risk factor of the risk event 	Threat detection, Security configuration, System configuration, Vulnerability detection, Anomaly detection, Account compromise, Cloud app activity, XDR detection
status 	The status of the risk event 	new, inProgress, remediated, dismissed, accepted, mitigated
detectedDateTime 	The time the event was detected 	Any value

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	-
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands with their correct operator 	-
gt 	Operator 'greater than' 	Only applicable to detectedDateTime
ge 	Operator 'greater than or equal' 	Only applicable to detectedDateTime
le 	Operator 'less than or equal' 	Only applicable to detectedDateTime
lt 	Operator 'less than' 	Only applicable to detectedDateTime
`

var FilterLocalAppDevices = `
string <= 1024 characters
Example: latestRiskScore eq 'high'

The filter for retrieving a subset of the device information list.

Supported fields:
Field 	Description 	Supported values
id 	The ID of the device 	Any value
name 	The name of the device 	Any value
latestRiskScore 	The most recent Risk Score of the device 	Any value

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	-
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands with their correct operator 	-
gt 	Operator 'greater than' 	Only applicable to latestRiskScore
ge 	Operator 'greater than or equal' 	Only applicable to latestRiskScore
le 	Operator 'less than or equal' 	Only applicable to latestRiskScore
lt 	Operator 'less than' 	Only applicable to latestRiskScore

Additional functions:
Function 	Description 	Notes
contains() 	Checks if the string contains the specified value 	Applicable to name only
`

var FilterLocalAppExecutables = `
string <= 1024 characters
Example: name eq 'abc.exe'

The filter for retrieving a subset of the executable file information list.

Supported fields:
Field 	Description 	Supported values
name 	The name of the executable file 	Any value
productName 	The name of the product associated with the executable file 	Any value
language 	The language of the executable file 	Any value
firstSeenDateTime 	The first time the executable file was detected, ISO 8601 format 	Any value
lastDetectedDateTime 	The last time the executable file was detected, ISO 8601 format 	Any value

Supported operators:
Operator 	Description 	Notes
eq 	Operator 'equal to' 	-
and 	Operator 'and' 	-
or 	Operator 'or' 	-
not 	Operator 'not' 	-
() 	Symbols for grouping operands with their correct operator 	-
gt 	Operator 'greater than' 	Only applicable to firstSeenDateTime, lastDetectedDateTime
ge 	Operator 'greater than or equal' 	Only applicable to firstSeenDateTime, lastDetectedDateTime
le 	Operator 'less than or equal' 	Only applicable to firstSeenDateTime, lastDetectedDateTime
lt 	Operator 'less than' 	Only applicable to firstSeenDateTime, lastDetectedDateTime

Additional functions:
Function 	Description 	Notes
contains() 	Checks if the string contains the specified value 	Applicable to productName only
`

var FilterCustomTags = `
string <= 1024 characters
Example: id eq 'qVQz+Y3HL1GQ56qTeSKhtFxYAIM=-01'

The filter for retrieving a subset of the cloud asset information list.

Supported fields:
Field 	Description 	Supported values
id 	The tag ID of each asset in assetCustomTags 	Any value in field
key 	The key of each asset in assetCustomTags 	Any value in field
value 	The tag value of each asset in assetCustomTags 	Any value in field
Supported operators: 		
Operator 	Description
eq 	Operator 'equal to'
and 	Operator 'and'
or 	Operator 'or'
not 	Operator 'not'
() 	Symbols for grouping operands with their correct operator
`

var FilterEmailAccounts = `
string <= 512 characters
Example: sensorDetectionStatus eq 'Enabled' and mailService eq 'Exchange Online'

The filter used to retrieve a subset of email accounts from a generated paginated list.

Supported fields:
Field 	Description 	Supported values
sensorDetectionStatus 	The account's email sensor detection status. 	Enabled, Disabled
protectionPolicyStatus 	The account's Cloud Email and Collaboration Protection policy status. 	'Disabled', 'Fully enabled', 'Partially enabled'
mailService 	The account's mail service (iam) type. 	'Exchange Online', 'Gmail', 'Unknown'

Supported operators:
Operator 	Description 	Notes
eq 	Operator "equal to" 	-
and 	Operator "and" 	-

Only support eq, and operators.
`

var FilterWorkbenchNotes = `
string <= 5000 characters
Example: creatorName eq 'John Doe'

Filter for retrieving a subset of Workbench alert notes. Supported fields and operators:

    id - Numeric string that identifies a Workbench alert note
    creatorMailAddress - Email address of the user that created a Workbench alert note
    creatorName - User that created a Workbench alert note
    lastUpdatedBy - Parameter that indicates the user who last modified a Workbench alert note
    'eq' - Abbreviation of the operator 'equal to'
    'and' - Operator 'and'
    'or' - Operator 'or'
    'not' - Operator 'not'
    '( )' - Symbols for grouping operands with their correct operator.

Note: Include this parameter in every request that generates paginated output.
`

var FilterContainerVuln = `
string <= 1024 characters
Example: riskLevel eq 'high'

The filter for retrieving a subset of the image vulnerabilities list. Include this header in every request that generates paginated output.

Supported fields and operators:
Field 	Description 	Allowed values
clusterType 	The type of cluster 	kubernetes, amazonecs
name 	The name of the vulnerability 	Any value
clusterId 	The ID of the cluster 	Any value
imageId 	The The ID of the container image 	Any value
riskLevel 	The risk level of the vulnerability 	high, medium, low

Supported operators:

    eq - Operator "equal to"
    and - Operator "and"
`

var FilterK8s = `
string <= 1024 characters
Example: name eq 'example_cluster'

The filter for retrieving a subset of the Kubernetes cluster list, which is included in every request that generates paginated output.

Supported fields:

    launchType - The launch type of the cluster. Support values: [ EC2, FARGATE, NODEPOOL, VIRTUALNODE, STANDARD, AUTOPILOT, MANAGED ]
    name - The name of the Kubernetes cluster
    orchestrator - The orchestrator of the Cluster. Support values: [ Self-managed, Amazon EKS, Microsoft AKS, Google GKE, Alibaba Cloud ACK ]
    policyId - The ID of the policy associated with the cluster
    protectionStatus - The protection status of the cluster. Supported values: [ UNKNOWN, HEALTHY, UNHEALTHY, WARNING ]
    runtimeSecurityEnabled - Whether Runtime Security is enabled for the cluster. Supported values: [ true, false ]
    vulnerabilityScanEnabled - Whether Runtime Vulnerability Scanning is enabled for the cluster. Supported values: [ true, false ]
    malwareScanEnabled - Whether Runtime Malware Scanning is enabled for the cluster. Supported values: [ true, false ]
    groupId | The ID of the group associated with the cluster.

Supported operators:

    eq - Operator "equal to"
    contains - String partial match
    not - Operator "not"
    or - Operator "or"
    () - Symbols for grouping operands
`

var FilterECS = `
string <= 1024 characters
Example: name eq 'example_cluster'

The filter for retrieving a subset of Amazon ECS clusters. Include this parameter in every request that generates paginated output

Supported fields:

    name - The name of the Amazon ECS cluster

Supported operators:

    eq - Operator "equal to"
    contains - String partial match
    not - Operator "not"
    or - Operator "or"
`

var FilterK8Images = `
string <= 1024 characters
Example: id eq 'imageId_1' and clusterId eq 'clusterId_1'

The filter for retrieving a subset of the Kubernetes image list. Include this parameter in every request that generates paginated output

Supported fields:
Field 	Description
id 	The ID of the Kubernetes image
clusterid 	The ID of the Kubernetes cluster
digest 	The container image digest
repository 	The repository of the container image
registry 	The registry of the container image

Supported operators:

    eq - Operator "equal to"
    and - Operator "and"
    or - Operator "or"
    not - Operator "not"
    () - Symbols for grouping operands
`

var FilterEndpoints = `
string <= 1024 characters
Example: not (osName eq 'Windows') and eppAgentAntiMalwareScans eq 'enabled'

Filter for retrieving a subset of the endpoint information list.

Supported fields:

Field
	

Description
	

Supported values

endpointName
	

The name of the endpoint.
	

Any value

agentGuid
	

The ID of the endpoint on the Trend Vision One platform.
	

Any value

type
	

The type of endpoint.
	

desktop,server

edrSensorAdvancedRiskTelemetryStatus
	

The status of the advanced risk telemetry on the endpoint.
	

enabled,disabled,enabling,disabling,unknown

edrSensorConnectivity
	

The connectivity of the sensor installed on the endpoint.
	

connected,disconnected

eppAgentProtectionManager
	

The name of your protection manager.
	

Any value

eppAgentEndpointGroup
	

The name of the endpoint group.
	

Any value

osName
	

The operating system of the endpoint.
	

Any value

osPlatform
	

The platform of the operating system of the endpoint.
	

windows,mac,linux,unix,unknown

serviceGatewayOrProxy
	

The endpoints routed through a Service Gateway or Proxy
	

Any value that follows the format described in the 'serviceGatewayOrProxy' field

osArchitecture
	

The type of operating system running on the endpoint.
	

x86,x86_64,ppc64le,sparc,powerpc,aarch64

versionControlPolicy
	

The name of the version control policy.
	

Any value

agentUpdateStatus
	

The status of the agent update policy.
	

onSchedule,pause,disable,notSupported

agentUpdatePolicy
	

The agent update policy version used for version control policy APIs.
	

Any value

creditAllocatedLicenses
	

The features enabled on the endpoint
	

Endpoint sensor detection and response,Advanced Endpoint Security,Advanced Server & Workload Protection,SAP Scanner for Trend Vision One - Endpoint Security (Pro)

securityPolicy
	

The name of the security policy applied to the endpoint
	

Any value

securityPolicyOverriddenStatus
	

Whether the endpoint security policy was overridden
	

enabled,disabled

edrSensorStatus
	

The status of the endpoint sensor detection and response.
	

enabled,disabled,enabling,disabling,unknown

edrSensorComponentUpdatePolicy
	

The version control update policy for the module/pattern of the sensor installed on the endpoint.
	

nrepresents the latest version.n - xrepresents x snapshots prior (Example:n - 2).

edrSensorComponentUpdateStatus
	

The status of the module/pattern updates of the sensor installed on the endpoint.
	

pause,onSchedule,notSupported

eppAgentStatus
	

The connectivity status of the endpoint protection agent.
	

on,off,unknown

eppAgentPolicyName
	

The name of a policy from your protection manager.
	

Any value

isolationStatus
	

Indicates if an endpoint is isolated.
	

on,off,unknown

eppAgentComponentVersion
	

The agent component version.
	

outdatedVersion,latestVersion,unknownVersions,controlledLatestVersion

eppAgentComponentUpdatePolicy
	

The update policy for the module/pattern of the agent installed on the endpoint.
	

nrepresents the latest version.n - xrepresents x snapshots prior (Example:n - 2).

eppAgentComponentUpdateStatus
	

The status of the module/pattern updates of the agent installed on the endpoint.
	

pause,onSchedule,notSupported

    eppAgentAntiMalwareScans

    eppAgentBehaviorMonitoring

    eppAgentPredictiveMachineLearning

    eppAgentWebReputation

    eppAgentSuspiciousConnectionSettings

    eppAgentVulnerabilityProtection

    eppAgentDeviceControl

    eppAgentApplicationControl

    eppAgentFirewall

    eppAgentIntegratedEndpointSensor

    eppAgentDataLossPrevention

    eppAgentSmartFeedback

    eppAgentAntiMalware

    eppAgentActivityMonitoring

    eppAgentIntrusionPreventionSystem

    eppAgentLogInspection

    eppAgentIntegrityMonitoring

    eppAgentAgentSelfProtection

    eppAgentSAPScanner

    eppAgentSecurityAgentPasswordUnlock

	

The status of the features supported by the endpoint protection agent.
	

enabled,disabled,notSupported

availableActions
	

Allows you to select the discoverable endpoints that may require attention.
	

    immediateActionRequired: An issue occurred on the endpoint that requires user intervention.

    unmanaged: The endpoints are discoverable on your network but do not have any available protection or sensor agent program installed.

    sensorUpdateRequired: The endpoint has an older version of the Endpoint Sensor component installed (including Activity Monitoring and Apex One Endpoint Sensor) and should update to the latest Trend Vision One version.

    sepMaintenanceRecommended: The Standard Endpoint Protection endpoint does not have XDR Endpoint Sensor. The Detection and Response feature is off.

    a1MaintenanceRecommended: The Apex One endpoint does not have XDR Endpoint Sensor. The Detection and Response feature is off.

    swpMaintenanceRecommended: The Server & Workload Protection endpoint does not have XDR Endpoint Sensor. The Detection and Response feature is off.

    c1MaintenanceRecommended: The Cloud One endpoint does not have XDR Endpoint Sensor. The Detection and Response feature is off.

    sensorDisabled: The Managed Endpoint and Detection and Response features are off. The endpoint has XDR Endpoint Sensor.

securityDeployment
	

Allows you to select the endpoints that have a protection or sensor agent installed.
	

    managed: All endpoints on your network that have a protection or sensor agent installed

    userProtection: All endpoints that have the Standard Endpoint Protection agent installed (includes endpoints that also have the Endpoint sensor detection and response feature enabled)

    workloadProtection: All endpoints that have the Server & Workload agent installed (includes endpoints that also have the Endpoint sensor detection and response feature enabled)

    sensorOnly: All endpoints that only have the Sensor agent installed (endpoints that are not protected by a Standard Endpoint Protection or Server & Workload Protection agent)

    connectedEndpointProtection: All endpoints that have protection or sensor agent installed from a connected product.

Supported operators:
Operator 	Description
eq 	Operator 'equal to'
and 	Operator 'and'
or 	Operator 'or'
not 	Operator 'not'
() 	Symbols for grouping operands

Additional functions:
Function 	Description 	Notes
hassubset() 	Checks if the array contains a subset 	Applicable to creditAllocatedLicenses only

Note:

    for edrSensorStatus and edrSensorAdvancedRiskTelemetryStatus, When filtering for all xxabled endpoionts as the same as the UI, you can filter by status eq xxabled or status eq xxabling.
    immediateActionRequired is mutually exclusive with other filters.
    You can't use availableActions and securityDeployment together.
    If a query contains both eppAgent<FeatureName> and eppComponentVersion, only one value of eppComponentVersion is included in the results.
    availableActions and securityDeployment are not returned in response.

`

var FilterEndpointTasks = `
string <= 1024 characters
Example: status eq 'succeeded'

The filter for retrieving a subset of the task list.

Supported fields:
Field 	Description 	Supported values
id 	The ID of the task 	Any value
agentGuid 	The ID of the endpoint on the Trend Vision One platform 	Any value
status 	The status of the task 	running,succeeded,failed
action 	The type of the task 	export,delete

Supported operators:
Operator 	Description
eq 	Operator 'equal to'
and 	Operator 'and'
or 	Operator 'or'
not 	Operator 'not'
() 	Symbols for grouping operands
`
