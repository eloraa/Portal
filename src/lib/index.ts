// place files you want to import through the `$lib` alias in this folder.
export function lerp(start: number, end: number, factor: number) {
	return start + (end - start) * factor;
}

export const backendHealthChecker = async (API_URL: string) => {
	try {
		const response = await fetch(`${API_URL}/health`, {
			headers: {
				'Content-Type': 'application/json'
			}
		});

		if ((await response.json())?.status === 'ok') return true;

		return false;
	} catch {
		return false;
	}
};
