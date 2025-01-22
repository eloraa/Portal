<script lang="ts">
	import '../app.css';
	import { theme } from '@/stores/theme';
	import type { Theme } from '@/stores/theme';
	import { browser } from '$app/environment';

	interface LayoutData {
		theme: Theme;
		username: string;
		avatarId: string;
	}

	let { data, children } = $props<{ data: LayoutData }>();

	$effect(() => {
		if (browser) {
			const root = document.documentElement;
			root.classList.remove('light', 'dark');

			if (data.theme === 'system') {
				const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches
					? 'dark'
					: 'light';
				root.classList.add(systemTheme);
				root.style.colorScheme = systemTheme;
			} else {
				root.classList.add(data.theme);
				root.style.colorScheme = data.theme;
			}
		}
		theme.set(data.theme);
	});
</script>

<svelte:head>
	<title>Portal</title>
	<meta name="description" content="Portal is a file sharing app" />
	<link rel="manifest" href="/site.webmanifest" />
	<meta name="theme-color" content="#000000" />
	<link rel="apple-touch-icon" href="/apple-touch-icon.png" />
</svelte:head>

{@render children()}
