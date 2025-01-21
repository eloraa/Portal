// place files you want to import through the `$lib` alias in this folder.
export function lerp(start: number, end: number, factor: number) {
	return start + (end - start) * factor;
}
