import React from 'react';
import '../../styles/layouts/accesibilidad.css';
import useAccesibilidad from '../../hooks/useAccesibilidad';

// Usar rutas públicas

function Accesibilidad() {
	useAccesibilidad();
	return (
		<div className="accessibility">
			<div className="accessibility__container">
				<button className="accessibility__button" id="toggle-contrast">
					<img loading="lazy" src="/img/icons/contraste-icon.png" alt="contraste icono" className="accessibility__icon" />
				</button>
				<span className="accessibility__text">Contraste</span>
			</div>
			<div className="accessibility__container">
				<button className="accessibility__button" id="toggle-zoom-in">
					<img loading="lazy" src="/img/icons/letra-mas-icon.png" alt="aumento letra icono" className="accessibility__icon" />
				</button>
				<span className="accessibility__text">Aumentar letra</span>
			</div>
			<div className="accessibility__container">
				<button className="accessibility__button" id="toggle-zoom-out">
					<img loading="lazy" src="/img/icons/letra-menos-icon.png" alt="disminuir letra icono" className="accessibility__icon" />
				</button>
				<span className="accessibility__text">Disminuir letra</span>
			</div>
			<div className="accessibility__container">
				<button className="accessibility__button">
						<a href="#inicio"><img loading="lazy" src="/img/icons/arriba-icon.png" alt="flecha arriba icono" className="accessibility__icon" /></a>
				</button>
				<span className="accessibility__text">Subir</span>
			</div>
			<div className="accessibility__container">
				<button className="accessibility__button">
						<a href="#final"><img loading="lazy" src="/img/icons/abajo-icon.png" alt="flecha abajo icono" className="accessibility__icon" /></a>
				</button>
				<span className="accessibility__text">Bajar</span>
			</div>
		</div>
	);
}

export default Accesibilidad;
