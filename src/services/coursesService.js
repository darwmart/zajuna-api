import apiClientCourses from "../api/apiClientCourses";

// Helper: try to extract a course-detail object from various possible wrappers
const extractDetail = (data) => {
  if (!data) return null;
  // If it's an array, try to find the first item that looks like a course
  if (Array.isArray(data)) {
    for (const it of data) {
      if (it && (it.id || it.fullname || it.shortname)) return it;
    }
    return data[0] || null;
  }

  // Known wrapper keys that may contain the actual object
  const wrapperKeys = ['course', 'item', 'data', 'result', 'payload', 'body'];
  for (const k of wrapperKeys) {
    if (data[k]) {
      const found = extractDetail(data[k]);
      if (found) return found;
    }
  }

  // Arrays under common properties
  if (Array.isArray(data.courses) && data.courses.length) return extractDetail(data.courses);
  if (Array.isArray(data.items) && data.items.length) return extractDetail(data.items);

  // If the object itself looks like a course, return it
  if (data.id || data.fullname || data.shortname) return data;

  // Otherwise return the original payload as fallback
  return data;
};

export const coursesService = {
  getAll: async () => {
    // Try a few common variants to request the full course list from backends
    // that paginate by default. Many Moodle-like APIs support a 'perpage' or
    // similar param (0 = all) or an 'includehidden' flag. Try them in order
    // and fall back to the simplest call.
    const attempts = [
      { params: { perpage: 0 } },
      { params: { limit: 0 } },
      { params: { includehidden: true, perpage: 0 } },
      { params: { includehidden: true } },
      {},
    ];
    for (const opt of attempts) {
      try {
        const res = await apiClientCourses.get("/courses", opt);
        if (!res || !res.data) continue;
        const data = res.data;

        // Normalizar respuesta: { courses: [...], pagination: {...} } → array
        const arr = Array.isArray(data) ? data : (data.courses || data.items || data || []);

        // If we got more than a small page (e.g. > 50) or any non-empty result,
        // accept it as the full list. Otherwise continue trying other params.
        if (Array.isArray(arr) && arr.length > 0) return arr;
        // If it's an object but has items/courses property as empty, still accept
        if (data && (data.items || data.courses)) return arr;
      } catch (err) {
        // ignore and try next
      }
    }
    // As a last resort, do the plain call (shouldn't normally reach here)
    const res = await apiClientCourses.get("/courses");

    // Normalizar respuesta
    const data = res.data;
    return Array.isArray(data) ? data : (data.courses || data.items || []);
  },

  getEnrolledUsers: async (courseId, options = {}) => {
    // Endpoint correcto: /enrollments/course/:courseid
    const res = await apiClientCourses.get(`/enrollments/course/${courseId}`, {
      params: options, // Permite pasar opciones como sortby, limitnumber, onlyactive, etc.
    });
    return res.data;
  },

  getUserCourses: async (userId) => {
    const res = await apiClientCourses.get("/courses", {
      params: { userid: userId },
    });
    return res.data;
  },

  getCategories: async () => {
    const res = await apiClientCourses.get("/categories");

    // Normalizar respuesta del backend
    // El backend puede devolver: { categories: [...] } o un array directo
    const response = res.data;

    // Si el backend devuelve un objeto con categories, devolver ese array
    if (response && response.categories) {
      return response.categories;
    }

    // Si es un array directo, devolverlo tal cual
    if (Array.isArray(response)) {
      return response;
    }

    // Fallback: devolver respuesta original
    return response;
  },

  getCoursesByField: async (field, value) => {
    // backend expects specific query params based on field type
    const params = {};

    // Map field names to backend-expected parameter names
    if (field === "category") {
      params.categoryid = value;
    } else {
      // For other fields, pass as-is
      // eslint-disable-next-line no-unused-expressions
      params[field] = value; incl
    }

    const res = await apiClientCourses.get("/courses", { params });

    // Normalizar respuesta del backend
    // El backend devuelve: { courses: [...], pagination: {...} }
    const response = res.data;

    // Si el backend devuelve un objeto con courses, devolver ese array
    if (response && response.courses) {
      return response.courses;
    }

    // Si es un array directo, devolverlo tal cual
    if (Array.isArray(response)) {
      return response;
    }

    // Fallback: devolver respuesta original
    return response;
  },
  // Fetch detailed course information from backend
  getCourseDetail: async (courseId) => {
    // Accept either a plain idnumber or an object (course) and extract idnumber
    const resolveIdNumber = (arg) => {
      if (arg === null || arg === undefined) return null;
      if (typeof arg === 'object') {
        // Priorizar idnumber sobre id
        return arg.idnumber ?? arg.idNumber ?? arg.id ?? arg.courseid ?? arg.courseId ?? null;
      }
      return arg;
    };
    const idnumber = resolveIdNumber(courseId);
    if (!idnumber) {
      throw new Error('idnumber requerido');
    }

    // Construir URL con el IDNumber (Axios no reemplaza :idnumber automáticamente)
    const url = `/courses/${idnumber}/details`;

    try {
      const res = await apiClientCourses.get(url);
      // Normalize and return a sensible object (or the raw payload as fallback)
      return extractDetail(res.data) || res.data;
    } catch (err) {
      // Re-throw error to be handled by caller
      throw err;
    }
  },

  // Try a server-side search endpoint first; if unavailable, fall back to fetching all and filtering client-side
  search: async (q) => {
    // Prefer a server-side search endpoint when available, but always ensure
    // a global result by falling back to fetching all courses and filtering
    // client-side. This guarantees the search input can find any course in
    // the database even if the backend search is restricted.
    const ql = (q || '').toLowerCase();
    try {
      // Use new search endpoint with criterianame and criteriavalue
      const res = await apiClientCourses.get("/courses/search", {
        params: {
          criterianame: 'search',
          criteriavalue: q
        }
      });
      const data = res.data;
      // Normalizar: { courses: [...], total: N } → array
      const list = Array.isArray(data) ? data : (data.courses || data.items || data || []);
      // If server returned nothing, fall back to client-side global filtering
      if (!list || list.length === 0) {
        throw new Error('server-empty');
      }
      return list;
    } catch (err) {
      // Fallback: retrieve all and filter client-side, request includehidden if supported
      const all = await apiClientCourses.get("/courses", { params: { includehidden: true } });
      const response = all.data || [];
      // Normalizar: { courses: [...] } → array
      const arr = Array.isArray(response) ? response : (response.courses || response.items || response || []);
      return arr.filter((c) => {
        const name = (c.fullname || c.displayname || c.shortname || c.idnumber || '').toLowerCase();
        return name.includes(ql);
      });
    }
  },

  // Advanced search with specific criteria (categoryid, id, idnumber, search)
  searchCourses: async (criteriaName, criteriaValue, page = 0, perPage = 0) => {
    const res = await apiClientCourses.get("/courses/search", {
      params: {
        criterianame: criteriaName,
        criteriavalue: criteriaValue,
        page,
        perpage: perPage
      }
    });
    return res.data;
  },

  toggleVisibility: async (courseId) => {
    const res = await apiClientCourses.patch(`/courses/${courseId}/visibility`);
    return res.data;
  },


  // Ocultar curso - cambia visible a 0 usando PUT
  hideCourse: async (courseId) => {
    const res = await apiClientCourses.put("/courses", {
      courses: [{ id: courseId, visible: 0 }]
    });
    return res.data;
  },

  // Mostrar curso  cambia visible a 1 usando PUT
  showCourse: async (courseId) => {
    const res = await apiClientCourses.put("/courses", {
      courses: [{ id: courseId, visible: 1 }]
    });
    return res.data;
  },

  // Mover categoría - reordena categorías
  moveCategory: async (categoryId, beforeId) => {
    const res = await apiClientCourses.post("/categories/move", {
      id: categoryId,
      beforeid: beforeId
    });
    return res.data;
  },

  // Mover curso - reordena cursos
  moveCourse: async (courseId, categoryId, beforeId) => {
    const res = await apiClientCourses.post("/courses/move", {
      courses: [{
        id: courseId,
        categoryid: categoryId,
        beforeid: beforeId
      }]
    });
    return res.data;
  },

  // Crear categoría - crea una o más categorías
  createCategory: async (data) => {
    const res = await apiClientCourses.post("/categories", data);
    return res.data;
  },
};

