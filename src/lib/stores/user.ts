import { writable } from 'svelte/store';

interface User {
    username: string;
    avatarId: string;
    userId: string | null;
}

function createUserStore() {
    const { subscribe, set, update } = writable<User>({
        username: '',
        avatarId: '',
        userId: null
    });

    return {
        subscribe,
        set,
        update,
        setUsername: (username: string) => {
            update(user => ({ ...user, username }));
            document.cookie = `username=${username};path=/;max-age=31536000`;
        },
        setAvatarId: (avatarId: string) => {
            update(user => ({ ...user, avatarId }));
            document.cookie = `avatarId=${avatarId};path=/;max-age=31536000`;
        },
        setUserId: (userId: string) => {
            update(user => ({ ...user, userId }));
        }
    };
}

export const user = createUserStore();
