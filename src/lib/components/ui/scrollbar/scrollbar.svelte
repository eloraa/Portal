<script lang="ts">
	import { onMount } from 'svelte';
	import { cn } from '@/utils';
	import { lerp } from '@/index';

	interface Props {
		container: HTMLElement | null;
		orientation?: 'horizontal' | 'vertical';
		class?: string;
	}

	let { container, orientation = 'horizontal', class: className = '' } = $props();

	let scrollbarThumb = $state<HTMLElement | null>(null);
	let isScrolling = $state(false);
	let thumbOffsetX = $state(0);
	let scrollPercentage = $state(0);
	let targetPercentage = $state(0);
	let animationFrame: number;

	function updateThumbPosition() {
		if (Math.abs(targetPercentage - scrollPercentage) < 0.1) {
			scrollPercentage = targetPercentage;
			animationFrame = 0;
			return;
		}

		scrollPercentage = lerp(scrollPercentage, targetPercentage, 0.2);
		animationFrame = requestAnimationFrame(updateThumbPosition);
	}

	function setScrollPercentage(newPercentage: number) {
		targetPercentage = Math.max(0, Math.min(100, newPercentage));

		if (animationFrame) {
			cancelAnimationFrame(animationFrame);
		}

		animationFrame = requestAnimationFrame(updateThumbPosition);
	}

	function updateScrollPercentage() {
		if (container) {
			const maxScroll =
				orientation === 'horizontal'
					? container.scrollWidth - container.clientWidth
					: container.scrollHeight - container.clientHeight;
			const currentScroll =
				orientation === 'horizontal' ? container.scrollLeft : container.scrollTop;

			if (maxScroll > 0) {
				const newPercentage = (currentScroll / maxScroll) * 100;
				setScrollPercentage(Math.min(newPercentage, 100));
			} else {
				setScrollPercentage(0);
			}
		}
	}

	function handleScrollbarClick(e: MouseEvent) {
		if (!container || e.target === scrollbarThumb || !scrollbarThumb) return;

		const scrollbar = scrollbarThumb.parentElement as HTMLElement;
		const bounds = scrollbar.getBoundingClientRect();
		const clickPosition =
			orientation === 'horizontal'
				? (e.clientX - bounds.left) / bounds.width
				: (e.clientY - bounds.top) / bounds.height;

		const maxScroll =
			orientation === 'horizontal'
				? container.scrollWidth - container.clientWidth
				: container.scrollHeight - container.clientHeight;

		if (orientation === 'horizontal') {
			container.scrollLeft = maxScroll * clickPosition;
		} else {
			container.scrollTop = maxScroll * clickPosition;
		}
		setScrollPercentage(clickPosition * 100);
	}

	function handleScrollbarMouseDown(e: MouseEvent | TouchEvent) {
		if (!scrollbarThumb) return;
		isScrolling = true;
		const thumbRect = scrollbarThumb.getBoundingClientRect();
		const client =
			'touches' in e
				? { x: e.touches[0].clientX, y: e.touches[0].clientY }
				: { x: e.clientX, y: e.clientY };

		thumbOffsetX =
			orientation === 'horizontal' ? client.x - thumbRect.left : client.y - thumbRect.top;

		document.body.style.userSelect = 'none';
	}

	function handleScrollbarMouseMove(e: MouseEvent | TouchEvent) {
		if (!isScrolling || !container || !scrollbarThumb) return;

		const scrollbar = scrollbarThumb.parentElement as HTMLElement;
		const bounds = scrollbar.getBoundingClientRect();
		const client =
			'touches' in e
				? { x: e.touches[0].clientX, y: e.touches[0].clientY }
				: { x: e.clientX, y: e.clientY };

		const position =
			orientation === 'horizontal'
				? (client.x - bounds.left - thumbOffsetX) / bounds.width
				: (client.y - bounds.top - thumbOffsetX) / bounds.height;

		const clampedPosition = Math.max(0, Math.min(1, position));

		const maxScroll =
			orientation === 'horizontal'
				? container.scrollWidth - container.clientWidth
				: container.scrollHeight - container.clientHeight;

		if (orientation === 'horizontal') {
			container.scrollLeft = maxScroll * clampedPosition;
		} else {
			container.scrollTop = maxScroll * clampedPosition;
		}
		setScrollPercentage(clampedPosition * 100);
	}

	function handleScrollbarMouseUp() {
		isScrolling = false;
		document.body.style.userSelect = '';
	}

	$effect(() => {
		if (container) {
			updateScrollPercentage(); // Initial calculation

			container.addEventListener('scroll', updateScrollPercentage, { passive: true });

			return () => {
				container.removeEventListener('scroll', updateScrollPercentage);
			};
		}
	});

	onMount(() => {
		updateScrollPercentage(); // Initial calculation

		document.addEventListener('mousemove', handleScrollbarMouseMove);
		document.addEventListener('mouseup', handleScrollbarMouseUp);
		document.addEventListener('mouseleave', handleScrollbarMouseUp);
		document.addEventListener('touchmove', handleScrollbarMouseMove);
		document.addEventListener('touchend', handleScrollbarMouseUp);
		document.addEventListener('touchcancel', handleScrollbarMouseUp);

		return () => {
			document.removeEventListener('mousemove', handleScrollbarMouseMove);
			document.removeEventListener('mouseup', handleScrollbarMouseUp);
			document.removeEventListener('mouseleave', handleScrollbarMouseUp);
			document.removeEventListener('touchmove', handleScrollbarMouseMove);
			document.removeEventListener('touchend', handleScrollbarMouseUp);
			document.removeEventListener('touchcancel', handleScrollbarMouseUp);
			if (animationFrame) {
				cancelAnimationFrame(animationFrame);
			}
		};
	});
