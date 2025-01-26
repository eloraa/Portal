<script lang="ts">
	import { page } from '$app/state';
	import { AvatarFallback, AvatarImage } from '@/components/ui/avatar';
	import Avatar from '@/components/ui/avatar/avatar.svelte';
	import Button from '@/components/ui/button/button.svelte';
	import { Dialog } from '@/components/ui/dialog';
	import DialogContent from '@/components/ui/dialog/dialog-content.svelte';
	import DialogDescription from '@/components/ui/dialog/dialog-description.svelte';
	import DialogHeader from '@/components/ui/dialog/dialog-header.svelte';
	import DialogTitle from '@/components/ui/dialog/dialog-title.svelte';
	import { Tooltip, TooltipContent, TooltipTrigger } from '@/components/ui/tooltip';
	import { rooms } from '@/consts';
	import { cn } from '@/utils';
	import Hash from 'lucide-svelte/icons/hash';
	import CaseSensitive from 'lucide-svelte/icons/case-sensitive';
	import User from 'lucide-svelte/icons/user';
	import Smile from 'lucide-svelte/icons/smile';
	import Lock from 'lucide-svelte/icons/lock';
	import Globe from 'lucide-svelte/icons/globe';
	import GlobeLock from 'lucide-svelte/icons/globe-lock';
	import Asterisk from 'lucide-svelte/icons/asterisk';

	import { user } from '@/stores/user';
	import { DropdownMenu, DropdownMenuTrigger } from '@/components/ui/dropdown-menu';
	import ChevronDown from 'lucide-svelte/icons/chevron-down';
	import DropdownMenuContent from '@/components/ui/dropdown-menu/dropdown-menu-content.svelte';
	import DropdownMenuRadioGroup from '@/components/ui/dropdown-menu/dropdown-menu-radio-group.svelte';
	import DropdownMenuRadioItem from '@/components/ui/dropdown-menu/dropdown-menu-radio-item.svelte';

	let isEditing = $state(false);
	let editableDiv: HTMLDivElement;
	let isDialogOpen = $state(false);
	let visibility = $state('public');

	let currentRoom = $state(rooms.find((room) => room.id === page.params.id));

	$effect(() => {
		currentRoom = rooms.find((room) => room.id === page.params.id);
	});

	function handleEdit(event: MouseEvent) {
		if (!isEditing) {
			isEditing = true;
			setTimeout(() => {
				editableDiv?.focus();
				// Select all text only when entering edit mode
				const range = document.createRange();
				const selection = window.getSelection();
				range.selectNodeContents(editableDiv);
				selection?.removeAllRanges();
				selection?.addRange(range);
			}, 0);
		}
		// When already editing, let the default selection behavior work
		event.stopPropagation();
	}

	function handleBlur() {
		isEditing = false;
		// Here you would typically save the new name
		// For now just logging it
		console.log('New name:', editableDiv.textContent);
	}

	function handleKeyDown(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			event.preventDefault();
			editableDiv.blur();
		}
		if (event.key === 'Escape') {
			editableDiv.textContent = currentRoom?.name || '';
			editableDiv.blur();
		}
	}
</script>

<header
	class="sticky top-0 flex w-[calc(100%-6.5rem)] items-center justify-between max-md:left-10 md:w-[calc(100%-4rem)]"
>
	<div class="flex items-center gap-4 text-sm">
		<!-- <Tooltip openDelay={0}>
			<TooltipTrigger
				class="flex h-11 w-11 items-center justify-center rounded-full transition-colors hover:bg-popover/50"
			>
				<Avatar
					class="flex h-8 w-8 items-center justify-center overflow-visible bg-muted/85 backdrop-blur-xl"
				>
					<AvatarImage src="/shotgun.png" class="absolute h-14 w-14 min-w-14 object-contain" />
					<AvatarFallback>
						{rooms.find((room) => room.id === page.params.id)?.name[0]}
					</AvatarFallback>
				</Avatar>
			</TooltipTrigger>
			<TooltipContent>{rooms.find((room) => room.id === page.params.id)?.name}</TooltipContent>
		</Tooltip> -->
		<div class="relative max-w-24 pr-6 md:max-w-64">
			<Button
				class="absolute right-0 top-0 h-5 w-5 overflow-hidden"
				size="icon"
				onclick={() => (isDialogOpen = true)}
			>
				<span class="font-mono text-xs">i</span>
				<span class="sr-only">info</span>
			</Button>
			<div
				bind:this={editableDiv}
				contenteditable={isEditing}
				onclick={handleEdit}
				onblur={handleBlur}
				onkeydown={handleKeyDown}
				role="textbox"
				tabindex="0"
				aria-label="Room name"
				class={cn(
					'truncate text-wrap border-b border-transparent outline-none',
					isEditing && 'border-accent/20'
				)}
			>
				{currentRoom?.name}
			</div>
			<h2 class="font-mono text-xs text-foreground/80 dark:text-foreground/60">
				{currentRoom?.id}
			</h2>
		</div>
	</div>
