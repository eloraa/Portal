import { writable } from 'svelte/store';

export type Theme = 'dark' | 'light' | 'system';

function createThemeStore() {
	const { subscribe, set } = writable<Theme>('dark');

	return {
		subscribe,
		set: (theme: Theme) => {
			document.documentElement.classList.remove('light', 'dark');
			if (theme === 'system') {
				const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches
					? 'dark'
					: 'light';
				document.documentElement.classList.add(systemTheme);
				document.documentElement.style.colorScheme = systemTheme;
			} else {
				document.documentElement.classList.add(theme);
				document.documentElement.style.colorScheme = theme;
			}
			document.cookie = `theme=${theme};path=/;max-age=31536000`;

			set(theme);
		}
	};
}

export const theme = createThemeStore();
