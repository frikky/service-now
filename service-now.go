package servicenow

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	"time"
)

type multiChange struct {
	Detail []ChangeRequest `json:"result"`
	Raw    []byte          `json:"-"`
}

type resultWrapper struct {
	Detail ChangeRequest `json:"result"`
	Raw    []byte        `json:"-"`
}

type CMDBItem struct {
	Status              string `json:"__status,omitempty"`
	Asset               string `json:"asset,omitempty"`
	AssetTag            string `json:"asset_tag,omitempty"`
	Assigned            string `json:"assigned,omitempty"`
	AssignedTo          string `json:"assigned_to,omitempty"`
	AssignmentGroup     string `json:"assignment_group,omitempty"`
	Attributes          string `json:"attributes,omitempty"`
	CanPrint            string `json:"can_print,omitempty"`
	Category            string `json:"category,omitempty"`
	ChangeControl       string `json:"change_control,omitempty"`
	CheckedIn           string `json:"checked_in,omitempty"`
	CheckedOut          string `json:"checked_out,omitempty"`
	Comments            string `json:"comments,omitempty"`
	Company             string `json:"company,omitempty"`
	CorrelationID       string `json:"correlation_id,omitempty"`
	Cost                string `json:"cost,omitempty"`
	CostCc              string `json:"cost_cc,omitempty"`
	CostCenter          string `json:"cost_center,omitempty"`
	DeliveryDate        string `json:"delivery_date,omitempty"`
	Department          string `json:"department,omitempty"`
	DiscoverySource     string `json:"discovery_source,omitempty"`
	DNSDomain           string `json:"dns_domain,omitempty"`
	Due                 string `json:"due,omitempty"`
	DueIn               string `json:"due_in,omitempty"`
	FaultCount          string `json:"fault_count,omitempty"`
	FirstDiscovered     string `json:"first_discovered,omitempty"`
	Fqdn                string `json:"fqdn,omitempty"`
	GlAccount           string `json:"gl_account,omitempty"`
	InstallDate         string `json:"install_date,omitempty"`
	InstallStatus       string `json:"install_status,omitempty"`
	InvoiceNumber       string `json:"invoice_number,omitempty"`
	IPAddress           string `json:"ip_address,omitempty"`
	Justification       string `json:"justification,omitempty"`
	LastDiscovered      string `json:"last_discovered,omitempty"`
	LeaseID             string `json:"lease_id,omitempty"`
	Location            string `json:"location,omitempty"`
	MacAddress          string `json:"mac_address,omitempty"`
	MaintenanceSchedule string `json:"maintenance_schedule,omitempty"`
	ManagedBy           string `json:"managed_by,omitempty"`
	Manufacturer        string `json:"manufacturer,omitempty"`
	ModelID             string `json:"model_id,omitempty"`
	ModelNumber         string `json:"model_number,omitempty"`
	Monitor             string `json:"monitor,omitempty"`
	Name                string `json:"name,omitempty"`
	OperationalStatus   string `json:"operational_status,omitempty"`
	OrderDate           string `json:"order_date,omitempty"`
	OwnedBy             string `json:"owned_by,omitempty"`
	PoNumber            string `json:"po_number,omitempty"`
	PurchaseDate        string `json:"purchase_date,omitempty"`
	Schedule            string `json:"schedule,omitempty"`
	SerialNumber        string `json:"serial_number,omitempty"`
	ShortDescription    string `json:"short_description,omitempty"`
	SkipSync            string `json:"skip_sync,omitempty"`
	StartDate           string `json:"start_date,omitempty"`
	Subcategory         string `json:"subcategory,omitempty"`
	SupportGroup        string `json:"support_group,omitempty"`
	SupportedBy         string `json:"supported_by,omitempty"`
	SysClassName        string `json:"sys_class_name,omitempty"`
	SysClassPath        string `json:"sys_class_path,omitempty"`
	SysCreatedBy        string `json:"sys_created_by,omitempty"`
	SysCreatedOn        string `json:"sys_created_on,omitempty"`
	SysDomain           string `json:"sys_domain,omitempty"`
	SysDomainPath       string `json:"sys_domain_path,omitempty"`
	SysID               string `json:"sys_id,omitempty"`
	SysModCount         string `json:"sys_mod_count,omitempty"`
	SysTags             string `json:"sys_tags,omitempty"`
	SysUpdatedBy        string `json:"sys_updated_by,omitempty"`
	SysUpdatedOn        string `json:"sys_updated_on,omitempty"`
	Unverified          string `json:"unverified,omitempty"`
	Vendor              string `json:"vendor,omitempty"`
	WarrantyExpiration  string `json:"warranty_expiration,omitempty"`
}

