<script>
	import GradientBg from '../gradient-bg/gradient-bg.svelte';
</script>

<footer
	class="relative mt-auto flex h-12 items-center overflow-hidden border-t border-accent/20 text-sm"
>
	<div class="absolute inset-0 -z-[1]">
		<GradientBg class="scale-150 opacity-30 dark:opacity-15" />
		<div class="absolute inset-0 bg-white/5"></div>
	</div>

	<div
		class="container flex items-center justify-between font-mono text-foreground/80 dark:text-foreground/60"
	>
		&copy; 2025

		<p>
			Made with &lt;3 <a
				href="https://github.com/eloraa"
				target="_blank"
				class="font-semibold text-foreground">Elora</a
			>
		</p>
	</div>
</footer>
