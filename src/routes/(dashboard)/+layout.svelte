<script lang="ts">
	import GradientBg from '@/components/common/gradient-bg/gradient-bg.svelte';
	import Header from './header.svelte';
	import Sidebar from '@/components/common/sidebar/sidebar.svelte';

	let { children } = $props();
</script>

<main class="relative h-full w-full overflow-hidden">
	<div class="inset fixed -z-[1] h-screen min-h-screen w-screen min-w-full">
		<GradientBg class="absolute inset-0 -z-[1] h-full w-full min-w-[100vh]" />
	</div>
	<div class="relative h-full w-full bg-background">
		<Header />
		<div class="absolute flex h-full w-full px-4">
			<Sidebar />
			<div class="flex-1">
				{@render children()}
			</div>
		</div>
	</div>
</main>