</script>

{#if orientation === 'horizontal'}
	<div
		class={cn(
			'relative h-5 w-full cursor-pointer rounded-xl bg-muted hover:bg-muted/80',
			className
		)}
		role="scrollbar"
		tabindex="0"
		aria-controls={container?.id || 'scrollable-content'}
		aria-valuenow={scrollPercentage}
		aria-valuemin="0"
		aria-valuemax="100"
		onmousedown={handleScrollbarClick}
	>
		<button
			class="absolute top-1/2 -translate-y-1/2 cursor-grab select-none items-center justify-center rounded-xl bg-accent/5 px-2 text-xs uppercase backdrop-blur-xl active:cursor-grabbing active:bg-primary/15"
			style="left: {scrollPercentage}%; transform: translateX(-{scrollPercentage}%) translateY(-50%)"
			bind:this={scrollbarThumb}
			onmousedown={handleScrollbarMouseDown}
			ontouchstart={handleScrollbarMouseDown}
		>
			scroll
		</button>
		{#each Array(11) as _, i}
			<div
				class={cn(
					'absolute top-1/2 w-[1px] -translate-y-1/2 bg-foreground/20 transition-all duration-75'
				)}
				style="
                    left: {i * 10}%;
                    height: {(() => {
					const offset = 5;
					const distance = Math.abs(i * 10 - (scrollPercentage + offset));
					if (distance === 0) return '24px';
					if (distance <= 10) return '12px';
					if (distance <= 20) return '8px';
					return '4px';
				})()};
                "
			></div>
		{/each}
	</div>
{:else}
	<div
		class={cn(
			'relative h-full w-3.5 cursor-pointer rounded-xl bg-muted/60 hover:bg-muted/80',
			className
		)}
		role="scrollbar"
		tabindex="0"
		aria-controls={container?.id || 'scrollable-content'}
		aria-valuenow={scrollPercentage}
		aria-valuemin="0"
		aria-valuemax="100"
		onmousedown={handleScrollbarClick}
	>
		<button
			class="absolute left-1/2 flex h-auto w-auto -translate-x-1/2 cursor-grab select-none items-center justify-center rounded-xl bg-accent/5 px-2 py-8 text-xs uppercase backdrop-blur-xl active:cursor-grabbing active:bg-primary/15"
			style="top: {scrollPercentage}%; transform: translateY(-{scrollPercentage}%) translateX(-50%)"
			bind:this={scrollbarThumb}
			onmousedown={handleScrollbarMouseDown}
			ontouchstart={handleScrollbarMouseDown}
		>
			<span class="absolute block rotate-90">scroll</span>
		</button>
		{#each Array(11) as _, i}
			<div
				class={cn(
					'absolute left-1/2 h-[1px] -translate-x-1/2 bg-foreground/20 transition-all duration-75'
				)}
				style="
                    top: {i * 10}%;
                    width: {(() => {
					const offset = 5;
					const distance = Math.abs(i * 10 - (scrollPercentage + offset));
					if (distance === 0) return '24px';
					if (distance <= 10) return '12px';
					if (distance <= 20) return '8px';
					return '4px';
				})()};
                "
			></div>
		{/each}
	</div>
{/if}
