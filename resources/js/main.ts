import 'flowbite';
import './dark-mode';
import './sidebar';
import './page/projects/new';

if (
  localStorage.getItem('color-theme') === 'dark'
  || (!('color-theme' in localStorage)
  && window.matchMedia('(prefers-color-scheme: dark)').matches)
) {
  document.documentElement.classList.add('dark');
} else {
  document.documentElement.classList.remove('dark')
}
