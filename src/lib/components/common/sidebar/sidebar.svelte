<script lang="ts">
	import Button from '@/components/ui/button/button.svelte';
	import Home from 'lucide-svelte/icons/house';
	import Settings from 'lucide-svelte/icons/settings';
	import Plus from 'lucide-svelte/icons/plus';
	import { cn } from '@/utils';
	import { page } from '$app/state';
	import { rooms } from '@/consts';
	import { Tooltip, TooltipTrigger } from '@/components/ui/tooltip';
	import { buttonVariants } from '@/components/ui/button';
	import TooltipContent from '@/components/ui/tooltip/tooltip-content.svelte';
	import ArrowDownToLine from 'lucide-svelte/icons/arrow-down-to-line';
	import { pwa, installPWA } from '@/stores/pwa';
	import Spinner from '@/components/ui/spinner/spinner.svelte';
	import Scrollbar from '@/components/ui/scrollbar/scrollbar.svelte';

	let hoveringSiblings = $state(false);
	let roomListContainer = $state<HTMLElement | null>(null);
</script>

<div class="w-sidebar relative top-16 flex h-[calc(100%-4rem)] flex-col px-4 max-md:hidden">
	<ul class="-mx-3">
		<li>
			<a href="/dashboard">
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
			</a>
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
	<ul
		bind:this={roomListContainer}
		class="no-scrollbar relative -mx-3 flex-1 overflow-y-auto overflow-x-hidden pr-4 text-sm"
	>
		<div class="sticky top-0 float-right -mr-4 h-full">
			<div class="absolute right-0 top-0 h-full px-1 py-2">
				<Scrollbar
					container={roomListContainer}
					orientation="vertical"
					class="scale-y-105 rounded-none opacity-60 hover:opacity-100 focus-visible:opacity-100 [&_button]:!rounded-none"
				/>
			</div>
		</div>
		{#each rooms as room}
			<li class="flex overflow-hidden py-0.5">
				<a href={'/room/' + room.id} class="flex w-full">
					<Tooltip openDelay={200}>
						<TooltipTrigger
							class={cn(
								buttonVariants({ size: 'sm', variant: 'ghost' }),
								'flex max-w-full items-center gap-3 overflow-hidden truncate rounded-lg pr-3.5 font-normal dark:font-light',
								page.params?.id === room.id && 'bg-accent/5'
							)}
						>
							<span class={cn('max-w-full', room.name.length > 14 && 'truncate')}>{room.name}</span>
							<span
								class="min-w-[1ch] truncate border p-0.5 font-mono text-xs leading-tight text-foreground/80 dark:text-foreground/60"
								>{room.id}</span
							>
						</TooltipTrigger>
						<TooltipContent>{room.id}</TooltipContent>
					</Tooltip>
				</a>
			</li>
		{/each}
	</ul>
	{#if true}
		<div class="-mx-3 mt-auto py-4">
			<Button
				class="gap-1 bg-transparent hover:bg-accent/10"
				size="sm"
				onclick={async () => {
					pwa.setInstalling(true);
					try {
						await installPWA();
					} finally {
						pwa.setInstalling(false);
					}
				}}
				disabled={$pwa.isInstalling}
			>
				{#if $pwa.isInstalling}
					<Spinner class="h-4 w-4" />
					Installing...
				{:else}
					<ArrowDownToLine class="h-4 w-4" />
					Install Portal
				{/if}
			</Button>
		</div>
	{/if}
</div>
