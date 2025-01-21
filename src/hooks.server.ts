import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const theme = event.cookies.get('theme') || 'dark';
	const actualTheme = theme === 'system' ? 'dark' : theme;

	return await resolve(event, {
		transformPageChunk: ({ html }) => {
			return html.replace(
				'<html lang="en"',
				`<html lang="en" class="${actualTheme}" style="color-scheme: ${actualTheme}"`
			);
		}
	});
};
