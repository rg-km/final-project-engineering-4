import { SERVICE_AUTH } from '@services/auth.service';
import create from 'zustand';

const useAuthStore = create((set) => ({
  token: null,
  user: null,
  isLoading: false,
  error: null,
  action: {
    login: async (role, email, password) => {
      set({ isLoading: true });

      const response = await SERVICE_AUTH.login(role, email, password);

      if (response.success) set({ token: response.payload.token, user: response.payload.user, isLoading: false });
      else set({ error: response, isLoading: false });
    },
  },
}));

export { useAuthStore };
