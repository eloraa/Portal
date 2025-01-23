<script lang="ts">
	import { Dialog, DialogContent, DialogHeader, DialogDescription } from '@/components/ui/dialog';
	import { Button } from '@/components/ui/button';
	import { Input } from '@/components/ui/input';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuTrigger,
		DropdownMenuRadioGroup,
		DropdownMenuRadioItem
	} from '@/components/ui/dropdown-menu';
	import Avatar from '@/components/ui/avatar/avatar.svelte';
	import AvatarFallback from '@/components/ui/avatar/avatar-fallback.svelte';
	import AvatarImage from '@/components/ui/avatar/avatar-image.svelte';
	import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover';
	import Scrollbar from '@/components/ui/scrollbar/scrollbar.svelte';
	import UserAvatar from './user-avatar.svelte';

	import User from 'lucide-svelte/icons/user';
	import PaintBrushVertical from 'lucide-svelte/icons/paintbrush-vertical';
	import ChevronDown from 'lucide-svelte/icons/chevron-down';
	import SunMoon from 'lucide-svelte/icons/sun-moon';
	import Sun from 'lucide-svelte/icons/sun';
	import Eclipse from 'lucide-svelte/icons/eclipse';
	import Bell from 'lucide-svelte/icons/bell';

	import { cn } from '@/utils';
	import { user } from '@/stores/user';
	import { avatars } from '@/consts';
	import { page } from '$app/state';
	import { onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { theme } from '@/stores/theme';

	let isDialogOpen = $state(false);
	let username = $state($user.username);
	let selectedTheme = $state(page.data.theme);
	let usernameError = $state('');
	let showUsernameError = $state(false);

	let scrollContainer = $state<HTMLElement | null>(null);
	let isDragging = $state(false);
	let startX = $state(0);
	let scrollLeft = $state(0);

	function getAvatarSrcFromId(id: string) {
		const avatar = avatars.find((opt) => opt.id === id);
		return avatar?.src || '/kazuha.png';
	}

	let selectedAvatar = $state(getAvatarSrcFromId($user.avatarId));

	function validateUsername(value: string) {
		if (value.length < 1) {
			return 'Username must be at least 1 character';
		}
		if (value.length > 25) {
			return 'Username must be less than 25 characters';
		}
		return '';
	}

	function handleAvatarSelect(src: string) {
		selectedAvatar = src;
		const avatar = avatars.find((opt) => opt.src === src);
		if (avatar) {
			user.setAvatarId(avatar.id);
		}
	}

	function handleUsernameChange(e: Event) {
		const input = e.target as HTMLInputElement;
		const error = validateUsername(input.value);

		if (input.value.length > 25) {
			input.value = input.value.slice(0, 25);
			username = input.value;
			usernameError = error;
			showUsernameError = true;
		} else {
			if (showUsernameError) {
				showUsernameError = false;
				usernameError = '';
			}
		}
	}

	function handleUsernameBlur(e: FocusEvent) {
		const error = validateUsername(username);
		if (error) {
			usernameError = error;
			showUsernameError = true;
			setTimeout(() => {
				const input = e.target as HTMLInputElement;
				input.focus();
			}, 0);
		} else {
			usernameError = '';
			showUsernameError = false;
			user.setUsername(username);
		}
	}

	function handleMouseDown(e: MouseEvent | TouchEvent) {
		if (!scrollContainer) return;

		isDragging = true;
		const clientX = e instanceof MouseEvent ? e.pageX : e.touches[0].clientX;
		startX = clientX - scrollContainer.offsetLeft;
		scrollLeft = scrollContainer.scrollLeft;
		scrollContainer.style.cursor = 'grabbing';
	}

	function handleMouseMove(e: MouseEvent | TouchEvent) {
		if (!isDragging || !scrollContainer) return;

		e.preventDefault();
		const clientX = e instanceof MouseEvent ? e.pageX : e.touches[0].clientX;
		const x = clientX - scrollContainer.offsetLeft;
		const walk = (x - startX) * 2;
		scrollContainer.scrollLeft = scrollLeft - walk;
	}

	function handleMouseUp() {
		if (!scrollContainer) return;

		isDragging = false;
		scrollContainer.style.cursor = 'grab';
	}

	$effect(() => {
		username = $user.username;
		selectedAvatar = getAvatarSrcFromId($user.avatarId);

		if (browser) {
			const root = document.documentElement;
			root.classList.remove('light', 'dark');

			if (selectedTheme === 'system') {
				const systemTheme = window.matchMedia('(prefers-color-scheme: dark)').matches
					? 'dark'
					: 'light';
				root.classList.add(systemTheme);
				root.style.colorScheme = systemTheme;
			} else {
				root.classList.add(selectedTheme);
				root.style.colorScheme = selectedTheme;
			}
		}
		theme.set(selectedTheme);
	});

	onMount(() => {
		document.addEventListener('mousemove', handleMouseMove);
		document.addEventListener('mouseup', handleMouseUp);
		document.addEventListener('mouseleave', handleMouseUp);

		return () => {
			document.removeEventListener('mousemove', handleMouseMove);
			document.removeEventListener('mouseup', handleMouseUp);
			document.removeEventListener('mouseleave', handleMouseUp);
		};
	});
</script>

<UserAvatar onClick={() => (isDialogOpen = true)} />

<Dialog bind:open={isDialogOpen}>
	<DialogContent class="md:max-w-sm">
		<DialogHeader>Profile</DialogHeader>
		<DialogDescription class="sr-only">Edit your profile</DialogDescription>

		<div class="flex flex-col gap-6 overflow-hidden p-1">
			<div class="relative">
				<div
					bind:this={scrollContainer}
					role="listbox"
					tabindex="0"
					aria-label="Avatar options"
					class={cn(
						'scrollable-avatars flex cursor-grab gap-8 px-4 py-4',
						'select-none',
						'overflow-x-auto',
						'no-scrollbar',
						isDragging && 'cursor-grabbing'
					)}
					style="touch-action: pan-y pinch-zoom;"
					onmousedown={handleMouseDown}
					ontouchstart={handleMouseDown}
					ontouchmove={handleMouseMove}
					ontouchend={handleMouseUp}
					onwheel={(e) => {
						if (!scrollContainer) return;
						e.preventDefault();
						scrollContainer.scrollLeft += e.deltaY;
					}}
					onkeydown={(e) => {
						if (!scrollContainer) return;
						if (e.key === 'ArrowRight') {
							scrollContainer.scrollBy({ left: 100, behavior: 'smooth' });
						} else if (e.key === 'ArrowLeft') {
							scrollContainer.scrollBy({ left: -100, behavior: 'smooth' });
						}
					}}
				>
					{#each avatars as option}
						<div
							class="first:pl-2 last:pr-2"
							role="option"
							aria-selected={selectedAvatar === option.src}
						>
							<button
								class={cn(
									'rounded-full border p-1 ring-0 ring-inset ring-foreground transition-all',
									'hover:bg-muted hover:ring-[40px]',
									selectedAvatar === option.src && 'ring-[40px]'
								)}
								onclick={() => handleAvatarSelect(option.src)}
								aria-label={`Select ${option.id} avatar`}
							>
								<Avatar
									class={cn('relative flex h-10 w-10 items-center justify-center overflow-visible')}
									style={'background-color:' + option.color}
								>
									<AvatarImage
										src={option.src}
										class="absolute h-20 w-auto min-w-20 object-cover"
									/>
									<AvatarFallback>{username[0]}</AvatarFallback>
								</Avatar>
							</button>
						</div>
					{/each}
				</div>

				<div class="flex w-full justify-between">
					<Scrollbar container={scrollContainer} />
				</div>
			</div>

			<ul class="text-sm">
				<li class="flex h-10 items-center border-b border-accent/15">
					<div class="flex w-3/5 items-center gap-1">
						<User class="w-4" />
						<h2>Username</h2>
					</div>

					<div class="w-2/5 min-w-0">
						<Input
							type="text"
							placeholder="Name"
							bind:value={username}
							maxlength={25}
							class={cn('w-full min-w-0 text-sm', usernameError && 'border-destructive')}
							oninput={handleUsernameChange}
							onblur={handleUsernameBlur}
						/>
						<Popover open={showUsernameError} closeOnOutsideClick={false} disableFocusTrap>
							<PopoverTrigger class="sr-only">Username error</PopoverTrigger>
							<PopoverContent
								class="w-fit bg-destructive px-3 py-1.5 text-xs text-destructive-foreground"
								side="bottom"
								align="start"
							>
								{usernameError}
							</PopoverContent>
						</Popover>
					</div>
				</li>
				<li class="flex h-10 items-center">
					<div class="flex w-3/5 items-center gap-1">
						<PaintBrushVertical class="w-4" />
						<h2>Theme</h2>
					</div>

					<div class="w-2/5 min-w-0">
						<DropdownMenu>
							<DropdownMenuTrigger class="flex w-full items-center gap-1 text-start">
								{selectedTheme[0].toUpperCase() + selectedTheme.slice(1)}
								<ChevronDown class="mt-1 w-4" />
							</DropdownMenuTrigger>
							<DropdownMenuContent class="w-full max-w-52">
								<DropdownMenuRadioGroup bind:value={selectedTheme}>
									<DropdownMenuRadioItem value="system">
										<div class="flex items-center gap-2">
											<SunMoon class="w-5" />
											System
										</div>
									</DropdownMenuRadioItem>
									<DropdownMenuRadioItem value="light">
										<div class="flex items-center gap-2">
											<Sun class="w-5" />
											Light
										</div>
									</DropdownMenuRadioItem>
									<DropdownMenuRadioItem value="dark" class="flex items-center gap-2">
										<div class="flex items-center gap-2">
											<Eclipse class="mr-0.5 w-[18px]" />
											Dark
										</div>
									</DropdownMenuRadioItem>
								</DropdownMenuRadioGroup>
							</DropdownMenuContent>
						</DropdownMenu>
					</div>
				</li>
				<!-- <li class="flex h-10 items-center">
					<div class="flex w-3/5 items-center gap-1">
						<Bell class="w-4" />
						<h2>Notification</h2>
					</div>
					<div class="w-2/5"></div>
				</li> -->
			</ul>

			<Button class="w-full" onclick={() => (isDialogOpen = false)}>Done</Button>
		</div>
	</DialogContent>
</Dialog>
