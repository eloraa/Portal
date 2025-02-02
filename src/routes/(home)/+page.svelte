<script lang="ts">
	import GradientBg from '@/components/common/gradient-bg/gradient-bg.svelte';
	import { Button } from '@/components/ui/button';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { PUBLIC_API_URL } from '$env/static/public';
	import { user } from '@/stores/user';

	let inputRefs = Array(8).fill(null);
	let values = $state(Array(8).fill(''));

	function handlePaste(event: ClipboardEvent, index: number) {
		event.preventDefault();
		const pastedText = event.clipboardData?.getData('text') || '';
		const chars = pastedText.split('').slice(0, 8 - index);

		chars.forEach((char, i) => {
			if (index + i < 8) {
				values[index + i] = char;
			}
		});
		values = values;

		const nextEmptyIndex = values.findIndex((v, i) => i > index && !v);
		if (nextEmptyIndex !== -1 && nextEmptyIndex < 8) {
			inputRefs[nextEmptyIndex].focus();
		} else {
			const lastFilledIndex = Math.min(index + chars.length - 1, 7);
			inputRefs[lastFilledIndex].focus();
		}
	}

	function handleInput(
		event: Event & { currentTarget: EventTarget & HTMLInputElement },
		index: number
	) {
		values[index] = event.currentTarget.value.slice(-1);
		values = values;

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

	async function createRoom() {
		try {
			const response = await fetch(`${PUBLIC_API_URL}/api/rooms`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					isPublic: true,
					userId: $user.userId
				})
			});

			if (!response.ok) {
				throw new Error('Failed to create room');
			}

			const data = await response.json();
			goto(`/room/${data.roomId}`);
		} catch (error) {
			console.error('Error creating room:', error);
		}
	}

	onMount(() => {
		if (inputRefs[0]) {
			inputRefs[0].focus();
		}
	});

	const { hideFeature } = $props<{ hideFeature: boolean }>();
</script>

<main class="container relative my-auto space-y-6 py-6">
	<h1 class="relative text-2xl md:text-6xl">
		Easily and Securely connect <span>
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

	<div>
		<div>
			<div class="inline-flex">
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
									onpaste={(e) => handlePaste(e, i)}
									onkeydown={(e) => handleKeydown(e, i)}
									maxlength={1}
									type="text"
									class="h-full w-full bg-transparent text-center font-mono outline-none focus:outline-none focus:ring-0 focus-visible:ring-0 active:outline-none"
								/>
							</li>
						{/each}
					</ul>
				</div>
			</div>
		</div>
		<div>
			<Button
				class="relative mt-4 h-12 w-full overflow-hidden rounded-md bg-transparent text-base text-white hover:bg-accent/15 md:max-w-md"
			>
				<span
					class="absolute inset-0 -z-[1] rounded-md"
					style="background: linear-gradient(87.81deg, #E6315C 1.77%, #FE013C 24.55%, #FE0140 57.03%, #FE0177 73.99%, #F5466F 98.71%);"
				></span>

				<!-- <GradientBg class="absolute inset-0 -z-[1] rounded-md" /> -->
				Join a Portal
			</Button>
		</div>
		<div>
			<Button
				class="inine-flex relative mt-2 h-12 w-full overflow-hidden rounded-md text-base md:max-w-md"
				on:click={createRoom}
			>
				Start a New Portal
			</Button>
		</div>
	</div>

	{#if !hideFeature}
		<div class="grid text-sm md:max-w-4xl md:grid-cols-3">
			<div class="flex h-full flex-col border-accent/15 max-md:py-4 md:border-r md:p-4 md:pl-0">
				<div class="flex items-center gap-2">
					<figure class="h-8 w-8">
						<img src="/lock.png" alt="Secure" class="h-full w-full" />
					</figure>
					<h1 class="font-mono font-medium uppercase">Secure Connections</h1>
				</div>
				<p class="ml-0.5 mt-2 text-foreground/60 dark:font-light">
					Protect your data with end-to-end encryption, ensuring total privacy.
				</p>
			</div>

			<div class="flex h-full flex-col border-accent/15 max-md:py-4 md:border-r md:p-4">
				<div class="flex items-center gap-2">
					<figure class="h-8 w-8">
						<img src="/zap.png" alt="Zap" class="h-full w-full" />
					</figure>
					<h1 class="font-mono font-medium uppercase">Lightning-Fast Sharing</h1>
				</div>
				<p class="ml-0.5 mt-2 text-foreground/60 dark:font-light">
					Share files instantly without delays, no matter the size.
				</p>
			</div>

			<div class="flex h-full flex-col max-md:py-4 md:p-4">
				<div class="flex items-center gap-2">
					<figure class="h-8 w-8">
						<img src="/globe.png" alt="Globe" class="h-full w-full" />
					</figure>
					<h1 class="font-mono font-medium uppercase">Connect Anywhere, Anytime</h1>
				</div>
				<p class="ml-0.5 mt-2 text-foreground/60 dark:font-light">
					Seamlessly link with peers worldwide need. flex-col
				</p>
			</div>
		</div>
	{/if}
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
