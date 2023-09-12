package gologin

type GologinRaw struct {
	UserAgent string
	Path      string
	Os        string
}

type Fingerprint struct {
	NewTabPage                                  NewTabPage          `json:"NewTabPage"`
	AccountIdMigrationState                     int                 `json:"account_id_migration_state"`
	AccountTrackerServiceLastUpdate             string              `json:"account_tracker_service_last_update"`
	ACKExistingNTPExtensions                    bool                `json:"ack_existing_ntp_extensions"`
	AlternateErrorPages                         AlternateErrorPages `json:"alternate_error_pages"`
	AnnouncementNotificationServiceFirstRunTime string              `json:"announcement_notification_service_first_run_time"`
	Apps                                        FingerprintApps     `json:"apps"`
	Autocomplete                                Autocomplete        `json:"autocomplete"`
	Autofill                                    Autofill            `json:"autofill"`
	BookmarkBar                                 BookmarkBar         `json:"bookmark_bar"`
	BookmarkEditor                              BookmarkEditor      `json:"bookmark_editor"`
	Bookmarks                                   Bookmarks           `json:"bookmarks"`
	Browser                                     Browser             `json:"browser"`
	CachedFonts                                 CachedFonts         `json:"cached_fonts"`
	CountryidAtInstall                          int64               `json:"countryid_at_install"`
	CredentialsEnableService                    bool                `json:"credentials_enable_service"`
	CustomLinks                                 CustomLinks         `json:"custom_links"`
	DefaultAppsInstallState                     int64               `json:"default_apps_install_state"`
	Devtools                                    Devtools            `json:"devtools"`
	DipsTimerLastUpdate                         string              `json:"dips_timer_last_update"`
	DomainDiversity                             DomainDiversity     `json:"domain_diversity"`
	Extensions                                  Extensions          `json:"extensions"`
	GaiaCookie                                  GaiaCookie          `json:"gaia_cookie"`
	Gcm                                         Gcm                 `json:"gcm"`
	Gologin                                     Gologin             `json:"gologin"`
	Google                                      Google              `json:"google"`
	History                                     History             `json:"history"`
	Intl                                        Intl                `json:"intl"`
	Invalidation                                Invalidation        `json:"invalidation"`
	Media                                       Media               `json:"media"`
	MediaRouter                                 MediaRouter         `json:"media_router"`
	NTP                                         NTP                 `json:"ntp"`
	OptimizationGuide                           OptimizationGuide   `json:"optimization_guide"`
	//Prefetch                                    Prefetch                       `json:"prefetch"`
	PrivacySandbox                 PrivacySandbox                 `json:"privacy_sandbox"`
	Profile                        Profile                        `json:"profile"`
	Safebrowsing                   Safebrowsing                   `json:"safebrowsing"`
	SegmentationPlatform           SegmentationPlatform           `json:"segmentation_platform"`
	Sessions                       Sessions                       `json:"sessions"`
	Settings                       FingerprintSettings            `json:"settings"`
	Signin                         Signin                         `json:"signin"`
	Spellcheck                     Spellcheck                     `json:"spellcheck"`
	SupervisedUser                 SupervisedUser                 `json:"supervised_user"`
	Sync                           Sync                           `json:"sync"`
	TranslateSiteBlacklist         []interface{}                  `json:"translate_site_blacklist"`
	TranslateSiteBlacklistWithTime TranslateSiteBlacklistWithTime `json:"translate_site_blacklist_with_time"`
	TranslateSiteBlocklistWithTime TranslateSiteBlacklistWithTime `json:"translate_site_blocklist_with_time"`
	UnifiedConsent                 UnifiedConsent                 `json:"unified_consent"`
	Updateclientdata               Updateclientdata               `json:"updateclientdata"`
	WebApps                        WebApps                        `json:"web_apps"`
	Webauthn                       Webauthn                       `json:"webauthn"`
	Zerosuggest                    Zerosuggest                    `json:"zerosuggest"`
}

type AlternateErrorPages struct {
	Backup bool `json:"backup"`
}

type FingerprintApps struct {
	ShortcutsArch    string `json:"shortcuts_arch"`
	ShortcutsVersion int64  `json:"shortcuts_version"`
}

type Autocomplete struct {
	RetentionPolicyLastVersion int64 `json:"retention_policy_last_version"`
}

type Autofill struct {
	OrphanRowsRemoved bool `json:"orphan_rows_removed"`
}

type BookmarkBar struct {
	ShowOnAllTabs bool `json:"show_on_all_tabs"`
}

type BookmarkEditor struct {
	ExpandedNodes []interface{} `json:"expanded_nodes"`
}

type Bookmarks struct {
	EditingEnabled bool `json:"editing_enabled"`
}

type Browser struct {
	EnableSpellchecking bool            `json:"enable_spellchecking"`
	HasSeenWelcomePage  bool            `json:"has_seen_welcome_page"`
	WindowPlacement     WindowPlacement `json:"window_placement"`
}

