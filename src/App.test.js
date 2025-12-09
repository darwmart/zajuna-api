import { render, screen } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import App from './App';

test('renders learn react link', () => {
  render(
    <MemoryRouter>
      <App />
    </MemoryRouter>
  );
  // The app header contains 'SENA - Zajuna' in the top bar
  const headerText = screen.getByText(/SENA - Zajuna/i);
  expect(headerText).toBeInTheDocument();
});
