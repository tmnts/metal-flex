import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import tailwindcss from '@tailwindcss/vite'; // <--- ВОТ ЭТА ШТУКА!

export default defineConfig({
	plugins: [
		tailwindcss(), // <--- ПЛАГИН ДОЛЖЕН БЫТЬ ТУТ
		sveltekit()
	]
});
