package postgres

//go:generate pg-table-bindings-wrapper --type=storage.ComplianceOperatorScanConfigurationV2 --search-category COMPLIANCE_SCAN_CONFIG --references=storage.ComplianceOperatorProfileV2 --get-all-func --feature-flag ComplianceEnhancements