</header>

<Dialog bind:open={isDialogOpen}>
	<DialogContent>
		<DialogHeader>
			<DialogTitle>Room Details</DialogTitle>
			<DialogDescription class="sr-only">Edit room details</DialogDescription>
		</DialogHeader>

		<div>
			<ul class="text-sm">
				<li class="flex h-10 items-center border-b border-accent/15">
					<div class="flex w-3/5 items-center gap-1">
						<Hash class="w-4" />
						<h2>Room Id</h2>
					</div>
					<div class="w-2/5">
						<span class="font-mono text-foreground/80 dark:text-foreground/60"
							>{currentRoom?.id}</span
						>
					</div>
				</li>
				<li class="flex h-10 items-center border-b border-accent/15">
					<div class="flex w-3/5 items-center gap-1">
						<CaseSensitive class="w-4" />
						<h2>Room Name</h2>
					</div>
					<div class="w-2/5">{currentRoom?.name}</div>
				</li>
				<li class="flex h-10 items-center border-b border-accent/15">
					<div class="flex w-3/5 items-center gap-1">
						<User class="w-4" />
						<h2>Total Members</h2>
					</div>
					<div class="w-2/5">
						<span class="font-mono text-foreground/80 dark:text-foreground/60">10</span>
					</div>
				</li>
				<li class="flex h-10 items-center border-b border-accent/15">
					<div class="flex w-3/5 items-center gap-1">
						<Smile class="w-4" />
						<h2>Owner</h2>
					</div>
					<div class="w-2/5">
						<span class="text-primary-light dark:text-primary">{$user.username}</span>
					</div>
				</li>
				<li
					class={cn(
						'flex h-10 items-center',
						visibility === 'private' && 'border-b border-accent/15'
					)}
				>
					<div class="flex w-3/5 items-center gap-1">
						<Lock class="w-4" />
						<h2>Visibility</h2>
					</div>
					<div class="w-2/5 min-w-0">
						<DropdownMenu>
							<DropdownMenuTrigger class="flex w-full items-center gap-1 text-start">
								{visibility[0].toUpperCase() + visibility.slice(1)}
								<ChevronDown class="w-4" />
							</DropdownMenuTrigger>
							<DropdownMenuContent class="w-full max-w-52">
								<DropdownMenuRadioGroup bind:value={visibility}>
									<DropdownMenuRadioItem value="public">
										<div class="flex items-center gap-2">
											<Globe class="w-5" />
											Public
										</div>
									</DropdownMenuRadioItem>
									<DropdownMenuRadioItem value="private" class="flex items-center gap-2">
										<div class="flex items-center gap-2">
											<GlobeLock class="mr-0.5 w-[18px]" />
											Private
										</div>
									</DropdownMenuRadioItem>
								</DropdownMenuRadioGroup>
							</DropdownMenuContent>
						</DropdownMenu>
					</div>
				</li>

				{#if visibility === 'private'}
					<li class="flex h-10 items-center">
						<div class="flex w-3/5 items-center gap-1">
							<Asterisk class="w-4" />
							<h2>Password</h2>
						</div>
						<div class="w-2/5">*******</div>
					</li>
				{/if}
			</ul>
			<Button class="mt-2 w-full" onclick={() => (isDialogOpen = false)}>Done</Button>
		</div>
	</DialogContent>
</Dialog>
