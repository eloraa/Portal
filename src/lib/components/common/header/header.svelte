<script lang="ts">
	import Logo from '../logo/logo.svelte';
	import { Button, buttonVariants } from '@/components/ui/button';
	import ArrowDownToLine from 'lucide-svelte/icons/arrow-down-to-line';

	import { cn } from '@/utils';
	import Spinner from '@/components/ui/spinner/spinner.svelte';
	import { Drawer, DrawerHeader, DrawerTrigger } from '@/components/ui/drawer';
	import DrawerContent from '@/components/ui/drawer/drawer-content.svelte';
	import DrawerTitle from '@/components/ui/drawer/drawer-title.svelte';
	import DrawerDescription from '@/components/ui/drawer/drawer-description.svelte';
	import DrawerFooter from '@/components/ui/drawer/drawer-footer.svelte';
	import GradientBg from '../gradient-bg/gradient-bg.svelte';
	import User from '../user/user.svelte';

	let isInstalling = $state(false);
	let deferredPrompt: any;
	let showInstallButton = $state(false);

	$effect(() => {
		if (window.matchMedia('(display-mode: standalone)').matches) {
			showInstallButton = false;
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

<header class="container relative z-50 flex items-center justify-between py-2">
	<a href="/" class="h-16 w-16"><Logo class="-ml-2 h-full w-full" /></a>

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

						<p class="text-sm font-normal text-muted-foreground">
							Install now to experience instant file sharing and messaging powered by secure
							peer-to-peer technology.
						</p>
					</DrawerDescription>
					<DrawerFooter class="p-0 py-4">
						<Button
							class="relative overflow-hidden rounded-xl bg-transparent text-white hover:bg-accent/15"
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
							<!-- <span
								class="absolute inset-0 -z-[1] rounded-xl"
								style="background: linear-gradient(87.81deg, #E6315C 1.77%, #FE013C 24.55%, #FE0140 57.03%, #FE0177 73.99%, #F5466F 98.71%);"
							></span>
							 -->
							<GradientBg class="absolute inset-0 -z-[1] rounded-xl" />
							<span class="absolute inset-0 -z-[1] rounded-xl bg-accent/15"></span>
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
		<User />
	</div>
</header>
