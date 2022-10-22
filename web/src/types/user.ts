export interface IUser {
  userId: number;
  email: string;
  password: string;
  role: number;
}

export interface IUserResponsePayload {
  userId: number;
  email: string;
  role: number;
  createdAt: string;
}
