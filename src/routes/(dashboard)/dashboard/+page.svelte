<script>
	import { Home } from '../../(home)';
</script>

<div class="relative flex h-full flex-col">
	<div
		class="pointer-events-none fixed inset-0 -z-[1] ml-auto w-[calc(100%-195px)] origin-top scale-90 xl:scale-150"
	>
		<div class="rayOne bg-light-ray-primary absolute rounded-full"></div>
		<div class="rayTwo bg-light-ray absolute rounded-full"></div>
		<div class="rayThree bg-light-ray absolute rounded-full"></div>
		<div class="rayFour bg-light-ray absolute rounded-full"></div>
		<div class="rayFive bg-light-ray absolute rounded-full"></div>
	</div>
	<div class="my-auto text-center">
		<Home hideFeature />
	</div>
</div>

<style>
	.rayOne {
		width: 480px;
		height: 680px;
		transform: rotate(80deg);
		top: -540px;
		left: calc(50% - 240px);
		filter: blur(110px);
	}
	.rayTwo {
		width: 110px;
		height: 400px;
		transform: rotate(-20deg);
		top: -280px;
		left: calc(50% + 100px);
		mix-blend-mode: overlay;
		opacity: 0.6;
		filter: blur(60px);
	}
	.rayThree {
		width: 400px;
		height: 370px;
		top: -350px;
		left: calc(50% - 100px);
		mix-blend-mode: overlay;
		opacity: 0.6;
		filter: blur(21px);
	}
	.rayFour {
		width: 330px;
		height: 370px;
		top: -330px;
		left: calc(50% - 240px);
		mix-blend-mode: overlay;
		opacity: 0.5;
		filter: blur(21px);
	}
	.rayFive {
		width: 110px;
		height: 400px;
		transform: rotate(-40deg);
		top: -280px;
		left: calc(50% - 325px);
		mix-blend-mode: overlay;
		opacity: 0.8;
		filter: blur(60px);
	}
</style>