type WindowPlacement struct {
	Bottom         int64 `json:"bottom"`
	Left           int64 `json:"left"`
	Maximized      bool  `json:"maximized"`
	Right          int64 `json:"right"`
	Top            int64 `json:"top"`
	WorkAreaBottom int64 `json:"work_area_bottom"`
	WorkAreaLeft   int64 `json:"work_area_left"`
	WorkAreaRight  int64 `json:"work_area_right"`
	WorkAreaTop    int64 `json:"work_area_top"`
}

type CachedFonts struct {
	SearchResultsPage SearchResultsPage `json:"search_results_page"`
}

type SearchResultsPage struct {
	Fallback []interface{} `json:"fallback"`
	Primary  []string      `json:"primary"`
}

type CustomLinks struct {
	Initialized bool   `json:"initialized"`
	List        []List `json:"list"`
}

type List struct {
	IsMostVisited bool   `json:"isMostVisited"`
	Title         string `json:"title"`
	URL           string `json:"url"`
}

type Devtools struct {
	ADBKey                        string                        `json:"adb_key"`
	Preferences                   Preferences                   `json:"preferences"`
	SyncedPreferencesSyncDisabled SyncedPreferencesSyncDisabled `json:"synced_preferences_sync_disabled"`
}

type Preferences struct {
	InspectorDrawerSplitViewState              string `json:"Inspector.drawerSplitViewState"`
	InspectorViewSplitViewState                string `json:"InspectorView.splitViewState"`
	StylesPaneSidebarTabOrder                  string `json:"Styles-pane-sidebar-tabOrder"`
	CloseableTabs                              string `json:"closeableTabs"`
	ConsoleSidebarWidth                        string `json:"console.sidebar.width"`
	ConsoleSidebarSelectedFilter               string `json:"console.sidebarSelectedFilter"`
	CurrentDockState                           string `json:"currentDockState"`
	DataGridCookiesTableColumnWeights          string `json:"dataGrid-cookiesTable-columnWeights"`
	DrawerViewTabOrder                         string `json:"drawer-view-tabOrder"`
	ElementsStylesSidebarWidth                 string `json:"elements.styles.sidebar.width"`
	InspectorVersion                           string `json:"inspectorVersion"`
	NetworkBlockedPatterns                     string `json:"networkBlockedPatterns"`
	NetworkPanelSidebarState                   string `json:"networkPanelSidebarState"`
	NetworkPanelSplitViewState                 string `json:"networkPanelSplitViewState"`
	NetworkPanelSplitViewWaterfall             string `json:"networkPanelSplitViewWaterfall"`
	NetworkResourceTypeFilters                 string `json:"networkResourceTypeFilters"`
	NetworkTextFilter                          string `json:"networkTextFilter"`
	PanelSelectedTab                           string `json:"panel-selectedTab"`
	RequestInfoFormDataCategoryExpanded        string `json:"request-info-formData-category-expanded"`
	RequestInfoGeneralCategoryExpanded         string `json:"request-info-general-category-expanded"`
	RequestInfoQueryStringCategoryExpanded     string `json:"request-info-queryString-category-expanded"`
	RequestInfoRequestHeadersCategoryExpanded  string `json:"request-info-requestHeaders-category-expanded"`
	RequestInfoRequestPayloadCategoryExpanded  string `json:"request-info-requestPayload-category-expanded"`
	RequestInfoResponseHeadersCategoryExpanded string `json:"request-info-responseHeaders-category-expanded"`
	ResourceViewTab                            string `json:"resourceViewTab"`
	ResourcesCookiesExpanded                   string `json:"resourcesCookiesExpanded"`
	ResourcesLastSelectedElementPath           string `json:"resourcesLastSelectedElementPath"`
	ResourcesLocalStorageExpanded              string `json:"resourcesLocalStorageExpanded"`
	SourcesPanelNavigatorSplitViewState        string `json:"sourcesPanelNavigatorSplitViewState"`
	SourcesPanelSplitViewState                 string `json:"sourcesPanelSplitViewState"`
	UndefinedTabOrder                          string `json:"undefined-tabOrder"`
}

type SyncedPreferencesSyncDisabled struct {
	AdornerSettings string `json:"adornerSettings"`
	ColorFormat     string `json:"colorFormat"`
}

type DomainDiversity struct {
	LastReportingTimestamp string `json:"last_reporting_timestamp"`
}

type Extensions struct {
	Alerts             Alerts                         `json:"alerts"`
	ChromeURLOverrides TranslateSiteBlacklistWithTime `json:"chrome_url_overrides"`
	Commands           map[string]Command             `json:"commands"`
	LastChromeVersion  string                         `json:"last_chrome_version"`
	Settings           ExtensionsSettings             `json:"settings"`
}

type Alerts struct {
	Initialized bool `json:"initialized"`
}

type TranslateSiteBlacklistWithTime struct {
}

type Command struct {
	CommandName string `json:"command_name"`
	Extension   string `json:"extension"`
	Global      bool   `json:"global"`
}

