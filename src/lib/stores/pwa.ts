import { writable, get } from 'svelte/store';

// Define the BeforeInstallPromptEvent type since it's not in standard lib
export interface BeforeInstallPromptEvent extends Event {
    prompt: () => Promise<void>;
    userChoice: Promise<{ outcome: 'accepted' | 'dismissed' }>;
}

type PWAStore = {
    showInstallButton: boolean;
    isInstalling: boolean;
    deferredPrompt: BeforeInstallPromptEvent | null;
};

function createPWAStore() {
    const { subscribe, set, update } = writable<PWAStore>({
        showInstallButton: false,
        isInstalling: false,
        deferredPrompt: null
    });

    return {
        subscribe,
        set,
        update,
        setInstalling: (isInstalling: boolean) => {
            update(store => ({ ...store, isInstalling }));
        },
        setShowInstallButton: (showInstallButton: boolean) => {
            update(store => ({ ...store, showInstallButton }));
        },
        setDeferredPrompt: (deferredPrompt: BeforeInstallPromptEvent | null) => {
            update(store => ({ ...store, deferredPrompt }));
        }
    };
}

export const pwa = createPWAStore();

export async function installPWA() {
    const store = get(pwa);

    if (store.deferredPrompt) {
        await store.deferredPrompt.prompt();
        const { outcome } = await store.deferredPrompt.userChoice;
        pwa.setDeferredPrompt(null);
        if (outcome === 'accepted') {
            pwa.setShowInstallButton(false);
        }
    }
} 