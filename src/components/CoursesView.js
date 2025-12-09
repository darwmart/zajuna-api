import { useEffect, useState, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { CircularProgress, Alert } from "@mui/material";
import { FaEye, FaEyeSlash, FaCog, FaSearch, FaArrowUp, FaArrowDown, FaTrash, FaCopy } from "react-icons/fa";
import { coursesService } from "../services/coursesService";

export default function CoursesView({ initialCategories }) {
  const navigate = useNavigate();
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // Track which course row is selected (for checkbox + green left border)
  const [selectedCourseId, setSelectedCourseId] = useState(null);

  const [categories, setCategories] = useState(initialCategories || []);
  const [selectedCategory, setSelectedCategory] = useState("");
  const [topFilter, setTopFilter] = useState("");
  const [searchValue, setSearchValue] = useState("");
  const [debouncedSearch, setDebouncedSearch] = useState("");
  const [allCourses, setAllCourses] = useState([]);
  const [loadingAllCourses, setLoadingAllCourses] = useState(false);
  const [searchResults, setSearchResults] = useState(null); // null = no active search, [] = searched but no results
  const [searchMeta, setSearchMeta] = useState({ count: 0, raw: null });
  // Track an expanded course whose details panel should appear below both panels
  const [expandedCourse, setExpandedCourse] = useState(null);
  const [courseDetail, setCourseDetail] = useState(null);
  const [loadingCourseDetail, setLoadingCourseDetail] = useState(false);
  const [courseDetailError, setCourseDetailError] = useState(null);
  const [hiddenCourses, setHiddenCourses] = useState(new Set());

  const toggleCourseDetails = (course) => {
    setExpandedCourse((prev) => (prev && String(prev.id) === String(course.id) ? null : course));
  };

  // When expandedCourse changes, fetch its detailed info from the backend
  useEffect(() => {
    let mounted = true;
    const loadDetail = async () => {
      if (!expandedCourse) {
        setCourseDetail(null);
        setCourseDetailError(null);
        setLoadingCourseDetail(false);
        return;
      }
      setLoadingCourseDetail(true);
      setCourseDetailError(null);
      try {
        // Pass the whole expandedCourse object to the service; it will resolve an id from multiple possible fields.
        const data = await coursesService.getCourseDetail(expandedCourse);
        if (!mounted) return;
        // backend may return object directly or wrapper
        const detail = Array.isArray(data) ? data[0] : (data && (data.course || data.item || data.data)) || data;
        setCourseDetail(detail || null);
      } catch (e) {
        console.error('Failed to load course detail', e);
        const msg = (e && e.response && (e.response.data?.message || e.response.data?.error)) || e.message || 'No se pudieron cargar los detalles';
        if (mounted) setCourseDetailError(msg);
        if (mounted) setCourseDetail(null);
      } finally {
        if (mounted) setLoadingCourseDetail(false);
      }
    };
    loadDetail();
    return () => { mounted = false; };
  }, [expandedCourse]);

  const renderListNames = (arr, key = 'name') => {
    // If backend explicitly returned null, show 0 per requirement
    if (arr === null) return '0';
    if (!Array.isArray(arr) || arr.length === 0) return '-';
    const nodes = arr.map((x, idx) => {
      if (x === null || x === undefined) return null;
      let label = '';
      if (typeof x === 'string' || typeof x === 'number') label = String(x).trim();
      else {
        const keysToTry = [];
        if (key) keysToTry.push(key);
        keysToTry.push('name', 'enrol', 'method', 'method_name', 'shortname', 'username', 'roleshort', 'rolename', 'id');
        for (const k of keysToTry) {
          if (x[k] !== undefined && x[k] !== null && String(x[k]).trim() !== '') {
            label = String(x[k]).trim();
            break;
          }
        }
        if (!label) {
          try { label = JSON.stringify(x); } catch (e) { label = String(x); }
        }
      }
      return label ? <div className="detail-subitem" key={`li-${idx}`}>{label}</div> : null;
    }).filter(Boolean);
    if (nodes.length === 0) return '-';
    return <>{nodes}</>;
  };

  const renderModules = (mods) => {
    // Null from backend should be shown as 0
    if (mods === null) return '0';
    if (!Array.isArray(mods) || mods.length === 0) return '-';
    const nodes = mods.map((m, idx) => {
      if (!m) return null;
      const label = String(m.name ?? m).trim();
      return label ? <div className="detail-subitem" key={`mod-${idx}`}>{label}</div> : null;
    }).filter(Boolean);
    return nodes.length === 0 ? '-' : <>{nodes}</>;
  };

  // Render sections: show only the 'name' of each section
  const renderSections = (secs) => {
    if (secs === null) return '0';
    if (!Array.isArray(secs) || secs.length === 0) return '-';
    const nodes = secs.map((s, idx) => {
      if (!s) return null;
      // Si es un string directamente, usarlo. Si es objeto, extraer name/title/section
      const label = typeof s === 'string' ? s.trim() : String(s.name ?? s.title ?? s.section ?? '').trim();
      return label ? <div className="detail-subitem" key={`sec-${idx}`}>{label}</div> : null;
    }).filter(Boolean);
    return nodes.length === 0 ? '-' : <>{nodes}</>;
  };

  // Render role assignments: backend sends an object/map like {"student": 5, "teacher": 2}
  const renderRoleAssignments = (data) => {
    if (data === null || data === undefined) return '0';

    // If it's an object/map (expected format from backend)
    if (typeof data === 'object' && !Array.isArray(data)) {
      const entries = Object.entries(data).filter(([role, count]) => role && count > 0);
      if (entries.length === 0) return '-';
      return (
        <>
          {entries.map(([role, count], idx) => (
            <div className="detail-subitem" key={`${role}-${idx}`}>{`${role}: ${count}`}</div>
          ))}
        </>
      );
    }

    // Fallback: if it's an array (legacy format)
    if (Array.isArray(data)) {
      if (data.length === 0) return '-';
      const map = new Map();
      for (const a of data) {
        if (!a) continue;
        const roleId = a.roleid ?? a.roleId ?? null;
        const short = (a.roleshort || a.role_short || a.rolename || '').toString().trim();
        const label = short || (roleId !== null && roleId !== undefined ? String(roleId) : '');
        const key = roleId !== null && roleId !== undefined ? `id:${String(roleId)}` : `label:${label}`;
        if (!map.has(key)) map.set(key, { label, count: 0 });
        map.get(key).count += 1;
      }
      const items = Array.from(map.values()).filter(it => it && it.label);
      if (items.length === 0) return '-';
      return (
        <>
          {items.map((it, idx) => (
            <div className="detail-subitem" key={`${it.label}-${idx}`}>{`${it.label}: ${it.count}`}</div>
          ))}
        </>
      );
    }

    return '-';
  };

  const handleSearchAction = async () => {
    const q = (searchValue || '').trim();
    if (!q) {
      setSearchResults(null);
      return;
    }
    setLoadingAllCourses(true);
    try {
      // Use the service search which will prefer a server endpoint but will
      // always fall back to a global client-side filter to guarantee results
      // include any course in the database.
      const list = await coursesService.search(q);
      const arr = Array.isArray(list) ? list : (list.items || list.courses || list || []);
      setSearchMeta({ count: arr.length, raw: null });
      setSearchResults(arr);
    } catch (e) {
      console.error('search error', e);
      setSearchResults([]);
      setSearchMeta({ count: 0, raw: null });
    } finally {
      setLoadingAllCourses(false);
    }
  };
  const [coursesInCategory, setCoursesInCategory] = useState([]);
  const [loadingCategoryCourses, setLoadingCategoryCourses] = useState(false);
  const [moveTargetCategory, setMoveTargetCategory] = useState("");
  // Track selected course ids (for enabling the move filter when at least one is selected)
  const [selectedCourseIds, setSelectedCourseIds] = useState(new Set());
  // Track selected category ids in the left panel (for enabling ordering/move controls)
  const [selectedCategoryIds, setSelectedCategoryIds] = useState(new Set());
  // Ordering filters state inside the categories panel
  const [orderingFilter1, setOrderingFilter1] = useState("");
  const [orderingFilter2, setOrderingFilter2] = useState("");
  const [orderingFilter3, setOrderingFilter3] = useState("");
  // Derived flag: when true, ordering secondary filters and Ordenar button should be disabled
  const orderingDisabled = (orderingFilter1 === "" && selectedCategoryIds.size === 0);
  // Track which floating menu is open in the create-controls area: 'order' | 'perpage' | null
  const [openMenu, setOpenMenu] = useState(null);
  const createControlsRef = useRef(null);
  const [perPage, setPerPage] = useState(0); // 0 = Todos
  const [currentPage, setCurrentPage] = useState(1);


  useEffect(() => {
    const loadAll = async () => {
      try {
        // Si ya tenemos categorías iniciales, usarlas en lugar de hacer request
        let cats = [];
        if (initialCategories && initialCategories.length > 0) {
          cats = initialCategories;
        } else {
          const res = await coursesService.getCategories();
          // API might return an array or an object with items/categories
          if (Array.isArray(res)) cats = res;
          else cats = res.items || res.categories || res || [];
        }
        // If the API doesn't include a numeric count property for categories, fetch counts per category
        // Note: this triggers one request per category when needed; acceptable for small sets.
        const needCounts = cats.filter((c) => {
          const has = c.coursecount !== undefined && c.coursecount !== null;
          const hasAlt = c.count !== undefined && c.count !== null;
          return !(has || hasAlt || (c.courses && Array.isArray(c.courses)));
        });

        if (needCounts.length > 0) {
          // fetch counts in parallel
          const counts = await Promise.all(
            needCounts.map(async (c) => {
              try {
                const cr = await coursesService.getCoursesByField("category", c.id);
                const list = Array.isArray(cr) ? cr : (cr.items || cr.courses || cr || []);
                return { id: c.id, _count: list.length };
              } catch (e) {
                return { id: c.id, _count: 0 };
              }
            })
          );
          const countsMap = counts.reduce((acc, cur) => ({ ...acc, [String(cur.id)]: cur._count }), {});
          cats = cats.map((c) => ({ ...c, _count: countsMap[String(c.id)] ?? c._count }));
        }

        // Ensure a visible flag exists for each category (default: visible)
        cats = cats.map((c) => ({ ...c, visible: c.visible !== undefined ? c.visible : true }));
        setCategories(cats);
        if (cats.length > 0) setSelectedCategory((prev) => String(prev || cats[0].id));
      } catch (e) {
        setError("No se pudieron cargar las categorías");
      } finally {
        setLoading(false);
      }
    };
    loadAll();
  }, []);

  useEffect(() => {
    const load = async () => {
      if (!selectedCategory) {
        setCoursesInCategory([]);
        return;
      }
      setLoadingCategoryCourses(true);
      try {
        const res = await coursesService.getCoursesByField("category", selectedCategory);
        // getCoursesByField ya devuelve un array normalizado, no un objeto con propiedades
        const list = Array.isArray(res) ? res : (res?.courses || res?.items || []);
        setCoursesInCategory(list);
      } catch (e) {
        setCoursesInCategory([]);
      } finally {
        setLoadingCategoryCourses(false);
      }
    };
    load();
  }, [selectedCategory]);

  // Load all courses when topFilter == 'courses'
  useEffect(() => {
    let mounted = true;
    const loadAll = async () => {
      // Only load global list when user selected 'courses' and no specific category is selected
      if (topFilter !== 'courses' || (topFilter === 'courses' && selectedCategory)) return;
      setLoadingAllCourses(true);
      try {
        const res = await coursesService.getAll();
        const list = Array.isArray(res) ? res : (res.items || res.courses || res || []);
        if (mounted) {
          setAllCourses(list);
          // Inicializar cursos ocultos (visible=0)
          const hidden = new Set(list.filter(c => c.visible === 0).map(c => c.id));
          setHiddenCourses(hidden);
        }
      } catch (e) {
        if (mounted) setAllCourses([]);
      } finally {
        if (mounted) setLoadingAllCourses(false);
      }
    };
    loadAll();
    return () => { mounted = false; };
  }, [topFilter, selectedCategory]);

  // Debounce search input to avoid rapid filtering
  useEffect(() => {
    const t = setTimeout(() => setDebouncedSearch(searchValue), 250);
    return () => clearTimeout(t);
  }, [searchValue]);

  // Close dropdowns when clicking outside the create-controls area
  useEffect(() => {
    const onDocClick = (ev) => {
      if (!createControlsRef.current) return;
      if (!createControlsRef.current.contains(ev.target)) {
        setOpenMenu(null);
      }
    };
    document.addEventListener('mousedown', onDocClick);
    return () => document.removeEventListener('mousedown', onDocClick);
  }, []);

  // Close dropdowns on Escape key
  useEffect(() => {
    const onKey = (e) => {
      if (e.key === 'Escape' || e.key === 'Esc') setOpenMenu(null);
    };
    document.addEventListener('keydown', onKey);
    return () => document.removeEventListener('keydown', onKey);
  }, []);

  const getCategoryCount = (cat) => {
    const candidates = [
      cat._count,
      cat.coursecount,
      cat.count,
      cat.numcourses,
      cat.numbercourses,
      (cat.courses && Array.isArray(cat.courses) && cat.courses.length),
      cat.coursescount,
    ];
    for (const v of candidates) {
      if (v !== undefined && v !== null) {
        const n = Number(v);
        if (!Number.isNaN(n)) return n;
        return String(v);
      }
    }
    return 0;
  };

  const handleMoveSelectedCourses = () => {
    // Placeholder: implement actual move logic later. For now log the chosen category and selected ids.
    const ids = Array.from(selectedCourseIds);
    console.log('Mover cursos seleccionados a', moveTargetCategory, 'ids:', ids);
    // TODO: collect selected course ids and call backend endpoint to move them.
  };

  const handleMoveSelectedCategories = () => {
    // Placeholder for moving categories. For now just log the chosen target and selected category ids.
    const ids = Array.from(selectedCategoryIds);
    console.log('Mover categorías seleccionadas a', moveTargetCategory, 'ids:', ids);
    // TODO: call backend endpoint to move categories when API is available.
  };

  const toggleCourseSelected = (courseId) => {
    setSelectedCourseIds((prev) => {
      const copy = new Set(prev);
      const id = String(courseId);
      if (copy.has(id)) copy.delete(id);
      else copy.add(id);
      return copy;
    });
  };

  const toggleCategorySelected = (catId) => {
    setSelectedCategoryIds((prev) => {
      const copy = new Set(prev);
      const id = String(catId);
      if (copy.has(id)) copy.delete(id);
      else copy.add(id);
      return copy;
    });
  };

  // Move category up in the list (swap with previous index)
  const handleMoveCategoryUp = async (catId) => {
    const idx = categories.findIndex((c) => String(c.id) === String(catId));
    if (idx <= 0) return; // already first or not found

    // Calcular beforeid: el ID de la categoría que estará 2 posiciones arriba
    // Si movemos la categoría a la posición 0, beforeid será el ID de la categoría que actualmente está en la posición 0
    const beforeId = idx === 1 ? categories[0].id : categories[idx - 2].id;

    try {
      // Llamar al backend para actualizar el sortOrder
      await coursesService.moveCategory(catId, beforeId);

      // Actualizar el estado local solo después de que el backend confirme
      setCategories((prev) => {
        const copy = [...prev];
        const tmp = copy[idx - 1];
        copy[idx - 1] = copy[idx];
        copy[idx] = tmp;
        return copy;
      });
    } catch (error) {
      console.error("Error al mover la categoría:", error);
      alert("Error al mover la categoría. Por favor, intenta de nuevo.");
    }
  };

  // Move category down in the list (swap with next index)
  const handleMoveCategoryDown = async (catId) => {
    const idx = categories.findIndex((c) => String(c.id) === String(catId));
    if (idx === -1 || idx >= categories.length - 1) return; // not found or already last

    // Calcular beforeid: el ID de la categoría que estará después de la posición objetivo
    // Si movemos a la última posición, beforeid será 0 (mover al final)
    const beforeId = idx + 2 >= categories.length ? 0 : categories[idx + 2].id;

    try {
      // Llamar al backend para actualizar el sortOrder
      await coursesService.moveCategory(catId, beforeId);

      // Actualizar el estado local solo después de que el backend confirme
      setCategories((prev) => {
        const copy = [...prev];
        const tmp = copy[idx + 1];
        copy[idx + 1] = copy[idx];
        copy[idx] = tmp;
        return copy;
      });
    } catch (error) {
      console.error("Error al mover la categoría:", error);
      alert("Error al mover la categoría. Por favor, intenta de nuevo.");
    }
  };

  // Move a course up in the currently visible course list - search results, category list or global list
  const handleMoveCourseUp = async (courseId) => {
    try {
      const idx = coursesInCategory.findIndex((c) => String(c.id) === String(courseId));

      if (idx <= 0) {
        return; // already first or not found
      }

      // Get the ID of the course that is currently before this one
      const beforeId = coursesInCategory[idx - 1].id;

      // Call backend to persist the move
      await coursesService.moveCourse(courseId, Number(selectedCategory), beforeId);

      // Reload courses from backend to get updated sortorder
      const res = await coursesService.getCoursesByField("category", selectedCategory);
      const list = res.items || res.courses || res || [];
      setCoursesInCategory(list);
    } catch (error) {
      console.error("Error al mover el curso:", error);
    }
  };

  // Move a course down in the currently visible course list
  const handleMoveCourseDown = async (courseId) => {
    try {
      const idx = coursesInCategory.findIndex((c) => String(c.id) === String(courseId));

      if (idx === -1 || idx >= coursesInCategory.length - 1) {
        return; // not found or already last
      }

      // When moving down, we want to swap with the next course
      // To do this, we tell backend to place current course "before" the course at idx+2
      // This makes the backend swap current with next (idx+1)
      const beforeId = idx + 2 < coursesInCategory.length ? coursesInCategory[idx + 2].id : null;

      // Call backend to persist the move
      await coursesService.moveCourse(courseId, Number(selectedCategory), beforeId);

      // Reload courses from backend to get updated sortorder
      const res = await coursesService.getCoursesByField("category", selectedCategory);
      const list = res.items || res.courses || res || [];
      setCoursesInCategory(list);
    } catch (error) {
      console.error("Error al mover el curso:", error);
    }
  };

  // Toggle course visibility (visible: 0 = oculto, 1 = visible)
  const handleToggleCourseVisibility = async (course, e) => {
    if (e) e.stopPropagation();
    try {
      const isHidden = hiddenCourses.has(course.id);

      if (isHidden) {
        // Curso oculto (0) → Mostrar (1) con PUT
        await coursesService.showCourse(course.id);
        setHiddenCourses((prev) => {
          const newSet = new Set(prev);
          newSet.delete(course.id);
          return newSet;
        });
        console.log(`Curso "${course.fullname || course.shortname}" mostrado`);
      } else {
        // Curso visible (1) → Ocultar (0) con PUT
        await coursesService.hideCourse(course.id);
        setHiddenCourses((prev) => {
          const newSet = new Set(prev);
          newSet.add(course.id);
          return newSet;
        });
        console.log(`Curso "${course.fullname || course.shortname}" ocultado`);
      }
    } catch (error) {
      console.error("Error al cambiar visibilidad del curso:", error);
      alert("Error al cambiar la visibilidad del curso");
    }
  };

  // Toggle category visibility for other users (in-memory)
  const toggleCategoryVisibility = (catId) => {
    setCategories((prev) => prev.map((c) => {
      if (String(c.id) === String(catId)) {
        return { ...c, visible: !(c.visible === false) };
      }
      return c;
    }));
    // TODO: persist visibility to backend if API supports it
  };

  // Decide source of courses depending on topFilter and selectedCategory
  // If a category is selected, prefer coursesInCategory. If no category selected and topFilter === 'courses' show allCourses.
  let coursesSource;
  if (selectedCategory) {
    coursesSource = coursesInCategory;
  } else if (topFilter === 'courses') {
    coursesSource = allCourses;
  } else {
    coursesSource = coursesInCategory;
  }

  // Ensure coursesSource is an array (API may return an object with items/courses)
  const asArray = (v) => {
    if (!v) return [];
    if (Array.isArray(v)) return v;
    if (v.items && Array.isArray(v.items)) return v.items;
    if (v.courses && Array.isArray(v.courses)) return v.courses;
    return [];
  };

  const coursesSourceArray = asArray(coursesSource);

  // Ensure searchResults is treated as an array when rendering (API may return an object)
  const searchResultsArray = asArray(searchResults);

  // Decide whether to apply the live (debounced) filter:
  // - apply when we're in the global "courses" topFilter view (user expects live filtering there),
  // - or when searchResults has been activated (we're showing the search panel).
  // When browsing a specific category (default), typing in the search box should not
  // filter the current category list live; the user must press Enter or click the
  // magnifier to run a full database search.
  const shouldFilterLive = (topFilter === 'courses') || (searchResults !== null);

  const displayedCourses = coursesSourceArray.filter((course) => {
    if (!shouldFilterLive) return true; // don't filter while typing in category view
    if (!debouncedSearch) return true;
    const name = (course.fullname || course.displayname || course.shortname || "").toLowerCase();
    return name.includes(debouncedSearch.toLowerCase());
  });

  // Reset pagination when the data or filters change
  useEffect(() => {
    setCurrentPage(1);
  }, [searchResults, topFilter, selectedCategory, debouncedSearch]);

  // Compute paginated slice
  const totalItems = displayedCourses.length;
  const pageSize = Number(perPage) > 0 ? Number(perPage) : totalItems || 0;
  const totalPages = pageSize > 0 ? Math.max(1, Math.ceil(totalItems / pageSize)) : 1;
  const startIndex = pageSize > 0 ? (currentPage - 1) * pageSize : 0;
  const endIndex = pageSize > 0 ? startIndex + pageSize : totalItems;
  const visibleCourses = (Number(perPage) > 0) ? displayedCourses.slice(startIndex, Math.min(endIndex, totalItems)) : displayedCourses;

  // Pagination for search results (rendered separately)
  const searchTotalItems = Array.isArray(searchResultsArray) ? searchResultsArray.length : 0;
  const searchPageSize = Number(perPage) > 0 ? Number(perPage) : searchTotalItems || 0;
  const searchTotalPages = searchPageSize > 0 ? Math.max(1, Math.ceil(searchTotalItems / searchPageSize)) : 1;
  const searchStart = searchPageSize > 0 ? (currentPage - 1) * searchPageSize : 0;
  const searchEnd = searchPageSize > 0 ? searchStart + searchPageSize : searchTotalItems;
  const visibleSearchResults = (Number(perPage) > 0) ? searchResultsArray.slice(searchStart, Math.min(searchEnd, searchTotalItems)) : searchResultsArray;

  // Componente de paginación reutilizable
  const renderPaginator = (current, total, onPageChange, keyPrefix = 'p') => {
    if (total <= 1) return null;

    return (
      <div className="inline-pager">
        <div className="moodle-pager" role="navigation" aria-label="Paginador">
          {current > 1 && (
            <button onClick={() => onPageChange(p => Math.max(1, p - 1))} aria-label="Página anterior">«</button>
          )}
          {(() => {
            if (current <= 10) {
              return Array.from({ length: Math.min(10, total) }).map((_, i) => {
                const p = i + 1;
                return (
                  <button key={`${keyPrefix}-${p}`} className={p === current ? 'active' : ''} onClick={() => onPageChange(p)} aria-current={p === current}>{p}</button>
                );
              });
            } else {
              const pages = [];
              pages.push(
                <button key={`${keyPrefix}-1`} className={1 === current ? 'active' : ''} onClick={() => onPageChange(1)} aria-current={1 === current}>1</button>
              );
              pages.push(<span key={`${keyPrefix}-ellipsis-start`} style={{ padding: '0 8px', color: '#666' }}>...</span>);

              let startPage = Math.max(2, current - 4);
              let endPage = Math.min(total - 1, startPage + 9);

              if (endPage === total - 1) {
                startPage = Math.max(2, endPage - 9);
              }

              for (let p = startPage; p <= endPage; p++) {
                pages.push(
                  <button key={`${keyPrefix}-${p}`} className={p === current ? 'active' : ''} onClick={() => onPageChange(p)} aria-current={p === current}>{p}</button>
                );
              }

              return pages;
            }
          })()}
          {total > 10 && current < total - 10 && (
            <>
              <span style={{ padding: '0 8px', color: '#666' }}>...</span>
              <button key={`${keyPrefix}-last`} className={total === current ? 'active' : ''} onClick={() => onPageChange(total)} aria-current={total === current}>{total}</button>
            </>
          )}
          {total > 10 && current >= total - 10 && current > 10 && (
            <button key={`${keyPrefix}-last`} className={total === current ? 'active' : ''} onClick={() => onPageChange(total)} aria-current={total === current}>{total}</button>
          )}
          {current < total && (
            <button onClick={() => onPageChange(p => Math.min(total, p + 1))} aria-label="Página siguiente">»</button>
          )}
        </div>
      </div>
    );
  };

  return (
    <div className="courses-page-root courses-page-left" style={{ paddingTop: 0, paddingLeft: 0 }}>
      <div className="courses-manage-top">
        {searchResults === null && (
          <select
            className="moodle-select courses-top-select"
            value={topFilter}
            onChange={(e) => setTopFilter(String(e.target.value))}
          >
            <option value="">Categorías de cursos y cursos</option>
            <option value="categories">Categorías</option>
            <option value="courses">Cursos</option>
          </select>
        )}

        {/* When user selects 'Cursos' show a category selector between the top filter and the search box.
            It uses the existing `categories` state and binds to `selectedCategory`. */}
        {searchResults === null && topFilter === 'courses' && (
          <select
            className="moodle-select courses-category-select"
            value={selectedCategory || ''}
            onChange={(e) => setSelectedCategory(String(e.target.value))}
            title="Filtrar cursos por categoría"
          >
            {/* Informational placeholder: not selectable */}
            <option value="" disabled>Elegir...</option>
            {(categories || []).map((c) => (
              <option key={`catsel-${c.id}`} value={String(c.id)}>{c.name || c.fullname || c.text}</option>
            ))}
          </select>
        )}

        <div className="search-wrapper">
          <input
            className="moodle-search-input"
            placeholder="Buscar cursos"
            value={searchValue}
            onChange={(e) => setSearchValue(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === 'Enter') {
                e.preventDefault();
                handleSearchAction();
              }
            }}
          />
          <button className="moodle-search-btn" title="Buscar" onClick={handleSearchAction}>
            <FaSearch />
          </button>
          {searchResults !== null && (
            <button
              type="button"
              className="moodle-search-clear"
              title="Cerrar resultados"
              aria-label="Cerrar resultados"
              onClick={() => { setSearchResults(null); setSearchValue(''); setSearchMeta({ count: 0, raw: null }); }}
            >
              ×
            </button>
          )}
        </div>
      </div>

      {loading && <CircularProgress />}
      {error && <Alert severity="error">{error}</Alert>}

      {!loading && !error && (
        <>
          {searchResults === null && (
            <div className="manage-header">{
              topFilter === 'courses' ? 'Gestionar cursos' : (topFilter === 'categories' ? 'Administrar categorías de cursos' : 'Administrar categorías de cursos y cursos')
            }</div>
          )}
          {searchResults !== null ? (
            <div className="search-results-plain">
              <h1 style={{ marginTop: 8, marginBottom: 6 }}>Resultados de la búsqueda</h1>
              <h3 style={{ marginTop: 0, marginBottom: 16, fontWeight: 400 }}>Cursos</h3>
              {loadingAllCourses ? (
                <div style={{ textAlign: 'center', padding: 12 }}>Buscando...</div>
              ) : (
                (!searchResults || searchResultsArray.length === 0) ? (
                  <div className="courses-empty">No se encontraron cursos con las palabras "{(searchValue || '').trim()}"</div>
                ) : (
                  // Wrap rows in a table-like container so rows share borders like a table
                  <div className="courses-list search-results-list">
                    {visibleSearchResults.map((course) => (
                      <div className={`courses-row course-row ${String(course.id) === String(selectedCourseId) ? 'selected' : ''}`} key={course.id}>
                        <div className="row-left">
                          <input
                            type="checkbox"
                            style={{ marginLeft: 0 }}
                            checked={selectedCourseIds.has(String(course.id))}
                            onChange={(e) => { e.stopPropagation(); toggleCourseSelected(course.id); }}
                          />
                          <div className="course-main search-inline" style={{ marginLeft: 12 }}>
                            <a
                              className="course-link"
                              href={`#/course/${course.id}`}
                              onClick={(e) => { e.preventDefault(); e.stopPropagation(); toggleCourseDetails(course); setSelectedCourseId(String(course.id)); }}
                            >
                              {course.fullname || course.displayname || course.shortname}
                            </a>
                            <span className="course-meta">{course.idnumber || ''}</span>
                          </div>
                        </div>
                        <div className="row-right">
                          <button className="icon-btn" title="Editar" onClick={(e) => { e.stopPropagation(); /* open edit UI */ }}><FaCog /></button>
                          <button className="icon-btn" title="Eliminar" onClick={(e) => { e.stopPropagation(); /* delete course */ }}><FaTrash /></button>
                          <button
                            className="icon-btn"
                            title={hiddenCourses.has(course.id) ? 'Mostrar curso' : 'Ocultar curso'}
                            onClick={(e) => handleToggleCourseVisibility(course, e)}
                          >
                            {hiddenCourses.has(course.id) ? <FaEyeSlash /> : <FaEye />}
                          </button>
                        </div>
                      </div>
                    ))}
                  </div>
                )
              )}


              {/* Search meta shown below the results casillas. Only render when we have a valid positive count */}
              {(() => {
                const metaCount = Number.isFinite(Number(searchMeta.count)) ? Number(searchMeta.count) : (Array.isArray(searchResultsArray) ? searchResultsArray.length : 0);
                if (!(metaCount > 0)) return null;
                // compute range display for search results when paginated
                const isSearchPaginated = Number(perPage) > 0 && searchTotalItems > searchPageSize;
                const startDisplaySearch = isSearchPaginated ? (searchStart + 1) : 0;
                const endDisplaySearch = isSearchPaginated ? Math.min(searchEnd, searchTotalItems) : searchTotalItems;
                return (
                  <>
                    <div className="search-results-meta">{isSearchPaginated ? `Mostrando cursos ${startDisplaySearch} a ${endDisplaySearch} de ${searchTotalItems} cursos` : `Mostrando todos ${metaCount} cursos`}</div>
                    {/* show paginator for search results when perPage active and items exceed page size */}
                    {(Number(perPage) > 0 && searchTotalItems > searchPageSize) && renderPaginator(currentPage, searchTotalPages, setCurrentPage, 'sp')}
                  </>
                );
              })()}

              {/* Move controls for search results: same layout as in the courses panel */}
              {(() => {
                const metaCount = Number.isFinite(Number(searchMeta.count)) ? Number(searchMeta.count) : (Array.isArray(searchResultsArray) ? searchResultsArray.length : 0);
                return metaCount > 0 ? (
                  <div className="courses-move-area search-move-area">
                    <div className="move-top">
                      <div className="move-first">Mover los cursos seleccionados a...</div>
                      <div className="move-controls">
                        <select
                          className="move-select"
                          value={moveTargetCategory}
                          onChange={(e) => setMoveTargetCategory(String(e.target.value))}
                          disabled={selectedCourseIds.size === 0}
                          title={selectedCourseIds.size === 0 ? 'Selecciona al menos un curso para activar este filtro' : 'Selecciona la categoría de destino'}
                        >
                          <option value="">Elegir...</option>
                          {(categories || []).map((c) => (
                            <option key={c.id} value={String(c.id)}>{c.name || c.fullname || c.text}</option>
                          ))}
                        </select>
                        <button type="button" className="move-btn" onClick={handleMoveSelectedCourses}>Mover</button>
                      </div>
                    </div>
                  </div>
                ) : null;
              })()}
            </div>
          ) : (
            (() => {
              const onlyCategories = topFilter === 'categories';
              const onlyCourses = topFilter === 'courses';
              // When a category is selected we should reflect the category-loading flag; otherwise
              // use the global all-courses loading state when in the 'courses' topFilter.
              const isLoadingCourses = selectedCategory ? loadingCategoryCourses : (topFilter === 'courses' ? loadingAllCourses : loadingCategoryCourses);
              return (
                <div className={`manage-two-col ${onlyCategories || onlyCourses ? 'single' : ''}`}>
                  {onlyCourses ? null : (
                    <div className={`left-panel ${onlyCategories ? 'full-width' : ''}`}>
                      <div className="panel">
                        <div className="panel-title">Categorías</div>
                        <div className="panel-body categories-panel">
                          <div style={{ textAlign: 'center' }}>
                            <button
                              type="button"
                              className="create-category-btn"
                              onClick={() => {
                                // Solo pasar parentId si hay una categoría seleccionada y no es vacía
                                const state = selectedCategory && selectedCategory !== ''
                                  ? { parentId: selectedCategory }
                                  : {};
                                navigate('/courses/categories/new', { state });
                              }}
                            >
                              Crear nueva categoría
                            </button>
                          </div>
                          <div className="categories-list">
                            {categories.length === 0 ? (
                              <div className="courses-empty">No hay categorías en esta vista</div>
                            ) : (
                              categories.map((cat) => (
                                <div
                                  className={`courses-row category-row ${String(cat.id) === String(selectedCategory) ? 'selected' : ''}`}
                                  key={cat.id}
                                  onClick={() => { if (topFilter === 'categories') setTopFilter(''); setSelectedCategory(String(cat.id)); setExpandedCourse(null); }}
                                  role="button"
                                  tabIndex={0}>
                                  <div className="row-left">
                                    <input
                                      type="checkbox"
                                      checked={selectedCategoryIds.has(String(cat.id))}
                                      onChange={(e) => { e.stopPropagation(); toggleCategorySelected(cat.id); }}
                                      onClick={(e) => e.stopPropagation()}
                                    />
                                    <a
                                      className="course-link" href="#"
                                      onClick={(e) => { e.preventDefault(); e.stopPropagation(); if (topFilter === 'categories') setTopFilter(''); setSelectedCategory(String(cat.id)); setExpandedCourse(null); }}
                                      onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); if (topFilter === 'categories') setTopFilter(''); setSelectedCategory(String(cat.id)); setExpandedCourse(null); } }}
                                      role="button"
                                      tabIndex={0}
                                    >
                                      {cat.name || cat.fullname || cat.text}
                                    </a>
                                  </div>
                                  <div className="row-right">
                                    {/* Category visibility toggle: show/hide category to other users */}
                                    <button
                                      className="icon-btn"
                                      title={cat.visible === false ? 'Mostrar categoría' : 'Ocultar categoría'}
                                      aria-pressed={cat.visible === false ? 'false' : 'true'}
                                      onClick={(e) => { e.stopPropagation(); toggleCategoryVisibility(cat.id); }}
                                    >
                                      <FaEye />
                                    </button>
                                    <button className="icon-btn icon-arrow" title="Subir" onClick={(e) => { e.stopPropagation(); handleMoveCategoryUp(cat.id); }}><FaArrowUp /></button>
                                    <button className="icon-btn icon-arrow" title="Bajar" onClick={(e) => { e.stopPropagation(); handleMoveCategoryDown(cat.id); }}><FaArrowDown /></button>
                                    <button className="icon-btn" title="Editar" onClick={(e) => { e.stopPropagation(); /* open edit UI if implemented */ }}><FaCog /></button>
                                    {/* Show course count aligned to the right; accessible and indicates 'Cursos' on hover */}
                                    <div
                                      className="category-count"
                                      title={`Cursos: ${getCategoryCount(cat)}`}
                                      tabIndex={0}
                                      role="button"
                                      aria-label={`Cursos: ${getCategoryCount(cat)}`}
                                      onClick={(e) => { e.stopPropagation(); /* no-op click - keeps accessibility affordance */ }}
                                      onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); e.stopPropagation(); } }}
                                    >
                                      {getCategoryCount(cat)}
                                    </div>
                                  </div>
                                </div>
                              ))
                            )}
                          </div>

                          {/* Bloque exactamente al final del contenedor blanco: título a la izquierda, filtros centrados y botón debajo */}
                          <div className="categories-ordering" aria-label="Ordenando" style={{ marginTop: 12, paddingLeft: 8 }}>
                            <div className="ordering-title">Ordenando</div>
                            <div className="ordering-controls">
                              <select
                                className="moodle-select ordering-select"
                                aria-label="Filtro 1"
                                value={orderingFilter1}
                                onChange={(e) => setOrderingFilter1(String(e.target.value))}
                              >
                                <option value="">Categorías seleccionadas</option>
                                {(categories || []).map((c) => (
                                  <option key={`o1-${c.id}`} value={String(c.id)}>{c.name || c.fullname || c.text}</option>
                                ))}
                              </select>

                              <select
                                className="moodle-select ordering-select"
                                aria-label="Filtro 2"
                                value={orderingFilter2}
                                onChange={(e) => setOrderingFilter2(String(e.target.value))}
                                disabled={orderingDisabled}
                                title={orderingDisabled ? 'Selecciona al menos una categoría o cambia el primer filtro' : ''}
                                style={orderingDisabled ? { opacity: 0.65, pointerEvents: 'none', background: '#efefef', color: '#8a8a8a' } : {}}
                              >
                                <option value="">Ordenar por Nombre de la categoría ascendente</option>
                                <option value="name_desc">Ordenar por Nombre de la categoría descendente</option>
                              </select>

                              <select
                                className="moodle-select ordering-select"
                                aria-label="Filtro 3"
                                value={orderingFilter3}
                                onChange={(e) => setOrderingFilter3(String(e.target.value))}
                                disabled={orderingDisabled}
                                title={orderingDisabled ? 'Selecciona al menos una categoría o cambia el primer filtro' : ''}
                                style={orderingDisabled ? { opacity: 0.65, pointerEvents: 'none', background: '#efefef', color: '#8a8a8a' } : {}}
                              >
                                <option value="">Ordenar por Nombre completo del curso ascendente</option>
                                <option value="course_desc">Ordenar por Nombre completo del curso descendente</option>
                              </select>

                              <div style={{ marginTop: 10 }}>
                                <button
                                  type="button"
                                  className="create-category-btn ordering-action-btn"
                                  title={'Ordenar'}
                                >
                                  Ordenar
                                </button>
                              </div>
                            </div>
                          </div>

                          {/* Bloque de mover categorías (mismo estilo que en la vista de cursos) */}
                          <div className="courses-move-area categories-move-area" style={{ marginTop: 12 }}>
                            <div className="move-top">
                              <div className="move-first">Mover las categorías seleccionadas a</div>
                              <div className="move-controls">
                                <select
                                  className="move-select"
                                  value={moveTargetCategory}
                                  onChange={(e) => setMoveTargetCategory(String(e.target.value))}
                                  disabled={selectedCategoryIds.size === 0}
                                  title={selectedCategoryIds.size === 0 ? 'Selecciona al menos una categoría para activar este filtro' : 'Selecciona la categoría de destino'}
                                >
                                  <option value="">Elegir...</option>
                                  {(categories || []).map((c) => (
                                    <option key={`mv-${c.id}`} value={String(c.id)}>{c.name || c.fullname || c.text}</option>
                                  ))}
                                </select>
                                <button type="button" className="move-btn" onClick={handleMoveSelectedCategories}>Mover</button>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                  )}
                  {onlyCategories ? null : (
                    <div className={`right-panel ${onlyCourses ? 'full-width' : ''}`}>
                      {/* If a specific category is selected and it has zero courses, don't render the white .panel box.
                          Instead show a small, muted message outside of the panel                                     <a
                                      className="course-link" href="#"
                                      onClick={(e) => { e.preventDefault(); e.stopPropagation(); if (topFilter === 'categories') setTopFilter(''); setSelectedCategory(String(cat.id)); setExpandedCourse(null); }}
                                      onKeyDown={(e) => { if (e.key === 'Enter' || e.key === ' ') { e.preventDefault(); if (topFilter === 'categories') setTopFilter(''); setSelectedCategory(String(cat.id)); setExpandedCourse(null); } }}
                                      role="button"
                                      tabIndex={0}
                                    >
structure. */}
                      <div className="panel">
                        <div className="panel-title">{categories.find(c => String(c.id) === String(selectedCategory))?.name || 'Cursos'}</div>
                        <div className="panel-body courses-panel">
                          {isLoadingCourses ? (
                            <div>Loading...</div>
                          ) : (selectedCategory && coursesSourceArray.length === 0) ? (
                            // Selected category but empty: show small message inside panel-body (no inner white box)
                            <div className="empty-category-message">No hay cursos en esta categoría</div>
                          ) : coursesSourceArray.length === 0 ? (
                            // Generic empty state when no category is selected (e.g. global courses list empty)
                            <div className="empty-category-message">No hay cursos</div>
                          ) : (
                            <>
                              <div className="create-controls" role="toolbar" aria-label="Crear cursos" ref={createControlsRef}>
                                <button
                                  type="button"
                                  className="create-category-btn create-course-btn"
                                  onClick={() => navigate('/courses/new', { state: { parentId: selectedCategory } })}
                                >
                                  Crear nuevo curso
                                </button>

                                <div className={`courses-order ${openMenu === 'order' ? 'open' : ''}`} style={{ position: 'relative' }}>
                                  <button
                                    type="button"
                                    className="orders-toggle"
                                    aria-expanded={openMenu === 'order'}
                                    aria-haspopup="true"
                                    onClick={(e) => { e.stopPropagation(); setOpenMenu(prev => prev === 'order' ? null : 'order'); }}
                                  >
                                    Ordenar cursos
                                    <svg className="caret-icon" width="12" height="8" viewBox="0 0 12 8" xmlns="http://www.w3.org/2000/svg" aria-hidden="true" focusable="false"><path d="M1 1 L6 6 L11 1" stroke="currentColor" strokeWidth="1.6" fill="none" strokeLinecap="round" strokeLinejoin="round" /></svg>
                                  </button>
                                  <div
                                    className="courses-order-menu"
                                    role="menu"
                                    aria-hidden={openMenu !== 'order'}
                                    onClick={(e) => e.stopPropagation()}
                                  >
                                    <button className="order-item" type="button" onClick={() => setOpenMenu(null)}>Ordenar por Nombre completo del curso ascendente</button>
                                    <button className="order-item" type="button" onClick={() => setOpenMenu(null)}>Ordenar por Nombre completo del curso descendente</button>
                                    <button className="order-item" type="button" onClick={() => setOpenMenu(null)}>Ordenar por Nombre corto del curso ascendente</button>
                                    <button className="order-item" type="button" onClick={() => setOpenMenu(null)}>Ordenar por Nombre corto del curso descendente</button>
                                    <button className="order-item" type="button" onClick={() => setOpenMenu(null)}>Ordenar por Número ID del curso ascendente</button>
                                    <button className="order-item" type="button" onClick={() => setOpenMenu(null)}>Ordenar por Número ID del curso descendente</button>
                                    <button className="order-item" type="button" onClick={() => setOpenMenu(null)}>Ordenar por Hora de creación del curso ascendente</button>
                                    <button className="order-item" type="button" onClick={() => setOpenMenu(null)}>Ordenar por Hora de creación del curso descendente</button>
                                  </div>
                                </div>

                                <div className={`courses-order ${openMenu === 'perpage' ? 'open' : ''}`} style={{ position: 'relative' }}>
                                  <button
                                    type="button"
                                    className="orders-toggle"
                                    aria-expanded={openMenu === 'perpage'}
                                    aria-haspopup="true"
                                    onClick={(e) => { e.stopPropagation(); setOpenMenu(prev => prev === 'perpage' ? null : 'perpage'); }}
                                  >
                                    Por página: {perPage > 0 ? perPage : 'Todos'}
                                    <svg className="caret-icon" width="12" height="8" viewBox="0 0 12 8" xmlns="http://www.w3.org/2000/svg" aria-hidden="true" focusable="false"><path d="M1 1 L6 6 L11 1" stroke="currentColor" strokeWidth="1.6" fill="none" strokeLinecap="round" strokeLinejoin="round" /></svg>
                                  </button>
                                  <div className="courses-order-menu" role="menu" aria-hidden={openMenu !== 'perpage'} onClick={(e) => e.stopPropagation()}>
                                    <button className="order-item" type="button" onClick={() => { setPerPage(5); setOpenMenu(null); }}>5</button>
                                    <button className="order-item" type="button" onClick={() => { setPerPage(10); setOpenMenu(null); }}>10</button>
                                    <button className="order-item" type="button" onClick={() => { setPerPage(20); setOpenMenu(null); }}>20</button>
                                    <button className="order-item" type="button" onClick={() => { setPerPage(50); setOpenMenu(null); }}>50</button>
                                    <button className="order-item" type="button" onClick={() => { setPerPage(100); setOpenMenu(null); }}>100</button>
                                    <button className="order-item" type="button" onClick={() => { setPerPage(0); setOpenMenu(null); }}>Todos</button>
                                  </div>
                                </div>
                              </div>
                              {/* Pagination control centered between the create-controls and the white courses-list */}
                              {(Number(perPage) > 0 && totalItems > pageSize) && renderPaginator(currentPage, totalPages, setCurrentPage, 'p')}
                              <div className="courses-list">
                                {visibleCourses.map((course, idx) => (
                                  <div className={`courses-row course-row ${String(course.id) === String(selectedCourseId) ? 'selected' : ''}`} key={course.id}>
                                    <div className="row-left">
                                      <button className="drag-handle custom-black" title="Mover" onClick={(e) => e.stopPropagation()}>
                                        {/* Custom black drag icon (inline SVG) matching provided asset */}
                                        <svg width="20" height="20" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg" aria-hidden="true" focusable="false">
                                          {/* Lines (cross) */}
                                          <path d="M12 6 L12 18" stroke="#000" strokeWidth="1.6" strokeLinecap="round" strokeLinejoin="round" fill="none" />
                                          <path d="M6 12 L18 12" stroke="#000" strokeWidth="1.6" strokeLinecap="round" strokeLinejoin="round" fill="none" />
                                          {/* Arrowheads as filled triangles (black) */}
                                          <path d="M12 3 L8.5 7 L15.5 7 Z" fill="#000" />
                                          <path d="M12 21 L8.5 17 L15.5 17 Z" fill="#000" />
                                          <path d="M3 12 L7 8.5 L7 15.5 Z" fill="#000" />
                                          <path d="M21 12 L17 8.5 L17 15.5 Z" fill="#000" />
                                        </svg>
                                      </button>
                                      <input
                                        type="checkbox"
                                        checked={selectedCourseIds.has(String(course.id))}
                                        onChange={(e) => { e.stopPropagation(); toggleCourseSelected(course.id); }}
                                        style={{ marginLeft: 8 }}
                                      />
                                      <div className="course-main" style={{ marginLeft: 12 }}>
                                        <a
                                          className="course-link"
                                          href={`#/course/${course.id}`}
                                          onClick={(e) => { e.preventDefault(); e.stopPropagation(); toggleCourseDetails(course); setSelectedCourseId(String(course.id)); }}
                                        >
                                          {course.fullname || course.displayname || course.shortname}
                                        </a>
                                      </div>
                                    </div>
                                    <div className="row-right">
                                      <span style={{ display: 'inline-block', width: '120px', textAlign: 'right', marginRight: 16, color: '#777', fontSize: 16 }}>{course.idnumber || ''}</span>
                                      <button className="icon-btn" title="Editar" onClick={(e) => { e.stopPropagation(); /* open edit UI */ }}><FaCog /></button>
                                      <button className="icon-btn" title="Copiar" onClick={(e) => { e.stopPropagation(); /* copy course */ }}><FaCopy /></button>
                                      <button className="icon-btn" title="Eliminar" onClick={(e) => { e.stopPropagation(); /* delete course */ }}><FaTrash /></button>
                                      <button
                                        className="icon-btn"
                                        title={hiddenCourses.has(course.id) ? 'Mostrar curso' : 'Ocultar curso'}
                                        onClick={(e) => handleToggleCourseVisibility(course, e)}
                                      >
                                        {hiddenCourses.has(course.id) ? <FaEyeSlash /> : <FaEye />}
                                      </button>
                                      <button
                                        className="icon-btn icon-arrow"
                                        title="Subir"
                                        onClick={(e) => { e.stopPropagation(); handleMoveCourseUp(course.id); }}
                                        disabled={idx === 0}
                                      ><FaArrowUp /></button>
                                      <button
                                        className="icon-btn icon-arrow"
                                        title="Bajar"
                                        onClick={(e) => { e.stopPropagation(); handleMoveCourseDown(course.id); }}
                                        disabled={idx === (displayedCourses.length - 1)}
                                      ><FaArrowDown /></button>
                                    </div>
                                  </div>
                                ))}
                              </div>
                              {/* Show total count for the selected category below the white courses-list, centered */}
                              {selectedCategory && (
                                (() => {
                                  const catTotal = Number.isFinite(Number(getCategoryCount(categories.find(c => String(c.id) === String(selectedCategory))))) ? Number(getCategoryCount(categories.find(c => String(c.id) === String(selectedCategory)))) : totalItems;
                                  const isCatPaginated = Number(perPage) > 0 && catTotal > pageSize;
                                  const startDisplay = isCatPaginated ? (startIndex + 1) : 0;
                                  const endDisplay = isCatPaginated ? Math.min(endIndex, catTotal) : catTotal;
                                  return (
                                    <div style={{ textAlign: 'center', marginTop: 12, color: '#777', fontWeight: 400 }}>
                                      {isCatPaginated ? `Mostrando cursos ${startDisplay} a ${endDisplay} de ${catTotal} cursos` : `Mostrando todos ${catTotal} cursos`}
                                    </div>
                                  );
                                })()
                              )}
                              {/* When viewing global courses (topFilter === 'courses' and no selectedCategory), show global totals */}
                              {topFilter === 'courses' && !selectedCategory && (
                                (() => {
                                  const globalArr = Array.isArray(allCourses) ? allCourses : asArray(allCourses);
                                  const globalTotal = Array.isArray(globalArr) ? globalArr.length : 0;
                                  const isGlobalPaginated = Number(perPage) > 0 && globalTotal > pageSize;
                                  const startDisplayGlobal = isGlobalPaginated ? (startIndex + 1) : 0;
                                  const endDisplayGlobal = isGlobalPaginated ? Math.min(endIndex, globalTotal) : globalTotal;
                                  return (
                                    <div style={{ textAlign: 'center', marginTop: 12, color: '#777', fontWeight: 400 }}>
                                      {isGlobalPaginated ? `Mostrando cursos ${startDisplayGlobal} a ${endDisplayGlobal} de ${globalTotal} cursos` : `Mostrando todos ${globalTotal} cursos`}
                                    </div>
                                  );
                                })()
                              )}
                              {/* Paginador para la vista de categoría: centrado entre el texto de "Mostrando..." y los filtros */}
                              {(Number(perPage) > 0 && totalItems > pageSize) && renderPaginator(currentPage, totalPages, setCurrentPage, 'cp')}
                              {/* Move control row: left title, centered select + button */}
                              <div className="courses-move-area">
                                <div className="move-top">
                                  <div className="move-first">Mover los cursos seleccionados a...</div>
                                  <div className="move-controls">
                                    <select
                                      className="move-select"
                                      value={moveTargetCategory}
                                      onChange={(e) => setMoveTargetCategory(String(e.target.value))}
                                      disabled={selectedCourseIds.size === 0}
                                      title={selectedCourseIds.size === 0 ? 'Selecciona al menos un curso para activar este filtro' : 'Selecciona la categoría de destino'}
                                    >
                                      <option value="">Elegir...</option>
                                      {(categories || []).map((c) => (
                                        <option key={c.id} value={String(c.id)}>{c.name || c.fullname || c.text}</option>
                                      ))}
                                    </select>
                                    <button type="button" className="move-btn" onClick={handleMoveSelectedCourses}>Mover</button>
                                  </div>
                                </div>
                              </div>
                            </>
                          )}
                        </div>
                      </div>
                    </div>
                  )}
                </div>
              );
            })()
          )}
          {/* Expanded course details panel spanning both columns */}
          {expandedCourse && (
            <div className="panel expanded-course-panel" style={{ marginTop: 16 }}>
              <div className="panel-title" style={{ fontWeight: 400 }}>{expandedCourse.fullname || expandedCourse.displayname || 'Detalle del curso'}</div>
              <div className="panel-body">
                {/* Centered action buttons (Vista, Editar, Usuarios matriculados, Borrar, Ocultar, Copia de seguridad, Restaurar) */}
                <div className="expanded-panel-inner">
                  <div className="expanded-actions" role="toolbar" aria-label="Acciones de curso">
                    <button type="button" className="action-pill" title="Vista">Vista</button>
                    <button type="button" className="action-pill" title="Editar">Editar</button>
                    <button type="button" className="action-pill" title="Usuarios matriculados" onClick={() => navigate(`/courses/${expandedCourse.id}/enrolled`)}>Usuarios matriculados</button>
                    <button type="button" className="action-pill" title="Borrar">Borrar</button>
                    <button type="button" className="action-pill" title="Ocultar">Ocultar</button>
                    <button type="button" className="action-pill" title="Copia de seguridad">Copia de seguridad</button>
                    <button type="button" className="action-pill" title="Restaurar">Restaurar</button>
                  </div>
                </div>

                {/* Two lines below the buttons: render each title and its value on the same line */}
                <div className="expanded-details" aria-hidden={false}>
                  {loadingCourseDetail ? (
                    <div style={{ paddingLeft: 8 }}><CircularProgress size={18} /></div>
                  ) : courseDetailError ? (
                    <div style={{ paddingLeft: 8, color: '#a00' }}>{courseDetailError}</div>
                  ) : courseDetail ? (
                    <>
                      <div className="detail-row">
                        <div className="detail-title">Nombre completo</div>
                        <div className="detail-value">{courseDetail.fullname || '-'}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Nombre corto</div>
                        <div className="detail-value">{courseDetail.shortname || '-'}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Número de ID</div>
                        <div className="detail-value">{courseDetail.idnumber ?? '-'}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Categoría</div>
                        <div className="detail-value">
                          {(() => {
                            const cat = categories.find(c => String(c.id) === String(courseDetail.category));
                            if (cat) {
                              return (
                                <button
                                  type="button"
                                  className="detail-category-link"
                                  onClick={() => {
                                    setTopFilter('');
                                    setSearchResults(null);
                                    setSearchValue('');
                                    setSearchMeta({ count: 0, raw: null });
                                    setSelectedCategory(String(cat.id));
                                    setExpandedCourse(null);
                                    setSelectedCourseId(null);
                                  }}
                                >
                                  {cat.name || cat.fullname}
                                </button>
                              );
                            }
                            // If backend returned a category id or name we can't resolve, still allow clicking to filter by that id
                            if (courseDetail.category !== undefined && courseDetail.category !== null) {
                              const raw = String(courseDetail.category);
                              return (
                                <button
                                  type="button"
                                  className="detail-category-link"
                                  onClick={() => {
                                    setTopFilter('');
                                    setSearchResults(null);
                                    setSearchValue('');
                                    setSearchMeta({ count: 0, raw: null });
                                    setSelectedCategory(raw);
                                    setExpandedCourse(null);
                                    setSelectedCourseId(null);
                                  }}
                                >
                                  {raw}
                                </button>
                              );
                            }
                            return '-';
                          })()}
                        </div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Agrupamientos</div>
                        <div className="detail-value">{courseDetail.groupings ?? 0}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Grupos</div>
                        <div className="detail-value">{courseDetail.groups ?? 0}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Asignaciones de roles</div>
                        <div className="detail-value">{renderRoleAssignments(courseDetail.role_assignments)}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Métodos de matriculación</div>
                        <div className="detail-value">{Array.isArray(courseDetail.enrol_methods) ? renderListNames(courseDetail.enrol_methods, 'enrol') : (courseDetail.enrol_methods === null ? '0' : (courseDetail.enrol_methods ? JSON.stringify(courseDetail.enrol_methods) : '-'))}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Formato</div>
                        <div className="detail-value">{courseDetail.format || '-'}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Secciones</div>
                        <div className="detail-value">{renderSections(courseDetail.sections)}</div>
                      </div>
                      <div className="detail-row">
                        <div className="detail-title">Módulos utilizados</div>
                        <div className="detail-value">{renderModules(courseDetail.modules)}</div>
                      </div>
                    </>
                  ) : (
                    <>
                      <div className="detail-row"><div className="detail-title">Nombre completo</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Nombre corto</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Número de ID</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Categoría</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Agrupamientos</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Grupos</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Asignaciones de roles</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Métodos de matriculación</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Formato</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Secciones</div><div className="detail-value">-</div></div>
                      <div className="detail-row"><div className="detail-title">Módulos utilizados</div><div className="detail-value">-</div></div>
                    </>
                  )}
                </div>
              </div>
            </div>
          )}
        </>
      )}

      {/* Modal removed: enrolled users are shown in their own route /courses/:id/enrolled */}
    </div>
  );
}
