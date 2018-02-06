package constant

//Defines the constants for the system
const (
	Version = 2.0

	TypeLookup    = "Lookups"
	TypePortlet   = "Portlets"
	TypeQuery     = "Queries"
	TypeProcess   = "Processes"
	TypePage      = "Pages"
	TypeMenu      = "Menus"
	TypeObject    = "Objects"
	TypeView      = "Views"
	TypeMigration = "Migrations"

	TypeCustomObjectInstance     = "CustomObjectInstances"
	TypeResourceClassInstance    = "ResourceClassInstances"
	TypeWipClassInstance         = "WipClassInstances"
	TypeInvestmentClassInstance  = "InvestmentClassInstances"
	TypeTransactionClassInstance = "TransactionClassInstances"
	TypeResourceInstance         = "ResourceInstances"
	TypeUserInstance             = "UserInstances"
	TypeProjectInstance          = "ProjectInstances"
	TypeProgramInstance          = "ProgramInstances"
	TypeIdeaInstance             = "IdeaInstances"
	TypeApplicationInstance      = "ApplicationInstances"
	TypeAssetInstance            = "AssetInstances"
	TypeOtherInvestmentInstance  = "OtherInvestmentInstances"
	TypeProductInstance          = "ProductInstances"
	TypeServiceInstance          = "ServiceInstances"
	TypeBenefitPlanInstance      = "BenefitPlanInstances"
	TypeBudgetPlanInstance       = "BudgetPlanInstances"
	TypeCategoryInstance         = "CategoryInstances"
	TypeChangeInstance           = "ChangeInstances"
	TypeChargeCodeInstance       = "ChargeCodeInstances"
	TypeCompanyClassInstance     = "CompanyClassInstances"
	TypeCostPlanInstance         = "CostPlanInstances"
	TypeCostPlusCodeInstance     = "CostPlusCodeInstances"
	TypeDepartmentInstance       = "DepartmentInstances"
	TypeEntityInstance           = "EntityInstances"
	TypeGroupInstance            = "GroupInstances"
	TypeIncidentInstance         = "IncidentInstances"
	TypeIssueInstance            = "IssueInstances"
	TypeOBSInstance              = "OBSInstances"
	TypePortfolioInstance        = "PortfolioInstances"
	TypeReleaseInstance          = "ReleaseInstances"
	TypeReleasePlanInstance      = "ReleasePlanInstances"
	TypeRequirementInstance      = "RequirementInstances"
	TypeRequisitionInstance      = "RequisitionInstances"
	TypeRiskInstance             = "RiskInstances"
	TypeRoleInstance             = "RoleInstances"
	TypeThemeInstance            = "ThemeInstances"
	TypeVendorInstance           = "VendorInstances"

	ActionReplace = "replace"
	ActionUpdate  = "update"
	ActionRemove  = "remove"
	ActionInsert  = "insert"

	Read    = "r"
	Write   = "w"
	Migrate = "m"

	FolderRead      = "_read/"
	FolderWrite     = "_write/"
	FolderMigration = "_migration/"
	FolderDebug     = "_debug/"
	FolderPackage   = "_packages/"
	FolderMock      = "mock/"

	Undefined     = ""
	OutputError   = "error"
	OutputWarning = "warning"
	OutputSuccess = "success"
	OutputIgnored = "ignored"

	PackageActionChangePartitionModel = "changePartitionModel"
	PackageActionChangePartition      = "changePartition"
	PackageActionReplaceString        = "replaceString"

	ColumnLeft  = "left"
	ColumnRight = "right"

	ElementTypeLink        = "link"
	ElementTypeAction      = "action"
	ElementTypeAttribute   = "attribute"
	ElementTypeActionGroup = "actionGroup"

	Source = "source"
	Target = "target"
)
