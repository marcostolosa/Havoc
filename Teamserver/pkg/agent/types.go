package agent

import (
    "Havoc/pkg/common/parser"
    "Havoc/pkg/packager"
)

const (
    typeGood  = "Good"
    typeInfo  = "Info"
    typeError = "Error"
)

type DemonInterface interface {
    // TODO: move everything from RoutineFunc here
}

type EventInterface interface {
}

type AgentHeader struct {
    Size       int
    MagicValue int
    AgentID    int
    Data       *parser.Parser
}

type ServiceAgentInterface interface {
    SendResponse(AgentInfo any, Header AgentHeader) []byte
    SendAgentBuildRequest(ClientID string, Config map[string]any, Listener map[string]any)
}

type RoutineFunc struct {
    CallbackSize func(DemonInstance *Agent, i int)
    CallbackTime func(DemonInstance *Agent)

    AgentGetInstance func(DemonID int) *Agent

    EventAppend        func(event packager.Package) []packager.Package
    EventBroadcast     func(ExceptClient string, pk packager.Package)
    EventNewDemon      func(DemonAgent *Agent) packager.Package
    EventAgentMark     func(AgentID, Mark string)
    EventListenerError func(ListenerName string, Error error)

    AppendListener func(FromUser string, Type int, Config any) packager.Package

    DemonCallback func(Command int, data []byte)
    AppendDemon   func(demon *Agent) []*Agent
    AgentExists   func(DemonID int) bool
    DemonOutput   func(DemonID string, CommandID int, Output map[string]string)
    AgentCallback func(DemonID string, Time string)

    ServiceAgentExits func(MagicValue int) bool
    ServiceAgentGet   func(MagicValue int) ServiceAgentInterface
}

type DemonJob struct {
    Command     uint32
    TaskID      uint32
    Data        []interface{}
    Payload     []byte
    Description string
}

type Agent struct {
    NameID     string
    JobQueue   []DemonJob
    SessionDir string
    Active     bool

    Encryption struct {
        AESKey []byte
        AESIv  []byte
    }

    Info   *AgentInfo
    Pivots struct {
        Parent *Agent
        Links  []*Agent
    }
}

type AgentInfo struct {
    // Connection Info
    Listener   any
    MagicValue int
    SleepDelay int

    // OS Info
    OSVersion string
    OSArch    string
    OSBuild   string

    InternalIP string
    ExternalIP string
    Hostname   string
    DomainName string
    Username   string

    // User Info
    Elevated string

    // Process Info
    ProcessArch string
    ProcessName string
    ProcessPID  int
    ProcessPPID int
    ProcessPath string

    // Call home from Demon
    FirstCallIn string
    LastCallIn  string
}

type Agents struct {
    Agents []*Agent
}

var InjectErrors = map[int]string{
    0x1001: "trying to inject a x64 payload into a x86 process",
    0x1002: "trying to inject a x86 payload into a x64 process",
}

