import type { SvelteHTMLElements } from 'svelte/elements';

export type WithElementRef<T> = T & {
    ref?: HTMLElement | null;
};

export type HTMLAttributes<T extends keyof SvelteHTMLElements> = SvelteHTMLElements[T]; 