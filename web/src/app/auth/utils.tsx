export function logged_in(s: number): boolean {
	return s === 200;
}

export function need_login(s: number): boolean {
	return s === 401;
}
