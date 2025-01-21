<script lang="ts">
	import { DropdownMenu as DropdownMenuPrimitive } from 'bits-ui';
	import Circle from 'lucide-svelte/icons/circle';
	import { cn } from '$lib/utils.js';

	type $$Props = DropdownMenuPrimitive.RadioItemProps;
	type $$Events = DropdownMenuPrimitive.RadioItemEvents;

	let className: $$Props['class'] = undefined;
	export let value: $$Props['value'];
	export { className as class };
</script>

<DropdownMenuPrimitive.RadioItem
	class={cn(
		'group relative flex w-full cursor-pointer select-none items-center justify-between gap-2 rounded-[10px] px-2 py-2.5 text-sm outline-none focus:bg-accent/15 data-[disabled]:pointer-events-none data-[disabled]:cursor-default data-[highlighted]:bg-accent/15 data-[highlighted]:text-accent-foreground data-[disabled]:opacity-50',
		className
	)}
	{value}
	{...$$restProps}
	on:click
	on:keydown
	on:focusin
	on:focusout
	on:pointerdown
	on:pointerleave
	on:pointermove
>
	<slot />
	<span class="flex h-3.5 w-3.5 items-center justify-center">
		<div
			class="flex h-4 w-4 shrink-0 items-center justify-center rounded-full border border-accent/15 group-data-[state=checked]:border-foreground group-data-[state=checked]:ring-4 group-data-[state=checked]:ring-inset group-data-[state=checked]:ring-foreground"
		></div>
	</span>
</DropdownMenuPrimitive.RadioItem>
