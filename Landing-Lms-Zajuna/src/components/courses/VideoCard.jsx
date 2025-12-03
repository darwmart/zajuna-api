import React from 'react';

function VideoCard({ videoSrc, alt = "Course video" }) {
  return (
    <video 
      className="cursos-virtuales__video"
      autoPlay 
      muted 
      loop 
      playsInline
      aria-label={alt}
    >
      <source src={videoSrc} type="video/mp4" />
    </video>
  );
}

// Memoizar para evitar re-renders innecesarios
export default React.memo(VideoCard);