type ExtensionsSettings struct {
	Ahfgeienlihckogmohjhadlkjgocpleb Ahfgeienlihckogmohjhadlkjgocpleb `json:"ahfgeienlihckogmohjhadlkjgocpleb"`
	Bkbngmckglogipmocjcncnhbikofbjbm Ahfgeienlihckogmohjhadlkjgocpleb `json:"bkbngmckglogipmocjcncnhbikofbjbm"`
	Lmccpkjihdfknjbpkilmbinhljdkcbak Lmccpkjihdfknjbpkilmbinhljdkcbak `json:"lmccpkjihdfknjbpkilmbinhljdkcbak"`
	Mhjfbmdgcfjbbpaeojofohoefgiehjai Ahfgeienlihckogmohjhadlkjgocpleb `json:"mhjfbmdgcfjbbpaeojofohoefgiehjai"`
	Neajdppkdcdipfabeoofebfddakdcjhd Ahfgeienlihckogmohjhadlkjgocpleb `json:"neajdppkdcdipfabeoofebfddakdcjhd"`
	Nkeimhogjdpnpccoofpliimaahmaaome Ahfgeienlihckogmohjhadlkjgocpleb `json:"nkeimhogjdpnpccoofpliimaahmaaome"`
}

type Ahfgeienlihckogmohjhadlkjgocpleb struct {
	ActivePermissions             Permissions                    `json:"active_permissions"`
	AppLauncherOrdinal            *string                        `json:"app_launcher_ordinal,omitempty"`
	Commands                      TranslateSiteBlacklistWithTime `json:"commands"`
	ContentSettings               []interface{}                  `json:"content_settings"`
	CreationFlags                 int64                          `json:"creation_flags"`
	Events                        []string                       `json:"events,omitempty"`
	FirstInstallTime              string                         `json:"first_install_time"`
	FromWebstore                  bool                           `json:"from_webstore"`
	IncognitoContentSettings      []interface{}                  `json:"incognito_content_settings"`
	IncognitoPreferences          TranslateSiteBlacklistWithTime `json:"incognito_preferences"`
	LastUpdateTime                string                         `json:"last_update_time"`
	Location                      int64                          `json:"location"`
	Manifest                      *Manifest                      `json:"manifest,omitempty"`
	NeedsSync                     *bool                          `json:"needs_sync,omitempty"`
	PageOrdinal                   *string                        `json:"page_ordinal,omitempty"`
	Path                          string                         `json:"path"`
	Preferences                   TranslateSiteBlacklistWithTime `json:"preferences"`
	RegularOnlyPreferences        TranslateSiteBlacklistWithTime `json:"regular_only_preferences"`
	State                         int64                          `json:"state"`
	WasInstalledByDefault         bool                           `json:"was_installed_by_default"`
	WasInstalledByOEM             bool                           `json:"was_installed_by_oem"`
	GrantedPermissions            *Permissions                   `json:"granted_permissions,omitempty"`
	NewAllowFileAccess            *bool                          `json:"newAllowFileAccess,omitempty"`
	ServiceWorkerRegistrationInfo *ServiceWorkerRegistrationInfo `json:"service_worker_registration_info,omitempty"`
	Serviceworkerevents           []string                       `json:"serviceworkerevents,omitempty"`
	WithholdingPermissions        *bool                          `json:"withholding_permissions,omitempty"`
}

type Permissions struct {
	API                 []string      `json:"api"`
	ExplicitHost        []string      `json:"explicit_host"`
	ManifestPermissions []interface{} `json:"manifest_permissions"`
	ScriptableHost      []string      `json:"scriptable_host"`
}

type Manifest struct {
	App                   *App                   `json:"app,omitempty"`
	Description           *string                `json:"description,omitempty"`
	Icons                 map[string]string      `json:"icons,omitempty"`
	Key                   string                 `json:"key"`
	Name                  string                 `json:"name"`
	Permissions           []PermissionElement    `json:"permissions"`
	Version               string                 `json:"version"`
	ContentSecurityPolicy *string                `json:"content_security_policy,omitempty"`
	Incognito             *string                `json:"incognito,omitempty"`
	ManifestVersion       *int64                 `json:"manifest_version,omitempty"`
	MIMETypes             []string               `json:"mime_types,omitempty"`
	MIMETypesHandler      *string                `json:"mime_types_handler,omitempty"`
	OfflineEnabled        *bool                  `json:"offline_enabled,omitempty"`
	Background            *Background            `json:"background,omitempty"`
	TTSEngine             *TTSEngine             `json:"tts_engine,omitempty"`
	ExternallyConnectable *ExternallyConnectable `json:"externally_connectable,omitempty"`
}

type App struct {
	Launch Launch   `json:"launch"`
	Urls   []string `json:"urls"`
}

type Launch struct {
	WebURL string `json:"web_url"`
}

type Background struct {
	Persistent bool     `json:"persistent"`
	Scripts    []string `json:"scripts,omitempty"`
	Page       *string  `json:"page,omitempty"`
}

type ExternallyConnectable struct {
	Matches []string `json:"matches"`
}

type PermissionClass struct {
	FileSystem []string `json:"fileSystem"`
}

type TTSEngine struct {
	Voices []Voice `json:"voices"`
}

type Voice struct {
	EventTypes []EventType `json:"event_types"`
	Gender     Gender      `json:"gender"`
	Lang       string      `json:"lang"`
	Remote     bool        `json:"remote"`
	VoiceName  string      `json:"voice_name"`
}

