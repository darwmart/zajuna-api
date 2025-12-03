import React from 'react';
import VideoCard from './VideoCard';
import CourseCardContent from './CourseCardContent';

function CourseCard({ 
  course, 
  isHovered, 
  onMouseEnter, 
  onMouseLeave 
}) {
  return (
    <div 
      className="cursos-virtuales__card"
      style={course.videoSrc ? {} : { 
        backgroundImage: `url('${course.img}')`,
        backgroundSize: 'cover',
        backgroundPosition: 'center',
        backgroundRepeat: 'no-repeat'
      }}
      onMouseEnter={onMouseEnter}
      onMouseLeave={onMouseLeave}
    >
      {/* Renderizar video si existe */}
      {course.videoSrc && <VideoCard videoSrc={course.videoSrc} alt={course.title} />}

      {/* Contenido del overlay */}
      <CourseCardContent course={course} isHovered={isHovered} />
    </div>
  );
}

// Memoizar para evitar re-renders innecesarios
export default React.memo(CourseCard);

