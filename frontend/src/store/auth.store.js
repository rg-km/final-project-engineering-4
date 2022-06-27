import { SERVICE_AUTH } from '@services/auth.service';
import create from 'zustand';
import { persist, devtools } from 'zustand/middleware';

const useAuthStore = create(
  devtools(
    persist(
      (set) => ({
        token: null,
        user: null,
        role: null,
        isLoading: false,
        handleLogin: async (role, email, password, callback) => {
          set({ isLoading: true });

          const response = await SERVICE_AUTH.login(role.value, email, password);

          set({
            token: response.success ? response.payload.token : null,
            user: response.success ? response.payload.user : null,
            role,
            isLoading: false,
          });

          callback(response.success, response.message);
        },
        handleRegister: async (role, data, callback) => {
          set({ isLoading: true });

          const response = await SERVICE_AUTH.register(role, data);

          set({ isLoading: false });

          callback(response.success, response.message);
        },
        handleLogout: () => set({ token: null, user: null, role: null }),
      }),
      {
        name: 'auth-store',
      }
    )
  )
);

export { useAuthStore };
