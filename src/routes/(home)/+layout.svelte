<script lang="ts">
	import { browser } from '$app/environment';
	import Footer from '@/components/common/footer/footer.svelte';
	import GradientBg from '@/components/common/gradient-bg/gradient-bg.svelte';
	import Header from '@/components/common/header/header.svelte';

	let { children } = $props();

	if (browser) {
		if ('serviceWorker' in navigator) {
			navigator.serviceWorker.register('/service-worker.js');
		}
	}
</script>

<div class="flex min-h-full flex-col">
	<div class="sticky top-0 z-50 overflow-hidden bg-background">
		<Header />
	</div>
	{@render children()}
	<Footer />
</div>
