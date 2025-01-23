<script lang="ts">
	import Avatar from '@/components/ui/avatar/avatar.svelte';
	import AvatarFallback from '@/components/ui/avatar/avatar-fallback.svelte';
	import AvatarImage from '@/components/ui/avatar/avatar-image.svelte';
	import { Tooltip, TooltipTrigger, TooltipContent } from '@/components/ui/tooltip';
	import { user } from '@/stores/user';
	import { avatars } from '@/consts';

	const { onClick = () => {} } = $props<{
		onClick?: () => void;
	}>();

	let username = $state<string>($user.username);
	let avatarSrc = $state<string>(getAvatarSrcFromId($user.avatarId));

	$effect(() => {
		username = $user.username;
		avatarSrc = getAvatarSrcFromId($user.avatarId);
	});

	function getAvatarSrcFromId(id: string) {
		const avatar = avatars.find((opt) => opt.id === id);
		return avatar?.src || '/kazuha.png';
	}
</script>

<Tooltip openDelay={0}>
	<TooltipTrigger
		class="flex h-11 w-11 items-center justify-center rounded-full transition-colors hover:bg-popover/50"
		onclick={onClick}
	>
		<Avatar
			class="flex h-8 w-8 items-center justify-center overflow-visible bg-muted/85 backdrop-blur-xl"
		>
			<AvatarImage src={avatarSrc} class="absolute h-14 w-14 min-w-14 object-contain" />
			<AvatarFallback>{username[0]}</AvatarFallback>
		</Avatar>
	</TooltipTrigger>
	<TooltipContent>{username}</TooltipContent>
</Tooltip>
