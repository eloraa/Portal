import type { LayoutServerLoad } from './$types';
import type { Theme } from '@/stores/theme';
import { generateRandomUsername, getRandomAvatarId } from '@/consts';

export const load: LayoutServerLoad = async ({ cookies }) => {
    const theme = (cookies.get('theme') || 'dark') as Theme;
    
    // Check if this is the first visit by looking for username cookie
    let username = cookies.get('username');
    let avatarId = cookies.get('avatarId');
    
    if (!username) {
        // First visit - generate random username and avatar
        username = generateRandomUsername();
        avatarId = getRandomAvatarId();
        
        // Set cookies with client-side access enabled
        cookies.set('username', username, {
            path: '/',
            maxAge: 60 * 60 * 24 * 365, // 1 year
            httpOnly: false, // Allow client-side access
            secure: false, // Allow non-HTTPS access
            sameSite: 'lax'
        });
        
        cookies.set('avatarId', avatarId, {
            path: '/',
            maxAge: 60 * 60 * 24 * 365, // 1 year
            httpOnly: false, // Allow client-side access
            secure: false, // Allow non-HTTPS access
            sameSite: 'lax'
        });
    }

    // Determine the actual theme class for server-side rendering
    let themeClass = theme;
    if (theme === 'system') {
        // Default to dark for system theme on server
        themeClass = 'dark';
    }

    return {
        theme,
        username,
        avatarId,
        themeClass
    };
}; 