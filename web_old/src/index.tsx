import ReactDOM from 'react-dom/client';
import './index.css';
import {
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import { IndexView } from './views';
import { SigninView } from './views/signin';
import { SignupView } from './views/signup';
import { AuthProvider } from './context/auth-context';
import { PersistLogin } from './components/persist-login';
import { RequireNoAuth } from './components/require-no-auth';
import { RequireAuth } from './components/require-auth';
import { CssBaseline } from '@mui/material';

import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import { AdminDrawer } from './features/menus/admin-drawer';
import { AdminOverviewView } from './views/admin-overview';
import { NotFoundView } from './views/notfound';
import { AdminUsersView } from './views/admin-users';
import { SnackbarProvider } from 'notistack';
import { AdminDecksView } from './views/admin-decks';
import { AdminCardsView } from './views/admin-cards';
import { UserDrawer } from './features/menus/user-drawer';
import { PremadeView } from './views/premade';
import { PremadeCardsView } from './views/premade-cards';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  //<React.StrictMode>
    <SnackbarProvider maxSnack={3}>
      <AuthProvider>
        <CssBaseline />
        <BrowserRouter>
          <Routes>
            <Route path="/">
              <Route element={<PersistLogin />} >
                <Route index element={<IndexView />} />
                <Route element={<RequireNoAuth />}>
                  <Route path="signin" element={<SigninView />} />
                  <Route path="signup" element={<SignupView />} />
                </Route>
                <Route element={<RequireAuth allowedRoles={[1]} />} >
                  <Route path="admin">
                    <Route element={<AdminDrawer />}>
                      <Route path="overview" element={<AdminOverviewView />} />
                      <Route path="users" element={<AdminUsersView />} />
                      <Route path="decks">
                        <Route index element={<AdminDecksView />} />
                        <Route path=":deckId" element={<AdminCardsView />} />
                      </Route>
                    </Route>
                  </Route>
                  <Route element={<UserDrawer />}>
                    <Route path="learn" element={<h1>learn</h1>} />
                    <Route path="statistics" element={<h1>statistics</h1>} />
                    <Route path="cards" element={<h1>cards</h1>} />
                    <Route path="premade">
                      <Route index element={<PremadeView />} />
                      <Route path=":deckId" element={<PremadeCardsView />} />
                    </Route>
                  </Route>
                </Route>
              </Route>
            </Route>
            <Route path="*" element={<NotFoundView />} />
          </Routes>
        </BrowserRouter>
      </AuthProvider>
    </SnackbarProvider>
  //</React.StrictMode>
);
