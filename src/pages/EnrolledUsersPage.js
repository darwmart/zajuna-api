import React, { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import EnrolledUsersView from '../components/EnrolledUsersView';
import { coursesService } from '../services/coursesService';

export default function EnrolledUsersPage() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [course, setCourse] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    let mounted = true;
    const load = async () => {
      try {
        const data = await coursesService.getCourseDetail({ id });
        const detail = Array.isArray(data) ? data[0] : (data && (data.course || data.item || data.data)) || data;
        if (!mounted) return;
        setCourse(detail || { id });
      } catch (e) {
        console.error('failed to load course', e);
        if (mounted) setCourse({ id });
      } finally { if (mounted) setLoading(false); }
    };
    load();
    return () => { mounted = false; };
  }, [id]);

  return (
    <div style={{ padding: 0, width: '100%', boxSizing: 'border-box' }}>
      <EnrolledUsersView course={course} />
    </div>
  );
}
