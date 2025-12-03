import React from 'react';
import '../../styles/titulada.css';
import Navbar from '../layouts/Navbar';
import Accesibilidad from '../layouts/Accesibilidad';
import Footer from '../layouts/Footer';
import ProgramCard from '../programs/ProgramCard';
import { getTecnologos, getTecnicos } from '../../data/programasTitulada';

function Titulada() {
    const tecnologos = getTecnologos();
    const tecnicos = getTecnicos();

    return (
        <>
            <Navbar />
            <Accesibilidad />
            <main className="main">
                <div className="main__container">
                    <img 
                        loading="lazy" 
                        src="/img/banners/formacion-titulada.webp" 
                        alt="banner formacion titulada" 
                        className="container__img" 
                    />
                    <div className="container__form">
                        <img 
                            loading="lazy" 
                            src="/img/banners/formacion-banner.webp" 
                            alt="Banner formación" 
                            className="container__form-img" 
                        />
                        <img 
                            loading="lazy" 
                            src="/img/banners/disponibilidad-banner.webp" 
                            alt="Banner disponibilidad" 
                            className="container__form-img" 
                        />
                    </div>

                    <div className="container__curso">
                        {/* Sección Tecnólogos */}
                        <div className="card_container">
                            <div className="programas__header">
                                <img 
                                    loading="lazy" 
                                    src="/img/banners/tecnologo-banner.webp" 
                                    alt="banner tecnologo" 
                                />
                            </div>

                            <div className="programas__body">
                                {tecnologos.map((programa) => (
                                    <ProgramCard
                                        key={programa.id}
                                        title={programa.title}
                                        sniesCode={programa.sniesCode}
                                        pdfUrl={programa.pdfUrl}
                                        videoUrl={programa.videoUrl}
                                    />
                                ))}
                            </div>
                        </div>

                        {/* Sección Técnicos */}
                        <div className="card_container card_container-2">
                            <div className="programas__header">
                                <img 
                                    loading="lazy" 
                                    src="/img/banners/tecnicos-banner.webp" 
                                    alt="banner tecnicos" 
                                />
                            </div>

                            <div className="programas__body">
                                {tecnicos.map((programa) => (
                                    <ProgramCard
                                        key={programa.id}
                                        title={programa.title}
                                        sniesCode={programa.sniesCode}
                                        pdfUrl={programa.pdfUrl}
                                        videoUrl={programa.videoUrl}
                                    />
                                ))}
                            </div>
                        </div>

                        {/* Sección de información adicional */}
                        <div className="certificado">
                            <img 
                                loading="lazy" 
                                src="/img/banners/certificado-banner.webp" 
                                alt="Certificado banner" 
                                className="certificado__img" 
                            />
                            <div className="virtual">
                                <div className="virtual__cont">
                                    <h5 className="virtual__text">
                                        ¿Que es Titulada Virtual y a Distancia?
                                    </h5>
                                    <iframe 
                                        width="100%" 
                                        style={{aspectRatio: '20 / 9', borderRadius: '16px'}} 
                                        src="https://www.youtube.com/embed/HhIH2sPGaNg?si=zl_j5EKEjL6uV2" 
                                        title="YouTube video player" 
                                        frameBorder="0" 
                                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
                                        allowFullScreen
                                    />
                                </div>
                                <div className="virtual__cont-titulada">
                                    <h5 className="virtual__text">
                                        ¿Como Inscribirse a una Titulada?
                                    </h5>
                                    <iframe 
                                        width="100%" 
                                        style={{aspectRatio: '20 / 9', borderRadius: '16px'}} 
                                        src="https://www.youtube.com/embed/CtdNmoWv9bM?si=NXCWBuEy18b2mqs_" 
                                        title="YouTube video player" 
                                        frameBorder="0" 
                                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
                                        allowFullScreen
                                    />
                                </div>
                                <div className="virtual__cont-programas">
                                    <h5 className="virtual__text">
                                        ¿Adecuación de programas?
                                    </h5>
                                    <iframe 
                                        width="100%" 
                                        style={{aspectRatio: '20 / 9', borderRadius: '16px'}} 
                                        src="https://www.youtube.com/embed/gDdA28I6Zsc?si=WtV9AgC6P8DLealL" 
                                        title="YouTube video player" 
                                        frameBorder="0" 
                                        allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" 
                                        allowFullScreen
                                    />
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </main>
            <Footer />
        </>
    );
}

export default Titulada;