type ServiceWorkerRegistrationInfo struct {
	Version string `json:"version"`
}

type Lmccpkjihdfknjbpkilmbinhljdkcbak struct {
	ActivePermissions             Permissions                    `json:"active_permissions"`
	Commands                      Commands                       `json:"commands"`
	ContentSettings               []interface{}                  `json:"content_settings"`
	CreationFlags                 int64                          `json:"creation_flags"`
	FilteredServiceWorkerEvents   FilteredServiceWorkerEvents    `json:"filtered_service_worker_events"`
	FirstInstallTime              string                         `json:"first_install_time"`
	FromWebstore                  bool                           `json:"from_webstore"`
	GrantedPermissions            Permissions                    `json:"granted_permissions"`
	IncognitoContentSettings      []interface{}                  `json:"incognito_content_settings"`
	IncognitoPreferences          TranslateSiteBlacklistWithTime `json:"incognito_preferences"`
	LastUpdateTime                string                         `json:"last_update_time"`
	Location                      int64                          `json:"location"`
	NewAllowFileAccess            bool                           `json:"newAllowFileAccess"`
	Path                          string                         `json:"path"`
	Preferences                   TranslateSiteBlacklistWithTime `json:"preferences"`
	RegularOnlyPreferences        TranslateSiteBlacklistWithTime `json:"regular_only_preferences"`
	ServiceWorkerRegistrationInfo ServiceWorkerRegistrationInfo  `json:"service_worker_registration_info"`
	Serviceworkerevents           []string                       `json:"serviceworkerevents"`
	State                         int64                          `json:"state"`
	WasInstalledByDefault         bool                           `json:"was_installed_by_default"`
	WasInstalledByOEM             bool                           `json:"was_installed_by_oem"`
	WithholdingPermissions        bool                           `json:"withholding_permissions"`
}

type Commands struct {
	AutomationTask AutomationTask `json:"automationTask"`
	HumanTyping    AutomationTask `json:"humanTyping"`
}

type AutomationTask struct {
	SuggestedKey string `json:"suggested_key"`
}

type FilteredServiceWorkerEvents struct {
	WindowsOnRemoved []TranslateSiteBlacklistWithTime `json:"windows.onRemoved"`
}

type GaiaCookie struct {
	ChangedTime          float64 `json:"changed_time"`
	Hash                 string  `json:"hash"`
	LastListAccountsData string  `json:"last_list_accounts_data"`
}

type Gcm struct {
	ProductCategoryForSubtypes string `json:"product_category_for_subtypes"`
}

type Gologin struct {
	AudioContext            AudioContext `json:"audioContext"`
	CanvasMode              string       `json:"canvasMode"`
	CanvasNoise             float64      `json:"canvasNoise"`
	ClientRectsNoiseEnable  bool         `json:"client_rects_noise_enable"`
	DeviceMemory            int64        `json:"deviceMemory"`
	DNS                     string       `json:"dns"`
	DoNotTrack              bool         `json:"doNotTrack"`
	GeoLocation             GeoLocation  `json:"geoLocation"`
	GetClientRectsNoice     float64      `json:"getClientRectsNoice"`
	GetClientRectsNoise     float64      `json:"get_client_rects_noise"`
	HardwareConcurrency     int64        `json:"hardwareConcurrency"`
	Icon                    Icon         `json:"icon"`
	IsM1                    bool         `json:"is_m1"`
	LangHeader              string       `json:"langHeader"`
	Languages               string       `json:"languages"`
	MediaDevices            MediaDevices `json:"mediaDevices"`
	Mobile                  Mobile       `json:"mobile"`
	Name                    string       `json:"name"`
	Navigator               Navigator    `json:"navigator"`
	Plugins                 Plugins      `json:"plugins"`
	ProfileID               string       `json:"profile_id"`
	Proxy                   Proxy        `json:"proxy"`
	ScreenHeight            int64        `json:"screenHeight"`
	ScreenWidth             int64        `json:"screenWidth"`
	StartupURL              string       `json:"startupUrl"`
	StartupUrls             []string     `json:"startup_urls"`
	Storage                 Storage      `json:"storage"`
	Timezone                Timezone     `json:"timezone"`
	UnpinableExtensionNames []string     `json:"unpinable_extension_names"`
	UserAgent               string       `json:"userAgent"`
	WebGl                   WebGl        `json:"webGl"`
	WebRTC                  WebRTC       `json:"webRtc"`
	Webgl                   Webgl        `json:"webgl"`
	WebglNoiceEnable        bool         `json:"webglNoiceEnable"`
	WebglNoiseValue         float64      `json:"webglNoiseValue"`
	WebglParams             WebglParams  `json:"webglParams"`
	GologinWebglNoiceEnable bool         `json:"webgl_noice_enable"`
	WebglNoiseEnable        bool         `json:"webgl_noise_enable"`
	GologinWebglNoiseValue  float64      `json:"webgl_noise_value"`
}

type AudioContext struct {
	Enable     bool    `json:"enable"`
	NoiseValue float64 `json:"noiseValue"`
}

