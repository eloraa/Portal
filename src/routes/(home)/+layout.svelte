<script lang="ts">
	import { browser } from '$app/environment';
	import Footer from '@/components/common/footer/footer.svelte';
	import Header from '@/components/common/header/header.svelte';
	import { backendAvailable } from '@/stores/backend';

	let { children } = $props();

	if (browser) {
		if ('serviceWorker' in navigator) {
			navigator.serviceWorker.register('/service-worker.js');
		}
	}
</script>

<div class="flex min-h-full flex-col">
	{#if !$backendAvailable}
		<div class="container bg-primary-light py-1 font-mono text-sm font-medium uppercase">
			Uh oh! Our server is unavailable. Website won't work as expected.
		</div>
	{/if}
	<div class="sticky top-0 z-50 overflow-hidden bg-background">
		<Header />
	</div>
	{@render children()}
	<Footer />
</div>
