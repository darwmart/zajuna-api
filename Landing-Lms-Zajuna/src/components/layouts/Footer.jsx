import React from 'react';
import '../../styles/layouts/footer.css';

// Rutas públicas

function Footer() {
	return (
		<footer>
			<div className="sector-empleo">
				<div className="sector-empleo__text">
                    <img loading="lazy" src="/img/icons/trabajo-icon.svg" alt="icono-trabajo" />
					<h3>Sector Empleo</h3>
				</div>
			</div>
			<div className="entidades">
                <a target="_blank" href="https://www.mintrabajo.gov.co/web/guest/inicio" className="entidades__link--ministerio" rel="noopener noreferrer"><img loading="lazy" src="/img/logos/ministerio-logo-color.svg" alt="ministerio logo" className="entidades__link-img " /></a>
                <a target="_blank" href="https://www.unidadsolidaria.gov.co/"><img loading="lazy" src="/img/logos/iOss-logo-color.svg" alt="oss logo" className="entidades__link-img " /></a>
                <a target="_blank" href="https://www.serviciodeempleo.gov.co/portada" className="entidades__link" rel="noopener noreferrer"><img loading="lazy" src="/img/logos/iEmpleo-logo-color.svg" alt="empleo logo" className="entidades__link-img " /></a>
                <a target="_blank" href="https://www.ssf.gov.co/" className="entidades__link" rel="noopener noreferrer"><img loading="lazy" src="/img/logos/iSuperSubsidio-logo-color.svg" alt="supersubsidio logo" className="entidades__link-img " /></a>
                <a target="_blank" href="https://www.colpensiones.gov.co/" className="entidades__link" rel="noopener noreferrer"><img loading="lazy" src="/img/logos/iColpensiones-logo-color.svg" alt="solpensiones logo" className="entidades__link-img " /></a>
			</div>
			<div className="gobierno">
				<div className="gobierno__container">
					<div className="gobierno__ministerio-container gobierno__ministerio-container--img">
                        <img loading="lazy" src="/img/logos/logoGovCol-logo.svg" alt="logogovcol logo" className="gobierno__img" />
					</div>
					<div className="gobierno__ministerio-container">
						<a href="https://presidencia.gov.co" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#c61720'}}></span>
							Presidencia</a>
						<a href="https://www.minjusticia.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#01630c'}}></span>
							MinJusticia</a>
						<a href="https://www.mininterior.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#3e6300'}}></span>
							MinInterior</a>
					</div>
					<div className="gobierno__ministerio-container">
						<a href="https://www.mintic.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#990001'}}></span> MinTic</a>
						<a href="https://www.minsalud.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#410e99'}}></span> MinSalud</a>
						<a href="https://www.mincultura.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#38170c'}}></span>
							MinCultura</a>
					</div>
					<div className="gobierno__ministerio-container">
						<a href="https://www.minminas.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#151f99'}}></span> MinMinas</a>
						<a href="https://www.mindefensa.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#531400'}}></span>
							MinDefensa</a>
						<a href="https://www.mineducacion.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#531400'}}></span>
							MinEducación</a>
					</div>
					<div className="gobierno__ministerio-container">
						<a href="https://www.mintrabajo.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#0e3e99'}}></span>
							MinTrabajo</a>
						<a href="https://www.mintransporte.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#5c8301'}}></span>
							MinTransporte</a>
						<a href="https://www.urnadecristal.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#2b1399'}}></span> Urna de
							Cristal</a>
					</div>
					<div className="gobierno__ministerio-container">
						<a href="https://www.minhacienda.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#996201'}}></span>
							MinHacienda</a>
						<a href="https://www.mincit.gov.co/inicio" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#1e7373'}}></span>
							MinComercio</a>
						<a href="https://www.minvivienda.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#992900'}}></span>
							MinVivienda</a>
					</div>
					<div className="gobierno__ministerio-container">
						<a href="https://www.minagricultura.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#3b9901'}}></span>
							MinAgricultura</a>
						<a href="https://www.vicepresidencia.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#919191'}}></span>
							Vicepresidencia</a>
						<a href="https://www.minambiente.gov.co/" className="gobierno__link" target="_blank" rel="noopener noreferrer"><span className="gobierno__ministerios-circle" style={{backgroundColor: '#990001'}}></span>
							MinAmbiente</a>
					</div>
				</div>

			</div>
			<div className="more-information">
				<div className="more-information__container">
					<div className="more-information__text">
						<h1>SERVICIO NACIONAL DE APRENDIZAJE SENA <br/>
						DIRECCIÓN GENERAL </h1>
						Calle 57 No. 8 - 69 Bogotá D.C. (Cundinamarca), Colombia<br />
						EL SENA brinda a la ciudadanía, atención presencial en las 33 Regionales y 117 Centros de Formación<br />
						<a href="http://www.sena.edu.co/es-co/sena/Paginas/directorio.aspx" target="_blank" rel="noopener noreferrer" className="more-information__link">Conozca aquí los puntos de atención</a><br />
						Líneas de atención al ciudadanos, empresarios y línea PQRS: Bogotá +(57) 601 736 60 60 - Línea gratuita y resto del país 018000 910270
					</div>
					<div className="more-information__logos">
						<img loading="lazy" className="more-information__sena-logo" src="/img/logos/sena-logo.svg" alt="Logo SENA" />
						<img loading="lazy" className="more-information__iso-logos" src="/img/logos/normas-iso-logos.svg" alt="Certificaciones ISO" title="Certificaciones ISO" />
					</div> 
				</div>
				<div className="footer__redes-container">
					<div className="footer__redes">
						<div className="footer-social__item">
							<svg xmlns="http://www.w3.org/2000/svg" height="14" width="14" viewBox="0 0 512 512" aria-hidden="true">
								<path fill="#159013" d="M512 256C512 114.6 397.4 0 256 0S0 114.6 0 256C0 376 82.7 476.8 194.2 504.5V334.2H141.4V256h52.8V222.3c0-87.1 39.4-127.5 125-127.5c16.2 0 44.2 3.2 55.7 6.4V172c-6-.6-16.5-1-29.6-1c-42 0-58.2 15.9-58.2 57.2V256h83.6l-14.4 78.2H287V510.1C413.8 494.8 512 386.9 512 256h0z" />
							</svg>
							<a href="https://www.facebook.com/SENA/" target="_blank" rel="noopener noreferrer" className="footer__redes-text">@SENA</a>
						</div>
						<div className="footer-social__item">
							<svg xmlns="http://www.w3.org/2000/svg" height="14" width="14" viewBox="0 0 512 512" aria-hidden="true">
								<path fill="#159013" d="M389.2 48h70.6L305.6 224.2 487 464H345L233.7 318.6 106.5 464H35.8L200.7 275.5 26.8 48H172.4L272.9 180.9 389.2 48zM364.4 421.8h39.1L151.1 88h-42L364.4 421.8z" />
							</svg>
							<a href="https://twitter.com/SENAComunica" target="_blank" rel="noopener noreferrer" className="footer__redes-text">@SENACOMUNICA</a>
						</div>
						<div className="footer-social__item">
							<svg xmlns="http://www.w3.org/2000/svg" height="14" width="14" viewBox="0 0 448 512" aria-hidden="true">
								<path fill="#159013" d="M224.1 141c-63.6 0-114.9 51.3-114.9 114.9s51.3 114.9 114.9 114.9c63.6 0 114.9-51.3 114.9-114.9s-51.3-114.9-114.9-114.9zm0 189.6c-41.1 0-74.7-33.5-74.7-74.7s33.5-74.7 74.7-74.7 74.7 33.5 74.7 74.7-33.6 74.7-74.7 74.7zm146.4-194.3c0 14.9-12 26.8-26.8 26.8-14.9 0-26.8-12-26.8-26.8s12-26.8 26.8-26.8 26.8 12 26.8 26.8zm76.1 27.2c-1.7-35.9-9.9-67.7-36.2-93.9-26.2-26.2-58-34.4-93.9-36.2-37-2.1-147.9-2.1-184.9 0-35.8 1.7-67.6 9.9-93.9 36.1s-34.4 58-36.2 93.9c-2.1 37-2.1 147.9 0 184.9 1.7 35.9 9.9 67.7 36.2 93.9s58 34.4 93.9 36.2c37 2.1 147.9 2.1 184.9 0 35.9-1.7 67.7-9.9 93.9-36.2 26.2-26.2 34.4-58 36.2-93.9 2.1-37 2.1-147.8 0-184.8zM398.8 388c-7.8 19.6-22.9 34.7-42.6 42.6-29.5 11.7-99.5 9-132.1 9s-102.7 2.6-132.1-9c-19.6-7.8-34.7-22.9-42.6-42.6-11.7-29.5-9-99.5-9-132.1s-2.6-102.7 9-132.1c7.8-19.6 22.9-34.7 42.6-42.6 29.5-11.7 99.5-9 132.1-9s102.7-2.6 132.1 9c19.6 7.8 34.7 22.9 42.6 42.6 11.7 29.5 9 99.5 9 132.1s2.7 102.7-9 132.1z" />
							</svg>
							<a href="https://www.instagram.com/senacomunica/" target="_blank" rel="noopener noreferrer" className="footer__redes-text">@SENACOMUNICA</a>
						</div>
						<div className="footer-social__item">
							<svg xmlns="http://www.w3.org/2000/svg" height="14" width="14" viewBox="0 0 576 512" aria-hidden="true">
								<path fill="#159013" d="M549.7 124.1c-6.3-23.7-24.8-42.3-48.3-48.6C458.8 64 288 64 288 64S117.2 64 74.6 75.5c-23.5 6.3-42 24.9-48.3 48.6-11.4 42.9-11.4 132.3-11.4 132.3s0 89.4 11.4 132.3c6.3 23.7 24.8 41.5 48.3 47.8C117.2 448 288 448 288 448s170.8 0 213.4-11.5c23.5-6.3 42-24.2 48.3-47.8 11.4-42.9 11.4-132.3 11.4-132.3s0-89.4-11.4-132.3zm-317.5 213.5V175.2l142.7 81.2-142.7 81.2z" />
							</svg>
							<a href="https://www.youtube.com/user/SENATV" target="_blank" rel="noopener noreferrer" className="footer__redes-text">@SENATV</a>
						</div>
						<div className="footer-social__item">
							<svg xmlns="http://www.w3.org/2000/svg" height="14" width="14" viewBox="0 0 448 512" aria-hidden="true">
								<path fill="#159013" d="M416 32H31.9C14.3 32 0 46.5 0 64.3v383.4C0 465.5 14.3 480 31.9 480H416c17.6 0 32-14.5 32-32.3V64.3c0-17.8-14.4-32.3-32-32.3zM135.4 416H69V202.2h66.5V416zm-33.2-243c-21.3 0-38.5-17.3-38.5-38.5S80.9 96 102.2 96c21.2 0 38.5 17.3 38.5 38.5 0 21.3-17.2 38.5-38.5 38.5zm282.1 243h-66.4V312c0-24.8-.5-56.7-34.5-56.7-34.6 0-39.9 27-39.9 54.9V416h-66.4V202.2h63.7v29.2h.9c8.9-16.8 30.6-34.5 62.9-34.5 67.2 0 79.7 44.3 79.7 101.9V416z" />
							</svg>
							<a href="https://www.linkedin.com/school/servicio-nacional-de-aprendizaje-sena-/" target="_blank" rel="noopener noreferrer" className="footer__redes-text">SENA</a>
						</div>
						<div className="footer-social__item">
							<svg xmlns="http://www.w3.org/2000/svg" height="14" width="14" viewBox="0 0 448 512" aria-hidden="true">
								<path fill="#159013" d="M448 209.9a210.1 210.1 0 0 1 -122.8-39.3V349.4A162.6 162.6 0 1 1 185 188.3V278.2a74.6 74.6 0 1 0 52.2 71.2V0l88 0a121.2 121.2 0 0 0 1.9 22.2h0A122.2 122.2 0 0 0 381 102.4a121.4 121.4 0 0 0 67 20.1z" />
							</svg>
							<a href="https://www.tiktok.com/@senacomunica_" target="_blank" rel="noopener noreferrer" className="footer__redes-text">@senacomunica</a>
						</div>
						<div className="footer-social__item">
							<svg xmlns="http://www.w3.org/2000/svg" height="14" width="14" viewBox="0 0 448 512" aria-hidden="true">
								<path fill="#159013" d="M380.9 97.1C339 55.1 283.2 32 223.9 32c-122.4 0-222 99.6-222 222 0 39.1 10.2 77.3 29.6 111L0 480l117.7-30.9c32.4 17.7 68.9 27 106.1 27h.1c122.3 0 224.1-99.6 224.1-222 0-59.3-25.2-115-67.1-157zm-157 341.6c-33.2 0-65.7-8.9-94-25.7l-6.7-4-69.8 18.3L72 359.2l-4.4-7c-18.5-29.4-28.2-63.3-28.2-98.2 0-101.7 82.8-184.5 184.6-184.5 49.3 0 95.6 19.2 130.4 54.1 34.8 34.9 56.2 81.2 56.1 130.5 0 101.8-84.9 184.6-186.6 184.6zm101.2-138.2c-5.5-2.8-32.8-16.2-37.9-18-5.1-1.9-8.8-2.8-12.5 2.8-3.7 5.6-14.3 18-17.6 21.8-3.2 3.7-6.5 4.2-12 1.4-32.6-16.3-54-29.1-75.5-66-5.7-9.8 5.7-9.1 16.3-30.3 1.8-3.7 .9-6.9-.5-9.7-1.4-2.8-12.5-30.1-17.1-41.2-4.5-10.8-9.1-9.3-12.5-9.5-3.2-.2-6.9-.2-10.6-.2-3.7 0-9.7 1.4-14.8 6.9-5.1 5.6-19.4 19-19.4 46.3 0 27.3 19.9 53.7 22.6 57.4 2.8 3.7 39.1 59.7 94.8 83.8 35.2 15.2 49 16.5 66.6 13.9 10.7-1.6 32.8-13.4 37.4-26.4 4.6-13 4.6-24.1 3.2-26.4-1.3-2.5-5-3.9-10.5-6.6z" />
							</svg>
							<a href="https://api.whatsapp.com/send/?phone=573168760255&text&app_absent=0" target="_blank" rel="noopener noreferrer" className="footer__redes-text">whatsapp: 3168760255</a>
						</div>
					</div>
					<br />
					<br />
					<div className="footer__terminos">
						<a href="https://www.sena.edu.co/es-co/sena/Paginas/directorio.aspx" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Directorio SENA</a>
						<a href="http://sciudadanos.sena.edu.co/SolicitudIndex.aspx" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">PQRS</a>
						<a href="https://www.sena.edu.co/es-co/ciudadano/Paginas/chat.aspx" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Chat en línea</a>
						<a href="https://www.sena.edu.co/es-co/ciudadano/Paginas/Denuncias_Corrupcion.aspx" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Denuncias por actos de corrupción</a>
						<a href="https://www.sena.edu.co/es-co/transparencia/Paginas/mecanismosContacto.aspx#notificacionesJudiciales" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Notificaciones judiciales</a>
						<a href="https://www.sena.edu.co/es-co/Paginas/politicasCondicionesUso.aspx" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Términos y condiciones del portal web</a>
						<a href="https://www.sena.edu.co/es-co/Paginas/politicasCondicionesUso.aspx#derechoAutor" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Derechos de autor y/o autorización de uso sobre contenidos</a>
						<a href="https://www.sena.edu.co/es-co/transparencia/Paginas/habeas_data.aspx" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Política de Tratamiento para Protección de Datos Personales</a>
						<a href="https://www.sena.edu.co/es-co/Paginas/mapaSitio.aspx" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Mapa del sitio</a>
						<a href="http://compromiso.sena.edu.co/index.php?text=inicio&id=27" target="_blank" rel="noopener noreferrer" className="footer__terminos-text">Política de seguridad y privacidad de la información</a>
					</div>
				</div>
			</div>

			<div className="gov" id="final">
				<div className="gov__container">
					<a href="https://www.gov.co/" target="_blank" rel="noopener noreferrer">
                        <img loading="lazy" src="/img/logos/gov-logo.svg" alt="Logo de la pagina gov.co" className="gov__img" />
					</a>

				</div>
			</div>
		</footer>
	);
}

export default Footer;
