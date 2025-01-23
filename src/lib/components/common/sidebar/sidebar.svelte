<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	import Home from 'lucide-svelte/icons/house';
	import Settings from 'lucide-svelte/icons/settings';
	import Plus from 'lucide-svelte/icons/plus';
	import { cn } from '@/utils';
	import { page } from '$app/state';

	let hoveringSiblings = $state(false);
</script>

<div class="w-sidebar relative top-16 h-[calc(100vh-4rem)] px-4 max-md:hidden">
	<ul class="-mx-3">
		<li>
			<Button
				variant="ghost"
				size="sm"
				class={cn(
					'flex items-center gap-1 pr-3.5',
					page.url?.pathname.includes('/dashboard') && !hoveringSiblings && 'bg-accent/10'
				)}
				onmouseover={() => (hoveringSiblings = true)}
				onmouseleave={() => (hoveringSiblings = false)}
			>
				<Home class="w-4" />
				Home
			</Button>
		</li>
		<li>
			<Button
				variant="ghost"
				size="sm"
				class="flex items-center gap-1 pr-3.5"
				onmouseover={() => (hoveringSiblings = true)}
				onmouseleave={() => (hoveringSiblings = false)}
			>
				<Settings class="w-4" />
				Setting
			</Button>
		</li>
	</ul>

	<div class="mt-6 flex items-center gap-2">
		<h2 class="font-mono text-xs uppercase text-foreground/80">Rooms</h2>
		<Button size="icon" variant="ghost" class="h-8 w-8">
			<Plus class="w-3.5" />
			<span class="sr-only">Add a new Room</span>
		</Button>
	</div>
	<ul class="-mx-2"></ul>
</div>
