<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { browser } from '$app/environment';
	import { cn } from '@/utils';

	export let speed = 10;
	export let moveSpeed = 0.2;
	export let className = '';

	let rotation = 0;
	let xPosition = 50;
	let yPosition = 50;
	let animationFrameId: number;
	let gradientDiv: HTMLDivElement;
	let resizeObserver: ResizeObserver;

	let xVelocity = moveSpeed;
	let yVelocity = moveSpeed * 0.75;

	function updateHeight() {
		if (!gradientDiv) return;
		const width = gradientDiv.offsetWidth;
		gradientDiv.style.height = `${width}px`;
	}

	function animate() {
		rotation = (rotation + speed * 0.016) % 360;

		xPosition += xVelocity;
		if (xPosition <= 0 || xPosition >= 100) {
			xVelocity = -xVelocity;
			xPosition += xVelocity;
		}

		yPosition += yVelocity;
		if (yPosition <= 0 || yPosition >= 100) {
			yVelocity = -yVelocity;
			yPosition += yVelocity;
		}

		animationFrameId = window.requestAnimationFrame(animate);
	}

	onMount(() => {
		if (browser) {
			animationFrameId = window.requestAnimationFrame(animate);

			resizeObserver = new ResizeObserver(updateHeight);
			resizeObserver.observe(gradientDiv.parentElement!);

			updateHeight();
		}
	});

	onDestroy(() => {
		if (browser) {
			window.cancelAnimationFrame(animationFrameId);
			if (resizeObserver) {
				resizeObserver.disconnect();
			}
		}
	});

	$: getCosScale = () =>
		Math.cos((Math.PI / 180) * 45) / Math.cos((Math.PI / 180) * ((rotation % 90) - 45));

	$: gradientBackground = Array.from({ length: 4 }, (_, i) => {
		const scale = getCosScale();
		const x = `calc(${xPosition}% + 50% * ${scale} * cos(${rotation + 90 * i}deg))`;
		const y = `calc(${yPosition}% + 50% * ${scale} * sin(${rotation + 90 * i}deg))`;
		return `radial-gradient(circle at ${x} ${y}, hsl(${90 * i - 15}, 100%, 50%), transparent)`;
	}).join(', ');
</script>

<div
	bind:this={gradientDiv}
	class={cn('gradient-div relative', className)}
	style:background={gradientBackground}
	style:transform="rotate(90deg)"
	{...$$restProps}
>
	<slot />
</div>
