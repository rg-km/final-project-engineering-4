import { SERVICE_KELAS } from '@services/kelas.service';
import create from 'zustand';
import { persist, devtools } from 'zustand/middleware';

const useKelasStore = create(
  devtools(
    persist(
      (set) => ({
        payload: null,
        isLoading: false,
        handleGetListKelas: async () => {
          set({ isLoading: true });

          const response = await SERVICE_KELAS.getListKelas();

          set({
            payload: response.payload,
            isLoading: false,
          });
        },
      }),
      {
        name: 'kelas-store',
      }
    )
  )
);

export { useKelasStore };
