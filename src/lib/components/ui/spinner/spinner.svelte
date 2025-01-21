<script lang="ts">
	import { cn } from "@/utils";

	let className = '';
	export { className as class };

</script>

<svg class={cn('h-11 w-11', className)} viewBox="0 0 100 100" {...$$restProps}>
	<polyline fill="none" points="0, 0, 100, 0, 100, 100" stroke-width="15" class="stroke-accent/15" />
	<polyline fill="none" points="0, 0, 0, 100, 100, 100" stroke-width="15" class="stroke-accent/15" />
	<polyline
		fill="none"
		points="0, 0, 100, 0, 100, 100"
		stroke-width="15"
		class="animate-loading origin-center transform"
	/>
	<polyline
		fill="none"
		points="0, 0, 0, 100, 100, 100"
		stroke-width="15"
		class="animate-loading origin-center transform"
	/>
</svg>