type InnerValue struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}

type ChangeRequest struct {
	Status                   string     `json:"__status,omitempty"`
	Active                   bool       `json:"active,string,omitempty"`
	ActivityDue              string     `json:"activity_due,omitempty"`
	AdditionalAssigneeList   string     `json:"additional_assignee_list,omitempty"`
	Approval                 string     `json:"approval,omitempty"`
	ApprovalHistory          string     `json:"approval_history,omitempty"`
	ApprovalSet              string     `json:"approval_set,omitempty"`
	AssignedTo               string     `json:"assigned_to,omitempty"`
	AssignmentGroup          InnerValue `json:"assignment_group,omitempty"`
	BackoutPlan              string     `json:"backout_plan,omitempty"`
	BusinessDuration         string     `json:"business_duration,omitempty"`
	BusinessService          string     `json:"business_service,omitempty"`
	BusinessConsequences     string     `json:"u_business_consequences,omitempty"`
	CabDelegate              string     `json:"cab_delegate,omitempty"`
	CabRecommendation        string     `json:"cab_recommendation,omitempty"`
	CabRequired              bool       `json:"cab_required,string,omitempty"`
	CalendarDuration         string     `json:"calendar_duration,omitempty"`
	Category                 string     `json:"category,omitempty"`
	ChangePlan               string     `json:"change_plan,omitempty"`
	CloseCode                string     `json:"close_code,omitempty"`
	CloseNotes               string     `json:"close_notes,omitempty"`
	ClosedBy                 string     `json:"closed_by,omitempty"`
	CmdbCi                   string     `json:"cmdb_ci,omitempty"`
	Comments                 string     `json:"comments,omitempty"`
	CommentsAndWorkNotes     string     `json:"comments_and_work_notes,omitempty"`
	Company                  string     `json:"company,omitempty"`
	ConflictLastRun          string     `json:"conflict_last_run,omitempty"`
	ConflictStatus           string     `json:"conflict_status,omitempty"`
	ContactType              string     `json:"contact_type,omitempty"`
	CorrelationDisplay       string     `json:"correlation_display,omitempty"`
	CorrelationID            string     `json:"correlation_id,omitempty"`
	DeliveryPlan             string     `json:"delivery_plan,omitempty"`
	DeliveryTask             string     `json:"delivery_task,omitempty"`
	Description              string     `json:"description,omitempty"`
	GroupList                string     `json:"group_list,omitempty"`
	ImplementationPlan       string     `json:"implementation_plan,omitempty"`
	Justification            string     `json:"justification,omitempty"`
	Knowledge                bool       `json:"knowledge,string,omitempty"`
	Location                 string     `json:"location,omitempty"`
	MadeSLA                  bool       `json:"made_sla,string,omitempty"`
	Number                   string     `json:"number,omitempty"`
	OnHold                   bool       `json:"on_hold,string,omitempty"`
	OnHoldReason             string     `json:"on_hold_reason,omitempty"`
	OpenedBy                 string     `json:"opened_by,omitempty"`
	Order                    string     `json:"order,omitempty"`
	OutsideMaintSchedule     bool       `json:"outside_maintenance_schedule,string,omitempty"`
	Parent                   string     `json:"parent,omitempty"`
	Phase                    string     `json:"phase,omitempty"`
	PhaseState               string     `json:"phase_state,omitempty"`
	ProductionSystem         bool       `json:"production_system,string,omitempty"`
	Reason                   string     `json:"reason,omitempty"`
	RequestedBy              string     `json:"requested_by,omitempty"`
	ReviewComments           string     `json:"review_comments,omitempty"`
	ReviewStatus             string     `json:"review_status,omitempty"`
	RiskImpactAnalysis       string     `json:"risk_impact_analysis,omitempty"`
	ShortDescription         string     `json:"short_description,omitempty"`
	SLADue                   string     `json:"sla_due,omitempty"`
	State                    string     `json:"state,omitempty"`
	StdChangeProducerVersion string     `json:"std_change_producer_version,omitempty"`
	SysClassName             string     `json:"sys_class_name,omitempty"`
	SysCreatedBy             string     `json:"sys_created_by,omitempty"`
	SysDomain                string     `json:"sys_domain,omitempty"`
	SysDomainPath            string     `json:"sys_domain_path,omitempty"`
	SysID                    string     `json:"sys_id,omitempty"`
	SysTags                  string     `json:"sys_tags,omitempty"`
	SysUpdatedBy             string     `json:"sys_updated_by,omitempty"`
	TestPlan                 string     `json:"test_plan,omitempty"`
	TimeWorked               string     `json:"time_worked,omitempty"`
	Type                     string     `json:"type,omitempty"`
	UponApproval             string     `json:"upon_approval,omitempty"`
	UponReject               string     `json:"upon_reject,omitempty"`
	Urgency                  string     `json:"urgency,omitempty"`
	UserInput                string     `json:"user_input,omitempty"`
	WatchList                string     `json:"watch_list,omitempty"`
	WorkEnd                  string     `json:"work_end,omitempty"`
	WorkNotes                string     `json:"work_notes,omitempty"`
	WorkNotesList            string     `json:"work_notes_list,omitempty"`
	WorkStart                string     `json:"work_start,omitempty"`
	FollowUp                 string     `json:"follow_up,omitempty"`
	Identifier               string     `json:"identifier,omitempty,omitempty"`
	Name                     string     `json:"name,omitempty"`
}

