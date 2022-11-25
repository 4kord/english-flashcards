import { PersistUser } from "@/components/persist-user";
import { MenuBar } from "@/features/menu/component";
import { AdminCardsPage } from "@/pages/admin-cards";
import { AdminDecksPage } from "@/pages/admin-decks";
import { AdminUsersPage } from "@/pages/admin-users";
import { HomePage } from "@/pages/home";
import { LoginPage } from "@/pages/login";
import { RegisterPage } from "@/pages/register";
import * as React from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

export const MainRoutes: React.FC = () => {
	return (
		<BrowserRouter>
			<Routes>

				<Route element={<PersistUser />}>

					<Route path="/">
						<Route index element={<HomePage />} />
						<Route path="register" element={<RegisterPage />} />
						<Route path="login" element={<LoginPage />} />
					</Route>

					<Route path="admin" element={<MenuBar />}>
						<Route path="overview" element={<h1>Overview</h1>} />
						<Route path="users" element={<AdminUsersPage />} />
						<Route path="decks" element={<AdminDecksPage /> } />
						<Route path="decks/:deckId" element={<AdminCardsPage /> } />
					</Route>

				</Route>

			</Routes>
		</BrowserRouter>
	);
}
