import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import './index.css';
import { HomePage } from './src/pages/home';
import { RegisterPage } from './src/pages/register';
import { LoginPage } from './src/pages/login';

import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import { SnackbarProvider } from 'notistack';
import { User } from './src/auth/user';
import { AuthProvider } from './src/auth/provider';
import { RequireAuth } from './src/auth/require-auth';
import { AdminUsersPage } from './src/pages/admin-users';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  // <React.StrictMode>
    <SnackbarProvider>
      <AuthProvider>
        <BrowserRouter>
          <Routes>
            <Route element={<User />}>
              <Route path="/">
                <Route index element={<HomePage />} />
                <Route path="register" element={<RegisterPage />} />
                <Route path="login" element={<LoginPage />} />
                <Route path="users" element={<AdminUsersPage />} />
              </Route>
            </Route>
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </SnackbarProvider>
  // </React.StrictMode>
)
