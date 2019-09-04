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

type Approvals struct {
	Result struct {
		UINotification []interface{} `json:"$$uiNotification"`
		Theme          struct {
			FooterFixed bool `json:"footer_fixed"`
			Footer      struct {
			} `json:"footer"`
			FooterDv     string `json:"footer_dv"`
			NavbarFixed  bool   `json:"navbar_fixed"`
			SysTags      string `json:"sys_tags"`
			HeaderDv     string `json:"header_dv"`
			SysClassName string `json:"sys_class_name"`
			CSSVariables string `json:"css_variables"`
			SysID        string `json:"sys_id"`
			Name         string `json:"name"`
			Header       struct {
				Template string `json:"template"`
				CSS      string `json:"css"`
				Data     struct {
					ConnectSupportQueueID interface{} `json:"connect_support_queue_id"`
					LoginPage             interface{} `json:"login_page"`
					Menu                  struct {
					} `json:"menu"`
					LoginWidget struct {
						Template string `json:"template"`
						CSS      string `json:"css"`
						Data     struct {
							UsernameMsg     string      `json:"usernameMsg"`
							IsLoggedIn      bool        `json:"is_logged_in"`
							PageURI         string      `json:"pageURI"`
							ErrorMsg2       string      `json:"errorMsg2"`
							ForgetMe        bool        `json:"forgetMe"`
							PasswordMsg     string      `json:"passwordMsg"`
							MultissoEnabled bool        `json:"multisso_enabled"`
							DefaultIdp      interface{} `json:"default_idp"`
							ErrorMsg        string      `json:"errorMsg"`
							ForgetMeDefault bool        `json:"forgetMeDefault"`
						} `json:"data"`
						RectangleID  string        `json:"rectangle_id"`
						OptionSchema string        `json:"option_schema"`
						SysClassName string        `json:"sys_class_name"`
						Dependencies []interface{} `json:"dependencies"`
						SysID        string        `json:"sys_id"`
						Public       bool          `json:"public"`
						ClientScript string        `json:"client_script"`
						Name         string        `json:"name"`
						Options      struct {
							Color          string `json:"color"`
							SpColumnDv     string `json:"sp_column_dv"`
							Active         bool   `json:"active"`
							ColorDv        string `json:"color_dv"`
							Title          string `json:"title"`
							SysTags        string `json:"sys_tags"`
							SysClassName   string `json:"sys_class_name"`
							SizeDv         string `json:"size_dv"`
							SpWidgetDv     string `json:"sp_widget_dv"`
							Size           string `json:"size"`
							SysName        string `json:"sys_name"`
							ID             string `json:"id"`
							SpWidget       string `json:"sp_widget"`
							SysClassNameDv string `json:"sys_class_name_dv"`
							Order          int    `json:"order"`
						} `json:"options"`
						Rectangle struct {
							Color          string `json:"color"`
							SpColumnDv     string `json:"sp_column_dv"`
							Active         bool   `json:"active"`
							ColorDv        string `json:"color_dv"`
							Title          string `json:"title"`
							SysTags        string `json:"sys_tags"`
							SysClassName   string `json:"sys_class_name"`
							SizeDv         string `json:"size_dv"`
							SpWidgetDv     string `json:"sp_widget_dv"`
							Size           string `json:"size"`
							SysName        string `json:"sys_name"`
							ID             string `json:"id"`
							SpWidget       string `json:"sp_widget"`
							SysClassNameDv string `json:"sys_class_name_dv"`
							Order          int    `json:"order"`
						} `json:"rectangle"`
						ID           string        `json:"id"`
						FieldList    string        `json:"field_list"`
						ControllerAs string        `json:"controller_as"`
						Providers    []interface{} `json:"providers"`
						ServerTime   string        `json:"_server_time"`
					} `json:"loginWidget"`
					Typeahead struct {
						Template    string `json:"template"`
						NgTemplates struct {
							SpTypeaheadHTML      string `json:"sp-typeahead.html"`
							SpTypeaheadPopupHTML string `json:"sp-typeahead-popup.html"`
						} `json:"ngTemplates"`
						CSS  string `json:"css"`
						Data struct {
							NavigationMsg      string      `json:"navigationMsg"`
							SearchMsg          string      `json:"searchMsg"`
							SearchType         interface{} `json:"searchType"`
							Limit              int         `json:"limit"`
							TypeaheadTemplates struct {
								SpTypeaheadScHTML          string `json:"sp-typeahead-sc.html"`
								SpTypeaheadScItserviceHTML string `json:"sp-typeahead-sc_itservice.html"`
							} `json:"typeaheadTemplates"`
							Results   []interface{} `json:"results"`
							ResultMsg string        `json:"resultMsg"`
						} `json:"data"`
						Link         string        `json:"link"`
						RectangleID  string        `json:"rectangle_id"`
						OptionSchema string        `json:"option_schema"`
						SysClassName string        `json:"sys_class_name"`
						Dependencies []interface{} `json:"dependencies"`
						SysID        string        `json:"sys_id"`
						Public       bool          `json:"public"`
						ClientScript string        `json:"client_script"`
						Name         string        `json:"name"`
						Options      struct {
							Color          string `json:"color"`
							SpColumnDv     string `json:"sp_column_dv"`
							Active         bool   `json:"active"`
							ColorDv        string `json:"color_dv"`
							Title          string `json:"title"`
							SysTags        string `json:"sys_tags"`
							SysClassName   string `json:"sys_class_name"`
							SizeDv         string `json:"size_dv"`
							Glyph          string `json:"glyph"`
							SpWidgetDv     string `json:"sp_widget_dv"`
							Size           string `json:"size"`
							SysName        string `json:"sys_name"`
							ID             string `json:"id"`
							SpWidget       string `json:"sp_widget"`
							SysClassNameDv string `json:"sys_class_name_dv"`
							Order          int    `json:"order"`
						} `json:"options"`
						Rectangle struct {
							Color          string `json:"color"`
							SpColumnDv     string `json:"sp_column_dv"`
							Active         bool   `json:"active"`
							ColorDv        string `json:"color_dv"`
							Title          string `json:"title"`
							SysTags        string `json:"sys_tags"`
							SysClassName   string `json:"sys_class_name"`
							SizeDv         string `json:"size_dv"`
							Glyph          string `json:"glyph"`
							SpWidgetDv     string `json:"sp_widget_dv"`
							Size           string `json:"size"`
							SysName        string `json:"sys_name"`
							ID             string `json:"id"`
							SpWidget       string `json:"sp_widget"`
							SysClassNameDv string `json:"sys_class_name_dv"`
							Order          int    `json:"order"`
						} `json:"rectangle"`
						ID           string        `json:"id"`
						FieldList    string        `json:"field_list"`
						ControllerAs string        `json:"controller_as"`
						Providers    []interface{} `json:"providers"`
						ServerTime   string        `json:"_server_time"`
					} `json:"typeahead"`
				} `json:"data"`
				Link         string        `json:"link"`
				OptionSchema string        `json:"option_schema"`
				SysClassName string        `json:"sys_class_name"`
				Dependencies []interface{} `json:"dependencies"`
				SysID        string        `json:"sys_id"`
				Public       bool          `json:"public"`
				ClientScript string        `json:"client_script"`
				Name         string        `json:"name"`
				Options      struct {
					SpWidgetDv string `json:"sp_widget_dv"`
					SpColumnDv string `json:"sp_column_dv"`
					Active     bool   `json:"active"`
					SysTags    string `json:"sys_tags"`
					Order      int    `json:"order"`
				} `json:"options"`
				Rectangle struct {
					SpWidgetDv string `json:"sp_widget_dv"`
					SpColumnDv string `json:"sp_column_dv"`
					Active     bool   `json:"active"`
					SysTags    string `json:"sys_tags"`
					Order      int    `json:"order"`
				} `json:"rectangle"`
				ID           string        `json:"id"`
				FieldList    string        `json:"field_list"`
				ControllerAs string        `json:"controller_as"`
				Providers    []interface{} `json:"providers"`
				ServerTime   string        `json:"_server_time"`
			} `json:"header"`
			SysName        string `json:"sys_name"`
			SysClassNameDv string `json:"sys_class_name_dv"`
		} `json:"theme"`
		Containers []struct {
			SysID              string `json:"sys_id"`
			BootstrapAlt       bool   `json:"bootstrap_alt"`
			Subheader          bool   `json:"subheader"`
			Background         string `json:"background"`
			Width              string `json:"width"`
			ContainerClassName string `json:"container_class_name"`
			Title              string `json:"title"`
			Rows               []struct {
				SysID   string `json:"sys_id"`
				Columns []struct {
					SysID       string        `json:"sys_id"`
					SizeClasses string        `json:"size_classes"`
					Rows        []interface{} `json:"rows"`
					Widgets     []struct {
						SysID  string `json:"sys_id"`
						Widget struct {
							Template string `json:"template"`
							CSS      string `json:"css"`
							Data     struct {
							} `json:"data"`
							RectangleID  string        `json:"rectangle_id"`
							OptionSchema string        `json:"option_schema"`
							SysClassName string        `json:"sys_class_name"`
							Dependencies []interface{} `json:"dependencies"`
							SysID        string        `json:"sys_id"`
							Public       bool          `json:"public"`
							ClientScript string        `json:"client_script"`
							Name         string        `json:"name"`
							Options      struct {
								Color          string `json:"color"`
								SpColumn       string `json:"sp_column"`
								SpColumnDv     string `json:"sp_column_dv"`
								Active         bool   `json:"active"`
								ColorDv        string `json:"color_dv"`
								SysTags        string `json:"sys_tags"`
								SysClassName   string `json:"sys_class_name"`
								SizeDv         string `json:"size_dv"`
								SpWidgetDv     string `json:"sp_widget_dv"`
								Size           string `json:"size"`
								SpWidget       string `json:"sp_widget"`
								SysClassNameDv string `json:"sys_class_name_dv"`
								Order          int    `json:"order"`
							} `json:"options"`
							Rectangle struct {
								Color          string `json:"color"`
								SpColumn       string `json:"sp_column"`
								SpColumnDv     string `json:"sp_column_dv"`
								Active         bool   `json:"active"`
								ColorDv        string `json:"color_dv"`
								SysTags        string `json:"sys_tags"`
								SysClassName   string `json:"sys_class_name"`
								SizeDv         string `json:"size_dv"`
								SpWidgetDv     string `json:"sp_widget_dv"`
								Size           string `json:"size"`
								SpWidget       string `json:"sp_widget"`
								SysClassNameDv string `json:"sys_class_name_dv"`
								Order          int    `json:"order"`
							} `json:"rectangle"`
							ID           string        `json:"id"`
							FieldList    string        `json:"field_list"`
							ControllerAs string        `json:"controller_as"`
							Providers    []interface{} `json:"providers"`
							ServerTime   string        `json:"_server_time"`
						} `json:"widget"`
						Title string `json:"title"`
					} `json:"widgets"`
					ClassName string `json:"class_name"`
					Order     int    `json:"order"`
				} `json:"columns"`
				ClassName string `json:"class_name"`
				Order     int    `json:"order"`
			} `json:"rows"`
			ClassName string `json:"class_name"`
			Order     int    `json:"order"`
		} `json:"containers"`
		Page struct {
			SysID          string `json:"sys_id"`
			Internal       bool   `json:"internal"`
			CSS            string `json:"css"`
			Public         bool   `json:"public"`
			Draft          bool   `json:"draft"`
			SysName        string `json:"sys_name"`
			ID             string `json:"id"`
			Title          string `json:"title"`
			SysTags        string `json:"sys_tags"`
			SysClassName   string `json:"sys_class_name"`
			SysClassNameDv string `json:"sys_class_name_dv"`
		} `json:"page"`
		Portal struct {
			SpRectangleMenu       string `json:"sp_rectangle_menu"`
			SqandaKnowledgeBaseDv string `json:"sqanda_knowledge_base_dv"`
			ScCategoryPageDv      string `json:"sc_category_page_dv"`
			Icon                  string `json:"icon"`
			ScCatalog             string `json:"sc_catalog"`
			KbKnowledgePageDv     string `json:"kb_knowledge_page_dv"`
			Title                 string `json:"title"`
			ScCatalogPage         string `json:"sc_catalog_page"`
			SysClassName          string `json:"sys_class_name"`
			SysID                 string `json:"sys_id"`
			ThemeDv               string `json:"theme_dv"`
			Default               bool   `json:"default"`
			HomepageDv            string `json:"homepage_dv"`
			NotfoundPageDv        string `json:"notfound_page_dv"`
			ScCatalogDv           string `json:"sc_catalog_dv"`
			Logo                  string `json:"logo"`
			SysName               string `json:"sys_name"`
			SpRectangleMenuDv     string `json:"sp_rectangle_menu_dv"`
			Theme                 string `json:"theme"`
			ScCatalogPageDv       string `json:"sc_catalog_page_dv"`
			SysClassNameDv        string `json:"sys_class_name_dv"`
			QuickStartConfig      string `json:"quick_start_config"`
			KbKnowledgePage       string `json:"kb_knowledge_page"`
			SysTags               string `json:"sys_tags"`
			LoginPageDv           string `json:"login_page_dv"`
			SpChatQueueDv         string `json:"sp_chat_queue_dv"`
			CSSVariables          string `json:"css_variables"`
			NotfoundPage          string `json:"notfound_page"`
			KbKnowledgeBaseDv     string `json:"kb_knowledge_base_dv"`
			KbKnowledgeBase       string `json:"kb_knowledge_base"`
			LoginPage             string `json:"login_page"`
			URLSuffix             string `json:"url_suffix"`
			Homepage              string `json:"homepage"`
		} `json:"portal"`
		User struct {
			CalendarIntegration     string `json:"calendar_integration"`
			LastLoginTime           string `json:"last_login_time"`
			LastLoginDevice         string `json:"last_login_device"`
			UBuOwnerGroupDv         string `json:"u_bu_owner_group_dv"`
			UCountryDv              string `json:"u_country_dv"`
			Notification            string `json:"notification"`
			SysDomain               string `json:"sys_domain"`
			Vip                     bool   `json:"vip"`
			LastLogin               string `json:"last_login"`
			LoggedIn                bool   `json:"logged_in"`
			CanDebug                bool   `json:"can_debug"`
			Active                  bool   `json:"active"`
			UInfonowNotification    string `json:"u_infonow_notification"`
			CalendarIntegrationDv   string `json:"calendar_integration_dv"`
			UManagerLevelDv         string `json:"u_manager_level_dv"`
			TimeSheetPolicyDv       string `json:"time_sheet_policy_dv"`
			LocationDv              string `json:"location_dv"`
			Name                    string `json:"name"`
			CanDebugAdmin           bool   `json:"can_debug_admin"`
			UBlockUpdates           bool   `json:"u_block_updates"`
			PasswordNeedsReset      bool   `json:"password_needs_reset"`
			DefaultPerspectiveDv    string `json:"default_perspective_dv"`
			UManagerLevel           string `json:"u_manager_level"`
			ScheduleDv              string `json:"schedule_dv"`
			UserName                string `json:"user_name"`
			FailedAttempts          int    `json:"failed_attempts"`
			EduStatus               string `json:"edu_status"`
			SysDomainDv             string `json:"sys_domain_dv"`
			EduStatusDv             string `json:"edu_status_dv"`
			SysClassName            string `json:"sys_class_name"`
			SysID                   string `json:"sys_id"`
			DepartmentDv            string `json:"department_dv"`
			InternalIntegrationUser bool   `json:"internal_integration_user"`
			ManagerDv               string `json:"manager_dv"`
			FirstName               string `json:"first_name"`
			SysClassNameDv          string `json:"sys_class_name_dv"`
			NotificationDv          string `json:"notification_dv"`
			UInfonowNotificationDv  string `json:"u_infonow_notification_dv"`
			Manager                 string `json:"manager"`
			UIsManager              bool   `json:"u_is_manager"`
			LockedOut               bool   `json:"locked_out"`
			Auditor                 bool   `json:"auditor"`
			LastName                string `json:"last_name"`
			Photo                   string `json:"photo"`
			SysTags                 string `json:"sys_tags"`
			CompanyDv               string `json:"company_dv"`
			LdapServerDv            string `json:"ldap_server_dv"`
			ULastUpdatedFromSource  string `json:"u_last_updated_from_source"`
			BuildingDv              string `json:"building_dv"`
			URolesSuspended         bool   `json:"u_roles_suspended"`
			CostCenterDv            string `json:"cost_center_dv"`
		} `json:"user"`
	} `json:"result"`
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
	User                     InnerValue `json:"user,omitempty"`
	BackoutPlan              string     `json:"backout_plan,omitempty"`
	BusinessDuration         string     `json:"business_duration,omitempty"`
	BusinessService          string     `json:"business_service,omitempty"`
	BusinessConsequences     string     `json:"u_business_consequences,omitempty"`
	CabDelegate              string     `json:"cab_delegate,omitempty"`
	CabRecommendation        string     `json:"cab_recommendation,omitempty"`
	CabRequired              bool       `json:"cab_required,string,omitempty"`
	CalendarDuration         string     `json:"calendar_duration,omitempty"`
	Category                 string     `json:"category,omitempty"`
	UserName                 string     `json:"user_name,omitempty"`
	CatItem                  InnerValue `json:"cat_item,omitempty"`
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
	Group                    InnerValue `json:"group,omitempty"`
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
	ClosedAt                 string     `json:"closed_at,omitempty"`
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
//CabDate                  SNTime      `json:"cab_date,omitempty"`
//DueDate                  SNTime      `json:"due_date,omitempty"`
//EndDate                  SNTime      `json:"end_date,omitempty"`
//Escalation               json.Number `json:"escalation,omitempty"`
//ExpectedStart            SNTime      `json:"expected_start,omitempty"`
//Priority                 json.Number `json:"priority,omitempty"`

type GlobalSearch struct {
	Raw    []byte `json:"-"`
	Result struct {
		Groups []struct {
			Description string `json:"description"`
			Tables      []struct {
				Filter         string `json:"filter"`
				Label          string `json:"label"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				ConditionQuery string `json:"condition_query"`
				LabelPlural    string `json:"label_plural"`
			} `json:"tables"`
			Name          string `json:"name"`
			ID            string `json:"id"`
			ResultCount   int    `json:"result_count"`
			SearchResults []struct {
				Records []interface{} `json:"records"`
				Label   string        `json:"label"`
				Name    string        `json:"name"`
				Fields  []struct {
					Label       string `json:"label"`
					Name        string `json:"name"`
					Type        string `json:"type"`
					LabelPlural string `json:"label_plural"`
					MaxLength   int    `json:"max_length"`
					Reference   string `json:"reference,omitempty"`
				} `json:"fields"`
				Query         string `json:"query"`
				LabelPlural   string `json:"label_plural"`
				RecordCount   int    `json:"record_count"`
				AllResultsURL string `json:"all_results_url"`
			} `json:"search_results"`
		} `json:"groups"`
	} `json:"result"`
}

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

// Replaces grequests for inner requests
func (client Logindata) GlobalSearchGet(url string) (*GlobalSearch, error) {
	ret, err := grequests.Get(url, &client.Ro)

	if err != nil {
		return nil, err
	}

	// Errors occur some places
	parsedRet := new(GlobalSearch)
	_ = json.Unmarshal(ret.Bytes(), parsedRet)
	parsedRet.Raw = ret.Bytes()

	return parsedRet, nil
}

// Replaces grequests for inner requests
func (client Logindata) GetApprovals(url string) (*Approvals, error) {
	ret, err := grequests.Get(url, &client.Ro)

	if err != nil {
		return nil, err
	}

	// Errors occur some places
	parsedRet := new(Approvals)
	err = json.Unmarshal(ret.Bytes(), parsedRet)
	if err != nil {
		return parsedRet, err
	}
	//parsedRet.Raw = ret.Bytes()

	return parsedRet, nil
}

// Replaces grequests for inner requests
func (client Logindata) Get(url string) (*resultWrapper, error) {
	ret, err := grequests.Get(url, &client.Ro)

	if err != nil {
		return nil, err
	}

	// Errors occur some places
	parsedRet := new(resultWrapper)
	_ = json.Unmarshal(ret.Bytes(), parsedRet)
	parsedRet.Raw = ret.Bytes()

	return parsedRet, nil
}

// Replaces grequests for inner requests
func (client Logindata) Patch(url string, data []byte) (*resultWrapper, error) {
	// Set data
	client.Ro.JSON = data

	ret, err := grequests.Patch(url, &client.Ro)
	fmt.Println(ret)

	if err != nil {
		return nil, err
	}

	// Errors occur some places
	parsedRet := new(resultWrapper)
	err = json.Unmarshal(ret.Bytes(), parsedRet)
	if err != nil {
		return parsedRet, err
	}
	parsedRet.Raw = ret.Bytes()

	return parsedRet, nil
}

// Replaces grequests for inner requests
func (client Logindata) Put(url string, data []byte) (*resultWrapper, error) {
	// Set data
	client.Ro.JSON = data

	ret, err := grequests.Put(url, &client.Ro)

	if err != nil {
		return nil, err
	}

	// Errors occur some places
	parsedRet := new(resultWrapper)
	err = json.Unmarshal(ret.Bytes(), parsedRet)
	if err != nil {
		return parsedRet, err
	}
	parsedRet.Raw = ret.Bytes()

	return parsedRet, nil
}