var Win32ErrorCodes = map[int]string{
    1:    "ERROR_INVALID_FUNCTION",
    2:    "ERROR_FILE_NOT_FOUND",
    3:    "ERROR_PATH_NOT_FOUND",
    4:    "ERROR_TOO_MANY_OPEN_FILES",
    5:    "ERROR_ACCESS_DENIED",
    6:    "ERROR_INVALID_HANDLE",
    7:    "ERROR_ARENA_TRASHED",
    8:    "ERROR_NOT_ENOUGH_MEMORY",
    9:    "ERROR_INVALID_BLOCK",
    10:   "ERROR_BAD_ENVIRONMENT",
    11:   "ERROR_BAD_FORMAT",
    12:   "ERROR_INVALID_ACCESS",
    13:   "ERROR_INVALID_DATA",
    14:   "ERROR_OUTOFMEMORY",
    15:   "ERROR_INVALID_DRIVE",
    16:   "ERROR_CURRENT_DIRECTORY",
    17:   "ERROR_NOT_SAME_DEVICE",
    18:   "ERROR_NO_MORE_FILES",
    19:   "ERROR_WRITE_PROTECT",
    20:   "ERROR_BAD_UNIT",
    21:   "ERROR_NOT_READY",
    22:   "ERROR_BAD_COMMAND",
    23:   "ERROR_CRC",
    24:   "ERROR_BAD_LENGTH",
    25:   "ERROR_SEEK",
    26:   "ERROR_NOT_DOS_DISK",
    27:   "ERROR_SECTOR_NOT_FOUND",
    28:   "ERROR_OUT_OF_PAPER",
    29:   "ERROR_WRITE_FAULT",
    30:   "ERROR_READ_FAULT",
    31:   "ERROR_GEN_FAILURE",
    32:   "ERROR_SHARING_VIOLATION",
    33:   "ERROR_LOCK_VIOLATION",
    34:   "ERROR_WRONG_DISK",
    36:   "ERROR_SHARING_BUFFER_EXCEEDED",
    38:   "ERROR_HANDLE_EOF",
    39:   "ERROR_HANDLE_DISK_FULL",
    50:   "ERROR_NOT_SUPPORTED",
    51:   "ERROR_REM_NOT_LIST",
    52:   "ERROR_DUP_NAME",
    53:   "ERROR_BAD_NETPATH",
    54:   "ERROR_NETWORK_BUSY",
    55:   "ERROR_DEV_NOT_EXIST",
    56:   "ERROR_TOO_MANY_CMDS",
    57:   "ERROR_ADAP_HDW_ERR",
    58:   "ERROR_BAD_NET_RESP",
    59:   "ERROR_UNEXP_NET_ERR",
    60:   "ERROR_BAD_REM_ADAP",
    61:   "ERROR_PRINTQ_FULL",
    62:   "ERROR_NO_SPOOL_SPACE",
    63:   "ERROR_PRINT_CANCELLED",
    64:   "ERROR_NETNAME_DELETED",
    65:   "ERROR_NETWORK_ACCESS_DENIED",
    66:   "ERROR_BAD_DEV_TYPE",
    67:   "ERROR_BAD_NET_NAME",
    68:   "ERROR_TOO_MANY_NAMES",
    69:   "ERROR_TOO_MANY_SESS",
    70:   "ERROR_SHARING_PAUSED",
    71:   "ERROR_REQ_NOT_ACCEP",
    72:   "ERROR_REDIR_PAUSED",
    80:   "ERROR_FILE_EXISTS",
    82:   "ERROR_CANNOT_MAKE",
    83:   "ERROR_FAIL_I24",
    84:   "ERROR_OUT_OF_STRUCTURES",
    85:   "ERROR_ALREADY_ASSIGNED",
    86:   "ERROR_INVALID_PASSWORD",
    87:   "ERROR_INVALID_PARAMETER",
    88:   "ERROR_NET_WRITE_FAULT",
    89:   "ERROR_NO_PROC_SLOTS",
    100:  "ERROR_TOO_MANY_SEMAPHORES",
    101:  "ERROR_EXCL_SEM_ALREADY_OWNED",
    102:  "ERROR_SEM_IS_SET",
    103:  "ERROR_TOO_MANY_SEM_REQUESTS",
    104:  "ERROR_INVALID_AT_INTERRUPT_TIME",
    105:  "ERROR_SEM_OWNER_DIED",
    106:  "ERROR_SEM_USER_LIMIT",
    107:  "ERROR_DISK_CHANGE",
    108:  "ERROR_DRIVE_LOCKED",
    109:  "ERROR_BROKEN_PIPE",
    110:  "ERROR_OPEN_FAILED",
    111:  "ERROR_BUFFER_OVERFLOW",
    112:  "ERROR_DISK_FULL",
    113:  "ERROR_NO_MORE_SEARCH_HANDLES",
    114:  "ERROR_INVALID_TARGET_HANDLE",
    117:  "ERROR_INVALID_CATEGORY",
    118:  "ERROR_INVALID_VERIFY_SWITCH",
    119:  "ERROR_BAD_DRIVER_LEVEL",
    120:  "ERROR_CALL_NOT_IMPLEMENTED",
    121:  "ERROR_SEM_TIMEOUT",
    122:  "ERROR_INSUFFICIENT_BUFFER",
    123:  "ERROR_INVALID_NAME",
    124:  "ERROR_INVALID_LEVEL",
    125:  "ERROR_NO_VOLUME_LABEL",
    126:  "ERROR_MOD_NOT_FOUND",
    127:  "ERROR_PROC_NOT_FOUND",
    128:  "ERROR_WAIT_NO_CHILDREN",
    129:  "ERROR_CHILD_NOT_COMPLETE",
    130:  "ERROR_DIRECT_ACCESS_HANDLE",
    131:  "ERROR_NEGATIVE_SEEK",
    132:  "ERROR_SEEK_ON_DEVICE",
    133:  "ERROR_IS_JOIN_TARGET",
    134:  "ERROR_IS_JOINED",
    135:  "ERROR_IS_SUBSTED",
    136:  "ERROR_NOT_JOINED",
    137:  "ERROR_NOT_SUBSTED",
    138:  "ERROR_JOIN_TO_JOIN",
    139:  "ERROR_SUBST_TO_SUBST",
    140:  "ERROR_JOIN_TO_SUBST",
    141:  "ERROR_SUBST_TO_JOIN",
    142:  "ERROR_BUSY_DRIVE",
    143:  "ERROR_SAME_DRIVE",
    144:  "ERROR_DIR_NOT_ROOT",
    145:  "ERROR_DIR_NOT_EMPTY",
    146:  "ERROR_IS_SUBST_PATH",
    147:  "ERROR_IS_JOIN_PATH",
    148:  "ERROR_PATH_BUSY",
    149:  "ERROR_IS_SUBST_TARGET",
    150:  "ERROR_SYSTEM_TRACE",
    151:  "ERROR_INVALID_EVENT_COUNT",
    152:  "ERROR_TOO_MANY_MUXWAITERS",
    153:  "ERROR_INVALID_LIST_FORMAT",
    154:  "ERROR_LABEL_TOO_LONG",
    155:  "ERROR_TOO_MANY_TCBS",
    156:  "ERROR_SIGNAL_REFUSED",
    157:  "ERROR_DISCARDED",
    158:  "ERROR_NOT_LOCKED",
    159:  "ERROR_BAD_THREADID_ADDR",
    160:  "ERROR_BAD_ARGUMENTS",
    161:  "ERROR_BAD_PATHNAME",
    162:  "ERROR_SIGNAL_PENDING",
    164:  "ERROR_MAX_THRDS_REACHED",
    167:  "ERROR_LOCK_FAILED",
    170:  "ERROR_BUSY",
    173:  "ERROR_CANCEL_VIOLATION",
    174:  "ERROR_ATOMIC_LOCKS_NOT_SUPPORTED",
    180:  "ERROR_INVALID_SEGMENT_NUMBER",
    182:  "ERROR_INVALID_ORDINAL",
    183:  "ERROR_ALREADY_EXISTS",
    186:  "ERROR_INVALID_FLAG_NUMBER",
    187:  "ERROR_SEM_NOT_FOUND",
    188:  "ERROR_INVALID_STARTING_CODESEG",
    189:  "ERROR_INVALID_STACKSEG",
    190:  "ERROR_INVALID_MODULETYPE",
    191:  "ERROR_INVALID_EXE_SIGNATURE",
    192:  "ERROR_EXE_MARKED_INVALID",
    193:  "ERROR_BAD_EXE_FORMAT",
    194:  "ERROR_ITERATED_DATA_EXCEEDS_64k",
    195:  "ERROR_INVALID_MINALLOCSIZE",
    196:  "ERROR_DYNLINK_FROM_INVALID_RING",
    197:  "ERROR_IOPL_NOT_ENABLED",
    198:  "ERROR_INVALID_SEGDPL",
    199:  "ERROR_AUTODATASEG_EXCEEDS_64k",
    200:  "ERROR_RING2SEG_MUST_BE_MOVABLE",
    201:  "ERROR_RELOC_CHAIN_XEEDS_SEGLIM",
    202:  "ERROR_INFLOOP_IN_RELOC_CHAIN",
    203:  "ERROR_ENVVAR_NOT_FOUND",
    205:  "ERROR_NO_SIGNAL_SENT",
    206:  "ERROR_FILENAME_EXCED_RANGE",
    207:  "ERROR_RING2_STACK_IN_USE",
    208:  "ERROR_META_EXPANSION_TOO_LONG",
    209:  "ERROR_INVALID_SIGNAL_NUMBER",
    210:  "ERROR_THREAD_1_INACTIVE",
    212:  "ERROR_LOCKED",
    214:  "ERROR_TOO_MANY_MODULES",
    215:  "ERROR_NESTING_NOT_ALLOWED",
    216:  "ERROR_EXE_MACHINE_TYPE_MISMATCH",
    217:  "ERROR_EXE_CANNOT_MODIFY_SIGNED_BINARY",
    218:  "ERROR_EXE_CANNOT_MODIFY_STRONG_SIGNED_BINARY",
    220:  "ERROR_FILE_CHECKED_OUT",
    221:  "ERROR_CHECKOUT_REQUIRED",
    222:  "ERROR_BAD_FILE_TYPE",
    223:  "ERROR_FILE_TOO_LARGE",
    224:  "ERROR_FORMS_AUTH_REQUIRED",
    225:  "ERROR_VIRUS_INFECTED",
    226:  "ERROR_VIRUS_DELETED",
    229:  "ERROR_PIPE_LOCAL",
    230:  "ERROR_BAD_PIPE",
    231:  "ERROR_PIPE_BUSY",
    232:  "ERROR_NO_DATA",
    233:  "ERROR_PIPE_NOT_CONNECTED",
    234:  "ERROR_MORE_DATA",
    240:  "ERROR_VC_DISCONNECTED",
    254:  "ERROR_INVALID_EA_NAME",
    255:  "ERROR_EA_LIST_INCONSISTENT",
    258:  "WAIT_TIMEOUT",
    259:  "ERROR_NO_MORE_ITEMS",
    266:  "ERROR_CANNOT_COPY",
    267:  "ERROR_DIRECTORY",
    275:  "ERROR_EAS_DIDNT_FIT",
    276:  "ERROR_EA_FILE_CORRUPT",
    277:  "ERROR_EA_TABLE_FULL",
    278:  "ERROR_INVALID_EA_HANDLE",
    282:  "ERROR_EAS_NOT_SUPPORTED",
    288:  "ERROR_NOT_OWNER",
    298:  "ERROR_TOO_MANY_POSTS",
    299:  "ERROR_PARTIAL_COPY",
    300:  "ERROR_OPLOCK_NOT_GRANTED",
    301:  "ERROR_INVALID_OPLOCK_PROTOCOL",
    302:  "ERROR_DISK_TOO_FRAGMENTED",
    303:  "ERROR_DELETE_PENDING",
    317:  "ERROR_MR_MID_NOT_FOUND",
    318:  "ERROR_SCOPE_NOT_FOUND",
    350:  "ERROR_FAIL_NOACTION_REBOOT",
    351:  "ERROR_FAIL_SHUTDOWN",
    352:  "ERROR_FAIL_RESTART",
    353:  "ERROR_MAX_SESSIONS_REACHED",
    400:  "ERROR_THREAD_MODE_ALREADY_BACKGROUND",
    401:  "ERROR_THREAD_MODE_NOT_BACKGROUND",
    402:  "ERROR_PROCESS_MODE_ALREADY_BACKGROUND",
    403:  "ERROR_PROCESS_MODE_NOT_BACKGROUND",
    487:  "ERROR_INVALID_ADDRESS",
    500:  "ERROR_USER_PROFILE_LOAD",
    534:  "ERROR_ARITHMETIC_OVERFLOW",
    535:  "ERROR_PIPE_CONNECTED",
    536:  "ERROR_PIPE_LISTENING",
    537:  "ERROR_VERIFIER_STOP",
    538:  "ERROR_ABIOS_ERROR",
    539:  "ERROR_WX86_WARNING",
    540:  "ERROR_WX86_ERROR",
    541:  "ERROR_TIMER_NOT_CANCELED",
    542:  "ERROR_UNWIND",
    543:  "ERROR_BAD_STACK",
    544:  "ERROR_INVALID_UNWIND_TARGET",
    545:  "ERROR_INVALID_PORT_ATTRIBUTES",
    546:  "ERROR_PORT_MESSAGE_TOO_LONG",
    547:  "ERROR_INVALID_QUOTA_LOWER",
    548:  "ERROR_DEVICE_ALREADY_ATTACHED",
    549:  "ERROR_INSTRUCTION_MISALIGNMENT",
    550:  "ERROR_PROFILING_NOT_STARTED",
    551:  "ERROR_PROFILING_NOT_STOPPED",
    552:  "ERROR_COULD_NOT_INTERPRET",
    553:  "ERROR_PROFILING_AT_LIMIT",
    554:  "ERROR_CANT_WAIT",
    555:  "ERROR_CANT_TERMINATE_SELF",
    556:  "ERROR_UNEXPECTED_MM_CREATE_ERR",
    557:  "ERROR_UNEXPECTED_MM_MAP_ERROR",
    558:  "ERROR_UNEXPECTED_MM_EXTEND_ERR",
    559:  "ERROR_BAD_FUNCTION_TABLE",
    560:  "ERROR_NO_GUID_TRANSLATION",
    561:  "ERROR_INVALID_LDT_SIZE",
    563:  "ERROR_INVALID_LDT_OFFSET",
    564:  "ERROR_INVALID_LDT_DESCRIPTOR",
    565:  "ERROR_TOO_MANY_THREADS",
    566:  "ERROR_THREAD_NOT_IN_PROCESS",
    567:  "ERROR_PAGEFILE_QUOTA_EXCEEDED",
    568:  "ERROR_LOGON_SERVER_CONFLICT",
    569:  "ERROR_SYNCHRONIZATION_REQUIRED",
    570:  "ERROR_NET_OPEN_FAILED",
    571:  "ERROR_IO_PRIVILEGE_FAILED",
    572:  "ERROR_CONTROL_C_EXIT",
    573:  "ERROR_MISSING_SYSTEMFILE",
    574:  "ERROR_UNHANDLED_EXCEPTION",
    575:  "ERROR_APP_INIT_FAILURE",
    576:  "ERROR_PAGEFILE_CREATE_FAILED",
    577:  "ERROR_INVALID_IMAGE_HASH",
    578:  "ERROR_NO_PAGEFILE",
    579:  "ERROR_ILLEGAL_FLOAT_CONTEXT",
    580:  "ERROR_NO_EVENT_PAIR",
    581:  "ERROR_DOMAIN_CTRLR_CONFIG_ERROR",
    582:  "ERROR_ILLEGAL_CHARACTER",
    583:  "ERROR_UNDEFINED_CHARACTER",
    584:  "ERROR_FLOPPY_VOLUME",
    585:  "ERROR_BIOS_FAILED_TO_CONNECT_INTERRUPT",
    586:  "ERROR_BACKUP_CONTROLLER",
    587:  "ERROR_MUTANT_LIMIT_EXCEEDED",
    588:  "ERROR_FS_DRIVER_REQUIRED",
    589:  "ERROR_CANNOT_LOAD_REGISTRY_FILE",
    590:  "ERROR_DEBUG_ATTACH_FAILED",
    591:  "ERROR_SYSTEM_PROCESS_TERMINATED",
    592:  "ERROR_DATA_NOT_ACCEPTED",
    593:  "ERROR_VDM_HARD_ERROR",
    594:  "ERROR_DRIVER_CANCEL_TIMEOUT",
    595:  "ERROR_REPLY_MESSAGE_MISMATCH",
    596:  "ERROR_LOST_WRITEBEHIND_DATA",
    597:  "ERROR_CLIENT_SERVER_PARAMETERS_INVALID",
    598:  "ERROR_NOT_TINY_STREAM",
    599:  "ERROR_STACK_OVERFLOW_READ",
    600:  "ERROR_CONVERT_TO_LARGE",
    601:  "ERROR_FOUND_OUT_OF_SCOPE",
    602:  "ERROR_ALLOCATE_BUCKET",
    603:  "ERROR_MARSHALL_OVERFLOW",
    604:  "ERROR_INVALID_VARIANT",
    605:  "ERROR_BAD_COMPRESSION_BUFFER",
    606:  "ERROR_AUDIT_FAILED",
    607:  "ERROR_TIMER_RESOLUTION_NOT_SET",
    608:  "ERROR_INSUFFICIENT_LOGON_INFO",
    609:  "ERROR_BAD_DLL_ENTRYPOINT",
    610:  "ERROR_BAD_SERVICE_ENTRYPOINT",
    611:  "ERROR_IP_ADDRESS_CONFLICT1",
    612:  "ERROR_IP_ADDRESS_CONFLICT2",
    613:  "ERROR_REGISTRY_QUOTA_LIMIT",
    614:  "ERROR_NO_CALLBACK_ACTIVE",
    615:  "ERROR_PWD_TOO_SHORT",
    616:  "ERROR_PWD_TOO_RECENT",
    617:  "ERROR_PWD_HISTORY_CONFLICT",
    618:  "ERROR_UNSUPPORTED_COMPRESSION",
    619:  "ERROR_INVALID_HW_PROFILE",
    620:  "ERROR_INVALID_PLUGPLAY_DEVICE_PATH",
    621:  "ERROR_QUOTA_LIST_INCONSISTENT",
    622:  "ERROR_EVALUATION_EXPIRATION",
    623:  "ERROR_ILLEGAL_DLL_RELOCATION",
    624:  "ERROR_DLL_INIT_FAILED_LOGOFF",
    625:  "ERROR_VALIDATE_CONTINUE",
    626:  "ERROR_NO_MORE_MATCHES",
    627:  "ERROR_RANGE_LIST_CONFLICT",
    628:  "ERROR_SERVER_SID_MISMATCH",
    629:  "ERROR_CANT_ENABLE_DENY_ONLY",
    630:  "ERROR_FLOAT_MULTIPLE_FAULTS",
    631:  "ERROR_FLOAT_MULTIPLE_TRAPS",
    632:  "ERROR_NOINTERFACE",
    633:  "ERROR_DRIVER_FAILED_SLEEP",
    634:  "ERROR_CORRUPT_SYSTEM_FILE",
    635:  "ERROR_COMMITMENT_MINIMUM",
    636:  "ERROR_PNP_RESTART_ENUMERATION",
    637:  "ERROR_SYSTEM_IMAGE_BAD_SIGNATURE",
    638:  "ERROR_PNP_REBOOT_REQUIRED",
    639:  "ERROR_INSUFFICIENT_POWER",
    640:  "ERROR_MULTIPLE_FAULT_VIOLATION",
    641:  "ERROR_SYSTEM_SHUTDOWN",
    642:  "ERROR_PORT_NOT_SET",
    643:  "ERROR_DS_VERSION_CHECK_FAILURE",
    644:  "ERROR_RANGE_NOT_FOUND",
    646:  "ERROR_NOT_SAFE_MODE_DRIVER",
    647:  "ERROR_FAILED_DRIVER_ENTRY",
    648:  "ERROR_DEVICE_ENUMERATION_ERROR",
    649:  "ERROR_MOUNT_POINT_NOT_RESOLVED",
    650:  "ERROR_INVALID_DEVICE_OBJECT_PARAMETER",
    651:  "ERROR_MCA_OCCURED",
    652:  "ERROR_DRIVER_DATABASE_ERROR",
    653:  "ERROR_SYSTEM_HIVE_TOO_LARGE",
    654:  "ERROR_DRIVER_FAILED_PRIOR_UNLOAD",
    655:  "ERROR_VOLSNAP_PREPARE_HIBERNATE",
    656:  "ERROR_HIBERNATION_FAILURE",
    665:  "ERROR_FILE_SYSTEM_LIMITATION",
    668:  "ERROR_ASSERTION_FAILURE",
    669:  "ERROR_ACPI_ERROR",
    670:  "ERROR_WOW_ASSERTION",
    671:  "ERROR_PNP_BAD_MPS_TABLE",
    672:  "ERROR_PNP_TRANSLATION_FAILED",
    673:  "ERROR_PNP_IRQ_TRANSLATION_FAILED",
    674:  "ERROR_PNP_INVALID_ID",
    675:  "ERROR_WAKE_SYSTEM_DEBUGGER",
    676:  "ERROR_HANDLES_CLOSED",
    677:  "ERROR_EXTRANEOUS_INFORMATION",
    678:  "ERROR_RXACT_COMMIT_NECESSARY",
    679:  "ERROR_MEDIA_CHECK",
    680:  "ERROR_GUID_SUBSTITUTION_MADE",
    681:  "ERROR_STOPPED_ON_SYMLINK",
    682:  "ERROR_LONGJUMP",
    683:  "ERROR_PLUGPLAY_QUERY_VETOED",
    684:  "ERROR_UNWIND_CONSOLIDATE",
    685:  "ERROR_REGISTRY_HIVE_RECOVERED",
    686:  "ERROR_DLL_MIGHT_BE_INSECURE",
    687:  "ERROR_DLL_MIGHT_BE_INCOMPATIBLE",
    688:  "ERROR_DBG_EXCEPTION_NOT_HANDLED",
    689:  "ERROR_DBG_REPLY_LATER",
    690:  "ERROR_DBG_UNABLE_TO_PROVIDE_HANDLE",
    691:  "ERROR_DBG_TERMINATE_THREAD",
    692:  "ERROR_DBG_TERMINATE_PROCESS",
    693:  "ERROR_DBG_CONTROL_C",
    694:  "ERROR_DBG_PRINTEXCEPTION_C",
    695:  "ERROR_DBG_RIPEXCEPTION",
    696:  "ERROR_DBG_CONTROL_BREAK",
    697:  "ERROR_DBG_COMMAND_EXCEPTION",
    698:  "ERROR_OBJECT_NAME_EXISTS",
    699:  "ERROR_THREAD_WAS_SUSPENDED",
    700:  "ERROR_IMAGE_NOT_AT_BASE",
    701:  "ERROR_RXACT_STATE_CREATED",
    702:  "ERROR_SEGMENT_NOTIFICATION",
    703:  "ERROR_BAD_CURRENT_DIRECTORY",
    704:  "ERROR_FT_READ_RECOVERY_FROM_BACKUP",
    705:  "ERROR_FT_WRITE_RECOVERY",
    706:  "ERROR_IMAGE_MACHINE_TYPE_MISMATCH",
    707:  "ERROR_RECEIVE_PARTIAL",
    708:  "ERROR_RECEIVE_EXPEDITED",
    709:  "ERROR_RECEIVE_PARTIAL_EXPEDITED",
    710:  "ERROR_EVENT_DONE",
    711:  "ERROR_EVENT_PENDING",
    712:  "ERROR_CHECKING_FILE_SYSTEM",
    713:  "ERROR_FATAL_APP_EXIT",
    714:  "ERROR_PREDEFINED_HANDLE",
    715:  "ERROR_WAS_UNLOCKED",
    716:  "ERROR_SERVICE_NOTIFICATION",
    717:  "ERROR_WAS_LOCKED",
    718:  "ERROR_LOG_HARD_ERROR",
    719:  "ERROR_ALREADY_WIN32",
    720:  "ERROR_IMAGE_MACHINE_TYPE_MISMATCH_EXE",
    721:  "ERROR_NO_YIELD_PERFORMED",
    722:  "ERROR_TIMER_RESUME_IGNORED",
    723:  "ERROR_ARBITRATION_UNHANDLED",
    724:  "ERROR_CARDBUS_NOT_SUPPORTED",
    725:  "ERROR_MP_PROCESSOR_MISMATCH",
    726:  "ERROR_HIBERNATED",
    727:  "ERROR_RESUME_HIBERNATION",
    728:  "ERROR_FIRMWARE_UPDATED",
    729:  "ERROR_DRIVERS_LEAKING_LOCKED_PAGES",
    730:  "ERROR_WAKE_SYSTEM",
    731:  "ERROR_WAIT_1",
    732:  "ERROR_WAIT_2",
    733:  "ERROR_WAIT_3",
    734:  "ERROR_WAIT_63",
    735:  "ERROR_ABANDONED_WAIT_0",
    736:  "ERROR_ABANDONED_WAIT_63",
    737:  "ERROR_USER_APC",
    738:  "ERROR_KERNEL_APC",
    739:  "ERROR_ALERTED",
    740:  "ERROR_ELEVATION_REQUIRED",
    741:  "ERROR_REPARSE",
    742:  "ERROR_OPLOCK_BREAK_IN_PROGRESS",
    743:  "ERROR_VOLUME_MOUNTED",
    744:  "ERROR_RXACT_COMMITTED",
    745:  "ERROR_NOTIFY_CLEANUP",
    746:  "ERROR_PRIMARY_TRANSPORT_CONNECT_FAILED",
    747:  "ERROR_PAGE_FAULT_TRANSITION",
    748:  "ERROR_PAGE_FAULT_DEMAND_ZERO",
    749:  "ERROR_PAGE_FAULT_COPY_ON_WRITE",
    750:  "ERROR_PAGE_FAULT_GUARD_PAGE",
    751:  "ERROR_PAGE_FAULT_PAGING_FILE",
    752:  "ERROR_CACHE_PAGE_LOCKED",
    753:  "ERROR_CRASH_DUMP",
    754:  "ERROR_BUFFER_ALL_ZEROS",
    755:  "ERROR_REPARSE_OBJECT",
    756:  "ERROR_RESOURCE_REQUIREMENTS_CHANGED",
    757:  "ERROR_TRANSLATION_COMPLETE",
    758:  "ERROR_NOTHING_TO_TERMINATE",
    759:  "ERROR_PROCESS_NOT_IN_JOB",
    760:  "ERROR_PROCESS_IN_JOB",
    761:  "ERROR_VOLSNAP_HIBERNATE_READY",
    762:  "ERROR_FSFILTER_OP_COMPLETED_SUCCESSFULLY",
    763:  "ERROR_INTERRUPT_VECTOR_ALREADY_CONNECTED",
    764:  "ERROR_INTERRUPT_STILL_CONNECTED",
    765:  "ERROR_WAIT_FOR_OPLOCK",
    766:  "ERROR_DBG_EXCEPTION_HANDLED",
    767:  "ERROR_DBG_CONTINUE",
    768:  "ERROR_CALLBACK_POP_STACK",
    769:  "ERROR_COMPRESSION_DISABLED",
    770:  "ERROR_CANTFETCHBACKWARDS",
    771:  "ERROR_CANTSCROLLBACKWARDS",
    772:  "ERROR_ROWSNOTRELEASED",
    773:  "ERROR_BAD_ACCESSOR_FLAGS",
    774:  "ERROR_ERRORS_ENCOUNTERED",
    775:  "ERROR_NOT_CAPABLE",
    776:  "ERROR_REQUEST_OUT_OF_SEQUENCE",
    777:  "ERROR_VERSION_PARSE_ERROR",
    778:  "ERROR_BADSTARTPOSITION",
    779:  "ERROR_MEMORY_HARDWARE",
    780:  "ERROR_DISK_REPAIR_DISABLED",
    781:  "ERROR_INSUFFICIENT_RESOURCE_FOR_SPECIFIED_SHARED_SECTION_SIZE",
    782:  "ERROR_SYSTEM_POWERSTATE_TRANSITION",
    783:  "ERROR_SYSTEM_POWERSTATE_COMPLEX_TRANSITION",
    784:  "ERROR_MCA_EXCEPTION",
    785:  "ERROR_ACCESS_AUDIT_BY_POLICY",
    786:  "ERROR_ACCESS_DISABLED_NO_SAFER_UI_BY_POLICY",
    787:  "ERROR_ABANDON_HIBERFILE",
    788:  "ERROR_LOST_WRITEBEHIND_DATA_NETWORK_DISCONNECTED",
    789:  "ERROR_LOST_WRITEBEHIND_DATA_NETWORK_SERVER_ERROR",
    790:  "ERROR_LOST_WRITEBEHIND_DATA_LOCAL_DISK_ERROR",
    791:  "ERROR_BAD_MCFG_TABLE",
    994:  "ERROR_EA_ACCESS_DENIED",
    995:  "ERROR_OPERATION_ABORTED",
    996:  "ERROR_IO_INCOMPLETE",
    997:  "ERROR_IO_PENDING",
    998:  "ERROR_NOACCESS",
    999:  "ERROR_SWAPERROR",
    1001: "ERROR_STACK_OVERFLOW",
    1002: "ERROR_INVALID_MESSAGE",
    1003: "ERROR_CAN_NOT_COMPLETE",
    1004: "ERROR_INVALID_FLAGS",
    1005: "ERROR_UNRECOGNIZED_VOLUME",
    1006: "ERROR_FILE_INVALID",
    1007: "ERROR_FULLSCREEN_MODE",
    1008: "ERROR_NO_TOKEN",
    1009: "ERROR_BADDB",
    1010: "ERROR_BADKEY",
    1011: "ERROR_CANTOPEN",
    1012: "ERROR_CANTREAD",
    1013: "ERROR_CANTWRITE",
    1014: "ERROR_REGISTRY_RECOVERED",
    1015: "ERROR_REGISTRY_CORRUPT",
    1016: "ERROR_REGISTRY_IO_FAILED",
    1017: "ERROR_NOT_REGISTRY_FILE",
    1018: "ERROR_KEY_DELETED",
    1019: "ERROR_NO_LOG_SPACE",
    1020: "ERROR_KEY_HAS_CHILDREN",
    1021: "ERROR_CHILD_MUST_BE_VOLATILE",
    1022: "ERROR_NOTIFY_ENUM_DIR",
    1051: "ERROR_DEPENDENT_SERVICES_RUNNING",
    1052: "ERROR_INVALID_SERVICE_CONTROL",
    1053: "ERROR_SERVICE_REQUEST_TIMEOUT",
    1054: "ERROR_SERVICE_NO_THREAD",
    1055: "ERROR_SERVICE_DATABASE_LOCKED",
    1056: "ERROR_SERVICE_ALREADY_RUNNING",
    1057: "ERROR_INVALID_SERVICE_ACCOUNT",
    1058: "ERROR_SERVICE_DISABLED",
    1059: "ERROR_CIRCULAR_DEPENDENCY",
    1060: "ERROR_SERVICE_DOES_NOT_EXIST",
    1061: "ERROR_SERVICE_CANNOT_ACCEPT_CTRL",
    1062: "ERROR_SERVICE_NOT_ACTIVE",
    1063: "ERROR_FAILED_SERVICE_CONTROLLER_CONNECT",
    1064: "ERROR_EXCEPTION_IN_SERVICE",
    1065: "ERROR_DATABASE_DOES_NOT_EXIST",
    1066: "ERROR_SERVICE_SPECIFIC_ERROR",
    1067: "ERROR_PROCESS_ABORTED",
    1068: "ERROR_SERVICE_DEPENDENCY_FAIL",
    1069: "ERROR_SERVICE_LOGON_FAILED",
    1070: "ERROR_SERVICE_START_HANG",
    1071: "ERROR_INVALID_SERVICE_LOCK",
    1072: "ERROR_SERVICE_MARKED_FOR_DELETE",
    1073: "ERROR_SERVICE_EXISTS",
    1074: "ERROR_ALREADY_RUNNING_LKG",
    1075: "ERROR_SERVICE_DEPENDENCY_DELETED",
    1076: "ERROR_BOOT_ALREADY_ACCEPTED",
    1077: "ERROR_SERVICE_NEVER_STARTED",
    1078: "ERROR_DUPLICATE_SERVICE_NAME",
    1079: "ERROR_DIFFERENT_SERVICE_ACCOUNT",
    1080: "ERROR_CANNOT_DETECT_DRIVER_FAILURE",
    1081: "ERROR_CANNOT_DETECT_PROCESS_ABORT",
    1082: "ERROR_NO_RECOVERY_PROGRAM",
    1083: "ERROR_SERVICE_NOT_IN_EXE",
    1084: "ERROR_NOT_SAFEBOOT_SERVICE",
    1100: "ERROR_END_OF_MEDIA",
    1101: "ERROR_FILEMARK_DETECTED",
    1102: "ERROR_BEGINNING_OF_MEDIA",
    1103: "ERROR_SETMARK_DETECTED",
    1104: "ERROR_NO_DATA_DETECTED",
    1105: "ERROR_PARTITION_FAILURE",
    1106: "ERROR_INVALID_BLOCK_LENGTH",
    1107: "ERROR_DEVICE_NOT_PARTITIONED",
    1108: "ERROR_UNABLE_TO_LOCK_MEDIA",
    1109: "ERROR_UNABLE_TO_UNLOAD_MEDIA",
    1110: "ERROR_MEDIA_CHANGED",
    1111: "ERROR_BUS_RESET",
    1112: "ERROR_NO_MEDIA_IN_DRIVE",
    1113: "ERROR_NO_UNICODE_TRANSLATION",
    1114: "ERROR_DLL_INIT_FAILED",
    1115: "ERROR_SHUTDOWN_IN_PROGRESS",
    1116: "ERROR_NO_SHUTDOWN_IN_PROGRESS",
    1117: "ERROR_IO_DEVICE",
    1118: "ERROR_SERIAL_NO_DEVICE",
    1119: "ERROR_IRQ_BUSY",
    1120: "ERROR_MORE_WRITES",
    1121: "ERROR_COUNTER_TIMEOUT",
    1122: "ERROR_FLOPPY_ID_MARK_NOT_FOUND",
    1123: "ERROR_FLOPPY_WRONG_CYLINDER",
    1124: "ERROR_FLOPPY_UNKNOWN_ERROR",
    1125: "ERROR_FLOPPY_BAD_REGISTERS",
    1126: "ERROR_DISK_RECALIBRATE_FAILED",
    1127: "ERROR_DISK_OPERATION_FAILED",
    1128: "ERROR_DISK_RESET_FAILED",
    1129: "ERROR_EOM_OVERFLOW",
    1130: "ERROR_NOT_ENOUGH_SERVER_MEMORY",
    1131: "ERROR_POSSIBLE_DEADLOCK",
    1132: "ERROR_MAPPED_ALIGNMENT",
    1140: "ERROR_SET_POWER_STATE_VETOED",
    1141: "ERROR_SET_POWER_STATE_FAILED",
    1142: "ERROR_TOO_MANY_LINKS",
    1150: "ERROR_OLD_WIN_VERSION",
    1151: "ERROR_APP_WRONG_OS",
    1152: "ERROR_SINGLE_INSTANCE_APP",
    1153: "ERROR_RMODE_APP",
    1154: "ERROR_INVALID_DLL",
    1155: "ERROR_NO_ASSOCIATION",
    1156: "ERROR_DDE_FAIL",
    1157: "ERROR_DLL_NOT_FOUND",
    1158: "ERROR_NO_MORE_USER_HANDLES",
    1159: "ERROR_MESSAGE_SYNC_ONLY",
    1160: "ERROR_SOURCE_ELEMENT_EMPTY",
    1161: "ERROR_DESTINATION_ELEMENT_FULL",
    1162: "ERROR_ILLEGAL_ELEMENT_ADDRESS",
    1163: "ERROR_MAGAZINE_NOT_PRESENT",
    1164: "ERROR_DEVICE_REINITIALIZATION_NEEDED",
    1165: "ERROR_DEVICE_REQUIRES_CLEANING",
    1166: "ERROR_DEVICE_DOOR_OPEN",
    1167: "ERROR_DEVICE_NOT_CONNECTED",
    1168: "ERROR_NOT_FOUND",
    1169: "ERROR_NO_MATCH",
    1170: "ERROR_SET_NOT_FOUND",
    1171: "ERROR_POINT_NOT_FOUND",
    1172: "ERROR_NO_TRACKING_SERVICE",
    1173: "ERROR_NO_VOLUME_ID",
    2108: "ERROR_CONNECTED_OTHER_PASSWORD",
    2202: "ERROR_BAD_USERNAME",
    2250: "ERROR_NOT_CONNECTED",
    2401: "ERROR_OPEN_FILES",
    2402: "ERROR_ACTIVE_CONNECTIONS",
    2404: "ERROR_DEVICE_IN_USE",
    1200: "ERROR_BAD_DEVICE",
    1201: "ERROR_CONNECTION_UNAVAIL",
    1202: "ERROR_DEVICE_ALREADY_REMEMBERED",
    1203: "ERROR_NO_NET_OR_BAD_PATH",
    1204: "ERROR_BAD_PROVIDER",
    1205: "ERROR_CANNOT_OPEN_PROFILE",
    1206: "ERROR_BAD_PROFILE",
    1207: "ERROR_NOT_CONTAINER",
    1208: "ERROR_EXTENDED_ERROR",
    1209: "ERROR_INVALID_GROUPNAME",
    1210: "ERROR_INVALID_COMPUTERNAME",
    1211: "ERROR_INVALID_EVENTNAME",
    1212: "ERROR_INVALID_DOMAINNAME",
    1213: "ERROR_INVALID_SERVICENAME",
    1214: "ERROR_INVALID_NETNAME",
    1215: "ERROR_INVALID_SHARENAME",
    1216: "ERROR_INVALID_PASSWORDNAME",
    1217: "ERROR_INVALID_MESSAGENAME",
    1218: "ERROR_INVALID_MESSAGEDEST",
    1219: "ERROR_SESSION_CREDENTIAL_CONFLICT",
    1220: "ERROR_REMOTE_SESSION_LIMIT_EXCEEDED",
    1221: "ERROR_DUP_DOMAINNAME",
    1222: "ERROR_NO_NETWORK",
    1223: "ERROR_CANCELLED",
    1224: "ERROR_USER_MAPPED_FILE",
    1225: "ERROR_CONNECTION_REFUSED",
    1226: "ERROR_GRACEFUL_DISCONNECT",
    1227: "ERROR_ADDRESS_ALREADY_ASSOCIATED",
    1228: "ERROR_ADDRESS_NOT_ASSOCIATED",
    1229: "ERROR_CONNECTION_INVALID",
    1230: "ERROR_CONNECTION_ACTIVE",
    1231: "ERROR_NETWORK_UNREACHABLE",
    1232: "ERROR_HOST_UNREACHABLE",
    1233: "ERROR_PROTOCOL_UNREACHABLE",
    1234: "ERROR_PORT_UNREACHABLE",
    1235: "ERROR_REQUEST_ABORTED",
    1236: "ERROR_CONNECTION_ABORTED",
    1237: "ERROR_RETRY",
    1238: "ERROR_CONNECTION_COUNT_LIMIT",
    1239: "ERROR_LOGIN_TIME_RESTRICTION",
    1240: "ERROR_LOGIN_WKSTA_RESTRICTION",
    1241: "ERROR_INCORRECT_ADDRESS",
    1242: "ERROR_ALREADY_REGISTERED",
    1243: "ERROR_SERVICE_NOT_FOUND",
    1244: "ERROR_NOT_AUTHENTICATED",
    1245: "ERROR_NOT_LOGGED_ON",
    1246: "ERROR_CONTINUE",
    1247: "ERROR_ALREADY_INITIALIZED",
    1248: "ERROR_NO_MORE_DEVICES",
    1249: "ERROR_NO_SUCH_SITE",
    1250: "ERROR_DOMAIN_CONTROLLER_EXISTS",
    1251: "ERROR_DS_NOT_INSTALLED",
    1300: "ERROR_NOT_ALL_ASSIGNED",
    1301: "ERROR_SOME_NOT_MAPPED",
    1302: "ERROR_NO_QUOTAS_FOR_ACCOUNT",
    1303: "ERROR_LOCAL_USER_SESSION_KEY",
    1304: "ERROR_NULL_LM_PASSWORD",
    1305: "ERROR_UNKNOWN_REVISION",
    1306: "ERROR_REVISION_MISMATCH",
    1307: "ERROR_INVALID_OWNER",
    1308: "ERROR_INVALID_PRIMARY_GROUP",
    1309: "ERROR_NO_IMPERSONATION_TOKEN",
    1310: "ERROR_CANT_DISABLE_MANDATORY",
    1311: "ERROR_NO_LOGON_SERVERS",
    1312: "ERROR_NO_SUCH_LOGON_SESSION",
    1313: "ERROR_NO_SUCH_PRIVILEGE",
    1314: "ERROR_PRIVILEGE_NOT_HELD",
    1315: "ERROR_INVALID_ACCOUNT_NAME",
    1316: "ERROR_USER_EXISTS",
    1317: "ERROR_NO_SUCH_USER",
    1318: "ERROR_GROUP_EXISTS",
    1319: "ERROR_NO_SUCH_GROUP",
    1320: "ERROR_MEMBER_IN_GROUP",
    1321: "ERROR_MEMBER_NOT_IN_GROUP",
    1322: "ERROR_LAST_ADMIN",
    1323: "ERROR_WRONG_PASSWORD",
    1324: "ERROR_ILL_FORMED_PASSWORD",
    1325: "ERROR_PASSWORD_RESTRICTION",
    1326: "ERROR_LOGON_FAILURE",
    1327: "ERROR_ACCOUNT_RESTRICTION",
    1328: "ERROR_INVALID_LOGON_HOURS",
    1329: "ERROR_INVALID_WORKSTATION",
    1330: "ERROR_PASSWORD_EXPIRED",
    1331: "ERROR_ACCOUNT_DISABLED",
    1332: "ERROR_NONE_MAPPED",
    1333: "ERROR_TOO_MANY_LUIDS_REQUESTED",
    1334: "ERROR_LUIDS_EXHAUSTED",
    1335: "ERROR_INVALID_SUB_AUTHORITY",
    1336: "ERROR_INVALID_ACL",
    1337: "ERROR_INVALID_SID",
    1338: "ERROR_INVALID_SECURITY_DESCR",
    1340: "ERROR_BAD_INHERITANCE_ACL",
    1341: "ERROR_SERVER_DISABLED",
    1342: "ERROR_SERVER_NOT_DISABLED",
    1343: "ERROR_INVALID_ID_AUTHORITY",
    1344: "ERROR_ALLOTTED_SPACE_EXCEEDED",
    1345: "ERROR_INVALID_GROUP_ATTRIBUTES",
    1346: "ERROR_BAD_IMPERSONATION_LEVEL",
    1347: "ERROR_CANT_OPEN_ANONYMOUS",
    1348: "ERROR_BAD_VALIDATION_CLASS",
    1349: "ERROR_BAD_TOKEN_TYPE",
    1350: "ERROR_NO_SECURITY_ON_OBJECT",
    1351: "ERROR_CANT_ACCESS_DOMAIN_INFO",
    1352: "ERROR_INVALID_SERVER_STATE",
    1353: "ERROR_INVALID_DOMAIN_STATE",
    1354: "ERROR_INVALID_DOMAIN_ROLE",
    1355: "ERROR_NO_SUCH_DOMAIN",
    1356: "ERROR_DOMAIN_EXISTS",
    1357: "ERROR_DOMAIN_LIMIT_EXCEEDED",
    1358: "ERROR_INTERNAL_DB_CORRUPTION",
    1359: "ERROR_INTERNAL_ERROR",
    1360: "ERROR_GENERIC_NOT_MAPPED",
    1361: "ERROR_BAD_DESCRIPTOR_FORMAT",
    1362: "ERROR_NOT_LOGON_PROCESS",
    1363: "ERROR_LOGON_SESSION_EXISTS",
    1364: "ERROR_NO_SUCH_PACKAGE",
    1365: "ERROR_BAD_LOGON_SESSION_STATE",
    1366: "ERROR_LOGON_SESSION_COLLISION",
    1367: "ERROR_INVALID_LOGON_TYPE",
    1368: "ERROR_CANNOT_IMPERSONATE",
    1369: "ERROR_RXACT_INVALID_STATE",
    1370: "ERROR_RXACT_COMMIT_FAILURE",
    1371: "ERROR_SPECIAL_ACCOUNT",
    1372: "ERROR_SPECIAL_GROUP",
    1373: "ERROR_SPECIAL_USER",
    1374: "ERROR_MEMBERS_PRIMARY_GROUP",
    1375: "ERROR_TOKEN_ALREADY_IN_USE",
    1376: "ERROR_NO_SUCH_ALIAS",
    1377: "ERROR_MEMBER_NOT_IN_ALIAS",
    1378: "ERROR_MEMBER_IN_ALIAS",
    1379: "ERROR_ALIAS_EXISTS",
    1380: "ERROR_LOGON_NOT_GRANTED",
    1381: "ERROR_TOO_MANY_SECRETS",
    1382: "ERROR_SECRET_TOO_LONG",
    1383: "ERROR_INTERNAL_DB_ERROR",
    1384: "ERROR_TOO_MANY_CONTEXT_IDS",
    1385: "ERROR_LOGON_TYPE_NOT_GRANTED",
    1386: "ERROR_NT_CROSS_ENCRYPTION_REQUIRED",
    1387: "ERROR_NO_SUCH_MEMBER",
    1388: "ERROR_INVALID_MEMBER",
    1389: "ERROR_TOO_MANY_SIDS",
    1390: "ERROR_LM_CROSS_ENCRYPTION_REQUIRED",
    1391: "ERROR_NO_INHERITANCE",
    1392: "ERROR_FILE_CORRUPT",
    1393: "ERROR_DISK_CORRUPT",
    1394: "ERROR_NO_USER_SESSION_KEY",
    1395: "ERROR_LICENSE_QUOTA_EXCEEDED",
    1400: "ERROR_INVALID_WINDOW_HANDLE",
    1401: "ERROR_INVALID_MENU_HANDLE",
    1402: "ERROR_INVALID_CURSOR_HANDLE",
    1403: "ERROR_INVALID_ACCEL_HANDLE",
    1404: "ERROR_INVALID_HOOK_HANDLE",
    1405: "ERROR_INVALID_DWP_HANDLE",
    1406: "ERROR_TLW_WITH_WSCHILD",
    1407: "ERROR_CANNOT_FIND_WND_CLASS",
    1408: "ERROR_WINDOW_OF_OTHER_THREAD",
    1409: "ERROR_HOTKEY_ALREADY_REGISTERED",
    1410: "ERROR_CLASS_ALREADY_EXISTS",
    1411: "ERROR_CLASS_DOES_NOT_EXIST",
    1412: "ERROR_CLASS_HAS_WINDOWS",
    1413: "ERROR_INVALID_INDEX",
    1414: "ERROR_INVALID_ICON_HANDLE",
    1415: "ERROR_PRIVATE_DIALOG_INDEX",
    1416: "ERROR_LISTBOX_ID_NOT_FOUND",
    1417: "ERROR_NO_WILDCARD_CHARACTERS",
    1418: "ERROR_CLIPBOARD_NOT_OPEN",
    1419: "ERROR_HOTKEY_NOT_REGISTERED",
    1420: "ERROR_WINDOW_NOT_DIALOG",
    1421: "ERROR_CONTROL_ID_NOT_FOUND",
    1422: "ERROR_INVALID_COMBOBOX_MESSAGE",
    1423: "ERROR_WINDOW_NOT_COMBOBOX",
    1424: "ERROR_INVALID_EDIT_HEIGHT",
    1425: "ERROR_DC_NOT_FOUND",
    1426: "ERROR_INVALID_HOOK_FILTER",
    1427: "ERROR_INVALID_FILTER_PROC",
    1428: "ERROR_HOOK_NEEDS_HMOD",
    1429: "ERROR_GLOBAL_ONLY_HOOK",
    1430: "ERROR_JOURNAL_HOOK_SET",
    1431: "ERROR_HOOK_NOT_INSTALLED",
    1432: "ERROR_INVALID_LB_MESSAGE",
    1433: "ERROR_SETCOUNT_ON_BAD_LB",
    1434: "ERROR_LB_WITHOUT_TABSTOPS",
    1435: "ERROR_DESTROY_OBJECT_OF_OTHER_THREAD",
    1436: "ERROR_CHILD_WINDOW_MENU",
    1437: "ERROR_NO_SYSTEM_MENU",
    1438: "ERROR_INVALID_MSGBOX_STYLE",
    1439: "ERROR_INVALID_SPI_VALUE",
    1440: "ERROR_SCREEN_ALREADY_LOCKED",
    1441: "ERROR_HWNDS_HAVE_DIFF_PARENT",
    1442: "ERROR_NOT_CHILD_WINDOW",
    1443: "ERROR_INVALID_GW_COMMAND",
    1444: "ERROR_INVALID_THREAD_ID",
    1445: "ERROR_NON_MDICHILD_WINDOW",
    1446: "ERROR_POPUP_ALREADY_ACTIVE",
    1447: "ERROR_NO_SCROLLBARS",
    1448: "ERROR_INVALID_SCROLLBAR_RANGE",
    1449: "ERROR_INVALID_SHOWWIN_COMMAND",
    1450: "ERROR_NO_SYSTEM_RESOURCES",
    1451: "ERROR_NONPAGED_SYSTEM_RESOURCES",
    1452: "ERROR_PAGED_SYSTEM_RESOURCES",
    1453: "ERROR_WORKING_SET_QUOTA",
    1454: "ERROR_PAGEFILE_QUOTA",
    1455: "ERROR_COMMITMENT_LIMIT",
    1456: "ERROR_MENU_ITEM_NOT_FOUND",
    1457: "ERROR_INVALID_KEYBOARD_HANDLE",
    1458: "ERROR_HOOK_TYPE_NOT_ALLOWED",
    1459: "ERROR_REQUIRES_INTERACTIVE_WINDOWSTATION",
    1460: "ERROR_TIMEOUT",
    1461: "ERROR_INVALID_MONITOR_HANDLE",
    1462: "ERROR_INCORRECT_SIZE",
    1463: "ERROR_SYMLINK_CLASS_DISABLED",
    1464: "ERROR_SYMLINK_NOT_SUPPORTED",
    1465: "ERROR_XML_PARSE_ERROR",
    1466: "ERROR_XMLDSIG_ERROR",
    1467: "ERROR_RESTART_APPLICATION",
    1468: "ERROR_WRONG_COMPARTMENT",
    1469: "ERROR_AUTHIP_FAILURE",
    1500: "ERROR_EVENTLOG_FILE_CORRUPT",
    1501: "ERROR_EVENTLOG_CANT_START",
    1502: "ERROR_LOG_FILE_FULL",
    1503: "ERROR_EVENTLOG_FILE_CHANGED",
    1601: "ERROR_INSTALL_SERVICE",
    1602: "ERROR_INSTALL_USEREXIT",
    1603: "ERROR_INSTALL_FAILURE",
    1604: "ERROR_INSTALL_SUSPEND",
    1605: "ERROR_UNKNOWN_PRODUCT",
    1606: "ERROR_UNKNOWN_FEATURE",
    1607: "ERROR_UNKNOWN_COMPONENT",
    1608: "ERROR_UNKNOWN_PROPERTY",
    1609: "ERROR_INVALID_HANDLE_STATE",
    1610: "ERROR_BAD_CONFIGURATION",
    1611: "ERROR_INDEX_ABSENT",
    1612: "ERROR_INSTALL_SOURCE_ABSENT",
    1613: "ERROR_BAD_DATABASE_VERSION",
    1614: "ERROR_PRODUCT_UNINSTALLED",
    1615: "ERROR_BAD_QUERY_SYNTAX",
    1616: "ERROR_INVALID_FIELD",
    1617: "ERROR_DEVICE_REMOVED",
    1618: "ERROR_INSTALL_ALREADY_RUNNING",
    1619: "ERROR_INSTALL_PACKAGE_OPEN_FAILED",
    1620: "ERROR_INSTALL_PACKAGE_INVALID",
    1621: "ERROR_INSTALL_UI_FAILURE",
    1622: "ERROR_INSTALL_LOG_FAILURE",
    1623: "ERROR_INSTALL_LANGUAGE_UNSUPPORTED",
    1624: "ERROR_INSTALL_TRANSFORM_FAILURE",
    1625: "ERROR_INSTALL_PACKAGE_REJECTED",
    1626: "ERROR_FUNCTION_NOT_CALLED",
    1627: "ERROR_FUNCTION_FAILED",
    1628: "ERROR_INVALID_TABLE",
    1629: "ERROR_DATATYPE_MISMATCH",
    1630: "ERROR_UNSUPPORTED_TYPE",
    1631: "ERROR_CREATE_FAILED",
    1632: "ERROR_INSTALL_TEMP_UNWRITABLE",
    1633: "ERROR_INSTALL_PLATFORM_UNSUPPORTED",
    1634: "ERROR_INSTALL_NOTUSED",
    1635: "ERROR_PATCH_PACKAGE_OPEN_FAILED",
    1636: "ERROR_PATCH_PACKAGE_INVALID",
    1637: "ERROR_PATCH_PACKAGE_UNSUPPORTED",
    1638: "ERROR_PRODUCT_VERSION",
    1639: "ERROR_INVALID_COMMAND_LINE",
    1640: "ERROR_INSTALL_REMOTE_DISALLOWED",
    1641: "ERROR_SUCCESS_REBOOT_INITIATED",
    1642: "ERROR_PATCH_TARGET_NOT_FOUND",
    1643: "ERROR_PATCH_PACKAGE_REJECTED",
    1644: "ERROR_INSTALL_TRANSFORM_REJECTED",
    1645: "ERROR_INSTALL_REMOTE_PROHIBITED",
    1646: "ERROR_PATCH_REMOVAL_UNSUPPORTED",
    1647: "ERROR_UNKNOWN_PATCH",
    1648: "ERROR_PATCH_NO_SEQUENCE",
    1649: "ERROR_PATCH_REMOVAL_DISALLOWED",
    1650: "ERROR_INVALID_PATCH_XML",
    1651: "ERROR_PATCH_MANAGED_ADVERTISED_PRODUCT",
    1652: "ERROR_INSTALL_SERVICE_SAFEBOOT",
}