type GeoLocation struct {
	Accuracy  int64   `json:"accuracy"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Mode      string  `json:"mode"`
}

type Icon struct {
	Avatar Avatar `json:"avatar"`
	Text   string `json:"text"`
}

type Avatar struct {
	Enabled  bool `json:"enabled"`
	FullSize bool `json:"full_size"`
}

type MediaDevices struct {
	AudioInputs  int64  `json:"audioInputs"`
	AudioOutputs int64  `json:"audioOutputs"`
	Enable       bool   `json:"enable"`
	Uid          string `json:"uid"`
	VideoInputs  int64  `json:"videoInputs"`
}

type Mobile struct {
	DeviceScaleFactor float64 `json:"device_scale_factor"`
	Enable            bool    `json:"enable"`
	Height            int64   `json:"height"`
	Width             int64   `json:"width"`
}

type Navigator struct {
	MaxTouchPoints int64  `json:"max_touch_points"`
	Platform       string `json:"platform"`
}

type Plugins struct {
	AllEnable   bool `json:"all_enable"`
	FlashEnable bool `json:"flash_enable"`
}

type Proxy struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Storage struct {
	Enable bool `json:"enable"`
}

type Timezone struct {
	ID string `json:"id"`
}

type WebGl struct {
	Mode     bool   `json:"mode"`
	Renderer string `json:"renderer"`
	Vendor   string `json:"vendor"`
}

type WebRTC struct {
	FillBasedOnIP  bool   `json:"fill_based_on_ip"`
	LocalIPMasking bool   `json:"local_ip_masking"`
	Mode           string `json:"mode"`
	PublicIP       string `json:"public_ip"`
}

type Webgl struct {
	Metadata WebGl `json:"metadata"`
}

type WebglParams struct {
	Antialiasing            bool                `json:"antialiasing"`
	Extensions              []string            `json:"extensions"`
	GlCanvas                string              `json:"glCanvas"`
	GlParamValues           []GlParamValue      `json:"glParamValues"`
	ShaiderPrecisionFormat  string              `json:"shaiderPrecisionFormat"`
	SupportedFunctions      []SupportedFunction `json:"supportedFunctions"`
	TextureMaxAnisotropyEXT int64               `json:"textureMaxAnisotropyExt"`
}

type SupportedFunction struct {
	Name      string `json:"name"`
	Supported bool   `json:"supported"`
}

type Google struct {
	Services Services `json:"services"`
}

type Services struct {
	ConsentedToSync      bool   `json:"consented_to_sync"`
	SigninScopedDeviceID string `json:"signin_scoped_device_id"`
}

type History struct {
	SavingDisabled bool `json:"saving_disabled"`
}

type Intl struct {
	AcceptLanguages   string `json:"accept_languages"`
	SelectedLanguages string `json:"selected_languages"`
}

type Invalidation struct {
	PerSenderTopicsToHandler map[string]TranslateSiteBlacklistWithTime `json:"per_sender_topics_to_handler"`
}

type Media struct {
	DeviceIDSalt string     `json:"device_id_salt"`
	Engagement   Engagement `json:"engagement"`
}

type Engagement struct {
	SchemaVersion int64 `json:"schema_version"`
}

type MediaRouter struct {
	ReceiverIDHashToken string `json:"receiver_id_hash_token"`
}

type NTP struct {
	NumPersonalSuggestions int64 `json:"num_personal_suggestions"`
}

type NewTabPage struct {
	PrevNavigationTime string `json:"PrevNavigationTime"`
}

type OptimizationGuide struct {
	Hintsfetcher                          Fetcher                               `json:"hintsfetcher"`
	Predictionmodelfetcher                Fetcher                               `json:"predictionmodelfetcher"`
	PreviouslyRegisteredOptimizationTypes PreviouslyRegisteredOptimizationTypes `json:"previously_registered_optimization_types"`
	StoreFilePathsToDelete                TranslateSiteBlacklistWithTime        `json:"store_file_paths_to_delete"`
}

type Fetcher struct {
	LastFetchAttempt string `json:"last_fetch_attempt"`
}

type PreviouslyRegisteredOptimizationTypes struct {
	AboutThisSite   bool `json:"ABOUT_THIS_SITE"`
	HistoryClusters bool `json:"HISTORY_CLUSTERS"`
}

type Prefetch struct {
	SearchPrefetch SearchPrefetch `json:"search_prefetch"`
}

type SearchPrefetch struct {
	Cache Cache `json:"cache"`
}

type Cache struct {
	HTTPSWWWGoogleCOMSearchQCreepjsOqCreepjsAqsChrome69I57J0I19I512L2J0I10I19I512L2J0I19I512J46I19I512J0I19I512J46I10I19I512L2991J0J7SourceidChromeIeUTF8 []string `json:"https://www.google.com/search?q=creepjs&oq=creepjs+&aqs=chrome..69i57j0i19i512l2j0i10i19i512l2j0i19i512j46i19i512j0i19i512j46i10i19i512l2.991j0j7&sourceid=chrome&ie=UTF-8"`
}

type PrivacySandbox struct {
	AntiAbuseInitialized bool `json:"anti_abuse_initialized"`
}

type Profile struct {
	AvatarIndex                            int64                          `json:"avatar_index"`
	ContentSettings                        ContentSettings                `json:"content_settings"`
	CreatedByVersion                       string                         `json:"created_by_version"`
	CreationTime                           string                         `json:"creation_time"`
	ExitType                               string                         `json:"exit_type"`
	LastEngagementTime                     string                         `json:"last_engagement_time"`
	LastTimeObsoleteHTTPCredentialsRemoved float64                        `json:"last_time_obsolete_http_credentials_removed"`
	LastTimePasswordStoreMetricsReported   float64                        `json:"last_time_password_store_metrics_reported"`
	ManagedUserID                          string                         `json:"managed_user_id"`
	Name                                   string                         `json:"name"`
	PasswordAccountStorageSettings         TranslateSiteBlacklistWithTime `json:"password_account_storage_settings"`
	WereOldGoogleLoginsRemoved             bool                           `json:"were_old_google_logins_removed"`
}

type ContentSettings struct {
	EnableQuietPermissionUIEnablingMethod EnableQuietPermissionUIEnablingMethod `json:"enable_quiet_permission_ui_enabling_method"`
	Exceptions                            Exceptions                            `json:"exceptions"`
	PrefVersion                           int64                                 `json:"pref_version"`
}

type EnableQuietPermissionUIEnablingMethod struct {
	Notifications int64 `json:"notifications"`
}

type Exceptions struct {
	AccessibilityEvents                TranslateSiteBlacklistWithTime `json:"accessibility_events"`
	AntiAbuse                          TranslateSiteBlacklistWithTime `json:"anti_abuse"`
	AppBanner                          TranslateSiteBlacklistWithTime `json:"app_banner"`
	Ar                                 TranslateSiteBlacklistWithTime `json:"ar"`
	AutoSelectCertificate              TranslateSiteBlacklistWithTime `json:"auto_select_certificate"`
	AutomaticDownloads                 TranslateSiteBlacklistWithTime `json:"automatic_downloads"`
	Autoplay                           TranslateSiteBlacklistWithTime `json:"autoplay"`
	BackgroundSync                     TranslateSiteBlacklistWithTime `json:"background_sync"`
	BluetoothChooserData               TranslateSiteBlacklistWithTime `json:"bluetooth_chooser_data"`
	BluetoothGuard                     TranslateSiteBlacklistWithTime `json:"bluetooth_guard"`
	BluetoothScanning                  TranslateSiteBlacklistWithTime `json:"bluetooth_scanning"`
	CameraPanTiltZoom                  TranslateSiteBlacklistWithTime `json:"camera_pan_tilt_zoom"`
	ClientHints                        ClientHints                    `json:"client_hints"`
	Clipboard                          TranslateSiteBlacklistWithTime `json:"clipboard"`
	Cookies                            TranslateSiteBlacklistWithTime `json:"cookies"`
	DurableStorage                     TranslateSiteBlacklistWithTime `json:"durable_storage"`
	FedcmActiveSession                 TranslateSiteBlacklistWithTime `json:"fedcm_active_session"`
	FedcmIdpRegistration               TranslateSiteBlacklistWithTime `json:"fedcm_idp_registration"`
	FedcmIdpSignin                     TranslateSiteBlacklistWithTime `json:"fedcm_idp_signin"`
	FedcmShare                         TranslateSiteBlacklistWithTime `json:"fedcm_share"`
	FileSystemAccessChooserData        TranslateSiteBlacklistWithTime `json:"file_system_access_chooser_data"`
	FileSystemLastPickedDirectory      TranslateSiteBlacklistWithTime `json:"file_system_last_picked_directory"`
	FileSystemReadGuard                TranslateSiteBlacklistWithTime `json:"file_system_read_guard"`
	FileSystemWriteGuard               TranslateSiteBlacklistWithTime `json:"file_system_write_guard"`
	FormfillMetadata                   TranslateSiteBlacklistWithTime `json:"formfill_metadata"`
	Geolocation                        TranslateSiteBlacklistWithTime `json:"geolocation"`
	GetDisplayMediaSetSelectAllScreens TranslateSiteBlacklistWithTime `json:"get_display_media_set_select_all_screens"`
	HIDChooserData                     TranslateSiteBlacklistWithTime `json:"hid_chooser_data"`
	HIDGuard                           TranslateSiteBlacklistWithTime `json:"hid_guard"`
	HTTPAllowed                        TranslateSiteBlacklistWithTime `json:"http_allowed"`
	IdleDetection                      TranslateSiteBlacklistWithTime `json:"idle_detection"`
	Images                             TranslateSiteBlacklistWithTime `json:"images"`
	ImportantSiteInfo                  TranslateSiteBlacklistWithTime `json:"important_site_info"`
	InsecurePrivateNetwork             TranslateSiteBlacklistWithTime `json:"insecure_private_network"`
	IntentPickerAutoDisplay            TranslateSiteBlacklistWithTime `json:"intent_picker_auto_display"`
	Javascript                         TranslateSiteBlacklistWithTime `json:"javascript"`
	JavascriptJIT                      TranslateSiteBlacklistWithTime `json:"javascript_jit"`
	LegacyCookieAccess                 TranslateSiteBlacklistWithTime `json:"legacy_cookie_access"`
	LocalFonts                         TranslateSiteBlacklistWithTime `json:"local_fonts"`
	MediaEngagement                    map[string]MediaEngagement     `json:"media_engagement"`
	MediaStreamCamera                  TranslateSiteBlacklistWithTime `json:"media_stream_camera"`
	MediaStreamMic                     TranslateSiteBlacklistWithTime `json:"media_stream_mic"`
	MIDISysex                          TranslateSiteBlacklistWithTime `json:"midi_sysex"`
	MixedScript                        TranslateSiteBlacklistWithTime `json:"mixed_script"`
	NFCDevices                         TranslateSiteBlacklistWithTime `json:"nfc_devices"`
	NotificationInteractions           TranslateSiteBlacklistWithTime `json:"notification_interactions"`
	NotificationPermissionReview       TranslateSiteBlacklistWithTime `json:"notification_permission_review"`
	Notifications                      TranslateSiteBlacklistWithTime `json:"notifications"`
	PasswordProtection                 TranslateSiteBlacklistWithTime `json:"password_protection"`
	PaymentHandler                     TranslateSiteBlacklistWithTime `json:"payment_handler"`
	PermissionAutoblockingData         TranslateSiteBlacklistWithTime `json:"permission_autoblocking_data"`
	PermissionAutorevocationData       TranslateSiteBlacklistWithTime `json:"permission_autorevocation_data"`
	Popups                             TranslateSiteBlacklistWithTime `json:"popups"`
	PrivateNetworkChooserData          TranslateSiteBlacklistWithTime `json:"private_network_chooser_data"`
	PrivateNetworkGuard                TranslateSiteBlacklistWithTime `json:"private_network_guard"`
	ProtectedMediaIdentifier           TranslateSiteBlacklistWithTime `json:"protected_media_identifier"`
	ProtocolHandler                    TranslateSiteBlacklistWithTime `json:"protocol_handler"`
	ReducedAcceptLanguage              TranslateSiteBlacklistWithTime `json:"reduced_accept_language"`
	SafeBrowsingURLCheckData           TranslateSiteBlacklistWithTime `json:"safe_browsing_url_check_data"`
	Sensors                            TranslateSiteBlacklistWithTime `json:"sensors"`
	SerialChooserData                  TranslateSiteBlacklistWithTime `json:"serial_chooser_data"`
	SerialGuard                        TranslateSiteBlacklistWithTime `json:"serial_guard"`
	SiteEngagement                     map[string]SiteEngagement      `json:"site_engagement"`
	Sound                              TranslateSiteBlacklistWithTime `json:"sound"`
	SSLCERTDecisions                   TranslateSiteBlacklistWithTime `json:"ssl_cert_decisions"`
	StorageAccess                      TranslateSiteBlacklistWithTime `json:"storage_access"`
	SubresourceFilter                  TranslateSiteBlacklistWithTime `json:"subresource_filter"`
	SubresourceFilterData              TranslateSiteBlacklistWithTime `json:"subresource_filter_data"`
	ThirdPartyStoragePartitioning      TranslateSiteBlacklistWithTime `json:"third_party_storage_partitioning"`
	TopLevelStorageAccess              TranslateSiteBlacklistWithTime `json:"top_level_storage_access"`
	UnusedSitePermissions              TranslateSiteBlacklistWithTime `json:"unused_site_permissions"`
	USBChooserData                     TranslateSiteBlacklistWithTime `json:"usb_chooser_data"`
	USBGuard                           TranslateSiteBlacklistWithTime `json:"usb_guard"`
	VR                                 TranslateSiteBlacklistWithTime `json:"vr"`
	WebidAPI                           TranslateSiteBlacklistWithTime `json:"webid_api"`
	WebidAutoReauthn                   TranslateSiteBlacklistWithTime `json:"webid_auto_reauthn"`
	WindowPlacement                    TranslateSiteBlacklistWithTime `json:"window_placement"`
}

type ClientHints struct {
	HTTPSWWWGoogleCOM443 HTTPSWWWGoogleCOM443 `json:"https://www.google.com:443,*"`
}

type HTTPSWWWGoogleCOM443 struct {
	LastModified string                       `json:"last_modified"`
	Setting      HTTPSWWWGoogleCOM443_Setting `json:"setting"`
}

type HTTPSWWWGoogleCOM443_Setting struct {
	ClientHints []int64 `json:"client_hints"`
}

type MediaEngagement struct {
	Expiration   string                 `json:"expiration"`
	LastModified string                 `json:"last_modified"`
	Setting      MediaEngagementSetting `json:"setting"`
}

type MediaEngagementSetting struct {
	HasHighScore          bool    `json:"hasHighScore"`
	LastMediaPlaybackTime float64 `json:"lastMediaPlaybackTime"`
	MediaPlaybacks        int64   `json:"mediaPlaybacks"`
	Visits                int64   `json:"visits"`
}

type SiteEngagement struct {
	LastModified string                `json:"last_modified"`
	Setting      SiteEngagementSetting `json:"setting"`
}

type SiteEngagementSetting struct {
	LastEngagementTime     float64 `json:"lastEngagementTime"`
	LastShortcutLaunchTime float64 `json:"lastShortcutLaunchTime"`
	PointsAddedToday       float64 `json:"pointsAddedToday"`
	RawScore               float64 `json:"rawScore"`
}

type Safebrowsing struct {
	EventTimestamps    TranslateSiteBlacklistWithTime `json:"event_timestamps"`
	MetricsLastLogTime string                         `json:"metrics_last_log_time"`
}

type SegmentationPlatform struct {
	ClientResultPrefs    string             `json:"client_result_prefs"`
	DeviceSwitcherUtil   DeviceSwitcherUtil `json:"device_switcher_util"`
	LastDBCompactionTime string             `json:"last_db_compaction_time"`
	SegmentationResult   SegmentationResult `json:"segmentation_result"`
}

type DeviceSwitcherUtil struct {
	Result Result `json:"result"`
}

type Result struct {
	Labels []string `json:"labels"`
}

type SegmentationResult struct {
	CrossDeviceUser User `json:"cross_device_user"`
	ShoppingUser    User `json:"shopping_user"`
}

type User struct {
	InUse         bool   `json:"in_use"`
	SegmentID     int64  `json:"segment_id"`
	SegmentRank   int64  `json:"segment_rank"`
	SelectionTime string `json:"selection_time"`
}

type Sessions struct {
	EventLog          []EventLog `json:"event_log"`
	SessionDataStatus int64      `json:"session_data_status"`
}

type EventLog struct {
	Crashed             *bool  `json:"crashed,omitempty"`
	Time                string `json:"time"`
	Type                int64  `json:"type"`
	RestoreBrowser      *bool  `json:"restore_browser,omitempty"`
	Synchronous         *bool  `json:"synchronous,omitempty"`
	ErroredReading      *bool  `json:"errored_reading,omitempty"`
	TabCount            *int64 `json:"tab_count,omitempty"`
	WindowCount         *int64 `json:"window_count,omitempty"`
	DidScheduleCommand  *bool  `json:"did_schedule_command,omitempty"`
	FirstSessionService *bool  `json:"first_session_service,omitempty"`
}

type FingerprintSettings struct {
	A11Y A11Y `json:"a11y"`
}

type A11Y struct {
	ApplyPageColorsOnlyOnIncreasedContrast bool `json:"apply_page_colors_only_on_increased_contrast"`
}

type Signin struct {
	Allowed bool `json:"allowed"`
}

type Spellcheck struct {
	Dictionaries       []interface{} `json:"dictionaries"`
	Dictionary         string        `json:"dictionary"`
	UseSpellingService bool          `json:"use_spelling_service"`
}

type SupervisedUser struct {
	Metrics Metrics `json:"metrics"`
}

type Metrics struct {
	DayID int64 `json:"day_id"`
}

type Sync struct {
	Requested bool `json:"requested"`
}

type UnifiedConsent struct {
	MigrationState int64 `json:"migration_state"`
}

type Updateclientdata struct {
	Apps UpdateclientdataApps `json:"apps"`
}

type UpdateclientdataApps struct {
	Ihcjicgdanjaechkgeegckofjjedodee Ihcjicgdanjaechkgeegckofjjedodee `json:"ihcjicgdanjaechkgeegckofjjedodee"`
}

type Ihcjicgdanjaechkgeegckofjjedodee struct {
	Cohort      string `json:"cohort"`
	Cohortname  string `json:"cohortname"`
	Dlrc        int64  `json:"dlrc"`
	Installdate int64  `json:"installdate"`
	Pf          string `json:"pf"`
}

type WebApps struct {
	DidMigrateDefaultChromeApps       []string `json:"did_migrate_default_chrome_apps"`
	LastPreinstallSynchronizeVersion  string   `json:"last_preinstall_synchronize_version"`
	SystemWebAppFailureCount          int64    `json:"system_web_app_failure_count"`
	SystemWebAppLastAttemptedLanguage string   `json:"system_web_app_last_attempted_language"`
	SystemWebAppLastAttemptedUpdate   string   `json:"system_web_app_last_attempted_update"`
	SystemWebAppLastInstalledLanguage string   `json:"system_web_app_last_installed_language"`
	SystemWebAppLastUpdate            string   `json:"system_web_app_last_update"`
}

type Webauthn struct {
	Touchid Touchid `json:"touchid"`
}

type Touchid struct {
	MetadataSecret string `json:"metadata_secret"`
}

type Zerosuggest struct {
	Cachedresults string `json:"cachedresults"`
}

type EventType string

const (
	End   EventType = "end"
	Error EventType = "error"
	Start EventType = "start"
)

type Gender string

const (
	Female Gender = "female"
	Male   Gender = "male"
)

type PermissionElement struct {
	PermissionClass *PermissionClass
	String          *string
}

type GlParamValue struct {
	Name  any `json:"name"`
	Value any `json:"value"`
}
