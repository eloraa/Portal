<main class="container py-6">
	<h1 class="text-2xl md:text-6xl">
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
		color: #FFD700;
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
