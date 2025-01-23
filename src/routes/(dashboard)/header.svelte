<script lang="ts">
	import LogoAccent from '@/components/common/logo/logo-accent.svelte';
	import User from '@/components/common/user/user.svelte';
</script>

<header
	class="pointer-events-none sticky inset-x-0 top-0 z-40 flex h-16 w-full items-center justify-between px-4 md:fixed"
>
	<a
		href="/dashboard"
		class="pointer-events-auto transition hover:animate-[spin_2s_cubic-bezier(0.18,0.89,0.32,1.28)_infinite] hover:text-primary-light"
	>
		<figure class="h-10 w-10">
			<LogoAccent class="h-full w-full" />
		</figure>
		<span class="sr-only">Portal</span>
	</a>
	<div class="pointer-events-auto flex items-center gap-8">
		<User />
	</div>
</header>
