package schemas

// geoIndicatorsSchema defines the column names and types for the `geoIndicators` table.
var GeoIndicatorsSchema = map[string]string{
    "rid": "integer",
    "PrimaryKey": "text",
    "DBKey": "text",
    "ProjectKey": "text",
    "DateVisited": "date",
    "EcologicalSiteId": "text",
    "PercentCoveredByEcoSite": "double precision",
    "Latitude_NAD83": "double precision",
    "Longitude_NAD83": "double precision",
    "LocationStatus": "text",
    "LocationType": "text",
    "Latitude_NAD83_Actual": "double precision",
    "Longitude_NAD83_Actual": "double precision",
    "BareSoil": "double precision",
    "TotalFoliarCover": "double precision",
    "AH_AnnGraminoidCover": "double precision",
    "AH_ForbCover": "double precision",
    "AH_AnnForbGraminoidCover": "double precision",
    "AH_GraminoidCover": "double precision",
    "AH_PerenForbCover": "double precision",
    "AH_PerenForbGraminoidCover": "double precision",
    "AH_PerenGraminoidCover": "double precision",
    "AH_ShrubCover": "double precision",
    "FH_CyanobacteriaCover": "double precision",
    "FH_DepSoilCover": "double precision",
    "FH_DuffCover": "double precision",
    "FH_EmbLitterCover": "double precision",
    "FH_HerbLitterCover": "double precision",
    "FH_LichenCover": "double precision",
    "FH_MossCover": "double precision",
    "FH_RockCover": "double precision",
    "FH_TotalLitterCover": "double precision",
    "FH_VagrLichenCover": "double precision",
    "FH_WaterCover": "double precision",
    "FH_WoodyLitterCover": "double precision",
    "GapCover_101_200": "double precision",
    "GapCover_200_plus": "double precision",
    "GapCover_25_50": "double precision",
    "GapCover_25_plus": "double precision",
    "GapCover_51_100": "double precision",
    "Hgt_Forb_Avg": "double precision",
    "Hgt_Graminoid_Avg": "double precision",
    "Hgt_Herbaceous_Avg": "double precision",
    "Hgt_PerenForb_Avg": "double precision",
    "Hgt_PerenForbGraminoid_Avg": "double precision",
    "Hgt_PerenGraminoid_Avg": "double precision",
    "Hgt_Woody_Avg": "double precision",
    "RH_AnnualProd": "text",
    "RH_BareGround": "text",
    "RH_BioticIntegrity": "text",
    "RH_CommentsBI": "text",
    "RH_CommentsHF": "text",
    "RH_CommentsSS": "text",
    "RH_Compaction": "text",
    "RH_DeadDyingPlantParts": "text",
    "RH_FuncSructGroup": "text",
    "RH_Gullies": "text",
    "RH_HydrologicFunction": "text",
    "RH_InvasivePlants": "text",
    "RH_LitterAmount": "text",
    "RH_LitterMovement": "text",
    "RH_PedestalsTerracettes": "text",
    "RH_PlantCommunityComp": "text",
    "RH_ReprodCapabilityPeren": "text",
    "RH_Rills": "text",
    "RH_SoilSiteStability": "text",
    "RH_SoilSurfLossDeg": "text",
    "RH_SoilSurfResisErosion": "text",
    "RH_WaterFlowPatterns": "text",
    "RH_WindScouredAreas": "text",
    "SoilStability_All": "double precision",
    "SoilStability_Protected": "double precision",
    "SoilStability_Unprotected": "double precision",
    "DateLoadedInDb": "date",
    "mlra_name": "text",
    "mlrarsym": "text",
    "na_l1name": "text",
    "na_l2name": "text",
    "us_l3name": "text",
    "us_l4name": "text",
    "State": "text",
    "modis_landcover": "text",
    "wkb_geometry": "text",
    "Precipitation_Long_Term_MEAN": "double precision",
    "Runoff_Long_Term_MEAN": "double precision",
    "Sediment_Yield_Long_Term_MEAN": "double precision",
    "Soil_Loss_Long_Term_MEAN": "double precision",
    "Precipitation_Yearly_Maximum_Daily_2_YR": "double precision",
}
