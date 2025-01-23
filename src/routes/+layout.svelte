<script lang="ts">
	import '../app.css';
	import { theme } from '@/stores/theme';
	import { backendAvailable } from '@/stores/backend';
	import type { Theme } from '@/stores/theme';
	import { browser } from '$app/environment';
	import { user } from '@/stores/user';
	import { onMount } from 'svelte';

	interface LayoutData {
		theme: Theme;
		username: string;
		avatarId: string;
		userId: string;
		backendAvailable: boolean;
	}

	let { data, children } = $props<{ data: LayoutData }>();

	$effect(() => {
		if (browser) {
			// Initialize user store first
			user.set({
				username: data.username,
				avatarId: data.avatarId,
				userId: data.userId
			});

			// Then handle theme and backend
			backendAvailable.set(data.backendAvailable);

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

			theme.set(data.theme);
		}
	});

	onMount(() => {
		if ($theme === 'system') {
			const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
			const handleChange = (e: MediaQueryListEvent) => {
				document.documentElement.classList.remove('light', 'dark');
				document.documentElement.classList.add(e.matches ? 'dark' : 'light');
			};
			mediaQuery.addEventListener('change', handleChange);
			return () => mediaQuery.removeEventListener('change', handleChange);
		}
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
