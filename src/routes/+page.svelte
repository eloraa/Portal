<script lang="ts">
	import GradientBg from '@/components/common/gradient-bg/gradient-bg.svelte';
	import { onMount, tick } from 'svelte';

	let inputRefs = Array(8).fill(null);
	let values = Array(8).fill('');

	async function handleInput(event: Event & { currentTarget: EventTarget & HTMLInputElement }, index: number) {
		// Get the value and ensure it's only one character
		values[index] = event.currentTarget.value.slice(-1);
		values = values;

		// Wait for DOM update
		await tick();

		// Move to next input if we have a value
		if (values[index] && index < 7 && inputRefs[index + 1]) {
			inputRefs[index + 1].focus();
		}
	}

	async function handleKeydown(event: KeyboardEvent, index: number) {
		if (event.key === 'Backspace') {
			event.preventDefault();

			if (!values[index] && index > 0 && inputRefs[index - 1]) {
				inputRefs[index - 1].focus();
			} else {
				values[index] = '';
				values = values;
			}
		}
	}

	onMount(() => {
		if (inputRefs[0]) {
			inputRefs[0].focus();
		}
	});
</script>

<main class="container relative py-6">
	<h1 class="relative text-2xl md:text-6xl">
		Easily connect <span>
			<img
				src="/connect.png"
				alt="Connect"
				class="inline-block h-10 w-10 rounded-xl bg-foreground p-0.5 md:h-12 md:w-12 md:rounded-2xl md:align-baseline"
			/>
		</span>, share
		<span>
			<img
				src="/share.png"
				alt="Share"
				class="inline-block h-10 w-10 rounded-xl bg-foreground p-0.5 md:h-12 md:w-12 md:rounded-2xl md:align-baseline"
			/>
		</span>, and communicate
		<span>
			<img
				src="/communicate.png"
				alt="Communicate"
				class="inline-block h-10 w-10 rounded-xl bg-foreground p-0.5 md:h-12 md:w-12 md:rounded-2xl md:align-baseline"
			/>
		</span>
		with anyone, anytime
		<span class="sparkle-wrapper">
			<span>✦</span>
			<img
				src="/flower.png"
				alt="Flower"
				class="inline-block h-10 w-10 p-0.5 md:h-12 md:w-12 md:align-baseline"
			/>
		</span>
	</h1>

	<div class="mt-6 md:mt-16">
		<div class="flex">
			<div
				class="group relative flex h-12 items-center justify-center overflow-hidden rounded-md border border-accent/60 focus-within:ring-2 focus-within:ring-primary/15 md:h-14"
			>
				<GradientBg className="w-full h-[1001px] absolute opacity-15 -z-[1]" />

				<ul class="flex h-full w-full">
					{#each Array(8) as _, i}
						<li class="h-full {i < 7 ? 'border-r' : ''} border-accent/60 md:w-14">
							<input
								bind:this={inputRefs[i]}
								bind:value={values[i]}
								oninput={(e) => handleInput(e, i)}
								onkeydown={(e) => handleKeydown(e, i)}
								maxlength={1}
								type="text"
								inputmode="numeric"
								class="h-full w-full text-center font-mono outline-none focus:outline-none focus:ring-0 focus-visible:ring-0 active:outline-none bg-transparent"
							/>
						</li>
					{/each}
				</ul>
			</div>
		</div>
	</div>
</main>

<style>
	.sparkle-wrapper {
		position: relative;
	}

	.sparkle-wrapper::before,
	.sparkle-wrapper::after,
	.sparkle-wrapper > span {
		content: '✦';
		position: absolute;
		color: #ffd700;
		animation: sparkle 2s linear infinite;
		font-size: 0.4em;
	}

	.sparkle-wrapper::before {
		top: -1px;
		right: -1px;
		animation-delay: -1s;
	}

	.sparkle-wrapper::after {
		bottom: -1px;
		left: -1px;
		animation-delay: -0.5s;
	}

	@keyframes sparkle {
		0% {
			opacity: 0;
			transform: rotate(0deg) scale(0.5);
		}
		50% {
			opacity: 1;
			transform: rotate(180deg) scale(1);
		}
		100% {
			opacity: 0;
			transform: rotate(360deg) scale(0.5);
		}
	}
</style>
