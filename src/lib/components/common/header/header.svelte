<script lang="ts">
	import { onMount } from 'svelte';
	import Avatar from '@/components/ui/avatar/avatar.svelte';
	import Logo from '../logo/logo.svelte';
	import AvatarFallback from '@/components/ui/avatar/avatar-fallback.svelte';
	import AvatarImage from '@/components/ui/avatar/avatar-image.svelte';
	import { Tooltip, TooltipTrigger, TooltipContent } from '@/components/ui/tooltip';
	import {
		Dialog,
		DialogContent,
		DialogHeader,
		DialogDescription,
		DialogTrigger
	} from '@/components/ui/dialog';
	import { Button, buttonVariants } from '@/components/ui/button';
	import { Input } from '@/components/ui/input';
	import {
		DropdownMenu,
		DropdownMenuContent,
		DropdownMenuTrigger,
		DropdownMenuRadioGroup,
		DropdownMenuRadioItem
	} from '@/components/ui/dropdown-menu';

	import User from 'lucide-svelte/icons/user';
	import PaintBrushVertical from 'lucide-svelte/icons/paintbrush-vertical';
	import ChevronDown from 'lucide-svelte/icons/chevron-down';
	import SunMoon from 'lucide-svelte/icons/sun-moon';
	import Sun from 'lucide-svelte/icons/sun';
	import Eclipse from 'lucide-svelte/icons/eclipse';
	import ArrowDownToLine from 'lucide-svelte/icons/arrow-down-to-line';

	import { cn } from '@/utils';
	import { theme, type Theme } from '@/stores/theme';
	import { Popover } from '@/components/ui/popover';
	import { PopoverContent } from '@/components/ui/popover';
	import { PopoverTrigger } from '@/components/ui/popover';
	import Scrollbar from '@/components/ui/scrollbar/scrollbar.svelte';
	import Spinner from '@/components/ui/spinner/spinner.svelte';
	import { Drawer, DrawerHeader, DrawerTrigger } from '@/components/ui/drawer';
	import DrawerContent from '@/components/ui/drawer/drawer-content.svelte';
	import DrawerTitle from '@/components/ui/drawer/drawer-title.svelte';
	import DrawerDescription from '@/components/ui/drawer/drawer-description.svelte';
	import DrawerFooter from '@/components/ui/drawer/drawer-footer.svelte';

	interface Props {
		initialTheme: Theme;
		initialUsername: string;
		initialAvatarId: string;
		installPWA: () => Promise<void>;
		showInstallButton: boolean;
	}

	let { initialTheme, initialUsername, initialAvatarId, installPWA, showInstallButton } = $props();

	let isDialogOpen = $state(false);
	let username = $state(initialUsername);
	let selectedTheme = $state(initialTheme);
	let usernameError = $state('');
	let showUsernameError = $state(false);

	let scrollContainer = $state<HTMLElement | null>(null);
	let isDragging = $state(false);
	let startX = $state(0);
	let scrollLeft = $state(0);

	let isInstalling = $state(false);

	const avatarOptions = [
		{ id: 'kazuha', src: '/kazuha.png', color: '#ff5144' },
		{ id: 'diluc', src: '/diluc.png', color: '#6d2626' },
		{ id: 'ganyu', src: '/ganyu.png', color: '#3F51B5' },
		{ id: 'hutao', src: '/hutao.png', color: '#a51308' },
		{ id: 'shotgun', src: '/shotgun.png', color: '#673AB7' },
		{ id: 'shenhe', src: '/shenhe.png', color: '#175d7d' }
	];

	function getAvatarSrcFromId(id: string) {
		const avatar = avatarOptions.find((opt) => opt.id === id);
		return avatar?.src || '/kazuha.png';
	}

	let selectedAvatar = $state(getAvatarSrcFromId(initialAvatarId));

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
		const avatar = avatarOptions.find((opt) => opt.src === src);
		if (avatar) {
			document.cookie = `avatarId=${avatar.id};path=/;max-age=31536000`;
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
			document.cookie = `username=${username};path=/;max-age=31536000`;
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
		theme.set(selectedTheme);
	});

	onMount(() => {
		if (selectedTheme === 'system') {
			const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
			const handleChange = (e: MediaQueryListEvent) => {
				document.documentElement.classList.remove('light', 'dark');
				document.documentElement.classList.add(e.matches ? 'dark' : 'light');
			};
			mediaQuery.addEventListener('change', handleChange);
			return () => mediaQuery.removeEventListener('change', handleChange);
		}

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

<header class="container flex items-center justify-between py-2">
	<a href="/" class="h-16 w-16"><Logo class="h-full w-full" /></a>

	<div class="flex items-center gap-4">
		{#if true}
			<Drawer>
				<DrawerTrigger class={cn(buttonVariants({ size: 'icon' }))}>
					<ArrowDownToLine class="w-5" />
					<span class="sr-only">Install App</span>
				</DrawerTrigger>
				<DrawerContent
					class="mx-auto rounded-t-2xl bg-popover p-4 text-center shadow-2xl md:max-w-sm"
				>
					<div
						class="absolute inset-0 -z-[2] scale-75 rounded-t-2xl bg-primary-light blur-3xl"
					></div>
					<div class="absolute inset-0 -z-[1] rounded-t-2xl border bg-popover"></div>
					<DrawerHeader class="flex items-center justify-center">
						<figure class="h-16 w-16">
							<img src="/logo.png" alt="Portal" class="h-full w-full" />
						</figure>
					</DrawerHeader>
					<DrawerTitle class="sr-only">Portal</DrawerTitle>
					<DrawerDescription class="text-base font-semibold text-foreground">
						Ready to connect, share, and communicate like never before?

						<p class="text-sm text-muted-foreground">
							Install now to experience instant file sharing and messaging powered by secure
							peer-to-peer technology.
						</p>
					</DrawerDescription>
					<DrawerFooter>
						<Button
							class="relative rounded-xl bg-transparent text-white hover:bg-accent/15"
							onclick={async () => {
								isInstalling = true;
								try {
									await installPWA();
								} finally {
									isInstalling = false;
								}
							}}
							disabled={isInstalling}
						>
							<span
								class="absolute inset-0 -z-[1] rounded-xl"
								style="background: linear-gradient(87.81deg, #E6315C 1.77%, #FE013C 24.55%, #FE0140 57.03%, #FE0177 73.99%, #F5466F 98.71%);"
							></span>
							{#if !isInstalling}
								Install Now
							{/if}
							{#if isInstalling}
								<Spinner class="absolute h-4 w-4" />
							{/if}
						</Button>
					</DrawerFooter>
				</DrawerContent>
			</Drawer>
		{/if}
		<Tooltip openDelay={0}>
			<TooltipTrigger
				class="flex h-11 w-11 items-center justify-center rounded-full transition-colors hover:bg-popover/50"
				onclick={() => (isDialogOpen = true)}
			>
				<Avatar
					class="flex h-8 w-8 items-center justify-center overflow-visible bg-muted/85 backdrop-blur-xl"
				>
					<AvatarImage src={selectedAvatar} class="absolute h-14 w-14 min-w-14 object-contain" />
					<AvatarFallback>{username[0]}</AvatarFallback>
				</Avatar>
			</TooltipTrigger>
			<TooltipContent>{username}</TooltipContent>
		</Tooltip>
	</div>
</header>

<Dialog open={isDialogOpen} onOpenChange={(val) => (isDialogOpen = val)}>
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
					{#each avatarOptions as option}
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
			</ul>

			<Button class="w-full" onclick={() => (isDialogOpen = false)}>Done</Button>
		</div>
	</DialogContent>
</Dialog>
