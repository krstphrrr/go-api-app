package schemas

// aero_summarySchema defines the column names and types for the `aero_summary` table.
var DataAeroSummarySchema = map[string]string{
    "rid": "integer",
    "PrimaryKey": "text",
    "horizontal_flux_total_LPI": "double precision",
    "horizontal_flux_total_UPI": "double precision",
    "horizontal_flux_total_MD": "double precision",
    "horizontal_flux_total_MN": "double precision",
    "horizontal_flux_total_STD": "double precision",
    "vertical_flux_LPI": "double precision",
    "vertical_flux_UPI": "double precision",
    "vertical_flux_MD": "double precision",
    "vertical_flux_MN": "double precision",
    "vertical_flux_STD": "double precision",
    "PM1_LPI": "double precision",
    "PM1_UPI": "double precision",
    "PM1_MD": "double precision",
    "PM1_MN": "double precision",
    "PM1_STD": "double precision",
    "PM2_5_LPI": "double precision",
    "PM2_5_UPI": "double precision",
    "PM2_5_MD": "double precision",
    "PM2_5_MN": "double precision",
    "PM2_5_STD": "double precision",
    "PM10_LPI": "double precision",
    "PM10_UPI": "double precision",
    "PM10_MD": "double precision",
    "PM10_MN": "double precision",
    "PM10_STD": "double precision",
}