//SysCreatedOn             SNTime      `json:"sys_created_on,omitempty"`
//SysModCount              json.Number `json:"sys_mod_count,omitempty"`
//SysUpdatedOn             SNTime      `json:"sys_updated_on,omitempty"`
//ReassignmentCount        json.Number `json:"reassignment_count,omitempty"`
//RequestedByDate          SNTime      `json:"requested_by_date,omitempty"`
//ReviewDate               SNTime      `json:"review_date,omitempty"`
//Risk                     json.Number `json:"risk,omitempty"`
//Scope                    json.Number `json:"scope,omitempty"`
//StartDate                SNTime      `json:"start_date,omitempty"`
//OpenedAt                 SNTime      `json:"opened_at,omitempty"`
//Impact                   json.Number `json:"impact,omitempty"`
//ClosedAt                 SNTime      `json:"closed_at,omitempty"`
//CabDate                  SNTime      `json:"cab_date,omitempty"`
//DueDate                  SNTime      `json:"due_date,omitempty"`
//EndDate                  SNTime      `json:"end_date,omitempty"`
//Escalation               json.Number `json:"escalation,omitempty"`
//ExpectedStart            SNTime      `json:"expected_start,omitempty"`
//Priority                 json.Number `json:"priority,omitempty"`

type Logindata struct {
	Url string
	Ro  grequests.RequestOptions
}

// Defines the creation of requestoptions
// Takes 3 parameters:
//  1. Url string
//  2. Username string
//  3. Verify bool
// Returns a login as Logindata struct based on function parameters
func CreateLogin(inurl string, username string, password string, verify bool) Logindata {
	return Logindata{
		Url: inurl,
		Ro: grequests.RequestOptions{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Auth:               []string{username, password},
			RequestTimeout:     time.Duration(10) * time.Second,
			InsecureSkipVerify: !verify,
		},
	}
}

// Defines how to get tickets from a table in service-now
// Takes 3 parameters:
//  1. Table string
//  2. Limit int
//  3. Params string
// Returns multiple snow changes as struct multiChange and response error
func (client Logindata) GetTable(table string, limit int, params string) (*multiChange, error) {
	url := fmt.Sprintf("%s/api/now/table/%s?sysparm_limit=%d%s", client.Url, table, limit, params)
	ret, err := grequests.Get(url, &client.Ro)

	if err != nil {
		return nil, err
	}

	// Errors occur some places
	parsedRet := new(multiChange)
	_ = json.Unmarshal(ret.Bytes(), parsedRet)
	parsedRet.Raw = ret.Bytes()

	return parsedRet, nil
}
