import { useEffect } from 'react';

export default function useSlider() {
	useEffect(() => {
		// Variables as in original
		let slides = document.querySelectorAll('.slider__img');
		let totalSlides = slides.length;
		let currentIndex = 0;
		let intervalId = null;
		let touchStartX = 0;
		let touchEndX = 0;

		const indicatorsContainer = document.getElementById('slider__indicators');
		const itemsContainer = document.getElementById('slider__items');
		const sliderContainer = document.getElementById('slider__container');

		function updateIndicators() {
			const buttons = indicatorsContainer ? indicatorsContainer.querySelectorAll('button') : [];
			buttons.forEach((button, index) => {
				button.classList.toggle('active', index === currentIndex);
			});
		}

		function changeSlide(index) {
			currentIndex = index;
			if (!slides || slides.length === 0 || !itemsContainer) return;
			const slideWidth = slides[0].offsetWidth;
			const offset = -slideWidth * index;
			itemsContainer.style.transform = `translateX(${offset}px)`;
			updateIndicators();
			restartAutoPlayAfterDelay();
		}

		function startAutoPlay() {
			clearInterval(intervalId);
			intervalId = setInterval(() => {
				currentIndex = (currentIndex + 1) % totalSlides;
				changeSlide(currentIndex);
			}, 5000);
		}

		function restartAutoPlayAfterDelay() {
			clearTimeout(intervalId);
			setTimeout(() => {
				startAutoPlay();
			}, 2000);
		}

		function handleSwipe() {
			const sensitivity = 50;
			if (touchEndX < touchStartX - sensitivity) {
				currentIndex = (currentIndex + 1) % totalSlides;
				changeSlide(currentIndex);
			} else if (touchEndX > touchStartX + sensitivity) {
				currentIndex = (currentIndex - 1 + totalSlides) % totalSlides;
				changeSlide(currentIndex);
			}
		}

		function createIndicators() {
			if (!indicatorsContainer) return [];
			indicatorsContainer.innerHTML = '';
			const buttons = [];
			for (let i = 0; i < totalSlides; i++) {
				const button = document.createElement('button');
				button.addEventListener('click', () => changeSlide(i));
				indicatorsContainer.appendChild(button);
				buttons.push(button);
			}
			updateIndicators();
			return buttons;
		}

		// Initialize
		slides = document.querySelectorAll('.slider__img');
		totalSlides = slides.length;
		createIndicators();
		startAutoPlay();

		const onTouchStart = (e) => {
			touchStartX = e.changedTouches[0].clientX;
		};
		const onTouchEnd = (e) => {
			touchEndX = e.changedTouches[0].clientX;
			handleSwipe();
		};

		sliderContainer && sliderContainer.addEventListener('touchstart', onTouchStart);
		sliderContainer && sliderContainer.addEventListener('touchend', onTouchEnd);

		return () => {
			// Cleanup
			if (indicatorsContainer) {
				const buttons = indicatorsContainer.querySelectorAll('button');
				buttons.forEach((btn) => {
					btn.replaceWith(btn.cloneNode(true)); // remove listeners
				});
				indicatorsContainer.innerHTML = '';
			}
			clearInterval(intervalId);
			sliderContainer && sliderContainer.removeEventListener('touchstart', onTouchStart);
			sliderContainer && sliderContainer.removeEventListener('touchend', onTouchEnd);
		};
	}, []);
}
