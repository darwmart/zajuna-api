// src/components/UsersTableView.js
import React, { useEffect, useState, useCallback, useRef, useMemo } from "react";
// using font-awesome <i> elements instead of react-icons
import { FaTrash, FaEye, FaEyeSlash, FaCog } from "react-icons/fa";
import { getUsers, getUsersByField, deleteUsers, toggleUserStatus } from "../services/usersService";
import { coursesService } from "../services/coursesService";
import "../styles/moodle-theme.css";
import { useNavigate } from "react-router-dom";
import MoreUserFilters from "./MoreUserFilters";

export default function UsersTableView() {
  const navigate = useNavigate();
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(false);
  const [total, setTotal] = useState(0);
  const [fullTotal, setFullTotal] = useState(null); // total without filters (full dataset)
  const [currentPage, setCurrentPage] = useState(1);
  const [perPage] = useState(30);
  const [filters, setFilters] = useState([]);
  const [filterOp, setFilterOp] = useState("contains");
  const [filterValue, setFilterValue] = useState("");
  const [showFilterSection, setShowFilterSection] = useState(true);
  const [showMoreFilters, setShowMoreFilters] = useState(false);
  const [lastNameOp, setLastNameOp] = useState("contains");
  const [lastNameVal, setLastNameVal] = useState("");
  const [firstNameOp, setFirstNameOp] = useState("contains");
  const [firstNameVal, setFirstNameVal] = useState("");
  const [usernameOp, setUsernameOp] = useState("contains");
  const [usernameVal, setUsernameVal] = useState("");
  const [emailOp, setEmailOp] = useState("contains");
  const [emailVal, setEmailVal] = useState("");
  const [cityOp, setCityOp] = useState("contains");
  const [cityVal, setCityVal] = useState("");
  const [countryOp, setCountryOp] = useState("any");
  const [countryVal, setCountryVal] = useState("Colombia");
  const [confirmedOp, setConfirmedOp] = useState("any");
  const [suspendedOp, setSuspendedOp] = useState("any");
  const [profileFieldOp, setProfileFieldOp] = useState("contains");
  const [profileFieldVal, setProfileFieldVal] = useState("");
  const [profileFieldName, setProfileFieldName] = useState("any");
  const [roleFieldOne, setRoleFieldOne] = useState("any");
  const [roleFieldTwo, setRoleFieldTwo] = useState("any");
  const [roleFieldVal, setRoleFieldVal] = useState("");
  const [roleCategories, setRoleCategories] = useState([]);
  const [enrolledOp, setEnrolledOp] = useState("any");
  const [systemRoleOp, setSystemRoleOp] = useState("any");
  const [cohortOp, setCohortOp] = useState("any");
  const [cohortVal, setCohortVal] = useState("");
  const [firstAfterEnabled, setFirstAfterEnabled] = useState(false);
  const [firstAfterDay, setFirstAfterDay] = useState(() => String(new Date().getDate()));
  const [firstAfterMonth, setFirstAfterMonth] = useState(() => String(new Date().getMonth() + 1));
  const [firstAfterYear, setFirstAfterYear] = useState(() => String(new Date().getFullYear()));
  const [firstBeforeEnabled, setFirstBeforeEnabled] = useState(false);
  const [firstBeforeDay, setFirstBeforeDay] = useState(() => String(new Date().getDate()));
  const [firstBeforeMonth, setFirstBeforeMonth] = useState(() => String(new Date().getMonth() + 1));
  const [firstBeforeYear, setFirstBeforeYear] = useState(() => String(new Date().getFullYear()));
  // Último acceso states (similar to Primer acceso)
  const [lastAfterEnabled, setLastAfterEnabled] = useState(false);
  const [lastAfterDay, setLastAfterDay] = useState(() => String(new Date().getDate()));
  const [lastAfterMonth, setLastAfterMonth] = useState(() => String(new Date().getMonth() + 1));
  const [lastAfterYear, setLastAfterYear] = useState(() => String(new Date().getFullYear()));
  const [lastBeforeEnabled, setLastBeforeEnabled] = useState(false);
  const [lastBeforeDay, setLastBeforeDay] = useState(() => String(new Date().getDate()));
  const [lastBeforeMonth, setLastBeforeMonth] = useState(() => String(new Date().getMonth() + 1));
  const [lastBeforeYear, setLastBeforeYear] = useState(() => String(new Date().getFullYear()));
  const [lastNever, setLastNever] = useState(false);
  // Tiempo de creado (created time) states
  const [createdAfterEnabled, setCreatedAfterEnabled] = useState(false);
  const [createdAfterDay, setCreatedAfterDay] = useState(() => String(new Date().getDate()));
  const [createdAfterMonth, setCreatedAfterMonth] = useState(() => String(new Date().getMonth() + 1));
  const [createdAfterYear, setCreatedAfterYear] = useState(() => String(new Date().getFullYear()));
  const [createdBeforeEnabled, setCreatedBeforeEnabled] = useState(false);
  const [createdBeforeDay, setCreatedBeforeDay] = useState(() => String(new Date().getDate()));
  const [createdBeforeMonth, setCreatedBeforeMonth] = useState(() => String(new Date().getMonth() + 1));
  const [createdBeforeYear, setCreatedBeforeYear] = useState(() => String(new Date().getFullYear()));
  // Última modificación (modified time) states
  const [modifiedAfterEnabled, setModifiedAfterEnabled] = useState(false);
  const [modifiedAfterDay, setModifiedAfterDay] = useState(() => String(new Date().getDate()));
  const [modifiedAfterMonth, setModifiedAfterMonth] = useState(() => String(new Date().getMonth() + 1));
  const [modifiedAfterYear, setModifiedAfterYear] = useState(() => String(new Date().getFullYear()));
  const [modifiedBeforeEnabled, setModifiedBeforeEnabled] = useState(false);
  const [modifiedBeforeDay, setModifiedBeforeDay] = useState(() => String(new Date().getDate()));
  const [modifiedBeforeMonth, setModifiedBeforeMonth] = useState(() => String(new Date().getMonth() + 1));
  const [modifiedBeforeYear, setModifiedBeforeYear] = useState(() => String(new Date().getFullYear()));
  const [modifiedNever, setModifiedNever] = useState(false);
  // Identificación filter operator
  const [identificationOp, setIdentificationOp] = useState('');
  const [identificationTouched, setIdentificationTouched] = useState(false);
  // Proveedor de ID MNET
  const [mnetProviderOp, setMnetProviderOp] = useState('');
  const [mnetProviderTouched, setMnetProviderTouched] = useState(false);
  // Número de ID
  const [idNumberOp, setIdNumberOp] = useState('contains');
  const [idNumberVal, setIdNumberVal] = useState('');
  // Institución
  const [institutionOp, setInstitutionOp] = useState('contains');
  const [institutionVal, setInstitutionVal] = useState('');
  // Departamento
  const [departmentOp, setDepartmentOp] = useState('contains');
  const [departmentVal, setDepartmentVal] = useState('');
  // Última dirección IP
  const [lastIpOp, setLastIpOp] = useState('contains');
  const [lastIpVal, setLastIpVal] = useState('');

  const identificationOptionsMap = {
    '': 'cualquier valor',
    cas: 'Usar un servidor CAS (SSO)',
    db: 'Usar una base de datos externa',
    email: 'Identificación basada en Email',
    ldap: 'Usar un servidor LDAP',
    lti: 'LTI',
    manual: 'Cuentas manuales',
    mnet: "Identificación de la Red Moodle ('Moodle Network')",
    nologin: 'No hay sesión',
    none: 'Sin identificación',
    oauth2: 'OAuth 2',
    oidc: 'Conexión OpenID',
    shibboleth: 'Shibboleth',
    webservice: "Identificación de Servicios Web ('Web Services')",
  };
  const mnetProviderOptionsMap = {
    '': 'cualquier valor',
    'sena_zajuna': 'SENA - Zajuna (local)',
    '0': 'id: 0 (Error)'
  };
  const [filterLogic] = useState("AND"); // AND | OR
  const [showActiveFilters, setShowActiveFilters] = useState(true);
  const [selectedFiltersToRemove, setSelectedFiltersToRemove] = useState([]);
  const filterInputRef = useRef(null);
  const lastNameRef = useRef(null);
  const firstNameRef = useRef(null);
  const usernameRef = useRef(null);
  const emailRef = useRef(null);
  const cityRef = useRef(null);
  const idNumberRef = useRef(null);
  const institutionRef = useRef(null);
  const departmentRef = useRef(null);
  const lastIpRef = useRef(null);
  const profileFieldRef = useRef(null);
  const roleFieldRef = useRef(null);
  const cohortRef = useRef(null);
  const afterDateInputRef = useRef(null);
  const beforeDateInputRef = useRef(null);
  const afterCalendarIconRef = useRef(null);
  const beforeCalendarIconRef = useRef(null);
  const lastAfterCalendarIconRef = useRef(null);
  const lastBeforeCalendarIconRef = useRef(null);
  const createdAfterCalendarIconRef = useRef(null);
  const createdBeforeCalendarIconRef = useRef(null);
  const modifiedAfterCalendarIconRef = useRef(null);
  const modifiedBeforeCalendarIconRef = useRef(null);
  const currentRequestId = useRef(0); // request id to avoid applying stale responses
  const resultsCache = useRef(new Map());
  const containerRef = useRef(null);
  const [hiddenUsers, setHiddenUsers] = useState(new Set());

  const headerHeightRef = useRef(null);

  const detectHeaderHeight = () => {
    if (headerHeightRef.current != null) return headerHeightRef.current;
    try {
      // Try to find a sticky or fixed header at the top of the document
      const candidates = Array.from(document.querySelectorAll('body *'));
      let found = null;
      for (const el of candidates) {
        try {
          const style = window.getComputedStyle(el);
          if ((style.position === 'sticky' || style.position === 'fixed') && (style.top === '0px' || style.top === '0')) {
            const rect = el.getBoundingClientRect();
            if (rect.height > 0) { found = rect.height; break; }
          }
        } catch (_) {
          // ignore computed style access errors for some nodes
        }
      }
      headerHeightRef.current = Number(found || 0);
      return headerHeightRef.current;
    } catch (e) {
      headerHeightRef.current = 0;
      return 0;
    }
  };

  const scrollToTop = (smooth = true) => {
    const doScroll = () => {
      try {
        const targetEl = containerRef.current || document.documentElement || document.body;

        // Find nearest scrollable ancestor (including document scrolling)
        const findScrollParent = (node) => {
          let el = node;
          while (el && el !== document.body) {
            try {
              const style = window.getComputedStyle(el);
              const overflowY = style.overflowY;
              if ((overflowY === 'auto' || overflowY === 'scroll') && el.scrollHeight > el.clientHeight) return el;
            } catch (_) {}
            el = el.parentElement;
          }
          return null;
        };

        const scrollParent = findScrollParent(targetEl);
        if (scrollParent) {
          if (process.env.NODE_ENV !== 'production') console.debug('scrollToTop -> scrollParent.top=0');
          if (smooth && 'scrollTo' in scrollParent) {
            scrollParent.scrollTo({ top: 0, behavior: 'smooth' });
          } else {
            scrollParent.scrollTop = 0;
          }
        } else {
          if (process.env.NODE_ENV !== 'production') console.debug('scrollToTop -> window.top=0');
          window.scrollTo({ top: 0, behavior: smooth ? 'smooth' : 'auto' });
        }
      } catch (e) {
        // fallback
        window.scrollTo(0, 0);
      }
    };

    // Ensure we execute after layout/paint. Use rAF twice and a timeout fallback.
    try {
      requestAnimationFrame(() => requestAnimationFrame(() => doScroll()));
      // Fallback in case rAF doesn't run (rare in test envs)
      setTimeout(doScroll, 50);
    } catch (e) {
      doScroll();
    }
  };

  const [calendarState, setCalendarState] = useState({ visible: false, which: null, x: 0, y: 0, placement: 'bottom', month: null, year: null });

  const [filtersVersion, setFiltersVersion] = useState(0);
  
  // Reset advanced filter controls (used in multiple places: add, replace, page-change)
  const resetAdvancedFilters = useCallback(() => {
    const now = new Date();
    setFilterOp('contains');
    setFilterValue('');

    setLastNameOp("contains"); setLastNameVal("");
    setFirstNameOp("contains"); setFirstNameVal("");
    setUsernameOp("contains"); setUsernameVal("");
    setEmailOp("contains"); setEmailVal("");
    setCityOp("contains"); setCityVal("");
    setCountryOp("any"); setCountryVal('Colombia');
    setConfirmedOp("any"); setSuspendedOp("any");
    setProfileFieldName('any'); setProfileFieldOp("contains"); setProfileFieldVal("");
    setRoleFieldOne("any"); setRoleFieldTwo("any"); setRoleFieldVal("");
    setEnrolledOp("any"); setSystemRoleOp("any");
    setCohortOp("any"); setCohortVal("");

    setFirstAfterEnabled(false);
    setFirstAfterDay(String(now.getDate()));
    setFirstAfterMonth(String(now.getMonth() + 1));
    setFirstAfterYear(String(now.getFullYear()));
    setFirstBeforeEnabled(false);
    setFirstBeforeDay(String(now.getDate()));
    setFirstBeforeMonth(String(now.getMonth() + 1));
    setFirstBeforeYear(String(now.getFullYear()));

    setLastAfterEnabled(false);
    setLastAfterDay(String(now.getDate()));
    setLastAfterMonth(String(now.getMonth() + 1));
    setLastAfterYear(String(now.getFullYear()));
    setLastBeforeEnabled(false);
    setLastBeforeDay(String(now.getDate()));
    setLastBeforeMonth(String(now.getMonth() + 1));
    setLastBeforeYear(String(now.getFullYear()));
    setLastNever(false);

    setModifiedAfterEnabled(false);
    setModifiedAfterDay(String(now.getDate()));
    setModifiedAfterMonth(String(now.getMonth() + 1));
    setModifiedAfterYear(String(now.getFullYear()));
    setModifiedBeforeEnabled(false);
    setModifiedBeforeDay(String(now.getDate()));
    setModifiedBeforeMonth(String(now.getMonth() + 1));
    setModifiedBeforeYear(String(now.getFullYear()));
    setModifiedNever(false);

    setCreatedAfterEnabled(false);
    setCreatedAfterDay(String(now.getDate()));
    setCreatedAfterMonth(String(now.getMonth() + 1));
    setCreatedAfterYear(String(now.getFullYear()));
    setCreatedBeforeEnabled(false);
    setCreatedBeforeDay(String(now.getDate()));
    setCreatedBeforeMonth(String(now.getMonth() + 1));
    setCreatedBeforeYear(String(now.getFullYear()));

    setIdentificationOp(''); setIdentificationTouched(false);
    setMnetProviderOp(''); setMnetProviderTouched(false);
    setIdNumberOp('contains'); setIdNumberVal('');
    setInstitutionOp('contains'); setInstitutionVal('');
    setDepartmentOp('contains'); setDepartmentVal('');
    setLastIpOp('contains'); setLastIpVal('');

    setSelectedFiltersToRemove([]);
    setShowMoreFilters(false);

    // Clear uncontrolled DOM inputs
    try {
      if (filterInputRef.current) filterInputRef.current.value = '';
      if (lastNameRef.current) lastNameRef.current.value = '';
      if (firstNameRef.current) firstNameRef.current.value = '';
      if (usernameRef.current) usernameRef.current.value = '';
      if (emailRef.current) emailRef.current.value = '';
      if (cityRef.current) cityRef.current.value = '';
      if (idNumberRef.current) idNumberRef.current.value = '';
      if (institutionRef.current) institutionRef.current.value = '';
      if (departmentRef.current) departmentRef.current.value = '';
      if (lastIpRef.current) lastIpRef.current.value = '';
      if (profileFieldRef.current) profileFieldRef.current.value = '';
      if (roleFieldRef.current) roleFieldRef.current.value = '';
      if (cohortRef.current) cohortRef.current.value = '';
    } catch (e) {
      // ignore DOM clearing errors
    }
  }, []);
 
  const staticStyles = useMemo(() => `
          .moodle-filter select.moodle-search-input {
            border: 2px solid #39A900;
            background: #fff;
            color: #222;
            border-radius: 8px;
            font-size: 15px;
            padding: 8px 16px;
            box-shadow: none;
            outline: none;
            transition: border 0.2s, box-shadow 0.2s;
          }
          .moodle-filter select.moodle-search-input:focus {
            border-color: #39A900;
            box-shadow: 0 0 0 2px rgba(57,169,0,0.2);
          }
          /* Tamaño solicitado para los select en la vista de usuarios.
             Use la misma especificidad que la regla global y !important para
             asegurarnos de que esta medida prevalezca. */
          /* Increase spacing between filter rows to 18px */
          .moodle-filter .filter-row {
            margin-bottom: 18px !important;
          }
          /* Set font size for all filter elements to 15px */
          .moodle-filter .filter-row,
          .moodle-filter .filter-row label,
          .moodle-filter .filter-row .moodle-select,
          .moodle-filter .filter-row .moodle-search-input {
            font-size: 15px !important;
          }
          /* Active filters list: enforce the same font-size for visibility */
          #active-filters-body,
          #active-filters-body label,
          #active-filters-body span {
            font-size: 15px !important;
          }
          .moodle-filter .filter-row .moodle-select {
            width: 134.617px !important;
            height: 36.5px !important;
            min-width: 134.617px !important;
            min-height: 36.5px !important;
            border-radius: 12px !important;
            box-sizing: border-box;
          }
          /* Country select: override width only for País select */
          .moodle-filter .filter-row .moodle-select.country-select {
            width: 343.133px !important;
            min-width: 343.133px !important;
          }
          /* Identificación select: requested fixed width */
          .moodle-filter .filter-row .moodle-select.identification-select {
            width: 379.767px !important;
            min-width: 379.767px !important;
          }
          /* MNET provider select: requested fixed width */
          .moodle-filter .filter-row .moodle-select.mnet-provider-select {
            width: 192.133px !important;
            min-width: 192.133px !important;
          }
          /* More specific selector to win against operator-specific rules */
          .moodle-filter .filter-row .moodle-select.operator-select.identification-select {
            width: 379.767px !important;
            min-width: 379.767px !important;
          }
          /* Profile field name select: specific width requested */
          .moodle-filter .filter-row .moodle-select.profile-field-name-select {
            width: 199.683px !important;
            min-width: 199.683px !important;
          }
          /* Operator select for País should be slightly wider */
          .moodle-filter .filter-row .moodle-select.operator-select {
            width: 139.633px !important;
            min-width: 139.633px !important;
          }
          /* Stronger selector to enforce width for the profile-field-name when it also has operator-select */
          .moodle-filter .filter-row .moodle-select.operator-select.profile-field-name-select {
            width: 199.683px !important;
            min-width: 199.683px !important;
          }
          /* Enforce width for the first select of the 'Rol de curso' filter */
          .moodle-filter .filter-row .moodle-select.operator-select.role-field-one-select {
            width: 258.067px !important;
            min-width: 258.067px !important;
          }
          /* Enforce width for the second select of the 'Rol de curso' filter */
          .moodle-filter .filter-row .moodle-select.operator-select.role-field-two-select {
            width: 306.417px !important;
            min-width: 306.417px !important;
          }
          /* Very specific rule for MNET provider when operator-select is present */
          .moodle-filter .filter-row .moodle-select.operator-select.mnet-provider-select {
            width: 192.133px !important;
            min-width: 192.133px !important;
          }
          /* Ensure Número de ID operator select uses the small default width */
          .moodle-filter .filter-row .moodle-select.operator-select.idnumber-select {
            width: 134.617px !important;
            min-width: 134.617px !important;
          }
          /* Enforce width for the 'Departamento' operator select */
          .moodle-filter .filter-row .moodle-select.operator-select.department-select {
            width: 134.617px !important;
            min-width: 134.617px !important;
          }
          /* Enforce width for the 'Última dirección IP' operator select */
          .moodle-filter .filter-row .moodle-select.operator-select.lastip-select {
            width: 134.617px !important;
            min-width: 134.617px !important;
          }
          /* Enforce width for the 'Institución' operator select */
          .moodle-filter .filter-row .moodle-select.operator-select.institution-select {
            width: 134.617px !important;
            min-width: 134.617px !important;
          }
          .moodle-filter .filter-row .moodle-select:focus {
            outline: none !important;
            box-shadow: 0 0 0 2px rgba(57,169,0,0.12) !important;
          }
          /* Enforce width for the 'Rol del sistema' select */
          .moodle-filter .filter-row .moodle-select.operator-select.system-role-select {
            width: 157.983px !important;
            min-width: 157.983px !important;
          }
          /* Enforce width for the profile field operator select (second select) */
          .moodle-filter .filter-row .moodle-select.operator-select.profile-field-op-select {
            width: 148px !important;
            min-width: 148px !important;
          }
            /* Tamaño solicitado para los inputs de búsqueda en la vista de usuarios */
            .moodle-filter .filter-row .moodle-search-input {
              width: 218px !important;
              height: 36.5px !important;
              min-width: 218px !important;
              min-height: 36.5px !important;
              box-sizing: border-box;
              padding: 6px 12px !important;
            }
            .moodle-filter .filter-row .moodle-search-input:focus {
              outline: none !important;
              box-shadow: 0 0 0 2px rgba(57,169,0,0.12) !important;
            }
            /* Show-more button style + hover (cover both .filter-row and .filter-body placements) */
            .moodle-filter .filter-row .show-more-btn,
            .moodle-filter .filter-body .show-more-btn {
              background: none !important;
              border: none !important;
              color: #39A900 !important;
              font-size: .9375rem !important;
              text-decoration: none !important;
              cursor: pointer !important;
              margin-left: 0 !important;
              padding-left: 0 !important;
            }
            .moodle-filter .filter-row .show-more-btn:hover,
            .moodle-filter .filter-body .show-more-btn:hover {
              color: #1f5d00 !important;
            }

            /* Use the bulk-action select caret (double-triangle) used in EnrolledUsersView.
              Apply it only to wrappers that have the bulk-caret class so we don't
              accidentally hide other pseudo-elements. */
            .moodle-filter .filter-row .select-wrapper.bulk-caret {
              position: relative;
              display: inline-block;
            }
            .moodle-filter .filter-row .select-wrapper.bulk-caret::after {
              content: "";
              position: absolute;
              right: 10px;
              top: 50%;
              transform: translateY(-50%);
              width: 16px;
              height: 18px;
              pointer-events: none;
              /* remove any border-based triangle from the original rule */
              border-left: 0 !important;
              border-right: 0 !important;
              border-top: 0 !important;
              background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 14 16'%3e%3cpath fill='%23343a40' d='M7 3L3 7h8L7 3z'/%3e%3cpath fill='%23343a40' d='M7 13l4-4H3l4 4z'/%3e%3c/svg%3e") !important;
              background-size: 16px 18px;
              background-repeat: no-repeat;
            }
            .moodle-filter .filter-row .select-wrapper.bulk-caret .moodle-select {
              padding-right: 44px !important; /* reserve space for custom caret */
              background-image: none !important;
              appearance: none !important;
              -webkit-appearance: none !important;
              -moz-appearance: none !important;
              color: #495057 !important;
            }
            /* Primer acceso selects should use same text color as other selects */
            .moodle-filter .filter-row .moodle-select.primer-acceso-select {
              color: #495057 !important;
            }
            /* Day select small fixed width for Primer acceso */
            .moodle-filter .filter-row .moodle-select.primer-acceso-day-select {
              width: 58.7167px !important;
              min-width: 58.7167px !important;
              padding-left: 10px !important;
              padding-right: 12px !important;
              text-align: left !important;
              color: #495057 !important;
              background: #fff !important;
              box-sizing: border-box !important;
            }
            /* For the País select reduce right padding so the selected text fits */
            .moodle-filter .filter-row .select-wrapper.bulk-caret .moodle-select.country-select {
              padding-right: 20px !important;
            }
            /* Reduce padding-right for the operator select so its label fits (e.g. 'cualquier valor') */
            .moodle-filter .filter-row .select-wrapper.bulk-caret .moodle-select.operator-select {
              padding-right: 12px !important;
            }
            /* Ensure option text also uses the desired color */
            .moodle-filter .filter-row .select-wrapper.bulk-caret .moodle-select option {
              color: #495057 !important;
            }
            /* Hide IE/Edge native expand glyph just in case */
            .moodle-filter .filter-row .moodle-select::-ms-expand { display: none !important; }

            /* Pager controls: force exact box size for pager buttons */
            .moodle-pager button {
              width: 34.35px !important;
              min-width: 34.35px !important;
              height: 36.75px !important;
              line-height: 36.75px !important;
              padding: 0 !important;
              box-sizing: border-box !important;
              font-size: .9375rem !important;
            }
            .moodle-pager button:first-child {
              border-radius: .25rem 0 0 .25rem !important;
              height: 36.75px !important;
            }
            .moodle-pager button:last-child {
              border-radius: 0 .25rem .25rem 0 !important;
              height: 36.75px !important;
            }
            /* Table headers and cells: enforce requested font size */
            .table-moodle th, .table-moodle td {
              font-size: .9375rem !important;
            }
            /* Icons inside action buttons: ensure consistent size */
            .table-icons .icon,
            .icon-btn .icon {
              font-size: 16px !important;
              line-height: 1 !important;
            }
            /* Ensure the 'Suspender cuenta de usuario' (eye) button shows green by default */
            .table-icons .icon-btn.gray,
            .table-icons .icon-btn.gray .icon {
              color: var(--moodle-green) !important;
            }
            /* Ensure icons animate color changes smoothly */
            .table-icons .icon,
            .table-icons .icon-btn {
              transition: color 120ms ease, opacity 120ms ease, transform 120ms ease;
            }
            /* Hover: change color of action icons (borrar, editar, suspender) */
            .table-icons .icon-btn:hover .icon,
            .table-icons .icon-btn:hover,
            /* Extra specific rule so .icon-btn.gray hover can override the default !important color */
            .table-icons .icon-btn.gray:hover .icon,
            .table-icons .icon-btn.gray:hover {
              color: #1f5d00 !important;
            }
            /* Force table cell height to 33.6px and center content vertically */
            .table-moodle th, .table-moodle td {
              height: 33.6px !important;
              min-height: 33.6px !important;
              line-height: 33.6px !important;
              padding-top: 0 !important;
              padding-bottom: 0 !important;
              vertical-align: middle !important;
              box-sizing: border-box !important;
            }
        `, []);

    // Memoize commonly reused option lists to avoid rebuilding them on every render
    const days = useMemo(() => Array.from({ length: 31 }, (_, i) => i + 1), []);
    const months = useMemo(() => ['enero','febrero','marzo','abril','mayo','junio','julio','agosto','septiembre','octubre','noviembre','diciembre'], []);
    const years = useMemo(() => Array.from({ length: 2050 - 1900 + 1 }, (_, i) => 1900 + i), []);
  const fetchUsers = useCallback(async (silent = false) => {
    if (!silent) setLoading(true);
    // generate a request id to ignore stale responses
    const reqId = ++currentRequestId.current;
    try {
      // Si existe un filtro 'role' marcado como no encontrado, no mostrar datos
      if (filters.some(f => f.field === 'role' && f.isNotFound)) {
        setUsers([]);
        setTotal(0);
        setLoading(false);
        return;
      }
      // Predicados por operador
      const opMatch = (value, op, needle) => {
        const v = `${value ?? ""}`.toLowerCase();
        const n = `${needle ?? ""}`.toLowerCase();
        switch (op) {
          case "equals": return v === n;
          case "notequals": return v !== n;
            case "defined": return (v !== "" && v.trim() !== "");
            case "notdefined": return (v === "" || v.trim() === "");
          case "any": return (v !== "" && v.trim() !== "");
          case "starts": return v.startsWith(n);
          case "ends": return v.endsWith(n);
          case "notcontains": return !v.includes(n);
          case "empty": return (v === "" || v.trim() === "");
          case "contains":
          default: return v.includes(n);
        }
      };

      // Separar filtros por tipo
      const fieldSet = new Set(["lastname", "firstname", "username", "email", "city", "country", "confirmed", "suspended", "profile", "enrolled", "systemrole", "cohort", "firstaccess_after", "firstaccess_before", "lastaccess_after", "lastaccess_before", "lastaccess_never", "timecreated_after", "timecreated_before", "timemodified_after", "timemodified_before", "timemodified_never",
        // Add previously-ignored fields so they are queried via getUsersByField
        "identification", "mnet_provider", "idnumber", "institution", "department", "lastip"]);
      const fieldFilters = filters.filter((f) => fieldSet.has(f.field));
      const fullnameFilters = filters.filter((f) => f.field === "fullname");

      const normalizeList = (raw) => {
        if (Array.isArray(raw)) return raw;
        if (raw?.items && Array.isArray(raw.items)) return raw.items;
        if (raw?.users && Array.isArray(raw.users)) return raw.users;
        if (raw?.data && Array.isArray(raw.data)) return raw.data;
        if (raw?.result && Array.isArray(raw.result)) return raw.result;
        for (const k in raw || {}) { if (Array.isArray(raw[k])) return raw[k]; }
        return [];
      };

      // baseDataset not used in current implementation

      if (fieldFilters.length > 0) {
        // Agrupar por campo y obtener desde by-field en paralelo
        const groups = fieldFilters.reduce((acc, f) => {
          acc[f.field] = acc[f.field] || [];
          acc[f.field].push(f);
          return acc;
        }, {});

        const entries = Object.entries(groups);
        const results = await Promise.all(entries.map(async ([field, flist]) => {
          const values = Array.from(new Set(flist.map((x) => x.value)));
          const raw = await getUsersByField(field, values);
          return [field, normalizeList(raw)];
        }));

        // Combinar por lógica seleccionada (AND/OR) usando ids
        const lists = results.map(([, list]) => list);

        const idOf = (u) => u.id ?? `${u.userid ?? u.ID ?? JSON.stringify(u)}`; // heurística

        let combined;
        if (filterLogic === "AND") {
          // Intersección por id
          const idSets = lists.map((lst) => new Set(lst.map(idOf)));
          const intersectionIds = Array.from(idSets[0] || []).filter((id) => idSets.every((s) => s.has(id)));
          // Tomar desde el primer listado
          const refList = lists[0] || [];
          combined = refList.filter((u) => intersectionIds.includes(idOf(u)));
        } else {
          // Unión por id
          const map = new Map();
          lists.forEach((lst) => lst.forEach((u) => map.set(idOf(u), u)));
          combined = Array.from(map.values());
        }

        // Aplicar todos los predicados de filtros (campo + fullname) según lógica
        const allFilters = [...fieldFilters, ...fullnameFilters];
        const matchesFilter = (u, f) => {
          if (f.field === "fullname") {
            const full = `${u.firstname ?? ""} ${u.lastname ?? ""}`;
            return opMatch(full, f.op, f.value);
          }
          return opMatch(u[f.field], f.op, f.value);
        };

        const finalFiltered = combined.filter((u) => {
          if (allFilters.length === 0) return true;
          if (filterLogic === "AND") return allFilters.every((f) => matchesFilter(u, f));
          return allFilters.some((f) => matchesFilter(u, f));
        });

        // Paginar en cliente
        const totalCount = finalFiltered.length;
        const start = (currentPage - 1) * perPage;
        const pageSlice = finalFiltered.slice(start, start + perPage);
        // cache the result for this filter set
        try { resultsCache.current.set(JSON.stringify(filters), { users: finalFiltered, total: totalCount, fullTotal: fullTotal ?? totalCount }); } catch (e) { /* ignore cache errors */ }
        // ignore if a newer request started
        if (reqId !== currentRequestId.current) return;
        setUsers(pageSlice);
        setTotal(totalCount);

        // Inicializar usuarios suspendidos
        const suspended = new Set(finalFiltered.filter(u => u.suspended === 1).map(u => u.id));
        setHiddenUsers(suspended);
      } else {
        // Sin filtros por campo: usar búsqueda general. Usamos únicamente el filtro
        // `fullname` si fue añadido mediante "Añadir filtro"; no realizamos
        // búsquedas en vivo mientras el usuario escribe en el campo.
        let queryString = fullnameFilters[0]?.value || "";
        const data = await getUsers(queryString, currentPage, perPage, 0);
        let usersList = normalizeList(data);

        if (fullnameFilters.length > 0) {
          const combinedPred = (u) => {
            if (filterLogic === "AND") {
              return fullnameFilters.every((f) => opMatch(`${u.firstname ?? ""} ${u.lastname ?? ""}`, f.op, f.value));
            }
            return fullnameFilters.some((f) => opMatch(`${u.firstname ?? ""} ${u.lastname ?? ""}`, f.op, f.value));
          };
          usersList = usersList.filter(combinedPred);
        }

        try { resultsCache.current.set(JSON.stringify(filters), { users: usersList, total: computedTotal, fullTotal: fullTotal ?? computedTotal }); } catch (e) { /* ignore */ }
        if (reqId !== currentRequestId.current) return;
        setUsers(usersList);
        // Si el backend no nos da total, asumimos longitud
        const computedTotal = Array.isArray(data?.total) ? data.total : (data?.total ?? usersList.length);
        if (reqId !== currentRequestId.current) return;
        setTotal(computedTotal);
        // Capture full dataset total only cuando no hay queryString (búsqueda en curso).
        // Esto evita que el contador total completo sea sobrescrito mientras el usuario
        // escribe en el campo de búsqueda (live-search).
        if (!queryString) {
          setFullTotal((prev) => prev ?? computedTotal);
        }

        // Inicializar usuarios suspendidos (funcionalidad nueva de darwin)
        const suspended = new Set(usersList.filter(u => u.suspended === 1).map(u => u.id));
        setHiddenUsers(suspended);
      }
    } catch (error) {
      console.error("Error al cargar usuarios:", error);
      setUsers([]);
      setTotal(0);
    } finally {
      if (!silent) setLoading(false);
    }
  }, [currentPage, perPage, filterLogic, filtersVersion]);

  // Run fetchUsers on mount and whenever filtersVersion/currentPage/perPage change.
  useEffect(() => {
    fetchUsers();
  }, [fetchUsers]);

  // Scroll to top when entering the view and when page/filters change.
  useEffect(() => {
    // On initial mount, jump to top (no animation) so view behaves like "first entry"
    scrollToTop(false);
  }, []);

  useEffect(() => {
    // When page or filters change (add/remove/clear) bring header into view.
    scrollToTop(true);
  }, [currentPage, filtersVersion]);

  // When the user changes page, close the "mostrar más" advanced filters
  // and reset their UI controls to initial defaults so the panel opens fresh.
  useEffect(() => {
    // central reset for advanced controls when page changes
    resetAdvancedFilters();
  }, [currentPage]);

    // Fetch course categories to populate the role second select
    useEffect(() => {
      let mounted = true;
      (async () => {
        try {
          const res = await coursesService.getCategories();
          const cats = Array.isArray(res) ? res : (res?.items || res?.categories || []);
          if (mounted) setRoleCategories(cats);
        } catch (err) {
          console.error('Error fetching categories for role filter', err);
          if (mounted) setRoleCategories([]);
        }
      })();
      return () => { mounted = false; };
    }, []);

  // Ensure countryVal has a sensible default when operator requires a value
  useEffect(() => {
    if ((countryOp === 'equals' || countryOp === 'notequals') && (!countryVal || countryVal.trim() === '')) {
      setCountryVal('Colombia');
    }
  }, [countryOp, countryVal]);

  const handleAddFilter = async () => {
    // Do not set loading unless we're actually adding filters (avoid needless fetch)
    const newFilters = [];
    const currentFilterValue = filterInputRef.current ? (filterInputRef.current.value || '').trim() : filterValue.trim();
    if (currentFilterValue) {
      newFilters.push({ field: "fullname", op: filterOp, value: currentFilterValue });
    }
    const lastNameCurrent = lastNameRef.current ? (lastNameRef.current.value || '').trim() : lastNameVal.trim();
    if (lastNameCurrent) {
      newFilters.push({ field: "lastname", op: lastNameOp, value: lastNameCurrent });
    }
    const firstNameCurrent = firstNameRef.current ? (firstNameRef.current.value || '').trim() : firstNameVal.trim();
    if (firstNameCurrent) {
      newFilters.push({ field: "firstname", op: firstNameOp, value: firstNameCurrent });
    }
    const usernameCurrent = usernameRef.current ? (usernameRef.current.value || '').trim() : usernameVal.trim();
    if (usernameCurrent) {
      newFilters.push({ field: "username", op: usernameOp, value: usernameCurrent });
    }
    const emailCurrent = emailRef.current ? (emailRef.current.value || '').trim() : emailVal.trim();
    if (emailCurrent) {
      newFilters.push({ field: "email", op: emailOp, value: emailCurrent });
    }
    const cityCurrent = cityRef.current ? (cityRef.current.value || '').trim() : cityVal.trim();
    if (cityCurrent) {
      newFilters.push({ field: "city", op: cityOp, value: cityCurrent });
    }
    if ((countryOp === 'equals' || countryOp === 'notequals') && countryVal.trim()) {
      newFilters.push({ field: "country", op: countryOp, value: countryVal.trim() });
    }
    if (confirmedOp !== 'any') {
      newFilters.push({ field: "confirmed", op: "equals", value: confirmedOp });
    }
    if (suspendedOp !== 'any') {
      newFilters.push({ field: "suspended", op: "equals", value: suspendedOp });
    }
    const profileCurrent = profileFieldRef.current ? (profileFieldRef.current.value || '').trim() : profileFieldVal.trim();
    if (profileCurrent) {
      newFilters.push({ field: "profile", subfield: profileFieldName, op: profileFieldOp, value: profileCurrent });
    }
    // Allow adding a role filter when either select is chosen even if the text input is empty
    const roleCurrent = roleFieldRef.current ? (roleFieldRef.current.value || '').trim() : roleFieldVal.trim();
    if (roleFieldOne !== 'any' || roleFieldTwo !== 'any' || roleCurrent) {
      // If user entered a roleFieldVal, validate it exists (search courses)
      let skipNormalRolePush = false;
      if (roleCurrent) {
        try {
          const q = roleCurrent;
          const res = await coursesService.search(q);
          const list = Array.isArray(res) ? res : (res?.items || res?.courses || []);
          // If category selected, filter by category id/name
          const filtered = (roleFieldTwo && roleFieldTwo !== 'any') ? list.filter(c => {
            return String(c.category ?? c.categoryid ?? c.category_id ?? c.catid ?? '') === String(roleFieldTwo) || String(c.categoryname ?? c.catname ?? c.name ?? '') === String(roleFieldTwo);
          }) : list;
          if (!filtered || filtered.length === 0) {
            // push an active filter that indicates the course was not found
            newFilters.push({ field: "role", subfield1: roleFieldOne, subfield2: roleFieldTwo, value: q, isNotFound: true });
            skipNormalRolePush = true;
          }
        } catch (err) {
          console.error('Error buscando curso para filtro role:', err);
          // In case of error, allow user to proceed without marking not-found
        }
      }
      if (!skipNormalRolePush) {
        newFilters.push({ field: "role", subfield1: roleFieldOne, subfield2: roleFieldTwo, value: roleCurrent });
      }
    }
    if (enrolledOp !== 'any') {
      newFilters.push({ field: 'enrolled', op: 'equals', value: enrolledOp });
    }
    if (systemRoleOp !== 'any') {
      newFilters.push({ field: 'systemrole', op: 'equals', value: systemRoleOp });
    }
    // Only add cohort filter if a value was provided or the operator does not require a value (e.g. 'empty')
    const cohortCurrent = cohortRef.current ? (cohortRef.current.value || '').trim() : cohortVal.trim();
    if (cohortOp === 'empty' || cohortCurrent) {
      newFilters.push({ field: 'cohort', op: cohortOp, value: cohortCurrent });
    }
    // Only add identification filter when the user changed the select explicitly
    if (identificationOp !== '' && identificationTouched) {
      newFilters.push({ field: 'identification', op: 'equals', value: identificationOp });
    }
    // Only add mnet provider filter when user changed the select explicitly
    if (mnetProviderOp !== '' && mnetProviderTouched) {
      newFilters.push({ field: 'mnet_provider', op: 'equals', value: mnetProviderOp });
    }
    // Número de ID: only add if operator is 'empty' or a value was provided
    const idNumberCurrent = idNumberRef.current ? (idNumberRef.current.value || '').trim() : idNumberVal.trim();
    if (idNumberOp === 'empty' || idNumberCurrent) {
      newFilters.push({ field: 'idnumber', op: idNumberOp, value: idNumberCurrent });
    }
    // Institución: only add if operator is 'empty' or a value was provided
    const institutionCurrent = institutionRef.current ? (institutionRef.current.value || '').trim() : institutionVal.trim();
    if (institutionOp === 'empty' || institutionCurrent) {
      newFilters.push({ field: 'institution', op: institutionOp, value: institutionCurrent });
    }
    // Departamento: only add if operator is 'empty' or a value was provided
    const departmentCurrent = departmentRef.current ? (departmentRef.current.value || '').trim() : departmentVal.trim();
    if (departmentOp === 'empty' || departmentCurrent) {
      newFilters.push({ field: 'department', op: departmentOp, value: departmentCurrent });
    }
    // Última dirección IP: only add if operator is 'empty' or a value was provided
    const lastIpCurrent = lastIpRef.current ? (lastIpRef.current.value || '').trim() : lastIpVal.trim();
    if (lastIpOp === 'empty' || lastIpCurrent) {
      newFilters.push({ field: 'lastip', op: lastIpOp, value: lastIpCurrent });
    }
    // (duplicate old check removed) -- last IP already handled via ref-aware logic above
    if (firstAfterEnabled) {
      const m = String(firstAfterMonth).padStart(2,'0');
      const d = String(firstAfterDay).padStart(2,'0');
      newFilters.push({ field: 'firstaccess_after', value: `${firstAfterYear}-${m}-${d}` });
    }
    if (firstBeforeEnabled) {
      const m2 = String(firstBeforeMonth).padStart(2,'0');
      const d2 = String(firstBeforeDay).padStart(2,'0');
      newFilters.push({ field: 'firstaccess_before', value: `${firstBeforeYear}-${m2}-${d2}` });
    }
    // Último acceso
    if (lastAfterEnabled) {
      const m = String(lastAfterMonth).padStart(2,'0');
      const d = String(lastAfterDay).padStart(2,'0');
      newFilters.push({ field: 'lastaccess_after', value: `${lastAfterYear}-${m}-${d}` });
    }
    if (lastBeforeEnabled) {
      const m2 = String(lastBeforeMonth).padStart(2,'0');
      const d2 = String(lastBeforeDay).padStart(2,'0');
      newFilters.push({ field: 'lastaccess_before', value: `${lastBeforeYear}-${m2}-${d2}` });
    }
    if (lastNever) {
      newFilters.push({ field: 'lastaccess_never', value: 'true' });
    }
    // Última modificación
    if (modifiedAfterEnabled) {
      const m = String(modifiedAfterMonth).padStart(2,'0');
      const d = String(modifiedAfterDay).padStart(2,'0');
      newFilters.push({ field: 'timemodified_after', value: `${modifiedAfterYear}-${m}-${d}` });
    }
    if (modifiedBeforeEnabled) {
      const m2 = String(modifiedBeforeMonth).padStart(2,'0');
      const d2 = String(modifiedBeforeDay).padStart(2,'0');
      newFilters.push({ field: 'timemodified_before', value: `${modifiedBeforeYear}-${m2}-${d2}` });
    }
    if (modifiedNever) {
      newFilters.push({ field: 'timemodified_never', value: 'true' });
    }
    // Tiempo de creado
    if (createdAfterEnabled) {
      const m = String(createdAfterMonth).padStart(2,'0');
      const d = String(createdAfterDay).padStart(2,'0');
      newFilters.push({ field: 'timecreated_after', value: `${createdAfterYear}-${m}-${d}` });
    }
    if (createdBeforeEnabled) {
      const m2 = String(createdBeforeMonth).padStart(2,'0');
      const d2 = String(createdBeforeDay).padStart(2,'0');
      newFilters.push({ field: 'timecreated_before', value: `${createdBeforeYear}-${m2}-${d2}` });
    }
    if (newFilters.length > 0) {
      // show loading immediately to hide pager while fetching new results
      setLoading(true);
      setFilters([...filters, ...newFilters]);
      setFiltersVersion(v => v + 1);
      // central reset for advanced controls
      resetAdvancedFilters();
      setCurrentPage(1);
    }
  };

  const openCalendarFor = (which, anchorEl) => {
    // support primer acceso ('after'/'before') and último acceso ('last_after'/'last_before')
    // If the calendar is already open for the same control, toggle (close) it.
    if (calendarState.visible && calendarState.which === which) {
      closeCalendar();
      return;
    }
    // Do NOT open the calendar if the corresponding checkbox is not enabled.
    // The calendar should only be usable when the user has explicitly enabled the date filter.
    const enabledMap = {
      after: firstAfterEnabled,
      before: firstBeforeEnabled,
      last_after: lastAfterEnabled,
      last_before: lastBeforeEnabled,
      created_after: createdAfterEnabled,
      created_before: createdBeforeEnabled,
      modified_after: modifiedAfterEnabled,
      modified_before: modifiedBeforeEnabled,
    };
    if (enabledMap[which] === false) return;
    let el;
    if (anchorEl) el = anchorEl;
    else if (which === 'after') el = afterCalendarIconRef.current;
    else if (which === 'before') el = beforeCalendarIconRef.current;
    else if (which === 'last_after') el = lastAfterCalendarIconRef.current;
    else if (which === 'last_before') el = lastBeforeCalendarIconRef.current;
    if (!el) return;
    const rect = el.getBoundingClientRect();
    const vw = window.innerWidth;
    const vh = window.innerHeight;
    const preferBelow = rect.bottom + 260 < vh;
    const placement = preferBelow ? 'bottom' : 'top';
    let month, year;
    if (which === 'after') { month = Number(firstAfterMonth) - 1; year = Number(firstAfterYear); }
    else if (which === 'before') { month = Number(firstBeforeMonth) - 1; year = Number(firstBeforeYear); }
    else if (which === 'last_after') { month = Number(lastAfterMonth) - 1; year = Number(lastAfterYear); }
    else if (which === 'last_before') { month = Number(lastBeforeMonth) - 1; year = Number(lastBeforeYear); }
    else if (which === 'created_after') { month = Number(createdAfterMonth) - 1; year = Number(createdAfterYear); }
    else if (which === 'created_before') { month = Number(createdBeforeMonth) - 1; year = Number(createdBeforeYear); }
    else if (which === 'modified_after') { month = Number(modifiedAfterMonth) - 1; year = Number(modifiedAfterYear); }
    else { month = Number(modifiedBeforeMonth) - 1; year = Number(modifiedBeforeYear); }
    // If computed month/year are invalid (e.g. empty strings), default to today's date
    if (!Number.isFinite(month) || isNaN(month) || month < 0 || month > 11 || !Number.isFinite(year) || isNaN(year)) {
      const now = new Date();
      month = now.getMonth();
      year = now.getFullYear();
    }
    setCalendarState({ visible: true, which, x: rect.left + rect.width / 2, y: placement === 'bottom' ? rect.bottom : rect.top, placement, month, year });
  };

  const closeCalendar = () => setCalendarState((s) => ({ ...s, visible: false }));

  const handleCalendarSelect = (which, y, m, d) => {
    if (which === 'after') {
      setFirstAfterYear(String(y));
      setFirstAfterMonth(String(m + 1));
      setFirstAfterDay(String(d));
      setFirstAfterEnabled(true);
    } else if (which === 'before') {
      setFirstBeforeYear(String(y));
      setFirstBeforeMonth(String(m + 1));
      setFirstBeforeDay(String(d));
      setFirstBeforeEnabled(true);
    } else if (which === 'last_after') {
      setLastAfterYear(String(y));
      setLastAfterMonth(String(m + 1));
      setLastAfterDay(String(d));
      setLastAfterEnabled(true);
    } else if (which === 'last_before') {
      setLastBeforeYear(String(y));
      setLastBeforeMonth(String(m + 1));
      setLastBeforeDay(String(d));
      setLastBeforeEnabled(true);
    } else if (which === 'created_after') {
      setCreatedAfterYear(String(y));
      setCreatedAfterMonth(String(m + 1));
      setCreatedAfterDay(String(d));
      setCreatedAfterEnabled(true);
    } else if (which === 'created_before') {
      setCreatedBeforeYear(String(y));
      setCreatedBeforeMonth(String(m + 1));
      setCreatedBeforeDay(String(d));
      setCreatedBeforeEnabled(true);
    } else if (which === 'modified_after') {
      setModifiedAfterYear(String(y));
      setModifiedAfterMonth(String(m + 1));
      setModifiedAfterDay(String(d));
      setModifiedAfterEnabled(true);
    } else if (which === 'modified_before') {
      setModifiedBeforeYear(String(y));
      setModifiedBeforeMonth(String(m + 1));
      setModifiedBeforeDay(String(d));
      setModifiedBeforeEnabled(true);
    }
    closeCalendar();
  };

  useEffect(() => {
    if (!calendarState.visible) return;
    const onDocClick = (ev) => {
      const pop = document.getElementById('zajuna-calendar-popup');
      if (!pop) return;
      if (!pop.contains(ev.target)
        && !afterCalendarIconRef.current?.contains(ev.target)
        && !beforeCalendarIconRef.current?.contains(ev.target)
        && !lastAfterCalendarIconRef.current?.contains(ev.target)
        && !lastBeforeCalendarIconRef.current?.contains(ev.target)
        && !createdAfterCalendarIconRef.current?.contains(ev.target)
        && !createdBeforeCalendarIconRef.current?.contains(ev.target)
        && !modifiedAfterCalendarIconRef.current?.contains(ev.target)
        && !modifiedBeforeCalendarIconRef.current?.contains(ev.target)) {
        closeCalendar();
      }
    };
    const onKey = (ev) => { if (ev.key === 'Escape') closeCalendar(); };
    document.addEventListener('mousedown', onDocClick);
    document.addEventListener('keydown', onKey);
    return () => { document.removeEventListener('mousedown', onDocClick); document.removeEventListener('keydown', onKey); };
  }, [calendarState.visible]);

  function CalendarPopup({ state }) {
    if (!state.visible) return null;
    const w = 260;
    const h = 260;
    const left = Math.max(8, Math.min(window.innerWidth - w - 8, Math.round(state.x - w / 2)));
    const top = state.placement === 'bottom' ? Math.round(state.y + 6) : Math.round(state.y - h - 6);
    const month = state.month;
    const year = state.year;

    const firstOfMonth = new Date(year, month, 1);
    const startDay = firstOfMonth.getDay();
    const adjustedStart = (startDay + 6) % 7;
    const daysInMonth = new Date(year, month + 1, 0).getDate();
    const prevDays = adjustedStart;
    const cells = [];
    const prevMonthLast = new Date(year, month, 0).getDate();
    for (let i = prevMonthLast - prevDays + 1; i <= prevMonthLast; i++) cells.push({ d: i, inMonth: false });
    for (let d = 1; d <= daysInMonth; d++) cells.push({ d, inMonth: true });
    while (cells.length % 7 !== 0) cells.push({ d: cells.length - daysInMonth - prevDays + 1, inMonth: false });

    const prevMonth = () => {
      let m = month - 1; let y = year;
      if (m < 0) { m = 11; y -= 1; }
      setCalendarState((s) => ({ ...s, month: m, year: y }));
    };
    const nextMonth = () => {
      let m = month + 1; let y = year;
      if (m > 11) { m = 0; y += 1; }
      setCalendarState((s) => ({ ...s, month: m, year: y }));
    };

    const weekdayLabels = ['lun','mar','mié','jue','vie','sáb','dom'];

    return (
      <div id="zajuna-calendar-popup" className="zajuna-cal" style={{ position: 'fixed', left, top, width: w, height: h, zIndex: 9999 }}>
        <div className="zajuna-cal__header">
          <button type="button" className="zajuna-cal__nav" onClick={prevMonth} aria-label="Mes anterior">◀</button>
          <div className="zajuna-cal__title">{`${['January','February','March','April','May','June','July','August','September','October','November','December'][month]} ${year}`}</div>
          <button type="button" className="zajuna-cal__nav" onClick={nextMonth} aria-label="Mes siguiente">▶</button>
        </div>
        <div className="zajuna-cal__weekdays">
          {weekdayLabels.map((w) => <div key={w} className="zajuna-cal__weekday">{w}</div>)}
        </div>
        <div className="zajuna-cal__grid">
          {cells.map((c, idx) => {
            const isDisabled = !c.inMonth;
            let selectedDay = null, selectedMonth = null, selectedYear = null;
            if (state.which === 'after') { selectedDay = Number(firstAfterDay); selectedMonth = Number(firstAfterMonth) - 1; selectedYear = Number(firstAfterYear); }
            else if (state.which === 'before') { selectedDay = Number(firstBeforeDay); selectedMonth = Number(firstBeforeMonth) - 1; selectedYear = Number(firstBeforeYear); }
            else if (state.which === 'last_after') { selectedDay = Number(lastAfterDay); selectedMonth = Number(lastAfterMonth) - 1; selectedYear = Number(lastAfterYear); }
            else if (state.which === 'last_before') { selectedDay = Number(lastBeforeDay); selectedMonth = Number(lastBeforeMonth) - 1; selectedYear = Number(lastBeforeYear); }
            else if (state.which === 'created_after') { selectedDay = Number(createdAfterDay); selectedMonth = Number(createdAfterMonth) - 1; selectedYear = Number(createdAfterYear); }
            else if (state.which === 'created_before') { selectedDay = Number(createdBeforeDay); selectedMonth = Number(createdBeforeMonth) - 1; selectedYear = Number(createdBeforeYear); }
            else if (state.which === 'modified_after') { selectedDay = Number(modifiedAfterDay); selectedMonth = Number(modifiedAfterMonth) - 1; selectedYear = Number(modifiedAfterYear); }
            else if (state.which === 'modified_before') { selectedDay = Number(modifiedBeforeDay); selectedMonth = Number(modifiedBeforeMonth) - 1; selectedYear = Number(modifiedBeforeYear); }
            const isSelected = c.inMonth && selectedDay === c.d && selectedMonth === month && selectedYear === year;
            const cellClass = `zajuna-cal__cell${isDisabled ? ' zajuna-cal__cell--muted' : ''}${isSelected ? ' zajuna-cal__cell--selected' : ''}`;
            return (
              <button key={idx} type="button" onClick={() => handleCalendarSelect(state.which, year, month, c.d)} disabled={isDisabled} className={cellClass} aria-disabled={isDisabled}>{c.d}</button>
            );
          })}
        </div>
      </div>
    );
  }

  const handleDateInputChange = (which, ev) => {
    const val = ev.target.value; // YYYY-MM-DD
    if (!val) return;
    const parts = val.split('-');
    if (parts.length !== 3) return;
    const [y, m, d] = parts;
    if (which === 'after') {
      setFirstAfterYear(String(Number(y)));
      setFirstAfterMonth(String(Number(m)));
      setFirstAfterDay(String(Number(d)));
      setFirstAfterEnabled(true);
    } else {
      setFirstBeforeYear(String(Number(y)));
      setFirstBeforeMonth(String(Number(m)));
      setFirstBeforeDay(String(Number(d)));
      setFirstBeforeEnabled(true);
    }
  };

  // date input changes for created fields
  const handleCreatedDateInputChange = (which, ev) => {
    const val = ev.target.value; // YYYY-MM-DD
    if (!val) return;
    const parts = val.split('-');
    if (parts.length !== 3) return;
    const [y, m, d] = parts;
    if (which === 'after') {
      setCreatedAfterYear(String(Number(y)));
      setCreatedAfterMonth(String(Number(m)));
      setCreatedAfterDay(String(Number(d)));
      setCreatedAfterEnabled(true);
    } else {
      setCreatedBeforeYear(String(Number(y)));
      setCreatedBeforeMonth(String(Number(m)));
      setCreatedBeforeDay(String(Number(d)));
      setCreatedBeforeEnabled(true);
    }
  };

  // Helper: format YYYY-MM-DD -> "lunes, 1 de diciembre de 2025, 00:00"
  const formatFirstAccessDate = (val) => {
    const parts = String(val || '').split('-');
    if (parts.length !== 3) return val || '';
    const [y, m, d] = parts;
    const monthNames = ['enero','febrero','marzo','abril','mayo','junio','julio','agosto','septiembre','octubre','noviembre','diciembre'];
    const weekdayNames = ['domingo','lunes','martes','miércoles','jueves','viernes','sábado'];
    const dt = new Date(Number(y), Number(m) - 1, Number(d));
    const wd = weekdayNames[dt.getDay()];
    const mn = monthNames[Number(m) - 1] || m;
    return `${wd}, ${Number(d)} de ${mn} de ${y}, 00:00`;
  };

  // Memoized display items and date indices (avoid recalculating on unrelated state changes)
  const displayItems = useMemo(() => {
    const items = [];
    const consumed = new Set();
    for (let i = 0; i < filters.length; i++) {
      if (consumed.has(i)) continue;
      const f = filters[i];
      if (f.field === 'firstaccess_after') {
        const j = filters.findIndex((ff, k) => k !== i && !consumed.has(k) && ff.field === 'firstaccess_before');
        if (j !== -1) {
          items.push({ type: 'between', idxA: i, idxB: j, a: filters[i], b: filters[j] });
          consumed.add(i); consumed.add(j);
          continue;
        }
      }
      if (f.field === 'firstaccess_before') {
        const j = filters.findIndex((ff, k) => k !== i && !consumed.has(k) && ff.field === 'firstaccess_after');
        if (j !== -1) {
          items.push({ type: 'between', idxA: j, idxB: i, a: filters[j], b: filters[i] });
          consumed.add(i); consumed.add(j);
          continue;
        }
      }
      items.push({ type: 'single', idx: i, f });
      consumed.add(i);
    }
    return items;
  }, [filters]);

  const firstAfterIndex = useMemo(() => filters.findIndex(ff => ff.field === 'firstaccess_after'), [filters]);
  const firstBeforeIndex = useMemo(() => filters.findIndex(ff => ff.field === 'firstaccess_before'), [filters]);
  const lastAfterIndex = useMemo(() => filters.findIndex(ff => ff.field === 'lastaccess_after'), [filters]);
  const lastBeforeIndex = useMemo(() => filters.findIndex(ff => ff.field === 'lastaccess_before'), [filters]);
  const createdAfterIndex = useMemo(() => filters.findIndex(ff => ff.field === 'timecreated_after'), [filters]);
  const createdBeforeIndex = useMemo(() => filters.findIndex(ff => ff.field === 'timecreated_before'), [filters]);
  const modifiedAfterIndex = useMemo(() => filters.findIndex(ff => ff.field === 'timemodified_after'), [filters]);
  const modifiedBeforeIndex = useMemo(() => filters.findIndex(ff => ff.field === 'timemodified_before'), [filters]);

  const handleReplaceFilters = async () => {
    const newFilters = [];
    const currentFilterValue = filterInputRef.current ? (filterInputRef.current.value || '').trim() : filterValue.trim();
    if (currentFilterValue) {
      newFilters.push({ field: "fullname", op: filterOp, value: currentFilterValue });
    }
    if (lastNameVal.trim()) {
      newFilters.push({ field: "lastname", op: lastNameOp, value: lastNameVal.trim() });
    }
    if (firstNameVal.trim()) {
      newFilters.push({ field: "firstname", op: firstNameOp, value: firstNameVal.trim() });
    }
    if (usernameVal.trim()) {
      newFilters.push({ field: "username", op: usernameOp, value: usernameVal.trim() });
    }
    if (emailVal.trim()) {
      newFilters.push({ field: "email", op: emailOp, value: emailVal.trim() });
    }
    if (cityVal.trim()) {
      newFilters.push({ field: "city", op: cityOp, value: cityVal.trim() });
    }
    if (confirmedOp !== 'any') {
      newFilters.push({ field: "confirmed", op: "equals", value: confirmedOp });
    }
    if (suspendedOp !== 'any') {
      newFilters.push({ field: "suspended", op: "equals", value: suspendedOp });
    }
      const profileCurrent = profileFieldRef.current ? (profileFieldRef.current.value || '').trim() : profileFieldVal.trim();
      if (profileCurrent) {
        newFilters.push({ field: "profile", subfield: profileFieldName, op: profileFieldOp, value: profileCurrent });
      }
      const roleCurrent = roleFieldRef.current ? (roleFieldRef.current.value || '').trim() : roleFieldVal.trim();
      if (roleFieldOne !== 'any' || roleFieldTwo !== 'any' || roleCurrent) {
        if (roleCurrent) {
            try {
              const q = roleCurrent;
              const res = await coursesService.search(q);
              const list = Array.isArray(res) ? res : (res?.items || res?.courses || []);
              const filtered = (roleFieldTwo && roleFieldTwo !== 'any') ? list.filter(c => {
                return String(c.category ?? c.categoryid ?? c.category_id ?? c.catid ?? '') === String(roleFieldTwo) || String(c.categoryname ?? c.catname ?? c.name ?? '') === String(roleFieldTwo);
              }) : list;
              if (!filtered || filtered.length === 0) {
                newFilters.push({ field: "role", subfield1: roleFieldOne, subfield2: roleFieldTwo, value: q, isNotFound: true });
              } else {
                newFilters.push({ field: "role", subfield1: roleFieldOne, subfield2: roleFieldTwo, value: roleCurrent });
              }
            } catch (err) {
              console.error('Error buscando curso para filtro role:', err);
              newFilters.push({ field: "role", subfield1: roleFieldOne, subfield2: roleFieldTwo, value: roleCurrent });
            }
          } else {
            newFilters.push({ field: "role", subfield1: roleFieldOne, subfield2: roleFieldTwo, value: roleCurrent });
          }
      }
        // Only add cohort filter if a value was provided or the operator does not require a value (e.g. 'empty')
        const cohortCurrent = cohortRef.current ? (cohortRef.current.value || '').trim() : cohortVal.trim();
        if (cohortOp === 'empty' || cohortCurrent) {
          newFilters.push({ field: 'cohort', op: cohortOp, value: cohortCurrent });
        }
    // Include identification when user changed the select explicitly
    if (identificationOp !== '' && identificationTouched) {
      newFilters.push({ field: 'identification', op: 'equals', value: identificationOp });
    }
    // Include MNET provider when user changed the select explicitly (replace flow)
    if (mnetProviderOp !== '' && mnetProviderTouched) {
      newFilters.push({ field: 'mnet_provider', op: 'equals', value: mnetProviderOp });
    }
    // Número de ID in replace flow (prefer uncontrolled ref when present)
    const idNumberCurrent = idNumberRef.current ? (idNumberRef.current.value || '').trim() : idNumberVal.trim();
    if (idNumberOp === 'empty' || idNumberCurrent) {
      newFilters.push({ field: 'idnumber', op: idNumberOp, value: idNumberCurrent });
    }
    // Institución in replace flow
    const institutionCurrent = institutionRef.current ? (institutionRef.current.value || '').trim() : institutionVal.trim();
    if (institutionOp === 'empty' || institutionCurrent) {
      newFilters.push({ field: 'institution', op: institutionOp, value: institutionCurrent });
    }
    // Departamento in replace flow
    const departmentCurrent = departmentRef.current ? (departmentRef.current.value || '').trim() : departmentVal.trim();
    if (departmentOp === 'empty' || departmentCurrent) {
      newFilters.push({ field: 'department', op: departmentOp, value: departmentCurrent });
    }
    // Última dirección IP in replace flow
    const lastIpCurrentReplace = lastIpRef.current ? (lastIpRef.current.value || '').trim() : lastIpVal.trim();
    if (lastIpOp === 'empty' || lastIpCurrentReplace) {
      newFilters.push({ field: 'lastip', op: lastIpOp, value: lastIpCurrentReplace });
    }
    // Replace existing filters instead of appending
    setFilters(newFilters);
    setFiltersVersion(v => v + 1);
    if (filterInputRef.current) filterInputRef.current.value = '';
    // central reset for advanced controls (align with add/page-change)
    resetAdvancedFilters();
    setCurrentPage(1);
    // Clear uncontrolled inputs (DOM) to reflect reset without causing controlled re-renders
    try {
      if (lastNameRef.current) lastNameRef.current.value = '';
      if (firstNameRef.current) firstNameRef.current.value = '';
      if (usernameRef.current) usernameRef.current.value = '';
      if (emailRef.current) emailRef.current.value = '';
      if (cityRef.current) cityRef.current.value = '';
      if (idNumberRef.current) idNumberRef.current.value = '';
      if (institutionRef.current) institutionRef.current.value = '';
      if (departmentRef.current) departmentRef.current.value = '';
      if (lastIpRef.current) lastIpRef.current.value = '';
        if (profileFieldRef.current) profileFieldRef.current.value = '';
        if (roleFieldRef.current) roleFieldRef.current.value = '';
        if (cohortRef.current) cohortRef.current.value = '';
      } catch (e) {
        // ignore DOM clearing errors
      }
  };

  // individual filter removal is handled via selection and handleRemoveSelectedFilters

  const handleRemoveSelectedFilters = () => {
    const removedIndexes = new Set(selectedFiltersToRemove);
    const removedFilters = filters.filter((_, i) => removedIndexes.has(i));
    const newFilters = filters.filter((_, i) => !removedIndexes.has(i));

    // Reset UI state for each removed filter to its original default
    removedFilters.forEach((f) => {
      switch (f.field) {
        case 'fullname':
          setFilterOp('contains');
          setFilterValue('');
          if (filterInputRef.current) filterInputRef.current.value = '';
          break;
        case 'lastname':
          setLastNameOp('contains');
          setLastNameVal('');
          if (lastNameRef.current) lastNameRef.current.value = '';
          break;
        case 'firstname':
          setFirstNameOp('contains');
          setFirstNameVal('');
          if (firstNameRef.current) firstNameRef.current.value = '';
          break;
        case 'username':
          setUsernameOp('contains');
          setUsernameVal('');
          if (usernameRef.current) usernameRef.current.value = '';
          break;
        case 'email':
          setEmailOp('contains');
          setEmailVal('');
          if (emailRef.current) emailRef.current.value = '';
          break;
        case 'city':
          setCityOp('contains');
          setCityVal('');
          if (cityRef.current) cityRef.current.value = '';
          break;
        case 'country':
          setCountryOp('any');
          setCountryVal('Colombia');
          break;
        case 'confirmed':
          setConfirmedOp('any');
          break;
        case 'suspended':
          setSuspendedOp('any');
          break;
        case 'profile':
          setProfileFieldOp('contains');
          setProfileFieldVal('');
          setProfileFieldName('any');
          if (profileFieldRef.current) profileFieldRef.current.value = '';
          break;
        case 'role':
          setRoleFieldOne('any');
          setRoleFieldTwo('any');
          setRoleFieldVal('');
          if (roleFieldRef.current) roleFieldRef.current.value = '';
          break;
        case 'enrolled':
          setEnrolledOp('any');
          break;
        case 'systemrole':
          setSystemRoleOp('any');
          break;
        case 'cohort':
          setCohortOp('any');
          setCohortVal('');
          if (cohortRef.current) cohortRef.current.value = '';
          break;
        case 'identification':
          setIdentificationOp('');
          setIdentificationTouched(false);
          break;
        case 'mnet_provider':
          setMnetProviderOp('');
          setMnetProviderTouched(false);
          break;
        case 'idnumber':
          setIdNumberOp('contains');
          setIdNumberVal('');
          if (idNumberRef.current) idNumberRef.current.value = '';
          break;
        case 'institution':
          setInstitutionOp('contains');
          setInstitutionVal('');
          if (institutionRef.current) institutionRef.current.value = '';
          break;
        case 'department':
          setDepartmentOp('contains');
          setDepartmentVal('');
          if (departmentRef.current) departmentRef.current.value = '';
          break;
        case 'lastip':
          setLastIpOp('contains');
          setLastIpVal('');
          if (lastIpRef.current) lastIpRef.current.value = '';
          break;
        case 'firstaccess_after':
          setFirstAfterEnabled(false);
          setFirstAfterDay(String(new Date().getDate()));
          setFirstAfterMonth(String(new Date().getMonth() + 1));
          setFirstAfterYear(String(new Date().getFullYear()));
          break;
        case 'firstaccess_before':
          setFirstBeforeEnabled(false);
          setFirstBeforeDay(String(new Date().getDate()));
          setFirstBeforeMonth(String(new Date().getMonth() + 1));
          setFirstBeforeYear(String(new Date().getFullYear()));
          break;
        case 'lastaccess_after':
          setLastAfterEnabled(false);
          setLastAfterDay(String(new Date().getDate()));
          setLastAfterMonth(String(new Date().getMonth() + 1));
          setLastAfterYear(String(new Date().getFullYear()));
          break;
        case 'lastaccess_before':
          setLastBeforeEnabled(false);
          setLastBeforeDay(String(new Date().getDate()));
          setLastBeforeMonth(String(new Date().getMonth() + 1));
          setLastBeforeYear(String(new Date().getFullYear()));
          break;
        case 'lastaccess_never':
          setLastNever(false);
          break;
        case 'timemodified_after':
          setModifiedAfterEnabled(false);
          setModifiedAfterDay(String(new Date().getDate()));
          setModifiedAfterMonth(String(new Date().getMonth() + 1));
          setModifiedAfterYear(String(new Date().getFullYear()));
          break;
        case 'timemodified_before':
          setModifiedBeforeEnabled(false);
          setModifiedBeforeDay(String(new Date().getDate()));
          setModifiedBeforeMonth(String(new Date().getMonth() + 1));
          setModifiedBeforeYear(String(new Date().getFullYear()));
          break;
        case 'timemodified_never':
          setModifiedNever(false);
          break;
        case 'timecreated_after':
          setCreatedAfterEnabled(false);
          setCreatedAfterDay(String(new Date().getDate()));
          setCreatedAfterMonth(String(new Date().getMonth() + 1));
          setCreatedAfterYear(String(new Date().getFullYear()));
          break;
        case 'timecreated_before':
          setCreatedBeforeEnabled(false);
          setCreatedBeforeDay(String(new Date().getDate()));
          setCreatedBeforeMonth(String(new Date().getMonth() + 1));
          setCreatedBeforeYear(String(new Date().getFullYear()));
          break;
        default:
          // no-op for unrecognized fields
          break;
      }
    });

    // Check cache for immediate display
    const newKey = JSON.stringify(newFilters);
    if (resultsCache.current.has(newKey)) {
      const cached = resultsCache.current.get(newKey);
      setFilters(newFilters);
      setFiltersVersion(v => v + 1);
      setSelectedFiltersToRemove([]);
      setCurrentPage(1);
      // show cached page 1
      setUsers((cached.users || []).slice(0, perPage));
      setTotal(cached.total || 0);
      if (cached.fullTotal !== undefined) setFullTotal(cached.fullTotal);
      // refresh in background without toggling loading
      fetchUsers(true);
      return;
    }

    // No cache: set loading and apply filter changes (will trigger fetch)
    setLoading(true);

    // apply filters state and trigger fetch
    setFilters(newFilters);
    setFiltersVersion(v => v + 1);
    setSelectedFiltersToRemove([]);
    setCurrentPage(1);
  };

  const handleToggleFilterSelection = (idx) => {
    setSelectedFiltersToRemove((prev) =>
      prev.includes(idx) ? prev.filter((i) => i !== idx) : [...prev, idx]
    );
  };

  const handleClearFilters = () => {
    // If we have cached full-results for empty filters, show them immediately
    const emptyKey = JSON.stringify([]);
    if (resultsCache.current.has(emptyKey)) {
      const cached = resultsCache.current.get(emptyKey);
      setFilters([]);
      setFiltersVersion(v => v + 1);
      setSelectedFiltersToRemove([]);
      setCurrentPage(1);
      setUsers((cached.users || []).slice(0, perPage));
      setTotal(cached.total || 0);
      if (cached.fullTotal !== undefined) setFullTotal(cached.fullTotal);
      // refresh in background
      fetchUsers(true);
      return;
    }
    // ensure UI reflects loading immediately to avoid pager flicker
    setLoading(true);
    setFilters([]);
    setFiltersVersion(v => v + 1);
    setSelectedFiltersToRemove([]);
    setCurrentPage(1);
    // Reset common UI controls to defaults
    setFilterOp('contains');
    setFilterValue('');
    setLastNameOp('contains');
    setLastNameVal('');
    setFirstNameOp('contains');
    setFirstNameVal('');
    setUsernameOp('contains');
    setUsernameVal('');
    setEmailOp('contains');
    setEmailVal('');
    setCityOp('contains');
    setCityVal('');
    setCountryOp("any");
    setCountryVal('Colombia');
    setConfirmedOp("any");
    setSuspendedOp("any");
    setProfileFieldOp('contains');
    setProfileFieldVal('');
    setProfileFieldName('any');
    setRoleFieldOne('any');
    setRoleFieldTwo('any');
    setRoleFieldVal('');
    setEnrolledOp('any');
    setSystemRoleOp('any');
    setCohortOp('any');
    setCohortVal('');
    setIdentificationOp('');
    setIdentificationTouched(false);
    setMnetProviderOp('');
    setMnetProviderTouched(false);
    setIdNumberOp('contains');
    setIdNumberVal('');
    setInstitutionOp('contains');
    setInstitutionVal('');
    setDepartmentOp('contains');
    setDepartmentVal('');
    setLastIpOp('contains');
    setLastIpVal('');
    setFirstAfterEnabled(false);
    setFirstBeforeEnabled(false);
    setLastAfterEnabled(false);
    setLastBeforeEnabled(false);
    setLastNever(false);
    setModifiedAfterEnabled(false);
    setModifiedBeforeEnabled(false);
    setModifiedNever(false);
    setCreatedAfterEnabled(false);
    setCreatedBeforeEnabled(false);
    setShowMoreFilters(false);
    // Clear uncontrolled inputs (DOM)
    try {
      if (filterInputRef.current) filterInputRef.current.value = '';
      if (lastNameRef.current) lastNameRef.current.value = '';
      if (firstNameRef.current) firstNameRef.current.value = '';
      if (usernameRef.current) usernameRef.current.value = '';
      if (emailRef.current) emailRef.current.value = '';
      if (cityRef.current) cityRef.current.value = '';
      if (idNumberRef.current) idNumberRef.current.value = '';
      if (institutionRef.current) institutionRef.current.value = '';
      if (departmentRef.current) departmentRef.current.value = '';
      if (lastIpRef.current) lastIpRef.current.value = '';
      if (profileFieldRef.current) profileFieldRef.current.value = '';
      if (roleFieldRef.current) roleFieldRef.current.value = '';
      if (cohortRef.current) cohortRef.current.value = '';
    } catch (e) {
      // ignore DOM clearing errors
    }
  };

  const handleDelete = useCallback(async (user) => {
    if (window.confirm(`¿Eliminar usuario ${user.firstname} ${user.lastname}?`)) {
      try {
        await deleteUsers([user.id]);
        fetchUsers();
      } catch (error) {
        console.error("Error al eliminar usuario:", error);
      }
    }
  }, [fetchUsers]);

  const handleToggleVisibility = async (user) => {
    try {
      console.log(`Intentando cambiar estado del usuario ID: ${user.id}`);

      // Llamar al nuevo endpoint que hace el toggle automáticamente
      const response = await toggleUserStatus(user.id);

      console.log('Respuesta del servidor:', response);

      // Actualizar el estado local basado en la respuesta del servidor
      if (response.new_status === 1) {
        // Ahora está suspendido
        setHiddenUsers((prev) => {
          const newSet = new Set(prev);
          newSet.add(user.id);
          return newSet;
        });
        console.log(`✓ Usuario ${user.firstname} ${user.lastname} suspendido`);
      } else {
        // Ahora está activo
        setHiddenUsers((prev) => {
          const newSet = new Set(prev);
          newSet.delete(user.id);
          return newSet;
        });
        console.log(`✓ Usuario ${user.firstname} ${user.lastname} activado`);
      }
    } catch (error) {
      console.error("❌ Error al cambiar estado de suspensión:", error);
      console.error("Detalles del error:", error.response?.data || error.message);

      const errorMessage = error.response?.data?.message ||
                          error.message ||
                          "Error desconocido al cambiar el estado del usuario";

      alert(`Error: ${errorMessage}`);
    }
  };

  const totalPages = Math.ceil(total / perPage);

  const Pagination = ({ position }) => {
    const getPageNumbers = () => {
      const pages = [];
      // Mostrar todas las páginas sin ellipsis
      for (let i = 1; i <= totalPages; i++) {
        pages.push(i);
      }
      return pages;
    };

    // If a filter/search is active and results fit on one page, hide the pager
    if ((filters.length > 0 || (filterValue && filterValue.trim() !== '')) && total <= perPage) return null;
    // Do not render pager if there are no pages
    if (totalPages < 1) return null;

    return (
      <div className="moodle-pager">
        {currentPage > 1 && (
          <button aria-label="Página anterior" onClick={() => setCurrentPage((p) => Math.max(1, p - 1))}>
            «
          </button>
        )}
        {(() => {
          if (currentPage <= 10) {
            return Array.from({ length: Math.min(10, totalPages) }).map((_, i) => {
              const page = i + 1;
              return (
                <button
                  key={page}
                  onClick={() => setCurrentPage(page)}
                  className={currentPage === page ? "active" : ""}
                  aria-current={currentPage === page ? 'page' : undefined}
                >
                  {page}
                </button>
              );
            });
          } else {
            const pages = [];
            pages.push(
              <button
                key={1}
                onClick={() => setCurrentPage(1)}
                className={currentPage === 1 ? "active" : ""}
                aria-current={currentPage === 1 ? 'page' : undefined}
              >
                1
              </button>
            );
            pages.push(<span key="ellipsis-start" style={{ padding: '0 8px', color: '#666' }}>...</span>);

            let startPage = Math.max(2, currentPage - 4);
            let endPage = Math.min(totalPages - 1, startPage + 9);

            if (endPage === totalPages - 1) {
              startPage = Math.max(2, endPage - 9);
            }

            for (let p = startPage; p <= endPage; p++) {
              pages.push(
                <button
                  key={p}
                  onClick={() => setCurrentPage(p)}
                  className={currentPage === p ? "active" : ""}
                  aria-current={currentPage === p ? 'page' : undefined}
                >
                  {p}
                </button>
              );
            }

            return pages;
          }
        })()}
        {totalPages > 10 && currentPage < totalPages - 10 && (
          <>
            <span style={{ padding: '0 8px', color: '#666' }}>...</span>
            <button
              key={totalPages}
              onClick={() => setCurrentPage(totalPages)}
              className={currentPage === totalPages ? "active" : ""}
              aria-current={currentPage === totalPages ? 'page' : undefined}
            >
              {totalPages}
            </button>
          </>
        )}
        {totalPages > 10 && currentPage >= totalPages - 10 && currentPage > 10 && (
          <button
            key={totalPages}
            onClick={() => setCurrentPage(totalPages)}
            className={currentPage === totalPages ? "active" : ""}
            aria-current={currentPage === totalPages ? 'page' : undefined}
          >
            {totalPages}
          </button>
        )}
        {currentPage < totalPages && (
          <button aria-label="Página siguiente" onClick={() => setCurrentPage((p) => Math.min(totalPages, p + 1))}>
            »
          </button>
        )}
      </div>
    );
  };

  // Memoize the users table rendering to avoid recalculating large JSX on every keystroke
  const usersTable = useMemo(() => {
    return null;
  }, [users, loading, total, perPage, currentPage, handleDelete, navigate]);

  return (
    <div style={{ position: "relative" }}>
      <div className="moodle-side-marker" aria-hidden="true" />
      <div ref={containerRef} className="moodle-container" style={{ position: "relative" }}>
        <style>{staticStyles}</style>
        <div className="moodle-top-row">
          <h2 className="moodle-title h2">
            {filters.length > 0 ? `${total} / ${fullTotal ?? total} Usuarios` : `${fullTotal ?? total} Usuarios`}
          </h2>
        </div>

        <div style={{ marginTop: 8 }}>
          <Pagination position="top" />
        </div>

        <div className="moodle-filter" style={{ borderBottom: 'none', paddingLeft: 0, paddingBottom:0 }}>
          <div className="filter-header" style={{ marginLeft: 0 }}>
            <button
              className={`filter-toggle moodle-chevron-toggle${showFilterSection ? '' : ' closed'}`}
              onClick={() => setShowFilterSection((v) => !v)}
              aria-expanded={showFilterSection}
              aria-controls="filter-body"
              style={{ fontSize: 16, fontWeight: 400, color: '#333', display: 'flex', alignItems: 'center', gap: 0, background: 'none', border: 'none', padding: 0, cursor: 'pointer' }}
            >
              <span className="moodle-chevron-circle" tabIndex={-1} style={{ display: 'inline-flex', alignItems: 'center', justifyContent: 'center', width: 36, height: 36, borderRadius: '50%', transition: 'background 0.15s', marginRight: 5 }}>
                <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                  style={{ transform: showFilterSection ? 'rotate(0deg)' : 'rotate(-90deg)', transition: 'transform 0.2s' }}>
                  <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
                </svg>
              </span>
              <h3 className="filter-title" style={{ fontSize: '1.17rem', fontWeight: 400, color: '#333', margin: 0 }}>Nuevo filtro</h3>
            </button>
          </div>
          {showFilterSection && (
          <div id="filter-body" className="filter-body">
            <form onSubmit={(e) => { e.preventDefault(); handleAddFilter(); if (filterInputRef.current) filterInputRef.current.blur(); }}>
            <div className="filter-row">
              <label>Nombre completo del usuario</label>
              <div className="select-wrap select-wrapper form-inline bulk-caret">
                <select
                  value={filterOp}
                  onChange={(e) => setFilterOp(e.target.value)}
                  className="moodle-select"
                >
                  <option value="contains">contiene</option>
                  <option value="notcontains">no contiene</option>
                  <option value="equals">es igual a</option>
                  <option value="starts">comienza con</option>
                  <option value="ends">termina en</option>
                  <option value="empty">está vacío</option>
                </select>
              </div>
                  <input
                ref={filterInputRef}
                className="moodle-search-input"
                type="text"
                defaultValue={filterValue}
                placeholder=""
                onBlur={() => { if (filterInputRef.current) setFilterValue(filterInputRef.current.value || ''); }}
              />
            </div>
            <div style={{ display: 'flex', alignItems: 'center', marginLeft: 0, marginTop: 0, marginBottom: 0 }}>
                    <button
                      type="button"
                      onClick={() => setShowMoreFilters((v) => !v)}
                      className="show-more-btn"
                    >
                      {showMoreFilters ? 'Ver menos...' : 'Mostrar más...'}
                    </button>
            </div>
            {/* Pre-mounted, memoized more filters component: stays mounted but hidden to make opening instant */}
            <MoreUserFilters
              visible={showMoreFilters}
              days={days} months={months} years={years}
              lastNameOp={lastNameOp} setLastNameOp={setLastNameOp} lastNameRef={lastNameRef} lastNameVal={lastNameVal}
              firstNameOp={firstNameOp} setFirstNameOp={setFirstNameOp} firstNameRef={firstNameRef} firstNameVal={firstNameVal}
              usernameOp={usernameOp} setUsernameOp={setUsernameOp} usernameRef={usernameRef} usernameVal={usernameVal}
              emailOp={emailOp} setEmailOp={setEmailOp} emailRef={emailRef} emailVal={emailVal}
              cityOp={cityOp} setCityOp={setCityOp} cityRef={cityRef} cityVal={cityVal}
              countryOp={countryOp} countryVal={countryVal} setCountryOp={setCountryOp} setCountryVal={setCountryVal}
              confirmedOp={confirmedOp} setConfirmedOp={setConfirmedOp} suspendedOp={suspendedOp} setSuspendedOp={setSuspendedOp}
              profileFieldName={profileFieldName} setProfileFieldName={setProfileFieldName} profileFieldOp={profileFieldOp} setProfileFieldOp={setProfileFieldOp} profileFieldRef={profileFieldRef} profileFieldVal={profileFieldVal}
              roleFieldOne={roleFieldOne} setRoleFieldOne={setRoleFieldOne} roleFieldTwo={roleFieldTwo} setRoleFieldTwo={setRoleFieldTwo} roleFieldRef={roleFieldRef} roleFieldVal={roleFieldVal} roleCategories={roleCategories}
              enrolledOp={enrolledOp} setEnrolledOp={setEnrolledOp} systemRoleOp={systemRoleOp} setSystemRoleOp={setSystemRoleOp}
              cohortOp={cohortOp} setCohortOp={setCohortOp} cohortRef={cohortRef} cohortVal={cohortVal}
              firstAfterEnabled={firstAfterEnabled} setFirstAfterEnabled={setFirstAfterEnabled} firstAfterDay={firstAfterDay} setFirstAfterDay={setFirstAfterDay} firstAfterMonth={firstAfterMonth} setFirstAfterMonth={setFirstAfterMonth} firstAfterYear={firstAfterYear} setFirstAfterYear={setFirstAfterYear}
              firstBeforeEnabled={firstBeforeEnabled} setFirstBeforeEnabled={setFirstBeforeEnabled} firstBeforeDay={firstBeforeDay} setFirstBeforeDay={setFirstBeforeDay} firstBeforeMonth={firstBeforeMonth} setFirstBeforeMonth={setFirstBeforeMonth} firstBeforeYear={firstBeforeYear} setFirstBeforeYear={setFirstBeforeYear}
              lastAfterEnabled={lastAfterEnabled} setLastAfterEnabled={setLastAfterEnabled} lastAfterDay={lastAfterDay} setLastAfterDay={setLastAfterDay} lastAfterMonth={lastAfterMonth} setLastAfterMonth={setLastAfterMonth} lastAfterYear={lastAfterYear} setLastAfterYear={setLastAfterYear}
              lastBeforeEnabled={lastBeforeEnabled} setLastBeforeEnabled={setLastBeforeEnabled} lastBeforeDay={lastBeforeDay} setLastBeforeDay={setLastBeforeDay} lastBeforeMonth={lastBeforeMonth} setLastBeforeMonth={setLastBeforeMonth} lastBeforeYear={lastBeforeYear} setLastBeforeYear={setLastBeforeYear} lastNever={lastNever} setLastNever={setLastNever}
              createdAfterEnabled={createdAfterEnabled} setCreatedAfterEnabled={setCreatedAfterEnabled} createdAfterDay={createdAfterDay} setCreatedAfterDay={setCreatedAfterDay} createdAfterMonth={createdAfterMonth} setCreatedAfterMonth={setCreatedAfterMonth} createdAfterYear={createdAfterYear} setCreatedAfterYear={setCreatedAfterYear}
              createdBeforeEnabled={createdBeforeEnabled} setCreatedBeforeEnabled={setCreatedBeforeEnabled} createdBeforeDay={createdBeforeDay} setCreatedBeforeDay={setCreatedBeforeDay} createdBeforeMonth={createdBeforeMonth} setCreatedBeforeMonth={setCreatedBeforeMonth} createdBeforeYear={createdBeforeYear} setCreatedBeforeYear={setCreatedBeforeYear}
              modifiedAfterEnabled={modifiedAfterEnabled} setModifiedAfterEnabled={setModifiedAfterEnabled} modifiedAfterDay={modifiedAfterDay} setModifiedAfterDay={setModifiedAfterDay} modifiedAfterMonth={modifiedAfterMonth} setModifiedAfterMonth={setModifiedAfterMonth} modifiedAfterYear={modifiedAfterYear} setModifiedAfterYear={setModifiedAfterYear}
              modifiedBeforeEnabled={modifiedBeforeEnabled} setModifiedBeforeEnabled={setModifiedBeforeEnabled} modifiedBeforeDay={modifiedBeforeDay} setModifiedBeforeDay={setModifiedBeforeDay} modifiedBeforeMonth={modifiedBeforeMonth} setModifiedBeforeMonth={setModifiedBeforeMonth} modifiedBeforeYear={modifiedBeforeYear} setModifiedBeforeYear={setModifiedBeforeYear} modifiedNever={modifiedNever} setModifiedNever={setModifiedNever}
              identificationOp={identificationOp} setIdentificationOp={setIdentificationOp} setIdentificationTouched={setIdentificationTouched}
              mnetProviderOp={mnetProviderOp} setMnetProviderOp={setMnetProviderOp} setMnetProviderTouched={setMnetProviderTouched}
              idNumberOp={idNumberOp} setIdNumberOp={setIdNumberOp} idNumberRef={idNumberRef} idNumberVal={idNumberVal}
              institutionOp={institutionOp} setInstitutionOp={setInstitutionOp} institutionRef={institutionRef} institutionVal={institutionVal}
              departmentOp={departmentOp} setDepartmentOp={setDepartmentOp} departmentRef={departmentRef} departmentVal={departmentVal}
              lastIpOp={lastIpOp} setLastIpOp={setLastIpOp} lastIpRef={lastIpRef} lastIpVal={lastIpVal}
              afterCalendarIconRef={afterCalendarIconRef} beforeCalendarIconRef={beforeCalendarIconRef} lastAfterCalendarIconRef={lastAfterCalendarIconRef} lastBeforeCalendarIconRef={lastBeforeCalendarIconRef} createdAfterCalendarIconRef={createdAfterCalendarIconRef} createdBeforeCalendarIconRef={createdBeforeCalendarIconRef} modifiedAfterCalendarIconRef={modifiedAfterCalendarIconRef} modifiedBeforeCalendarIconRef={modifiedBeforeCalendarIconRef}
              openCalendarFor={openCalendarFor}
            />

            

            <div className="filter-actions actions-aligned" style={{ marginTop: 15, marginLeft: 350, display: 'flex', gap: 8 }}>
              {filters.length > 0 && (
                <button type="button" className="btn-moodle" onClick={handleReplaceFilters}>Reemplazar filtros</button>
              )}
              <button type="submit" className="btn-moodle">Añadir filtro</button>
            </div>
            </form>
          </div>
          )}
        </div>

        {filters.length > 0 && (
          <div style={{ padding: '18px 0 0 0' }}>
            <div style={{ borderTop: '1px solid #d6d6d6', marginBottom: 12 }}></div>
            <div className="filter-header">
              <button
                className={`filter-toggle moodle-chevron-toggle${showActiveFilters ? '' : ' closed'}`}
                onClick={() => setShowActiveFilters((v) => !v)}
                aria-expanded={showActiveFilters}
                aria-controls="active-filters-body"
                style={{ fontSize: 16, fontWeight: 400, color: '#333', display: 'flex', alignItems: 'center', gap: 0, background: 'none', border: 'none', padding: 0, cursor: 'pointer' }}
              >
                <span className="moodle-chevron-circle" tabIndex={-1} style={{ display: 'inline-flex', alignItems: 'center', justifyContent: 'center', width: 36, height: 36, borderRadius: '50%', transition: 'background 0.15s', marginRight: 5 }}>
                  <svg className="chevron" width="30" height="30" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"
                    style={{ transform: showActiveFilters ? 'rotate(0deg)' : 'rotate(-90deg)', transition: 'transform 0.2s' }}>
                    <path d="M6 8L10 12L14 8" stroke="currentColor" strokeWidth="1.5" strokeLinecap="round" strokeLinejoin="round"/>
                  </svg>
                </span>
                <h3 className="filter-title" style={{ fontSize: '1.17rem', fontWeight: 400, color: '#333', margin: 0 }}>Filtros activos</h3>
              </button>
            </div>
            {showActiveFilters && (
            <div id="active-filters-body" style={{ marginTop: 10 }}>
              <div style={{ display: 'flex', flexDirection: 'column', gap: 8, marginLeft: '350px' }}>
                {filters.map((f, idx) => (
                  <label
                    key={idx}
                    style={{
                      display: ((firstAfterIndex !== -1 && firstBeforeIndex !== -1 && idx === Math.max(firstAfterIndex, firstBeforeIndex)) || (lastAfterIndex !== -1 && lastBeforeIndex !== -1 && idx === Math.max(lastAfterIndex, lastBeforeIndex)) || (createdAfterIndex !== -1 && createdBeforeIndex !== -1 && idx === Math.max(createdAfterIndex, createdBeforeIndex)) || (modifiedAfterIndex !== -1 && modifiedBeforeIndex !== -1 && idx === Math.max(modifiedAfterIndex, modifiedBeforeIndex))) ? 'none' : 'flex',
                      alignItems: 'center',
                      gap: 8,
                      fontSize: 14,
                      color: '#333',
                      cursor: 'pointer',
                    }}
                  >
                    <input
                      type="checkbox"
                      checked={selectedFiltersToRemove.includes(idx)}
                      onChange={() => handleToggleFilterSelection(idx)}
                        style={{
                        width: 14,
                        height: 14,
                        cursor: 'pointer',
                        accentColor: '#d32f2f',
                      }}
                    />
                      <span style={f.isNotFound ? { color: '#f0ad4e' } : undefined}>
                        {(f.field === 'role' && f.isNotFound) ? '' : (
                          f.field === 'fullname' ? 'Nombre completo'
                          : f.field === 'lastname' ? 'Apellido(s)'
                          : f.field === 'firstname' ? 'Nombre'
                          : f.field === 'username' ? 'Nombre de usuario'
                          : f.field === 'email' ? 'Dirección de correo'
                          : f.field === 'city' ? 'Ciudad'
                          : f.field === 'country' ? 'País'
                          : f.field === 'confirmed' ? 'Confirmado'
                          : f.field === 'suspended' ? 'Cuenta suspendida'
                            : f.field === 'enrolled' ? 'Matriculado en cualquier curso'
                          : f.field === 'systemrole' ? 'Rol del sistema'
                          : f.field === 'cohort' ? 'ID de la cohorte'
                          : f.field === 'identification' ? 'Identificación'
                          : f.field === 'idnumber' ? 'Número de ID'
                          : f.field === 'institution' ? 'Institución'
                          : f.field === 'department' ? 'Departamento'
                          : f.field === 'lastip' ? 'Última dirección IP'
                          : f.field === 'mnet_provider' ? 'Proveedor de ID MNET'
                          : (f.field === 'firstaccess_after' || f.field === 'firstaccess_before') ? 'Primer acceso'
                          : (f.field === 'lastaccess_after' || f.field === 'lastaccess_before') ? 'Último acceso'
                          : (f.field === 'timecreated_after' || f.field === 'timecreated_before') ? 'Tiempo de creado'
                          : f.field === 'timemodified_never' ? 'Nunca se ha modificado'
                          : (f.field === 'timemodified_after' || f.field === 'timemodified_before') ? 'Última modificación'
                          : f.field === 'lastaccess_never' ? 'Nunca se ha accedido'
                          : f.field === 'profile' ? 'Campos de perfil del usuario'
                          : f.field === 'role' ? 'Rol de curso'
                          : f.field
                        )}
                        {f.field === 'profile' ? (
                          // Profile filters: show 'Campos de perfil del usuario: <subfield> <op> <value>'
                          <>
                            {': '}
                            {f.subfield === 'any' ? 'cualquier campo' : (
                              f.subfield === 'about' ? 'Acerca de mí'
                              : f.subfield === 'experience' ? 'Experiencia Profesional'
                              : f.subfield === 'profession' ? 'Profesión'
                              : f.subfield
                            )}
                            {' '}
                            {f.op === 'contains' ? 'contiene' : null}
                            {f.op === 'notcontains' ? 'no contiene' : null}
                            {f.op === 'equals' ? 'es igual a' : null}
                            {f.op === 'notequals' ? 'no es igual a' : null}
                            {f.op === 'starts' ? 'comienza con' : null}
                            {f.op === 'ends' ? 'termina en' : null}
                            {f.op === 'empty' ? 'está vacío' : null}
                            {f.op === 'defined' ? 'está definido' : null}
                            {f.op === 'notdefined' ? 'no está definido' : null}
                            {f.op === 'any' ? 'cualquier valor' : null}
                            {' '}
                            { (f.op === 'defined' || f.op === 'notdefined') ? null : (
                              <>
                                {`"${(f.subfield === 'any' ? f.value : f.value)}"`}
                              </>
                            ) }
                          </>
                          ) : f.field === 'systemrole' ? (
                            <>
                              {` es "${(f.value === 'manager' ? 'Gestor' : (f.value === 'coursecreator' ? 'Creador de curso' : f.value))}"`}
                            </>
                          ) : f.field === 'role' ? (
                          // Role filters: if not found, show explicit error sentence; otherwise render normal sentence
                          <>
                            {f.isNotFound ? (
                              /* Example: error Rol de curso: el curso "perro" no existe */
                              `error Rol de curso: el curso "${f.value}" no existe`
                            ) : (
                              (() => {
                                const roleLabel = (f.subfield1 === 'any') ? '' : (
                                  f.subfield1 === 'learner' ? 'Aprendiz'
                                  : f.subfield1 === 'instructor_no_edit' ? 'Instructor sin permiso de edición'
                                  : f.subfield1 === 'instructor' ? 'Instructor'
                                  : f.subfield1 === 'manager' ? 'Gestor'
                                  : f.subfield1
                                );
                                const cat = (f.subfield2 === 'any' || !f.subfield2) ? '' : (
                                  (() => {
                                    const found = roleCategories.find(rc => String(rc.id) === String(f.subfield2) || String(rc.categoryid) === String(f.subfield2) || String(rc.name) === String(f.subfield2));
                                    return found ? (found.name || found.fullname || found.displayname || found.shortname) : f.subfield2;
                                  })()
                                );
                                return ` es "${roleLabel}"` + ` en cualquier curso de "${cat}"`;
                              })()
                            )}
                          </>
                        ) : (
                          // Non-profile, non-role filters: original order (Field <op> <value>)
                          <>
                            {' '}
                            {f.op === 'contains' ? 'contiene'
                              : f.op === 'notcontains' ? 'no contiene'
                              : f.op === 'equals' ? 'es igual a'
                              : f.op === 'notequals' ? 'no es igual a'
                              : f.op === 'starts' ? 'comienza con'
                              : f.op === 'ends' ? 'termina en'
                              : f.op === 'empty' ? 'está vacío'
                              : f.op === 'defined' ? 'está definido'
                              : f.op === 'notdefined' ? 'no está definido'
                              : f.op === 'any' ? 'cualquier valor'
                              : f.op}
                            {' '}
                            { (f.op === 'defined' || f.op === 'notdefined') ? null : (
                              <>
                                {(() => {
                                  if (f.field === 'confirmed' || f.field === 'suspended' || f.field === 'enrolled') return `"${f.value === 'true' ? 'Sí' : (f.value === 'false' ? 'No' : f.value)}"`;
                                  if (f.field === 'identification') return `"${identificationOptionsMap[f.value] || f.value}"`;
                                  if (f.field === 'mnet_provider') return `"${mnetProviderOptionsMap[f.value] || f.value}"`;
                                  if (f.field === 'idnumber') return `"${f.value}"`;
                                  if (f.field === 'institution') return `"${f.value}"`;
                                  if (f.field === 'department') return `"${f.value}"`;
                                  if (f.field === 'lastip') return `"${f.value}"`;
                                  if (f.field === 'systemrole') return `"${f.value === 'manager' ? 'Gestor' : (f.value === 'coursecreator' ? 'Creador de curso' : f.value)}"`;
                                  if (f.field === 'cohort') return `"${f.value}"`;
                                                              if (['firstaccess_after','firstaccess_before','lastaccess_after','lastaccess_before','lastaccess_never','timecreated_after','timecreated_before','timemodified_after','timemodified_before','timemodified_never'].includes(f.field)) {
                                                                const key = String(f.field);
                                                                const isLast = key.includes('lastaccess');
                                                                const isCreated = key.includes('timecreated');
                                                                const isModified = key.includes('timemodified');
                                                                const afterKey = isLast ? 'lastaccess_after' : (isCreated ? 'timecreated_after' : (isModified ? 'timemodified_after' : 'firstaccess_after'));
                                                                const beforeKey = isLast ? 'lastaccess_before' : (isCreated ? 'timecreated_before' : (isModified ? 'timemodified_before' : 'firstaccess_before'));
                                    const aIndex = filters.findIndex(ff => ff.field === afterKey);
                                    const bIndex = filters.findIndex(ff => ff.field === beforeKey);
                                    if (aIndex !== -1 && bIndex !== -1) {
                                      const firstIdx = Math.min(aIndex, bIndex);
                                      if (idx === firstIdx) {
                                        const aVal = filters[aIndex]?.value || '';
                                        const bVal = filters[bIndex]?.value || '';
                                        return `está entre ${formatFirstAccessDate(aVal)} y ${formatFirstAccessDate(bVal)}`;
                                      }
                                      return '';
                                    }
                                    if (f.field === 'lastaccess_never' || f.field === 'timemodified_never') {
                                      return '';
                                    }
                                    // Single-case: render "está después de / está antes de" with formatted date
                                    const parts = String(f.value || '').split('-');
                                    if (parts.length === 3) {
                                      const [y, m, d] = parts;
                                      const monthNames = ['enero','febrero','marzo','abril','mayo','junio','julio','agosto','septiembre','octubre','noviembre','diciembre'];
                                      const weekdayNames = ['domingo','lunes','martes','miércoles','jueves','viernes','sábado'];
                                      const dt = new Date(Number(y), Number(m) - 1, Number(d));
                                      const wd = weekdayNames[dt.getDay()];
                                      const mn = monthNames[Number(m) - 1] || m;
                                      const when = (f.field === afterKey) ? 'está después de' : 'está antes de';
                                      return `${when} ${wd}, ${Number(d)} de ${mn} de ${y}, 00:00`;
                                    }
                                    return `${(f.field === afterKey) ? 'está después de' : 'está antes de'} ${f.value || ''}`;
                                  }
                                  return `"${f.value}"`;
                                })()}
                              </>
                            )}
                          </>
                        )}
                      </span>
                  </label>
                ))}
              </div>
              <div style={{ display: 'flex', gap: 10, marginTop: 16, marginLeft: '350px' }}>
                <button
                  className="btn-moodle"
                  style={{
                    fontWeight: 400,
                    fontSize: 14,
                    padding: '8px 17.6px',
                    borderRadius: 12,
                    cursor: selectedFiltersToRemove.length === 0 ? 'not-allowed' : 'pointer',
                    opacity: selectedFiltersToRemove.length === 0 ? 0.5 : 1,
                  }}
                  onClick={handleRemoveSelectedFilters}
                  disabled={selectedFiltersToRemove.length === 0}
                >
                  Eliminar seleccionados
                </button>
                <button
                  className="btn-moodle"
                  style={{
                    fontWeight: 400,
                    fontSize: 14,
                    padding: '8px 17.6px',
                    borderRadius: 12,
                    cursor: 'pointer',
                  }}
                  onClick={handleClearFilters}
                >
                  Eliminar todos los filtros
                </button>
              </div>
            </div>
            )}
          </div>
        )}

        <div style={{ borderTop: '1px solid #d6d6d6', marginTop: 18, marginBottom: 10 }}></div>

        {usersTable}

        <div style={{ marginTop: 18, display: 'flex', justifyContent: 'flex-start' }}>
          <button
            className="btn-moodle btn-nuevo-usuario"
            style={{
              background: '#cfd4db',
              color: '#222',
              fontWeight: 400,
              fontSize: 14,
              padding: '8px 17.6px',
              borderRadius: 12,
              border: 'none',
              boxShadow: 'none',
              cursor: 'pointer',
              transition: 'background 0.2s'
            }}
            onMouseOver={e => e.currentTarget.style.background = '#bfc3c7'}
            onMouseOut={e => e.currentTarget.style.background = '#cfd4db'}
            onClick={() => navigate('/users/new')}
          >
            Crear un nuevo usuario
          </button>
        </div>

        {/* Calendar popup rendered at root so it can position over other elements */}
        <CalendarPopup state={calendarState} />
      </div>
    </div>
  );
}
