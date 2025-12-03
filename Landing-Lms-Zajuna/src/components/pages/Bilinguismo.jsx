import React from 'react';
import '../../styles/bilinguismo.css';
import Navbar from '../layouts/Navbar';
import Accesibilidad from '../layouts/Accesibilidad';
import EnglishLevel from '../courses/EnglishLevel';
import { getEnglishLevels } from '../../data/englishLevels';

function Bilinguismo() {
	const englishLevels = getEnglishLevels();

	return (
		<>
			<Navbar />
			<Accesibilidad />
			<main className="bilinguismo">
				<div className="bilinguismo__container">
					{/* Banner de oportunidad */}
					<div className="bilinguismo__oportunidad">
						<img 
							loading="lazy" 
							src="/img/banners/oportunidad-banner.webp" 
							alt="oportunidad banner" 
							className="bilinguismo__oportunidad-img" 
						/>
					</div>
					
					{/* Grid de niveles de inglés */}
					<div className="bilinguismo__ingles-cards">
						{englishLevels.map((level) => (
							<EnglishLevel
								key={level.id}
								image={level.image}
								alt={level.alt}
								description={level.description}
								href={level.href}
							/>
						))}
					</div>
					<div className="bilinguismo__informacion">
						<img loading="lazy" src="/img/banners/ingles-banner.webp" alt="ingles oportunidad banner" className="bilinguismo__informacion-card bilinguismo__informacion-card--none" />
						<div className="bilinguismo__informacion-card">
							<div className="bilinguismo__informacion-card-header">
								<h3 className="bilinguismo__information-card-title">¿Qué es English Does Work?</h3>
							</div>
							<div className="bilinguismo__informacion-card-container">
								<img loading="lazy" src="/img/icons/conveersacion-icon.webp" alt="conversacion icono" className="bilinguismo__informacion-card-icon" />
								<p className="bilinguismo__informacion-card-text">Es el programa de formación de inglés virtual que
									ofrece el SENA, una alternativa en línea cuyo, objetivo principal es llegar a todas las
									regiones del país y fortalecer las habilidades lingüísticas y comunicativas en inglés de los
									colombianos.</p>
							</div>
						</div>
						<div className="bilinguismo__informacion-card">
							<div className="bilinguismo__informacion-card-header">
								<h3 className="bilinguismo__information-card-title">¿Quiénes pueden acceder a este Programa de
									Formación virtual?</h3>
							</div>
							<div className="bilinguismo__informacion-card-container">
								<img loading="lazy" src="/img/icons/personas-icon.webp" alt=" grupo icono" className="bilinguismo__informacion-card-icon" />
								<p className="bilinguismo__informacion-card-text">English Does Work está compuesto por 13 niveles
									que inician en level 1 y terminan con un level 13. Dichos niveles tienen relación con el
									Marco Común Europeo de Referencia para las Lenguas (MCERL) y al finalizar la totalidad del
									Programa de Formación, es decir los 13 niveles, se espera que el aprendiz pueda desarrollar
									un nivel intermedio de competencia en inglés.</p>
							</div>
						</div>
					</div>
				</div>
			</main>
		</>
	);
}

export default Bilinguismo;
