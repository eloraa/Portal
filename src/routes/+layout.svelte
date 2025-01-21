<script lang="ts">
	import Header from '@/components/common/header/header.svelte';
	import '../app.css';
	import { theme } from '@/stores/theme';
	import type { Theme } from '@/stores/theme';
	import { browser } from '$app/environment';
	import { onMount } from 'svelte';

	interface LayoutData {
		theme: Theme;
		username: string;
		avatarId: string;
	}

	let { data, children } = $props<{ data: LayoutData }>();
	let deferredPrompt: any;
	let showInstallButton = $state(false);

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

	onMount(() => {
		if (browser) {
			if (window.matchMedia('(display-mode: standalone)').matches) {
				showInstallButton = false;
			}

			if ('serviceWorker' in navigator) {
				navigator.serviceWorker.register('/service-worker.js');
			}

			window.addEventListener('beforeinstallprompt', (e) => {
				e.preventDefault();
				deferredPrompt = e;
				showInstallButton = true;
			});

			window.addEventListener('appinstalled', () => {
				showInstallButton = false;
				deferredPrompt = null;
			});
		}
	});

	async function installPWA() {
		if (deferredPrompt) {
			deferredPrompt.prompt();
			const { outcome } = await deferredPrompt.userChoice;
			deferredPrompt = null;
			if (outcome === 'accepted') {
				showInstallButton = false;
			}
		}
	}
</script>

<svelte:head>
	<title>Portal</title>
	<meta name="description" content="Portal is a file sharing app" />
	<link rel="manifest" href="/site.webmanifest" />
	<meta name="theme-color" content="#000000" />
	<link rel="apple-touch-icon" href="/apple-touch-icon.png" />
</svelte:head>

<Header
	initialTheme={data.theme}
	initialUsername={data.username}
	initialAvatarId={data.avatarId}
	{installPWA}
	{showInstallButton}
/>
{@render children()}
