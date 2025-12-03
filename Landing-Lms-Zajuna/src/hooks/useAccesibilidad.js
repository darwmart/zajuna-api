import { useEffect } from 'react';

export default function useAccesibilidad() {
	useEffect(() => {
		// Original variable names preserved
		const toggleButton = document.getElementById('toggle-contrast');
		const rootElement = document.documentElement;
		const zoomInButton = document.getElementById('toggle-zoom-in');
		const zoomOutButton = document.getElementById('toggle-zoom-out');

		function applyImageContrast() {
			const images = document.querySelectorAll('iframe, img:not(.entidades__link-img):not(.entidades__link-img:hover)');
			images.forEach(img => {
				img.style.filter = rootElement.classList.contains('contrast') ? 'grayscale(100%)' : 'none';
			});
		}

		// Initialize contrast setting from localStorage
		const onDOMContentLoaded = () => {
			const contrastSetting = localStorage.getItem('contrastSetting');
			if (contrastSetting === 'true') {
				rootElement.classList.add('contrast');
				applyImageContrast();
			}
		};
		document.addEventListener('DOMContentLoaded', onDOMContentLoaded);

		// Contrast toggle
		const onToggleContrast = () => {
			rootElement.classList.toggle('contrast');
			applyImageContrast();
			localStorage.setItem('contrastSetting', rootElement.classList.contains('contrast'));
		};
		toggleButton && toggleButton.addEventListener('click', onToggleContrast);

		// Zoom controls
		const MIN_FONT_SIZE = 13;
		const MAX_FONT_SIZE = 20;

		const onZoomIn = () => {
			let fontSize = parseFloat(getComputedStyle(document.documentElement).fontSize);
			fontSize += 1;
			fontSize = Math.min(MAX_FONT_SIZE, fontSize);
			document.documentElement.style.fontSize = fontSize + 'px';
		};
		const onZoomOut = () => {
			let fontSize = parseFloat(getComputedStyle(document.documentElement).fontSize);
			fontSize -= 1;
			fontSize = Math.max(MIN_FONT_SIZE, fontSize);
			document.documentElement.style.fontSize = fontSize + 'px';
		};
		zoomInButton && zoomInButton.addEventListener('click', onZoomIn);
		zoomOutButton && zoomOutButton.addEventListener('click', onZoomOut);

		return () => {
			document.removeEventListener('DOMContentLoaded', onDOMContentLoaded);
			toggleButton && toggleButton.removeEventListener('click', onToggleContrast);
			zoomInButton && zoomInButton.removeEventListener('click', onZoomIn);
			zoomOutButton && zoomOutButton.removeEventListener('click', onZoomOut);
		};
	}, []);
}
