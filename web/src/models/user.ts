export interface AddUser {
	email: string;
	password: string;
	is_admin: boolean;
}

export interface User {
	id: number;
	email: string;
	is_admin: boolean;
	created_at: string;
}
