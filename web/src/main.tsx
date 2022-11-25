import ReactDOM from 'react-dom/client';
import './index.css';
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import { SnackbarProvider } from 'notistack';
import { AuthProvider } from '@/providers/auth-provider';
import { MainRoutes } from '@/routes';

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  // <React.StrictMode>
    <SnackbarProvider>
      <AuthProvider>
        <MainRoutes />
      </AuthProvider>
    </SnackbarProvider>
  // </React.StrictMode>
)
