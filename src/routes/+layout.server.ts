import type { LayoutServerLoad } from './$types';
import type { Theme } from '@/stores/theme';
import { generateRandomUsername, getRandomAvatarId } from '@/consts';
import { env } from '$env/dynamic/private';

export const load: LayoutServerLoad = async ({ cookies }) => {
	const theme = (cookies.get('theme') || 'dark') as Theme;

	let username = cookies.get('username');
	let avatarId = cookies.get('avatarId');
	let userId = cookies.get('userId');

	if (!userId) {
		username = generateRandomUsername();
		avatarId = getRandomAvatarId();

		try {
			const response = await fetch(`${env.API_URL}/api/users`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({ username, avatarId })
			});

			if (!response.ok) {
				throw new Error('Failed to create user');
			}

			const data = await response.json();
			userId = data.userId;

			if (userId) {
				cookies.set('userId', userId, {
					path: '/',
					maxAge: 60 * 60 * 24 * 365, // 1 year
					httpOnly: false,
					secure: false,
					sameSite: 'lax'
				});
			} else {
				console.error('No userId returned from API');
			}
		} catch (error) {
			console.error('Error creating user:', error);
		}

		cookies.set('username', username, {
			path: '/',
			maxAge: 60 * 60 * 24 * 365,
			httpOnly: false,
			secure: false,
			sameSite: 'lax'
		});

		cookies.set('avatarId', avatarId, {
			path: '/',
			maxAge: 60 * 60 * 24 * 365,
			httpOnly: false,
			secure: false,
			sameSite: 'lax'
		});
	}

	let themeClass = theme;
	if (theme === 'system') {
		themeClass = 'dark';
	}

	return {
		theme,
		username,
		avatarId,
		userId,
		themeClass
	};
};
