import { SERVICE_AUTH } from '@services/auth.service';
import create from 'zustand';

const useAuthStore = create((set) => ({
  state: {
    token: null,
    user: null,
  },
  action: {
    login: async (role, email, password) => {
      const response = await SERVICE_AUTH.login(role, email, password);
      if (response.success) {
        set({ token: response.payload.token, user: response.payload.user });
      }
      console.log('zustand', { response });
    },
  },
}));

export { useAuthStore };
